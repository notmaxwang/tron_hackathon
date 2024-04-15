package middleware

import (
	"context"
	"keyfi-backend/util/cryptography"
	"keyfi-backend/util/persistence"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func SignatureValidator(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Some APIs don't need to validate signatures like the auth APIs
	if shouldSkipValidation(info.FullMethod) {
		return handler(ctx, req)
	}

	// Get metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	// Extract the signature and message from metadata
	signatureList, ok := md["signature"]
	if !ok || len(signatureList) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing signature")
	}
	signature := signatureList[0]
	walletAddressList, ok := md["walletAddress"]
	if !ok || len(walletAddressList) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing walletAddress")
	}
	walletAddress := walletAddressList[0]

	dao, err := persistence.GetMainTableDao()
	if err != nil {
		log.Fatalf("failed to initialize DB Dao\n")
		return nil, status.Errorf(codes.Internal, "failed to access DB")
	}

	item, err := dao.GetItem(walletAddress)
	if err != nil {
		log.Fatalf("failed to read item for\n")
		return nil, status.Errorf(codes.Internal, "failed to access DB")
	}

	if item == nil {
		log.Printf("wallet not found in DB\n")
		return nil, status.Errorf(codes.NotFound, "walletAddress not found in DB")
	}

	verify, err := cryptography.ValidateDefaultMessage(item.SignatureExpiry, signature, walletAddress)
	if err != nil {
		log.Printf("Error while verifying signature\n")
		return nil, status.Errorf(codes.Internal, "failed to verify signature")
	}

	if !verify {
		log.Printf("request denied because the signature is not verified\n")
		return nil, status.Errorf(codes.PermissionDenied, "login denied")
	}

	// Proceed with the gRPC handler
	return handler(ctx, req)
}

func shouldSkipValidation(method string) bool {
	// Skip this for auth because we expect the user to not have a valid login
	if method == "/AuthenticationService/Login" || method == "/AuthenticationService/Register" {
		return true
	}
	return false
}

// UnaryInterceptor returns a new unary server interceptor that includes the middleware.
func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return SignatureValidator
}
