package auth

import (
	"context"
	pb "keyfi-backend/protos/auth"
	"keyfi-backend/util/cryptography"
	"keyfi-backend/util/persistence"
	"keyfi-backend/util/persistence/models"
	"log"
	"time"

	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAuthenticationServiceServer
}

func (s *Server) Login(ctx context.Context, request *pb.AuthRequest) (*pb.AuthResponse, error) {
	if request == nil || request.WalletAddress == nil || request.Signature == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: walletAddress and/or signature are empty")
	}

	log.Printf("Incoming login for: %s\n", request.WalletAddress)

	dao, err := persistence.GetMainTableDao()
	if err != nil {
		log.Fatalf("failed to initialize DB Dao\n")
		return nil, status.Errorf(code.Code_INTERNAL, "failed to access DB")
	}

	item, err := dao.GetItem(request.WalletAddress)
	if err != nil {
		log.Fatalf("failed to read item for %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_INTERNAL, "failed to access DB")
	}

	if item == nil {
		log.Printf("wallet not found in DB: %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_NOT_FOUND, "walletAddress not found in DB")
	}

	if item.SignatureExpiry > request.SignatureExpiry {
		log.Printf("login attempt denied because expiry date less than saved: %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_PERMISSION_DENIED, "login denied")
	}

	verify, err := cryptography.ValidateDefaultMessage(item.SignatureExpiry, request.Signature, request.WalletAddress)
	if err != nil {
		log.Printf("Error while verifying signature for %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_INTERNAL, "failed to verify signature")
	}

	if !verify {
		log.Printf("login attempt denied because signature cannot be verified: %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_PERMISSION_DENIED, "login denied")
	}

	item.SignatureExpiry = request.SignatureExpiry
	err = dao.PutItem(item)
	if err != nil {
		log.Fatalf("failed to update item for %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_INTERNAL, "failed to update DB")
	}

	return &pb.AuthResponse{
		Success:       true,
	}, nil
}

func (s *Server) Register(ctx context.Context, request *pb.AuthRequest) (*pb.AuthResponse, error) {
	if request == nil || request.WalletAddress == nil || request.Signature == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: walletAddress and/or signature are empty")
	}

	if request.FirstName == nil || request.LastName == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: no name provided")
	}

	log.Printf("Incoming registration for: %s\n", request.WalletAddress)

	dao, err := persistence.GetMainTableDao()
	if err != nil {
		log.Fatalf("failed to initialize DB Dao\n")
		return nil, status.Errorf(code.Code_INTERNAL, "failed to access DB")
	}

	item, err := dao.GetItem(request.WalletAddress)
	if err != nil {
		log.Fatalf("failed to read item for %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_INTERNAL, "failed to access DB")
	}

	if item != nil {
		log.Printf("wallet already exists: %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_ALREADY_EXISTS, "an account for this address already exists")
	}

	now = time.Now().UTC().Unix()
	if request.SignatureExpiry <= now {
		log.Printf("login attempt denied because expiry date has passed: %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_PERMISSION_DENIED, "login denied")
	}

	verify, err := cryptography.ValidateDefaultMessage(request.SignatureExpiry, request.Signature, request.walletAddress)
	if err != nil {
		log.Printf("Error while verifying signature for %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_INTERNAL, "failed to verify signature")
	}

	if !verify {
		log.Printf("login attempt denied because signature cannot be verified: %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_PERMISSION_DENIED, "login denied")
	}

	item = models.UserProfileModel{
		WalletAddress:   request.WalletAddress,
		Signature:       request.Signature,
		SignatureExpiry: request.SignatureExpiry,
		FirstName:       request.FirstName,
		LastName:        request.LastName,
		MiddleName:      request.MiddleName,
		Role:            "user",
		CreationTime:    now,
	}

	err = dao.PutItem(item)
	if err != nil {
		log.Fatalf("failed to update item for %s\n", request.WalletAddress)
		return nil, status.Errorf(code.Code_INTERNAL, "failed to update DB")
	}

	return &pb.AuthResponse{
		Success:       true,
	}, nil
}
