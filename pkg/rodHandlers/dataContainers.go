package rodHandlers

type DataContainer struct {
	// https://api.target.com/location_proximities/v1/nearby_locations?
	Key        string
	Limit      int
	Unit       string
	Within     int
	Place      string
	NearbyType string

	// https://redsky.target.com/redsky_aggregations/v1/web/plp_search_v1?
	// key string
	Category                      string
	Channel                       string
	Count                         int
	Default_purchasability_filter bool
	Included_sponsored            bool
	Offset                        int
	Page                          string
	Platform                      string
	Useragent                     string

	Pricing_store_id            string // populated after store info retrieved
	Scheduled_delivery_store_id string // populated after store info retrieved
	Store_ids                   string // populated after store info retrieved
	Visitor_id                  string // populated after store info retrieved

	// https://redsky.target.com/redsky_aggregations/v1/web_platform/product_summary_with_fulfillment_v1?
	// key string
	Tcins string // populated after product info retrieved

	Zip                   string
	Has_required_store_id bool

	Store_id          string  // populated after store info retrieved
	State             string  // populated after store info retrieved
	Latitude          float64 // populated after store info retrieved
	Longitude         float64 // populated after store info retrieved
	Required_store_id string  // populated after store info retrieved
}

type Store struct {
	Company       string
	Zip_code      string
	Address_line1 string
	Address_line2 string
	City          string
	State         string
	Country       string
	Longitude     float64 // not part of final output
	Latitude      float64 // not part of final output
	Store_name    string  // not part of final output
	Store_id      float64 // not part of final output
	Store_items   []Product
}

type Product struct {
	Product_id     string // Tcin, which is not part of final output
	Product_family string
	Product        string
	Price          float64
	Availability   string
	Quantity       float64
	Product_url    string
}
