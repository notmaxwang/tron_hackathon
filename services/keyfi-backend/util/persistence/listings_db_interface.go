package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type ListingsDao struct {
	connection *sql.DB
}

type ListingObject struct {
	ListingId      string
	WalletAddress  string
	SchoolDistrict string
	StreetAddress  string
	City           string
	State          string
	Zipcode        int
	Area           int
	Beds           int
	Baths          int
	HouseType      string
	Price          int
	ImageKey       string
}

var (
	listingsDao *ListingsDao
	once        sync.Once
)

func GetListingsDao() (*ListingsDao, error) {
	once.Do(func()  {
		// Connection parameters
		user := os.Getenv("MYSQL_LOGIN")
		password := os.Getenv("MYSQL_PASS")
		host := "database-1.cvueska0quuw.us-east-1.rds.amazonaws.com"
		port := "3306" // Default MySQL port
		dbname := ""

		// Construct connection string
		connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

		// Open the database connection
		db, err := sql.Open("mysql", connectionString)
		if err != nil {
			log.Println("error while opening SQL connection", err)
			return
		}
		listingsDao = &ListingsDao{
			connection: db,
		}
	})
	return listingsDao, nil
}

func (dao *ListingsDao) QueryListingDetail(listingId string) (*ListingObject, error) {
	rows, err := dao.Query("SELECT * FROM your_table WHERE listing_id = ?", listingId)
	if err != nil {
		log.Println("couldnt query DB", err)
		return nil, err
	}
	defer rows.Close()

	// Slice to hold the retrieved listings
	var listings []ListingObject

	for rows.Next() {
		var listing ListingObject
		// Scan the values from the row into the struct fields
		if err := rows.Scan(&listing.ListingId, &listing.WalletAddress, &listing.SchoolDistrict, &listing.StreetAddress, &listing.City, &listing.State, &listing.Zipcode, &listing.Area, &listing.Beds, &listing.Baths, &listing.HouseType, &listing.Price, &listing.ImageKey); err != nil {
			log.Fatal(err)
		}
		// Append the filled struct to the slice
		listings = append(listings, listing)
	}
	if err := rows.Err(); err != nil {
		log.Println("error while reading from rows", err)
		return nil, err
	}

	var result *ListingObject
	if len(listings) > 0 {
		result = &listings[0]
	}

	return result, nil
}

func (dao *ListingsDao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return dao.connection.Query(query, args...)
}

func (dao *ListingsDao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return dao.connection.Exec(query, args...)
}
