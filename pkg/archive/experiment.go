package rodHandlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/go-rod/rod"
// 	"github.com/go-rod/rod/lib/proto"
// 	errorhandlers "github.com/stoic-llama/formula-scraper/pkg/errorHandlers"
// )

// // import (
// // 	"encoding/json"
// // 	"log"
// // 	"strconv"
// // 	"strings"
// // 	"time"

// // 	"github.com/go-rod/rod"
// // )

// // func GrabAvailability() {
// // 	browser := rod.New().MustConnect()
// // 	page := browser.MustPage("https://www.target.com/p/similac-go-38-grow-powder-toddler-formula-30-8oz/-/A-51479802#lnk=sametab")
// // 	elAvailability := page.MustElement(`div[data-test="soldOutBlock"]`)
// // 	log.Printf("Availability: %v", elAvailability.MustText())

// // 	elBool, elAvailability2, elError := page.Has(`div[data-test="soldOutBlock"]`)
// // 	log.Printf("Bool: %v", elBool)
// // 	log.Printf("Element: %v", elAvailability2)
// // 	log.Printf("Error: %v ", elError)
// // }

// // func GrabJSON() {
// // 	browser := rod.New().MustConnect()
// // 	page := browser.MustPage("https://www.target.com/p/similac-pure-bliss-non-gmo-powder-infant-formula-24-7oz/-/A-54556804#lnk=sametab")

// // 	time.Sleep(time.Second * 10)

// // 	elProduct := page.MustElement(`#json`).MustEval(`() => this.innerText`).String()
// // 	elProductFormatted := strings.ReplaceAll(elProduct, "@", "")

// // 	var result map[string]interface{}
// // 	json.Unmarshal([]byte(elProductFormatted), &result)

// // 	var product ProductItemJSON

// // 	product.graph.context = result["graph"].([]interface{})[0].(map[string]interface{})["context"].(string)
// // 	product.graph.productType = result["graph"].([]interface{})[0].(map[string]interface{})["type"].(string)
// // 	product.graph.name = result["graph"].([]interface{})[0].(map[string]interface{})["name"].(string)
// // 	product.graph.brand = result["graph"].([]interface{})[0].(map[string]interface{})["brand"].(string)
// // 	product.graph.image = result["graph"].([]interface{})[0].(map[string]interface{})["image"].(string)
// // 	product.graph.sku = result["graph"].([]interface{})[0].(map[string]interface{})["sku"].(string)
// // 	product.graph.description = result["graph"].([]interface{})[0].(map[string]interface{})["description"].(string)
// // 	product.graph.gtin13 = result["graph"].([]interface{})[0].(map[string]interface{})["gtin13"].(string)

// // 	product.graph.offers.offerType = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["type"].(string)

// // 	priceStr := result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["price"].(string)
// // 	price, _ := strconv.ParseFloat(priceStr[1:], 64) // remove "$" from "$XX.XX" so it can be type asserted to float64
// // 	product.graph.offers.price = price

// // 	product.graph.offers.priceCurrency = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["priceCurrency"].(string)
// // 	product.graph.offers.availability = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["availability"].(string)
// // 	product.graph.offers.availableAtOrFrom.availableAtOrFromType = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["availableAtOrFrom"].(map[string]interface{})["type"].(string)
// // 	product.graph.offers.availableAtOrFrom.branchCode = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["availableAtOrFrom"].(map[string]interface{})["branchCode"].(string)
// // 	product.graph.offers.availableDeliveryMethod = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["availableDeliveryMethod"].(string)
// // 	product.graph.offers.potentialAction.potentialActionType = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["potentialAction"].(map[string]interface{})["type"].(string)
// // 	product.graph.offers.url = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["url"].(string)

// // 	product.graph.aggregateRating.aggregateRatingType = result["graph"].([]interface{})[0].(map[string]interface{})["aggregateRating"].(map[string]interface{})["type"].(string)
// // 	product.graph.aggregateRating.bestRating = result["graph"].([]interface{})[0].(map[string]interface{})["aggregateRating"].(map[string]interface{})["bestRating"].(float64)
// // 	product.graph.aggregateRating.ratingValue = result["graph"].([]interface{})[0].(map[string]interface{})["aggregateRating"].(map[string]interface{})["ratingValue"].(float64)
// // 	product.graph.aggregateRating.reviewCount = result["graph"].([]interface{})[0].(map[string]interface{})["aggregateRating"].(map[string]interface{})["reviewCount"].(float64)
// // 	product.graph.aggregateRating.worstRating = result["graph"].([]interface{})[0].(map[string]interface{})["aggregateRating"].(map[string]interface{})["worstRating"].(float64)

// // 	review := result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})
// // 	totalNumberOfReviews := len(review)
// // 	product.graph.review = make([]Review, totalNumberOfReviews)
// // 	for i := 0; i < totalNumberOfReviews; i++ {
// // 		product.graph.review[i].reviewType = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["type"].(string)
// // 		product.graph.review[i].description = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["description"].(string)
// // 		product.graph.review[i].author = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["author"].(string)
// // 		product.graph.review[i].name = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["name"].(string)

// // 		product.graph.review[i].reviewRating.reviewRatingType = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["reviewRating"].(map[string]interface{})["type"].(string)
// // 		product.graph.review[i].reviewRating.worstRating = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["reviewRating"].(map[string]interface{})["worstRating"].(float64)
// // 		product.graph.review[i].reviewRating.ratingValue = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["reviewRating"].(map[string]interface{})["ratingValue"].(float64)
// // 		product.graph.review[i].reviewRating.bestRating = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["reviewRating"].(map[string]interface{})["bestRating"].(float64)
// // 	}

// // 	// log.Printf("product.graph: %v", product.graph)
// // 	log.Printf("product.graph.review[0].author: %v", product.graph.review[0].author)

// // }

// // /////////////////////////////////////////////
// // ////////////// ProductItemJSON //////////////
// // /////////////////////////////////////////////

// // type ProductItemJSON struct {
// // 	graph struct {
// // 		context         string          `json:"@context,omitempty"`
// // 		productType     string          `json:"@type:omitempty"`
// // 		name            string          `json:"name,omitempty"`
// // 		brand           string          `json:"brand,omitempty"`
// // 		image           string          `json:"image,omitempty"`
// // 		sku             string          `json:"sku,omitempty"`
// // 		description     string          `json:"description,omitempty"`
// // 		gtin13          string          `json:"gtin13,omitempty"`
// // 		offers          Offers          `json:"offers,omitempty"`
// // 		aggregateRating AggregateRating `json:"aggregateRating,omitempty"`
// // 		review          []Review        `json:"review,omitempty"`
// // 	}
// // }

// // ////////////////////////////////////
// // ////////////// Offers //////////////
// // ////////////////////////////////////

// // type Offers struct {
// // 	offerType               string            `json:"@type,omitempty"`
// // 	price                   float64           `json:"price,omitempty"`
// // 	priceCurrency           string            `json:"priceCurrency,omitempty"`
// // 	availability            string            `json:"availability,omitempty"`
// // 	availableAtOrFrom       AvailableAtOrFrom `json:"availableAtOrFrom,omitempty"`
// // 	availableDeliveryMethod string            `json:"availableDeliveryMethod,omitempty"`
// // 	deliveryLeadTime        DeliveryLeadTime  `json:"deliveryLeadTime,omitempty"`
// // 	potentialAction         PotentialAction   `json:"potentialAction,omitempty"`
// // 	url                     string            `json:"url,omitempty"`
// // }

// // type AvailableAtOrFrom struct {
// // 	availableAtOrFromType string `json:"@type,omitempty"`
// // 	branchCode            string `json:"branchCode,omitempty"`
// // }

// // type DeliveryLeadTime struct {
// // 	deliveryLeadTimeType string  `json:"@type,omitempty"`
// // 	value                float64 `json:"value,omitempty"`
// // }

// // type PotentialAction struct {
// // 	potentialActionType string `json:"@type,omitempty"`
// // }

// // ///////////////////////////////////
// // ///////// AggregateRating /////////
// // ///////////////////////////////////

// // type AggregateRating struct {
// // 	aggregateRatingType string  `json:"@type,omitempty"`
// // 	bestRating          float64 `json:"bestRating,omitempty"`
// // 	ratingValue         float64 `json:"ratingValue,omitempty"`
// // 	reviewCount         float64 `json:"reviewCount,omitempty"`
// // 	worstRating         float64 `json:"worstRating,omitempty"`
// // }

// // ////////////////////////////////////
// // ////////////// Review //////////////
// // ////////////////////////////////////

// // type Review struct {
// // 	reviewType   string       `json:"@type,omitempty"`
// // 	description  string       `json:"description,omitempty"`
// // 	author       string       `json:"author,omitempty"`
// // 	name         string       `json:"name,omitempty"`
// // 	reviewRating ReviewRating `json:"reviewRating,omitempty"`
// // }

// // type ReviewRating struct {
// // 	reviewRatingType string  `json:"@type,omitempty"`
// // 	worstRating      float64 `json:"worstRating,omitempty"`
// // 	ratingValue      float64 `json:"ratingValue,omitempty"`
// // 	bestRating       float64 `json:"bestRating,omitempty"`
// // }

// func GetTargetProducts() {
// 	browser := rod.New().MustConnect()
// 	rodPage := browser.MustPage("https://www.target.com/")

// 	rodPage.MustWaitLoad()

// 	time.Sleep(time.Second * 20)

// 	results, _ := proto.NetworkGetAllCookies{}.Call(rodPage)
// 	// var UserLocationExpires string
// 	// var fiatsCookieExpires string
// 	var visitorId string
// 	var fiatsCookie string
// 	for _, cookie := range results.Cookies {
// 		// log.Printf("#%v: %v's value is %v", k, cookie.Name, cookie.Value)

// 		// if cookie.Name == "UserLocation" {
// 		// 	UserLocationExpires = cookie.Expires.Time().String()
// 		// }
// 		// if cookie.Name == "fiatsCookie" {
// 		// 	fiatsCookieExpires = cookie.Expires.Time().String()
// 		// }
// 		if cookie.Name == "visitorId" {
// 			visitorId = cookie.Value
// 		}
// 		if cookie.Name == "fiatsCookie" {
// 			fiatsCookie = cookie.Value
// 		}
// 	}

// 	// Set cookie by first deleting the initial fiatsCookie and adding a new one
// 	// The fiatsCookie determines the store location of the rodPage session
// 	// Which would determine what inventory we see for the products (availability is location specific)
// 	initialLocation := `document.cookie='fiatsCookie=` + fiatsCookie + `;Domain=.target.com;Path=/;expires=Thu, 08 Jun 1970 00:00:00 GMT;SameSite=Lax;Secure'`
// 	rodPage.Eval(initialLocation)
// 	time.Sleep(time.Second * 2)

// 	targetFiatsCookie := `DSI_2380|DSN_Harlem|DSZ_10035` // hard coding for now - will be input parameter
// 	targetLocation := `document.cookie='fiatsCookie=` + targetFiatsCookie + `;Domain=.target.com;Path=/;expires=Thu, 08 Jun 1970 00:00:00 GMT;SameSite=Lax;Secure'`
// 	rodPage.Eval(targetLocation)
// 	// rodPage.Eval(`document.cookie='fiatsCookie=DSI_2380|DSN_Harlem|DSZ_10035;Domain=.target.com;Path=/;expires=Thu, 08 Jun 2023 00:00:00 GMT;SameSite=Lax;Secure'`)
// 	time.Sleep(time.Second * 2)

// 	rodPage.Navigate(`https://www.target.com/c/formula-nursing-feeding-baby/-/N-5xtkh`)

// 	time.Sleep(time.Second * 20)

// 	// finding total number of products
// 	titleElement, err := rodPage.Element(`h2[data-test="resultsHeading"]`)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	titleStr, err := titleElement.Text()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	titleArr := strings.Split(titleStr, " ")
// 	total, err := strconv.Atoi(titleArr[0])
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	log.Printf("total: %v", total)

// 	// calling product aggregation API to get product details for all products
// 	baseURL := `https://redsky.target.com/redsky_aggregations/v1/web/plp_search_v1?`
// 	var sb string
// 	key := `9f36aeafbe60771e321a7cc95a78140772ab3e96`
// 	category := `5xtkh`
// 	channel := `WEB`
// 	count := `24`
// 	default_purchasability_filter := `true`
// 	include_sponsored := `true`
// 	page := `%2Fc%2F5xtkh` // decoded: /c/5xtkh
// 	platform := `desktop`
// 	pricing_store_id := `2380`
// 	scheduled_delivery_store_id := `1289`
// 	store_ids := `2380%2C3387%2C2475%2C3394%2C3312` // decoded: 2380,3387,2475,3394,3312
// 	useragent := `Mozilla%2F5.0+%28Macintosh%3B+Intel+Mac+OS+X+11_0_0%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F87.0.4280.88+Safari%2F537.36`

// 	for offset := 0; offset <= total; offset = offset + 24 {
// 		// baseURL2 := `key=9f36aeafbe60771e321a7cc95a78140772ab3e96&category=5xtkh&channel=WEB&count=24&default_purchasability_filter=true&include_sponsored=true&page=%2Fc%2F5xtkh&platform=desktop&pricing_store_id=2380&scheduled_delivery_store_id=1289&store_ids=2380%2C3387%2C2475%2C3394%2C3312&useragent=Mozilla%2F5.0+%28Macintosh%3B+Intel+Mac+OS+X+11_0_0%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F87.0.4280.88+Safari%2F537.36&visitor_id=` + visitorId + `&offset=` + strconv.Itoa(offset)
// 		baseURL2 := `key=` + key + `&category=` + category + `&channel=` + channel + `&count=` + count + `&default_purchasability_filter=` + default_purchasability_filter + `&include_sponsored=` + include_sponsored + `&page=` + page + `&platform=` + platform + `&pricing_store_id=` + pricing_store_id + `&scheduled_delivery_store_id=` + scheduled_delivery_store_id + `&store_ids=` + store_ids + `&useragent=` + useragent + `&visitor_id=` + visitorId + `&offset=` + strconv.Itoa(offset)
// 		productSummaryURL := baseURL + baseURL2

// 		resp, err := http.Get(productSummaryURL)

// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		sb = sb + string(body)
// 		time.Sleep(time.Second * time.Duration(rand.Intn(10)))

// 	}

// 	AppendFile("./cmd/targetProducts.txt", sb)

// 	// log.Printf(sb)

// 	// key: 9f36aeafbe60771e321a7cc95a78140772ab3e96
// 	// tcins: 78616412,79329649,75665719,21538751,81624903,75665506,53272589,79329648,83041740,79384108,80858562,80858560,80858561,51327450,82104756,82104750,75666059,15723744,79551374,80769180,75665509,14599313
// 	// store_id: 1802
// 	// zip: 06052
// 	// state: CT
// 	// latitude: 41.660
// 	// longitude: -72.790
// 	// scheduled_delivery_store_id: 1289
// 	// required_store_id: 1802
// 	// has_required_store_id: true
// 	//https://redsky.target.com/redsky_aggregations/v1/web_platform/product_summary_with_fulfillment_v1?key=9f36aeafbe60771e321a7cc95a78140772ab3e96&tcins=78616412,79329649,75665719,21538751,81624903,75665506,53272589,79329648,83041740,79384108,80858562,80858560,80858561,51327450,82104756,82104750,75666059,15723744,79551374,80769180,75665509,14599313&store_id=1802&zip=06052&state=CT&latitude=41.660&longitude=-72.790&scheduled_delivery_store_id=1289&required_store_id=1802&has_required_store_id=true

// }

// func ReadFile(fileName string) {
// 	data, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		log.Panicf("failed reading data from file: %s", err)
// 	}
// 	fmt.Printf("\nLength: %d bytes", len(data))
// 	fmt.Printf("\nData: %s", data)
// 	fmt.Printf("\nError: %v", err)
// }

// func GetDateToday() {
// 	currentTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location()) // time.Now()

// 	if currentTime.Day() < 10 {
// 		weekday := currentTime.Weekday().String()[0:3]
// 		day := currentTime.Day()
// 		month := currentTime.Month().String()[0:3]
// 		year := currentTime.Year()
// 		fiatsDate := weekday + ", 0" + strconv.Itoa(day) + " " + month + " " + strconv.Itoa(year) + " 00:00:00 GMT"
// 		fmt.Println(fiatsDate)
// 		// fmt.Printf("Thu, 08 Jun 2023 00:00:00 GMT is = %v, 0%v %v %v 00:00:00 GMT", currentTime.Weekday().String()[0:3], currentTime.Day(), currentTime.Month().String()[0:3], currentTime.Year())
// 	}
// 	if currentTime.Day() >= 10 {
// 		weekday := currentTime.Weekday().String()[0:3]
// 		day := currentTime.Day()
// 		month := currentTime.Month().String()[0:3]
// 		year := currentTime.Year()
// 		fiatsDate := weekday + ", " + strconv.Itoa(day) + " " + month + " " + strconv.Itoa(year) + " 00:00:00 GMT"
// 		fmt.Println(fiatsDate)
// 		// fmt.Printf("Thu, 08 Jun 2023 00:00:00 GMT is = %v, %v %v %v 00:00:00 GMT", currentTime.Weekday().String()[0:3], currentTime.Day(), currentTime.Month().String()[0:3], currentTime.Year())
// 	}
// }

// func GetTargetProductsV2(zipcode string) {
// 	log.Printf("Starting job...")
// 	browser := rod.New().MustConnect()

// 	data := DataContainer{
// 		//////
// 		Key:        `9f36aeafbe60771e321a7cc95a78140772ab3e96`,
// 		Limit:      2000,
// 		Unit:       `mile`,
// 		Within:     25,
// 		Place:      zipcode,
// 		NearbyType: `store`,

// 		//////
// 		Category:                      `5xtkh`,
// 		Channel:                       `WEB`,
// 		Count:                         24,
// 		Default_purchasability_filter: true,
// 		Included_sponsored:            true,
// 		Offset:                        0,
// 		Page:                          `%2Fc%2F5xtkh`, // decoded: /c/5xtkh
// 		Platform:                      `desktop`,

// 		// decode: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36
// 		Tcins:     "",
// 		Useragent: `Mozilla%2F5.0+%28Macintosh%3B+Intel+Mac+OS+X+10_15_7%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F102.0.5005.61+Safari%2F537.36`,

// 		//////
// 		Zip:                   zipcode,
// 		Has_required_store_id: true,

// 		//////
// 		Start: 0,
// 	}

// 	/*
// 		0) Find 5 closest stores from zip code provided by user.
// 			https://api.target.com/location_proximities/v1/nearby_locations?limit=2000&unit=mile&within=25&place=10029&type=store&key=9f36aeafbe60771e321a7cc95a78140772ab3e96
// 			input:
// 				- unit: mile
// 				- within: 25
// 				- place: <place zip code provided by user>
// 				- type: store
// 				- key: static api key
// 			output:
// 				- ** company: "Target" (hard coded and not provided by the Target API)
// 				- location_id (store id)
// 				- location_name ("Harlem" - this is not in final output, but will be used in later steps)
// 				- address > address_line1
// 				- address > address_line2
// 				- address > city
// 				- address > region (i.e., state or "NY" initials)
// 				- address > postal_code (i.e., zip code)
// 				- ** country: "US" (hard coded and not provided by the Target API)
// 				- geographic specifications > latitude
// 				- geographic specifications > longitude
// 	*/

// 	stores := GetNearByStores(data)

// 	/*
// 		1) Go to https://www.target.com and let website load default cookie and store location
// 		2) Get fiatsCookie from default load, and delete it
// 		3) Set a new fiatsCookie to Store ID.  This would be the first of five stores from #1.

// 			rodPage.Eval(`document.cookie='
// 				fiatsCookie=DSI_2380|DSN_Harlem|DSZ_10035;
// 				Domain=.target.com;
// 				Path=/;
// 				expires=Thu, 08 Jun 2023 00:00:00 GMT;
// 				SameSite=Lax;
// 				Secure'`
// 			)
// 			input:
// 				- location_id ('DSI_' + location_id + '|')
// 				- location_name ('DSN_' + location_name + '|')
// 				- postal_code ('DSZ_' + postal_code + ';')
// 				- dateTomorrow ('expires=' + dateTomorrow + ';SameSite=Lax;Secure')

// 	*/

// 	SetStoreLocation(browser, stores[0]) // set store location to store closest to zipcode provided by user

// 	url := proto.TargetCreateTarget{
// 		URL: "https://www.target.com/c/formula-nursing-feeding-baby/-/N-5xtkh",
// 	}
// 	rodPage, err := browser.Page(url)
// 	errorhandlers.FatalError("Yeah yeah yeah yeah yeah", err)

// 	/*
// 		4) Let page load and grab total products, divided by 24 to get number of pages

// 		5) Get TCINs for the product summary

// 			https://redsky.target.com/redsky_aggregations/v1/web/plp_search_v1?key=9f36aeafbe60771e321a7cc95a78140772ab3e96&category=5xtkh&channel=WEB&count=24&default_purchasability_filter=true&include_sponsored=true&offset=96&page=%2Fc%2F5xtkh&platform=desktop&pricing_store_id=1802&scheduled_delivery_store_id=1289&store_ids=1802%2C1289%2C3333%2C2434%2C1373&useragent=Mozilla%2F5.0+%28Macintosh%3B+Intel+Mac+OS+X+10_15_7%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F102.0.5005.61+Safari%2F537.36&visitor_id=018146D20D8802019CB9140837058F93

// 			input:
// 				key: 9f36aeafbe60771e321a7cc95a78140772ab3e96
// 				category: 5xtkh
// 				channel: WEB
// 				count: 24
// 				default_purchasability_filter: true
// 				include_sponsored: true
// 				offset: 96 <this would be total pages (5) minus one (4) multiplied by count (24) = 96>
// 				page: /c/5xtkh
// 				platform: desktop
// 				pricing_store_id: 1802 <first of the stores from #1>
// 				scheduled_delivery_store_id: 1289 <second of the stores from #1>
// 				store_ids: 1802,1289,3333,2434,1373 <all stores ids from #1>
// 				useragent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36
// 				visitor_id: 018146D20D8802019CB9140837058F93 <from cookie>

// 			output:
// 				+++tcin: data > search > products > item > tcin <loop through all products>
// 				product_family: data > search > products > item > primary_brand > name

// 		6) Get product details

// 			https://redsky.target.com/redsky_aggregations/v1/web_platform/product_summary_with_fulfillment_v1?key=9f36aeafbe60771e321a7cc95a78140772ab3e96&tcins=78616412%2C53272589%2C79329649%2C21538751%2C75665719%2C75665506%2C79329648%2C83041740%2C79384108%2C80858562%2C80858560%2C80858561%2C51327450%2C82104756%2C82104750%2C75666059%2C15723744%2C79551374%2C80769180%2C75665509%2C14599313&store_id=1802&zip=06052&state=CT&latitude=41.660&longitude=-72.790&scheduled_delivery_store_id=1289&required_store_id=1802&has_required_store_id=true

// 			input:
// 				key: 9f36aeafbe60771e321a7cc95a78140772ab3e96
// 				tcins: <from prior step> 78616412,53272589,79329649,21538751,75665719,75665506,79329648,83041740,79384108,80858562,80858560,80858561,51327450,82104756,82104750,75666059,15723744,79551374,80769180,75665509,14599313
// 				store_id: <from prior step, first of the nearest stores> 1802
// 				zip: <from user input> 06052
// 				state: CT
// 				latitude: 41.660
// 				longitude: -72.790
// 				scheduled_delivery_store_id: 1289
// 				required_store_id: 1802
// 				has_required_store_id: true

// 			output:
// 				product: data > product summaries > item > production_description > title
// 				price: data > product summaries > price > current_retail
// 				availability: data > product_summaries > fulfillment > store_options > in_store_only > availability_status
// 				quantity: 'Null'
// 				product_url: data > product_summaries > item > enrichment > buy_url
// 				product_family: <from prior api call on plp_search_v>

// 		7) Paginate and rinse and repeat on TCINs and Prod Details

// 	*/

// 	data.Pricing_store_id = strconv.FormatFloat(stores[0].Store_id, 'f', -1, 64)
// 	data.Scheduled_delivery_store_id = strconv.FormatFloat(stores[1].Store_id, 'f', -1, 64)
// 	data.Store_ids = strconv.FormatFloat(stores[0].Store_id, 'f', -1, 64) + `%2C` + strconv.FormatFloat(stores[1].Store_id, 'f', -1, 64) + `%2C` + strconv.FormatFloat(stores[2].Store_id, 'f', -1, 64) + `%2C` + strconv.FormatFloat(stores[3].Store_id, 'f', -1, 64) + `%2C` + strconv.FormatFloat(stores[4].Store_id, 'f', -1, 64)
// 	data.Visitor_id = GetCookieValue(rodPage, `visitorId`).(string)
// 	data.Store_id = strconv.FormatFloat(stores[0].Store_id, 'f', -1, 64)
// 	data.State = stores[0].State
// 	data.Latitude = stores[0].Latitude
// 	data.Longitude = stores[0].Longitude
// 	data.Required_store_id = strconv.FormatFloat(stores[0].Store_id, 'f', -1, 64)

// 	log.Printf("Getting to GetTotalPagesByNav()...")
// 	totalPages := GetTotalPagesByNav(rodPage)
// 	var finalListProducts []Product

// 	log.Printf("Getting to page loop...")
// 	// go through each page
// 	for page := 1; page <= totalPages; page++ {
// 		time.Sleep(time.Second * time.Duration(rand.Intn(30)))

// 		products := GetTcinsAndProductFamily(data)

// 		tcinsArr := CreateTcinsArr(products)

// 		tcinsDivided := DivideIntoSubArrays(data, tcinsArr)

// 		for _, v := range tcinsDivided {
// 			data.Tcins = ConvertTcinsToString(v)
// 			time.Sleep(time.Second * time.Duration(rand.Intn(10)))
// 			products = GetProductDetails(data, products)
// 			for _, product := range products {
// 				finalListProducts = append(finalListProducts, product)
// 			}
// 		}

// 		data.Offset = data.Offset + 24
// 		log.Printf("##################################################")
// 		log.Printf("################ End of Loop %v ##################", page)
// 		log.Printf("##################################################")
// 	}

// 	log.Printf("Starting to add products to the store")
// 	stores[0].Store_items = finalListProducts

// 	store := &stores[0]
// 	x, err := json.Marshal(store)
// 	errorhandlers.FatalError("Oh, it's beautiful, but wait a minute, isn't this?", err)
// 	log.Printf("json.Marshal(store): %v", string(x))

// 	AppendFile("./cmd/targetProducts.txt", string(x))
// }

// func ReadCookie(rodPage *rod.Page, cookieName string) {
// 	results, _ := proto.NetworkGetAllCookies{}.Call(rodPage)

// 	for _, cookie := range results.Cookies {
// 		if cookie.Name == cookieName {
// 			log.Printf("%v's value is: %v", cookie.Name, cookie.Value)
// 		}
// 	}
// }
