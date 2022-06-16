package rodHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	errorhandlers "github.com/stoic-llama/formula-scraper/pkg/errorHandlers"
)

/////////////////////////////////
/////// GetNearByStores() ///////
/////////////////////////////////

func GetNearByStores(data DataContainer) []Store {
	nearbyStoresURL := `https://api.target.com/location_proximities/v1/nearby_locations?limit=` + strconv.Itoa(data.Limit) + `&unit=` + data.Unit + `&within=` + strconv.Itoa(data.Within) + `&place=` + data.Place + `&type=` + data.NearbyType + `&key=` + data.Key
	resp, err := http.Get(nearbyStoresURL)
	errorhandlers.FatalError("Oops I did it again", err)
	body, err := ioutil.ReadAll(resp.Body)
	errorhandlers.FatalError("Oops I did it again", err)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	// stores[0] would be the current location that is set
	// stores[1..n] are nearby stores to the current location that is set
	var stores = make([]Store, len(result["locations"].([]interface{})))
	for k := 0; k < len(stores); k++ {
		stores[k].Company = `Target`
		stores[k].Store_name = result["locations"].([]interface{})[k].(map[string]interface{})["location_names"].([]interface{})[0].(map[string]interface{})["name"].(string)
		stores[k].Store_id = result["locations"].([]interface{})[k].(map[string]interface{})["location_id"].(float64)
		stores[k].Address_line1 = result["locations"].([]interface{})[k].(map[string]interface{})["address"].(map[string]interface{})["address_line1"].(string)
		if result["locations"].([]interface{})[k].(map[string]interface{})["address"].(map[string]interface{})["address_line2"] != nil {
			stores[k].Address_line2 = result["locations"].([]interface{})[k].(map[string]interface{})["address"].(map[string]interface{})["address_line2"].(string)
		}
		stores[k].City = result["locations"].([]interface{})[k].(map[string]interface{})["address"].(map[string]interface{})["city"].(string)
		stores[k].State = result["locations"].([]interface{})[k].(map[string]interface{})["address"].(map[string]interface{})["region"].(string)         // state initials like so, "CT"
		stores[k].Zip_code = result["locations"].([]interface{})[k].(map[string]interface{})["address"].(map[string]interface{})["postal_code"].(string) // zip code
		stores[k].Country = `US`
		stores[k].Longitude = result["locations"].([]interface{})[k].(map[string]interface{})["geographic_specifications"].(map[string]interface{})["latitude"].(float64)
		stores[k].Latitude = result["locations"].([]interface{})[k].(map[string]interface{})["geographic_specifications"].(map[string]interface{})["longitude"].(float64)
	}

	return stores
}

//////////////////////////////////
/////// SetStoreLocation() ///////
//////////////////////////////////

func SetStoreLocation(browser *rod.Browser, store Store) {
	url := proto.TargetCreateTarget{
		URL: "https://www.target.com",
	}
	rodPage, err := browser.Page(url)
	errorhandlers.FatalError("I think I did it again", err)

	rodPage.WaitLoad()

	time.Sleep(time.Second * 10)

	fiatsCookie := GetCookieValue(rodPage, "fiatsCookie").(string)

	// Set cookie by first deleting the initial fiatsCookie and adding a new one
	// The fiatsCookie determines the store location of the rodPage session
	// Which would determine what inventory we see for the products (availability is location specific)
	initialLocation := `document.cookie='fiatsCookie=` + fiatsCookie + `;Domain=.target.com;Path=/;expires=Thu, 08 Jun 1970 00:00:00 GMT;SameSite=Lax;Secure'`
	rodPage.Eval(initialLocation)
	time.Sleep(time.Second * 2)

	targetFiatsCookie := `DSI_` + strconv.FormatFloat(store.Store_id, 'f', -1, 64) + `|DSN_` + store.Store_name + `|DSZ_` + store.Zip_code[0:5]
	targetFiatsDate := GetDateTomorrow()
	targetLocation := `document.cookie='fiatsCookie=` + targetFiatsCookie + `;Domain=.target.com;Path=/;expires=` + targetFiatsDate + `;SameSite=Lax;Secure'`
	rodPage.Eval(targetLocation)

	time.Sleep(time.Second * 2)
}

func GetDateTomorrow() string {
	futureTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location()) // time.Now()
	futureTime = futureTime.Add(time.Hour * 24)                                                                         // add one day

	var fiatsDate string
	if futureTime.Day() < 10 {
		weekday := futureTime.Weekday().String()[0:3]
		day := futureTime.Day()
		month := futureTime.Month().String()[0:3]
		year := futureTime.Year()
		fiatsDate = weekday + ", 0" + strconv.Itoa(day) + " " + month + " " + strconv.Itoa(year) + " 00:00:00 GMT"
		// fmt.Println(fiatsDate)
		// fmt.Printf("Thu, 08 Jun 2023 00:00:00 GMT is = %v, 0%v %v %v 00:00:00 GMT", futureTime.Weekday().String()[0:3], futureTime.Day(), futureTime.Month().String()[0:3], futureTime.Year())

		return fiatsDate
	}
	if futureTime.Day() >= 10 {
		weekday := futureTime.Weekday().String()[0:3]
		day := futureTime.Day()
		month := futureTime.Month().String()[0:3]
		year := futureTime.Year()
		fiatsDate = weekday + ", " + strconv.Itoa(day) + " " + month + " " + strconv.Itoa(year) + " 00:00:00 GMT"
		// fmt.Println(fiatsDate)
		// fmt.Printf("Thu, 08 Jun 2023 00:00:00 GMT is = %v, %v %v %v 00:00:00 GMT", futureTime.Weekday().String()[0:3], futureTime.Day(), futureTime.Month().String()[0:3], futureTime.Year())

		return fiatsDate
	}

	return "could not find date"
}

func GetCookieValue(rodPage *rod.Page, cookieName string) interface{} {
	results, _ := proto.NetworkGetAllCookies{}.Call(rodPage)
	var cookieValue string

	for _, cookie := range results.Cookies {
		if cookie.Name == cookieName {
			cookieValue = cookie.Value
		}
	}

	return cookieValue
}
