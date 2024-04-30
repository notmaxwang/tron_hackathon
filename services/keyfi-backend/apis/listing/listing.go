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

func (s *Server) GetListings(ctx context.Context, request *pb.GetListingsRequest) (*pb.GetListingsResponse, error) {
	if len(request.Cities) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "no query params")
	}

	listingsDao, err := persistence.GetListingsDao()
	if err != nil {
		log.Println("couldnt init listings DAO for GetListingsDetail", err)
		return nil, status.Errorf(codes.Internal, "could not init DB")
	}

	result := make([]*pb.Listing, 0)

	for _, city := range request.Cities {
		listings, err := listingsDao.QueryAllListingsInCity(city)
		if err != nil {
			log.Println("error querying DB")
			return nil, status.Errorf(codes.Internal, "failed to query DB")
		}
		for _, listing := range *listings {
			convertedListing := &pb.Listing{
				ListingId: listing.ListingId,
				Address: listing.StreetAddress,
				City: listing.City,
				State: listing.State,
				Zipcode: listing.Zipcode,
				Price: listing.Price,
				CoordLat: listing.CoordLat,
				CoordLong: listing.CoordLong,
				Area: listing.Area,
				SchoolDistrict: listing.SchoolDistrict,
				ImageKey: listing.ImageKey,
			}

			result = append(result, convertedListing)
		}
	}

	return &pb.GetListingsResponse{
		Listings: result,
	}, nil
}
