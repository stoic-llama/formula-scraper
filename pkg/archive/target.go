package rodHandlers

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

func BrowserSetUp() *rod.Browser {
	browser := rod.New().MustConnect().NoDefaultDevice()
	return browser
}

func PageSetUp(browser *rod.Browser, query string) *rod.Page {
	page := browser.MustPage(query)
	return page
}

func GetNearestStores(browser *rod.Browser, query string) []Store {
	page := browser.MustPage(query)

	page.MustElement("#web-store-id-msg-btn").MustClick()

	page.MustWaitLoad()

	time.Sleep(time.Second * 30)

	elAddresses := page.MustElements(`div.h-text-grayDark.h-margin-b-tight`)

	var addresses []string
	for _, v := range elAddresses {
		addresses = append(addresses, v.MustText())
	}

	stores := make([]Store, len(addresses))
	var address []string
	var zipstate []string
	for i, v := range addresses {
		address = strings.Split(v, ",")
		// log.Println("++++++++++++++++++++++++++++++++++++++")
		// log.Printf("Address #%v with length %v", i, len(address))
		// log.Printf(v)

		last := len(address) - 1

		zipstate = strings.Split(address[last], " ")

		log.Println(zipstate)
		log.Printf("length of zipstate %v", len(zipstate))
		stores[i].zipcode = zipstate[2]
		stores[i].state = zipstate[1]

		stores[i].city = strings.TrimSpace(address[last-1])
		// log.Println(stores[i].city)

		if len(address) > 3 {
			stores[i].street = address[0] + "," + address[1]
		}
		if len(address) == 3 {
			stores[i].street = address[0]
		}

		stores[i].country = "USA"
		stores[i].latitude = -1
		stores[i].longitude = -1

		// log.Printf("stores[%v]: %v", i, stores[i])

		// log.Println("++++++++++++++++++++++++++++++++++++++")
	}

	// Get list of 5 nearest stores
	var storesNearest []Store
	for i := 0; i < 5; i++ {
		storesNearest = append(storesNearest, stores[i])
	}

	log.Println("Loading storesNearest...")
	for _, store := range storesNearest {
		log.Println(store)
	}

	page.MustElement(`button[aria-label="close"]`).MustClick()

	return storesNearest
}

func SetLocation(page *rod.Page, zipcode string) {
	page.MustWaitLoad()

	time.Sleep(time.Second * 20)

	page.MustElement("#web-store-id-msg-btn").MustClick()

	time.Sleep(time.Second * 5)

	page.MustElement("input#zip-or-city-state").MustInput(zipcode)
	page.MustElement(`button[data-test="@web/StoreLocationSearch/button"]`).MustClick()

	// page.MustElement(`button[data-test="@web/StoreIdListItem/SetStoreButton"]`).MustClick()

	time.Sleep(time.Second * 5)

	page.MustElement(".h-flex-align-center.h-flex-justify-center > div:nth-child(1) > button").MustClick()

	time.Sleep(time.Second * 10)

	page.Eval(`window.scrollTo(0,0)`)
}

func SetLocationByStore(page *rod.Page, store Store) {
	page.MustWaitLoad()

	time.Sleep(time.Second * 20)

	page.MustElement("#web-store-id-msg-btn").MustClick()

	time.Sleep(time.Second * 5)

	page.MustElement("input#zip-or-city-state").MustInput(store.zipcode)
	page.MustElement(`button[data-test="@web/StoreLocationSearch/button"]`).MustClick()

	page.MustElement(".h-flex-align-center.h-flex-justify-center > div:nth-child(1) > button").MustClick()

	page.Eval(`window.scrollTo(0,0)`)
}

func GetFullProductListing(browser *rod.Browser, query string) rod.Elements { // *[]string {
	page := browser.MustPage(query)

	page.MustWaitLoad()

	time.Sleep(time.Second * 10)

	title := page.MustElement(`h2[data-test="resultsHeading"]`).MustText()
	titles := strings.Split(title, " ")
	total, _ := strconv.Atoi(titles[0])
	log.Printf("total: %v", total)

	pageCount := total / 24
	log.Printf("page count: %v", pageCount)

	var fullProductListing rod.Elements

	for i := 0; i < pageCount; i++ {
		time.Sleep(time.Second * 10)
		isBlocked, _, _ := page.Has(`#kampyleInvite`)
		if isBlocked {
			page.MustElement(`#kplDeferButton`).MustClick() // close the Modal
		}

		for i := 0; i <= 15; i++ {
			page.Eval(`
				window.scrollBy(0,600)
			`)
			time.Sleep(time.Second * 5)
			log.Printf("loop: %v", i)
		}

		listing := page.MustElements(`a.gCNFxQ`)

		for _, v := range listing {
			fullProductListing = append(fullProductListing, v)
		}

		page.MustElement(`button[data-test="next"]`).MustClick()

		page.MustWaitLoad()

	}

	log.Printf("fullProductListing length: %v", len(fullProductListing))

	return fullProductListing
}

// func GetFullProductListingSinglePage(browser *rod.Browser, query string) rod.Elements { // *[]string {
// 	page := browser.MustPage(query)
// 	page.MustWaitLoad()
// 	time.Sleep(time.Second * 10)
// 	title := page.MustElement(`h2[data-test="resultsHeading"]`).MustText()
// 	titles := strings.Split(title, " ")
// 	total, _ := strconv.Atoi(titles[0])
// 	log.Printf("total: %v", total)
// 	pageCount := total / 24
// 	log.Printf("page count: %v", pageCount)
// 	// productItemLinks := []string{}
// 	var fullProductListing rod.Elements
// 	// for i := 0; i < pageCount; i++ {
// 	// 	for i := 0; i <= 15; i++ {
// 	// 		page.Eval(`
// 	// 			window.scrollBy(0,600)
// 	// 		`)
// 	// 		time.Sleep(time.Second * 5)
// 	// 		log.Printf("loop: %v", i)
// 	// 	}
// 	// 	productItems = page.MustElements(`a.gCNFxQ`)
// 	// productItems := page.MustElements("a.gCNFxQ")
// 	// for _, v := range productItems {
// 	// 	x, _ := v.Property("href")
// 	// 	productItemLinks = append(productItemLinks, x.String())
// 	// }
// 	// log.Println(productItemLinks)
// 	// log.Println(len(productItems))
// 	// page.MustElement(`button[data-test="next"]`).MustClick()
// 	// page.MustWaitLoad()
// 	// time.Sleep(time.Second * 10)
// 	// }
// 	// log.Println("Full Product Listing Downloaded with URLs.")
// 	// for i := 0; i < pageCount; i++ {
// 	// 	for i := 0; i <= 15; i++ {
// 	// 		page.Eval(`
// 	// 			window.scrollBy(0,600)
// 	// 		`)
// 	// 		time.Sleep(time.Second * 5)
// 	// 		log.Printf("loop: %v", i)
// 	// 	}
// 	for i := 0; i <= 5; i++ {
// 		page.Eval(`
// 				window.scrollBy(0,600)
// 			`)
// 		time.Sleep(time.Second * 5)
// 		log.Printf("loop: %v", i)
// 	}
// 	fullProductListing = page.MustElements(`a.gCNFxQ`)
// 	// page.MustElement(`button[data-test="next"]`).MustClick()
// 	// page.MustWaitLoad()
// 	// time.Sleep(time.Second * 10)
// 	// }
// 	log.Println("Full Product Listing Downloaded with URLs.")
// 	return fullProductListing
// }

func GetProductItemDetails(browser *rod.Browser, store Store, fullProductListing rod.Elements) []ProductItem { // listing *[]string) {
	var productItems = make([]ProductItem, len(fullProductListing))
	var productItemsFiltered []ProductItem

	log.Printf("Store received in GetProductItemDetails(): %v", store)

	// Populate each Product Item
	for k, v := range fullProductListing {
		log.Println("++++++++++++++++++++++++++++++++++")
		log.Printf("Iteration #%v:", k)

		x, _ := v.Property(`href`)
		url := x.String()

		page := browser.MustPage(url)
		time.Sleep(time.Second * 10)
		isBlocked, _, _ := page.Has(`#kampyleInvite`)
		if isBlocked {
			page.MustElement(`#kplDeferButton`).MustClick() // close the Modal
		}

		product := GetProductJSON(page, url)

		// Populate company
		productItems[k].company = "Target"

		// Populate zipcode
		productItems[k].zipcode = store.zipcode

		// Populate street
		productItems[k].street = store.street

		// Populate city
		productItems[k].city = store.city

		// Populate state
		productItems[k].state = store.state

		// Populate country
		productItems[k].country = store.country

		// Populate Product Family
		productItems[k].productFamily = product.graph.brand

		// Populate Product
		productItems[k].product = product.graph.name

		// Populate url
		productItems[k].url = url
		log.Printf("url: %v", url)

		// Populate price
		productItems[k].price = product.graph.offers.price

		// randomWeightPrice, _, _ := page.Has(`span[data-test="product-random-weight-price"]`)
		// noWeightPrice, _, _ := page.Has(`span[data-test="product-price"]`)

		// if randomWeightPrice {
		// 	elPrice, _ := page.Element(`span[data-test="product-random-weight-price"]`)

		// 	price := elPrice.MustText()

		// 	priceFiltered := price[1:] // remove first char "$"

		// 	elPriceFloat64, err := strconv.ParseFloat(priceFiltered, 64)
		// 	productItems[k].price = elPriceFloat64
		// 	log.Printf("randomWeightPrice: %v", productItems[k].price)
		// 	log.Printf("randomWeightPrice Conversion Error: %v", err)
		// }

		// if noWeightPrice {
		// 	elPrice, _ := page.Element(`span[data-test="product-price"]`)

		// 	price := elPrice.MustText()

		// 	priceFiltered := price[1:] // remove first char "$"

		// 	elPriceFloat64, err := strconv.ParseFloat(priceFiltered, 64)
		// 	productItems[k].price = elPriceFloat64
		// 	log.Printf("noWeightPrice: %v", productItems[k].price)
		// 	log.Printf("noWeightPrice Conversion Error: %v", err)
		// }

		// Populate availability

		// // Scenario One
		// // Out of stock at selected store and all nearby stores
		// // Hide in-stock stores button
		// outOfStockAll, eloutOfStockAll, outOfStockAllerr := page.Has(`div[data-test="outOfStockNearbyMessage"]`)
		// temp1 := eloutOfStockAll.String()
		// log.Printf("outOfStockAll: %v", temp1)
		// log.Printf("outOfStockAll error: %v", outOfStockAllerr)

		// // Scenario Two
		// // Limited Stock at selected store
		// // Hide in-stock stores button
		// limitedStockSelected, ellimitedStockSelected, limitedStockSelectederr := page.Has(`div[data-test="inStoreOnlyMessage"]`)
		// temp2 := ellimitedStockSelected.String()
		// log.Printf("limitedStockSelected: %v", temp2)
		// log.Printf("limitedStockSelected error: %v", limitedStockSelectederr)

		// // Scenario Three
		// // Out of stock at selected store
		// // Show in-stock stores button
		// notSoldAtSelected, elnotSoldAtSelected, notSoldAtSelectederr := page.Has(`div[data-test="notSoldAtMessage"]`)
		// temp3 := elnotSoldAtSelected.String()
		// log.Printf("notSoldAtSelected: %v", temp3)
		// log.Printf("notSoldAtSelected error: %v", notSoldAtSelectederr)

		// // Scenario Four
		// // Sold out
		// // Hide in-stock stores button
		// soldOut, elsoldOut, soldOuterr := page.Has(`div[data-test="soldOutBlock"]`)
		// temp4 := elsoldOut.String()
		// log.Printf("soldOut: %v", temp4)
		// log.Printf("soldOut error: %v", soldOuterr)

		// // Scenario Five
		// // Out of Stock at selected store
		// // Show in-stock stores button
		// outOfStockSelected, eloutOfStockSelected, outOfStockSelectederr := page.Has(`div[data-test="outOfStockMessage"]`)
		// temp5 := eloutOfStockSelected.String()
		// log.Printf("outOfStockSelected: %v", temp5)
		// log.Printf("outOfStockSelected error: %v", outOfStockSelectederr)

		// if outOfStockAll || notSoldAtSelected || outOfStockSelected || soldOut {
		// 	productItems[k].availability = "Out of stock"
		// }

		// if limitedStockSelected {
		// 	productItems[k].availability = "Limited stock"
		// }

		// if productItems[k].availability == "" {
		// 	productItems[k].availability = "Did not find availability"
		// }

		productItems[k].availability = product.graph.offers.availability

		log.Printf("Availability: %v", productItems[k].availability)

		// Populate quantity - future enh
		productItems[k].quantity = -1

		// Populate longitude - future enh
		productItems[k].longitude = -1

		// Populate latitude - future enh
		productItems[k].latitude = -1

		// time.Sleep(time.Second * time.Duration(rand.Intn(5)))

		log.Println("++++++++++++++++++++++++++++++++++")

		page.MustClose()

	}

	// filter out elements with header and privacy policy links
	for _, v := range productItems {
		if strings.Contains(v.url, "https://www.target.com/p/") {
			productItemsFiltered = append(productItemsFiltered, v)
		}
	}

	log.Printf("Length of productItems: %v", len(productItemsFiltered))

	return productItemsFiltered
}

// func GetAvailability(page *rod.Page) string {
// 	var status = []string{
// 		`div[data-test="outOfStockNearbyMessage"]`,
// 		`div[data-test="inStoreOnlyMessage"]`,
// 		`div[data-test="notSoldAtMessage"]`,
// 		`div[data-test="soldOutBlock"]`,
// 		`div[data-test="outOfStockMessage"]`,
// 	}
// 	time.Sleep(time.Second * 30)
// 	for k := range status {
// 		elBool, elAvailability2, elError := page.Has(status[k])
// 		log.Printf("Bool: %v", elBool)
// 		log.Printf("Element: %v", elAvailability2)
// 		log.Printf("Error: %v ", elError)
// 		if elBool == true && status[k] == `div[data-test="inStoreOnlyMessage"]` {
// 			return "Limited Stock"
// 		}
// 		if elBool == true && status[k] != `div[data-test="inStoreOnlyMessage"]` {
// 			return "Out of Stock"
// 		}
// 	}
// 	return "Did not find availability"
// }

///////////////////////////////

func GetProductJSON(page *rod.Page, url string) ProductItemJSON {
	time.Sleep(time.Second * 10)

	var product ProductItemJSON

	if strings.Contains(url, "https://www.target.com/p/") {

		elProduct := page.MustElement(`#json`).MustEval(`() => this.innerText`).String()
		elProductFormatted := strings.ReplaceAll(elProduct, "@", "") // remove "@" because illegal to have it beginning of a field name in Golang

		log.Printf("elProductFormatted: %v", elProductFormatted)

		var result map[string]interface{}
		json.Unmarshal([]byte(elProductFormatted), &result)

		product.graph.context = result["graph"].([]interface{})[0].(map[string]interface{})["context"].(string)
		product.graph.productType = result["graph"].([]interface{})[0].(map[string]interface{})["type"].(string)
		product.graph.name = result["graph"].([]interface{})[0].(map[string]interface{})["name"].(string)
		product.graph.brand = result["graph"].([]interface{})[0].(map[string]interface{})["brand"].(string)
		product.graph.image = result["graph"].([]interface{})[0].(map[string]interface{})["image"].(string)
		product.graph.sku = result["graph"].([]interface{})[0].(map[string]interface{})["sku"].(string)
		product.graph.description = result["graph"].([]interface{})[0].(map[string]interface{})["description"].(string)
		product.graph.gtin13 = result["graph"].([]interface{})[0].(map[string]interface{})["gtin13"].(string)

		product.graph.offers.offerType = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["type"].(string)

		priceStr := result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["price"].(string)
		price, _ := strconv.ParseFloat(priceStr[1:], 64) // remove "$" from "$XX.XX" so it can be type asserted to float64
		product.graph.offers.price = price

		product.graph.offers.priceCurrency = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["priceCurrency"].(string)
		product.graph.offers.availability = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["availability"].(string)
		product.graph.offers.availableAtOrFrom.availableAtOrFromType = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["availableAtOrFrom"].(map[string]interface{})["type"].(string)
		product.graph.offers.availableAtOrFrom.branchCode = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["availableAtOrFrom"].(map[string]interface{})["branchCode"].(string)
		product.graph.offers.availableDeliveryMethod = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["availableDeliveryMethod"].(string)
		product.graph.offers.potentialAction.potentialActionType = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["potentialAction"].(map[string]interface{})["type"].(string)
		product.graph.offers.url = result["graph"].([]interface{})[0].(map[string]interface{})["offers"].(map[string]interface{})["url"].(string)

		product.graph.aggregateRating.aggregateRatingType = result["graph"].([]interface{})[0].(map[string]interface{})["aggregateRating"].(map[string]interface{})["type"].(string)
		product.graph.aggregateRating.bestRating = result["graph"].([]interface{})[0].(map[string]interface{})["aggregateRating"].(map[string]interface{})["bestRating"].(float64)
		product.graph.aggregateRating.ratingValue = result["graph"].([]interface{})[0].(map[string]interface{})["aggregateRating"].(map[string]interface{})["ratingValue"].(float64)
		product.graph.aggregateRating.reviewCount = result["graph"].([]interface{})[0].(map[string]interface{})["aggregateRating"].(map[string]interface{})["reviewCount"].(float64)
		product.graph.aggregateRating.worstRating = result["graph"].([]interface{})[0].(map[string]interface{})["aggregateRating"].(map[string]interface{})["worstRating"].(float64)

		review := result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})
		totalNumberOfReviews := len(review)
		product.graph.review = make([]Review, totalNumberOfReviews)
		for i := 0; i < totalNumberOfReviews; i++ {
			product.graph.review[i].reviewType = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["type"].(string)
			product.graph.review[i].description = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["description"].(string)
			product.graph.review[i].author = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["author"].(string)
			product.graph.review[i].name = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["name"].(string)

			product.graph.review[i].reviewRating.reviewRatingType = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["reviewRating"].(map[string]interface{})["type"].(string)
			product.graph.review[i].reviewRating.worstRating = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["reviewRating"].(map[string]interface{})["worstRating"].(float64)
			product.graph.review[i].reviewRating.ratingValue = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["reviewRating"].(map[string]interface{})["ratingValue"].(float64)
			product.graph.review[i].reviewRating.bestRating = result["graph"].([]interface{})[0].(map[string]interface{})["review"].([]interface{})[i].(map[string]interface{})["reviewRating"].(map[string]interface{})["bestRating"].(float64)
		}

		log.Printf("product.graph: %v", product.graph)
		// log.Printf("product.graph.review[0].author: %v", product.graph.review[0].author)

		return product
	}

	return product
}

/////////////////////////////////////////////

type ProductItem struct {
	company       string  `json:"company,omitempty"`
	zipcode       string  `json:"zip code,omitempty"`
	street        string  `json:"street,omitempty"`
	city          string  `json:"city,omitempty"`
	state         string  `json:"state,omitempty"`
	country       string  `json:"country,omitempty"`
	productFamily string  `json:"productFamily,omitempty"`
	product       string  `json:"product,omitempty"`
	price         float64 `json:"price,omitempty"`
	availability  string  `json:"availability,omitempty"`
	quantity      int     `json:"quantity,omitempty"`
	url           string  `json:"url,omitempty"`
	longitude     float32 `json:"longitude,omitempty"`
	latitude      float32 `json:"latitude,omitempty"`
}

type Store struct {
	zipcode   string  `json:"zip code,omitempty"`
	street    string  `json:"street,omitempty"`
	city      string  `json:"city,omitempty"`
	state     string  `json:"state,omitempty"`
	country   string  `json:"country,omitempty"`
	longitude float32 `json:"longitude,omitempty"`
	latitude  float32 `json:"latitude,omitempty"`
}

/////////////////////////////////////////////
////////////// ProductItemJSON //////////////
/////////////////////////////////////////////

type ProductItemJSON struct {
	graph struct {
		context         string          `json:"@context,omitempty"`
		productType     string          `json:"@type:omitempty"`
		name            string          `json:"name,omitempty"`
		brand           string          `json:"brand,omitempty"`
		image           string          `json:"image,omitempty"`
		sku             string          `json:"sku,omitempty"`
		description     string          `json:"description,omitempty"`
		gtin13          string          `json:"gtin13,omitempty"`
		offers          Offers          `json:"offers,omitempty"`
		aggregateRating AggregateRating `json:"aggregateRating,omitempty"`
		review          []Review        `json:"review,omitempty"`
	}
}

////////////////////////////////////
////////////// Offers //////////////
////////////////////////////////////

type Offers struct {
	offerType               string            `json:"@type,omitempty"`
	price                   float64           `json:"price,omitempty"`
	priceCurrency           string            `json:"priceCurrency,omitempty"`
	availability            string            `json:"availability,omitempty"`
	availableAtOrFrom       AvailableAtOrFrom `json:"availableAtOrFrom,omitempty"`
	availableDeliveryMethod string            `json:"availableDeliveryMethod,omitempty"`
	deliveryLeadTime        DeliveryLeadTime  `json:"deliveryLeadTime,omitempty"`
	potentialAction         PotentialAction   `json:"potentialAction,omitempty"`
	url                     string            `json:"url,omitempty"`
}

type AvailableAtOrFrom struct {
	availableAtOrFromType string `json:"@type,omitempty"`
	branchCode            string `json:"branchCode,omitempty"`
}

type DeliveryLeadTime struct {
	deliveryLeadTimeType string  `json:"@type,omitempty"`
	value                float64 `json:"value,omitempty"`
}

type PotentialAction struct {
	potentialActionType string `json:"@type,omitempty"`
}

///////////////////////////////////
///////// AggregateRating /////////
///////////////////////////////////

type AggregateRating struct {
	aggregateRatingType string  `json:"@type,omitempty"`
	bestRating          float64 `json:"bestRating,omitempty"`
	ratingValue         float64 `json:"ratingValue,omitempty"`
	reviewCount         float64 `json:"reviewCount,omitempty"`
	worstRating         float64 `json:"worstRating,omitempty"`
}

////////////////////////////////////
////////////// Review //////////////
////////////////////////////////////

type Review struct {
	reviewType   string       `json:"@type,omitempty"`
	description  string       `json:"description,omitempty"`
	author       string       `json:"author,omitempty"`
	name         string       `json:"name,omitempty"`
	reviewRating ReviewRating `json:"reviewRating,omitempty"`
}

type ReviewRating struct {
	reviewRatingType string  `json:"@type,omitempty"`
	worstRating      float64 `json:"worstRating,omitempty"`
	ratingValue      float64 `json:"ratingValue,omitempty"`
	bestRating       float64 `json:"bestRating,omitempty"`
}
