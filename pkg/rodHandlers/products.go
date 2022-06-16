package rodHandlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/go-rod/rod"
	errorhandlers "github.com/stoic-llama/formula-scraper/pkg/errorHandlers"
)

////////////////////////////////////
/////// GetTotalPagesByNav() ///////
////////////////////////////////////

func GetTotalPagesByNav(page *rod.Page) int {
	time.Sleep(time.Second * 20)

	element, err := page.Element(`.dPMkLN`)
	errorhandlers.FatalError("It might seem like a crush", err)

	elText, err := element.Text()
	errorhandlers.FatalError("But it doesn't meant that I'm serious", err)

	totalPages, err := strconv.Atoi(elText[len(elText)-1:])
	errorhandlers.FatalError("Oh baby, baby", err)

	log.Printf("GetTotalPagesByNav(): %v", totalPages)

	return totalPages // starting at position len(val)-1, give me everything to the end
}

//////////////////////////////////////////
/////// GetTcinsAndProductFamily() ///////
//////////////////////////////////////////

func GetTcinsAndProductFamily(data DataContainer) []Product {
	getProductsURL := `https://redsky.target.com/redsky_aggregations/v1/web/plp_search_v1?key=` + data.Key + `&category=` + data.Category + `&channel=` + data.Channel + `&count=` + strconv.Itoa(data.Count) + `&default_purchasability_filter=` + strconv.FormatBool(data.Default_purchasability_filter) + `&include_sponsored=` + strconv.FormatBool(data.Included_sponsored) + `&offset=` + strconv.Itoa(data.Offset) + `&page=` + data.Page + `&platform=` + data.Platform + `&pricing_store_id=` + data.Pricing_store_id + `&scheduled_delivery_store_id=` + data.Scheduled_delivery_store_id + `&store_ids=` + data.Store_ids + `&useragent=` + data.Useragent + `&visitor_id=` + data.Visitor_id
	resp, err := http.Get(getProductsURL)
	errorhandlers.FatalError("I played with your heart, got lost in the game", err)
	body, err := ioutil.ReadAll(resp.Body)
	errorhandlers.FatalError("I played with your heart, got lost in the game", err)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	var products []Product

	numberOfProducts := len(result["data"].(map[string]interface{})["search"].(map[string]interface{})["products"].([]interface{}))

	for i := 0; i < numberOfProducts; i++ {
		var productNew Product = Product{}
		products = append(products, productNew)
	}

	// go through each product on the current page
	for i := 0; i < numberOfProducts; i++ {
		products[i].Product_id = result["data"].(map[string]interface{})["search"].(map[string]interface{})["products"].([]interface{})[i].(map[string]interface{})["tcin"].(string)
		products[i].Product_family = result["data"].(map[string]interface{})["search"].(map[string]interface{})["products"].([]interface{})[i].(map[string]interface{})["item"].(map[string]interface{})["primary_brand"].(map[string]interface{})["name"].(string)
	}

	return products
}

func CreateTcinsArr(products []Product) []string {
	var tcinsArr []string
	for _, product := range products {
		tcinsArr = append(tcinsArr, product.Product_id)
	}

	return tcinsArr
}

func DivideIntoSubArrays(data DataContainer, tcinsArr []string) [][]string {
	var divided [][]string

	chunkSize := len(tcinsArr)
	if len(tcinsArr) >= 30 {
		x := float64(len(tcinsArr)) / float64(data.Count)
		// this allows math.Ceil to round up to next integer
		// math.Ceil(1.06) -> 1, not 2.  So we need to add 1 to x to round up.
		x = x + 1.0
		chunkSize = len(tcinsArr) / int(math.Ceil(x))
	}

	log.Printf("chunkSize: %v", chunkSize)

	for i := 0; i < chunkSize; i += chunkSize {
		end := i + chunkSize

		if end > len(tcinsArr) {
			end = len(tcinsArr)
		}

		divided = append(divided, tcinsArr[i:end])
	}

	log.Printf("divided array: %#v", divided)

	return divided
}

func ConvertTcinsToString(tcins []string) string {
	var tcinsStr string
	for i := 0; i < len(tcins); i++ {
		if i != len(tcins)-1 {
			tcinsStr = tcinsStr + tcins[i] + `%2C`
		}
		if i == len(tcins)-1 {
			tcinsStr = tcinsStr + tcins[i]
		}
	}

	return tcinsStr
}

///////////////////////////////////
/////// GetProductDetails() ///////
///////////////////////////////////

func GetProductDetails(data DataContainer, products []Product) []Product {
	getProductDetailsURL := `https://redsky.target.com/redsky_aggregations/v1/web_platform/product_summary_with_fulfillment_v1?key=` + data.Key + `&tcins=` + data.Tcins + `&store_id=` + data.Store_id + `&zip=` + data.Zip + `&state=` + data.State + `&latitude=` + strconv.FormatFloat(data.Latitude, 'f', -1, 64) + `&longitude=` + strconv.FormatFloat(data.Longitude, 'f', -1, 64) + `&scheduled_delivery_store_id=` + data.Scheduled_delivery_store_id + `&required_store_id=` + data.Required_store_id + `&has_required_store_id=` + strconv.FormatBool(data.Has_required_store_id)
	resp, err := http.Get(getProductDetailsURL)
	errorhandlers.FatalError("I played with your heart, got lost in the game", err)
	body, err := ioutil.ReadAll(resp.Body)
	errorhandlers.FatalError("I played with your heart, got lost in the game", err)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	for i := 0; i < len(products); i++ {
		products[i].Product = result["data"].(map[string]interface{})["product_summaries"].([]interface{})[i].(map[string]interface{})["item"].(map[string]interface{})["product_description"].(map[string]interface{})["title"].(string)
		products[i].Price = result["data"].(map[string]interface{})["product_summaries"].([]interface{})[i].(map[string]interface{})["price"].(map[string]interface{})["current_retail"].(float64)
		products[i].Availability = result["data"].(map[string]interface{})["product_summaries"].([]interface{})[i].(map[string]interface{})["fulfillment"].(map[string]interface{})["store_options"].([]interface{})[0].(map[string]interface{})["in_store_only"].(map[string]interface{})["availability_status"].(string)
		products[i].Quantity = -1 // hard code for now
		products[i].Product_url = result["data"].(map[string]interface{})["product_summaries"].([]interface{})[i].(map[string]interface{})["item"].(map[string]interface{})["enrichment"].(map[string]interface{})["buy_url"].(string)
	}

	return products
}
