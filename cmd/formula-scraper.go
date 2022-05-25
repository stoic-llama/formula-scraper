package main

import (
	"time"

	rodhandlers "github.com/stoic-llama/formula-scraper/pkg/rodHandlers"
)

func main() {
	/* 	1) user visits custom search page
	   	2) user types in zip code, baby formula name into custom search field
	   	3) program gets https://www.target.com/s?searchTerm= the baby formula name query from Step 2
	   	4) program stores search results (product listings) in a list of key-value pairs (maps)
		   Note: This will later go to database and in cache.

			item data structure
			- company name (target)
				- product (Enfamil)
					- product item (A.R. Powder Formula)
						- price
						- count
						- URL
						- store location (newington)


	   	5) for each product item in the list, program starts a thread to flag item as out-of-stock or in-stock
	   		a) if product item is out-of-stock (count = 0), skip
			b) if product item is in-stock (count > 0), find all the stores that are closest zip code

		6) program consolidates results for list of product items
			a) fill in information for price, count, URL, store location

	*/

	// x := interfaces.DefaultCompany{
	// 	Name:    "Target",
	// 	BaseURL: "https://www.target.com/",
	// }
	// log.Println(x)

	// resp := internal.DoGet()
	// // strings.Split(resp, "styles__StyledCol-sc-ct8kx6-0")

	// log.Println(resp)

	// for i := 0; i <= 100; i++ {
	// 	go doSomething(i)
	// }
	// rodhandlers.PageLoader()

	browser := rodhandlers.BaseSetUp()

	// rodhandlers.SetLocation("10029", browser)

	rodhandlers.GetFullProductListing(browser)

	time.Sleep(time.Hour)
}

// func doSomething(x int) {
// 	fmt.Printf("Hi %v", x)
// }
