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

	listing := &pb.Listing{
		ListingId:      listingObject.ListingId,
		Address:        listingObject.StreetAddress,
		City:           listingObject.City,
		State:          listingObject.State,
		Zipcode:        listingObject.Zipcode,
		Price:          listingObject.Price,
		CoordLat:       listingObject.CoordLat,
		CoordLong:      listingObject.CoordLong,
		Area:           listingObject.Area,
		SchoolDistrict: listingObject.SchoolDistrict,
		ImageKey:       listingObject.ImageKey,
	}

	return &pb.GetListingDetailResponse{
		Listing: listing,
	}, nil
}

func (s *Server) GetListingByAddress(ctx context.Context, request *pb.GetListingByAddressRequest) (*pb.GetListingByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
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
				ListingId:      listing.ListingId,
				Address:        listing.StreetAddress,
				City:           listing.City,
				State:          listing.State,
				Zipcode:        listing.Zipcode,
				Price:          listing.Price,
				CoordLat:       listing.CoordLat,
				CoordLong:      listing.CoordLong,
				Area:           listing.Area,
				SchoolDistrict: listing.SchoolDistrict,
				ImageKey:       listing.ImageKey,
			}

			result = append(result, convertedListing)
		}
	}

	return &pb.GetListingsResponse{
		Listings: result,
	}, nil
}
