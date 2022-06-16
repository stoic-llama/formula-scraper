package rodHandlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	errorhandlers "github.com/stoic-llama/formula-scraper/pkg/errorHandlers"
)

func TargetScrape(zipcode string) {
	log.Printf("Starting job...")
	browser := rod.New().MustConnect()

	data := DataContainer{
		/*
			https://api.target.com/location_proximities/v1/nearby_locations?
		*/
		Key:        `9f36aeafbe60771e321a7cc95a78140772ab3e96`,
		Limit:      2000,
		Unit:       `mile`,
		Within:     25,
		Place:      zipcode,
		NearbyType: `store`,

		/*
			https://redsky.target.com/redsky_aggregations/v1/web/plp_search_v1?
		*/
		Category:                      `5xtkh`,
		Channel:                       `WEB`,
		Count:                         24,
		Default_purchasability_filter: true,
		Included_sponsored:            true,
		Offset:                        0,
		Page:                          `%2Fc%2F5xtkh`, // decoded: /c/5xtkh
		Platform:                      `desktop`,
		Useragent:                     `Mozilla%2F5.0+%28Macintosh%3B+Intel+Mac+OS+X+10_15_7%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F102.0.5005.61+Safari%2F537.36`,
		// decode: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36
		Tcins: "",

		/*
			https://redsky.target.com/redsky_aggregations/v1/web_platform/product_summary_with_fulfillment_v1?
		*/
		Zip:                   zipcode,
		Has_required_store_id: true,
	}

	/*
		1) Find 5 closest stores from zip code provided by user.
	*/

	stores := GetNearByStores(data)

	/*
		2) Set Default Store Location for inventory information. (first of five stores)
	*/

	SetStoreLocation(browser, stores[0])

	/*
		3) Populate the rest of the Datacontainer fields for API calls from nearby stores information.
	*/

	// Going to Target Site to get dynamically generated Visitor Id to pass onto Target APIs later.
	url := proto.TargetCreateTarget{
		URL: "https://www.target.com/c/formula-nursing-feeding-baby/-/N-5xtkh",
	}
	rodPage, err := browser.Page(url)
	errorhandlers.FatalError("Yeah yeah yeah yeah yeah", err)

	// https://redsky.target.com/redsky_aggregations/v1/web/plp_search_v1?
	data.Pricing_store_id = strconv.FormatFloat(stores[0].Store_id, 'f', -1, 64)
	data.Scheduled_delivery_store_id = strconv.FormatFloat(stores[1].Store_id, 'f', -1, 64)
	data.Store_ids = strconv.FormatFloat(stores[0].Store_id, 'f', -1, 64) + `%2C` + strconv.FormatFloat(stores[1].Store_id, 'f', -1, 64) + `%2C` + strconv.FormatFloat(stores[2].Store_id, 'f', -1, 64) + `%2C` + strconv.FormatFloat(stores[3].Store_id, 'f', -1, 64) + `%2C` + strconv.FormatFloat(stores[4].Store_id, 'f', -1, 64)
	data.Visitor_id = GetCookieValue(rodPage, `visitorId`).(string)

	// https://redsky.target.com/redsky_aggregations/v1/web_platform/product_summary_with_fulfillment_v1?
	data.Store_id = strconv.FormatFloat(stores[0].Store_id, 'f', -1, 64)
	data.State = stores[0].State
	data.Latitude = stores[0].Latitude
	data.Longitude = stores[0].Longitude
	data.Required_store_id = strconv.FormatFloat(stores[0].Store_id, 'f', -1, 64)

	/*
		4) Get list of products and their details in relation to default store (e.g., inventory, etc.)
	*/
	totalPages := GetTotalPagesByNav(rodPage)

	var finalListProducts []Product

	// go through each page to get products
	for page := 1; page <= totalPages; page++ {
		time.Sleep(time.Second * time.Duration(rand.Intn(30)))

		products := GetTcinsAndProductFamily(data)

		tcinsArr := CreateTcinsArr(products)

		tcinsDivided := DivideIntoSubArrays(data, tcinsArr)

		for _, v := range tcinsDivided {
			data.Tcins = ConvertTcinsToString(v)
			time.Sleep(time.Second * time.Duration(rand.Intn(10)))
			products = GetProductDetails(data, products)
			for _, product := range products {
				finalListProducts = append(finalListProducts, product)
			}
		}

		data.Offset = data.Offset + 24 // need to add 24 to offset to send to API to advance to next page

		log.Printf("##################################################")
		log.Printf("################ End of Loop %v ###################", page)
		log.Printf("##################################################")
	}

	log.Printf("Adding products to the store...")
	stores[0].Store_items = finalListProducts

	/*
		5) Load product details into JSON to send to database API.
	*/
	store := &stores[0]
	x, err := json.Marshal(store)
	errorhandlers.FatalError("Oh, it's beautiful, but wait a minute, isn't this?", err)
	// log.Printf("json.Marshal(store): %v", string(x))

	AppendFile("./cmd/targetProductsV2.0.txt", string(x))

	log.Printf("Ending job...")
}

func AppendFile(fileName string, fileContent string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	len, err := file.WriteString(fileContent)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	log.Printf("Length: %d bytes", len)
	log.Printf("File Name: %s", file.Name())
}
