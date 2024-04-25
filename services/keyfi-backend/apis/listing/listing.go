package listing

import (
	"context"
	pb "keyfi-backend/protos/listing"
	"keyfi-backend/util/persistence"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedListingServiceServer
}

func (s *Server) GetListingDetail(ctx context.Context, request *pb.GetListingDetailRequest) (*pb.GetListingDetailResponse, error) {
	if request.ListingId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing listingId")
	}

	listingsDao, err := persistence.GetListingsDao()
	if err != nil {
		log.Println("couldnt init listings DAO for GetListingsDetail", err)
		return nil, status.Errorf(codes.Internal, "could not init DB")
	}

	listingObject, err := listingsDao.QueryListingDetail(request.ListingId)
	if err != nil {
		log.Println("error querying DB")
		return nil, status.Errorf(codes.Internal, "failed to query DB")
	}

	listingDetails := &pb.ListingDetail{
		ListingId:      request.ListingId,
		WalletAddress:  listingObject.WalletAddress,
		Area:           string(listingObject.Area),
		SchoolDistrict: listingObject.SchoolDistrict,
		Bed:            string(listingObject.Beds),
		Bath:           string(listingObject.Baths),
		HouseType:      listingObject.HouseType,
	}

	return &pb.GetListingDetailResponse{
		ListingDetails: listingDetails,
	}, nil
}
