package main

import (
	"time"

	rodhandlers "github.com/stoic-llama/formula-scraper/pkg/rodHandlers"
)

func main() {
	// browser := rodhandlers.BrowserSetUp()
	// page := rodhandlers.PageSetUp(browser, "https://www.target.com")

	// rodhandlers.SetLocation(page, "10029")

	// // Get 5 nearest stores to location that was set earlier
	// stores := rodhandlers.GetNearestStores(browser, "https://www.target.com")

	// // Get all the infant formula product inventory near the selected zip code
	// listing := rodhandlers.GetFullProductListing(browser, "https://www.target.com/c/formula-nursing-feeding-baby/-/N-5xtkh")

	// var productItems []rodhandlers.ProductItem
	// for _, store := range stores {
	// 	for i := 0; i < len(listing); i++ {
	// 		productItems = append(productItems, rodhandlers.GetProductItemDetails(browser, stores[i], listing)...)
	// 		rodhandlers.SetLocationByStore(page, store)
	// 	}
	// }

	// log.Printf("Product Items from %v: %v", "10029", productItems)

	// rodhandlers.GetTargetProducts()

	// rodhandlers.GetDateToday()

	// rodhandlers.GetDateTomorrow()

	rodhandlers.TargetScrape("10029")

	time.Sleep(time.Hour)
}
