package rodhandlers

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

func BaseSetUp() *rod.Browser {
	browser := rod.New().MustConnect().NoDefaultDevice()
	return browser
}

func SetLocation(zipcode string, browser *rod.Browser) {
	page := browser.MustPage("https://www.target.com")

	page.MustWaitLoad()

	time.Sleep(time.Second * 20)

	page.MustElement("#web-store-id-msg-btn").MustClick()

	time.Sleep(time.Second * 5)

	page.MustElement("input#zip-or-city-state").MustInput(zipcode)
	page.MustElement(`button[data-test="@web/StoreLocationSearch/button"]`).MustClick()

	page.MustElement(".h-flex-align-center.h-flex-justify-center > div:nth-child(1) > button").MustClick()

	page.Eval(`window.scrollTo(0,0)`)

}

func GetFullProductListing(browser *rod.Browser) {
	page := browser.MustPage("https://www.target.com/s?searchTerm=enfamil+baby+formula").MustWindowFullscreen()

	page.MustWaitLoad()

	time.Sleep(time.Second * 10)

	title := page.MustElement(`h2[data-test="resultsHeading"]`).MustText()
	titles := strings.Split(title, " ")
	total, _ := strconv.Atoi(titles[0])
	log.Printf("total: %v", total)

	pageCount := total / 24
	log.Printf("page count: %v", pageCount)

	productItemLinks := []string{}

	for i := 0; i < pageCount; i++ {
		for i := 0; i <= 15; i++ {
			page.Eval(`	
				window.scrollBy(0,600)
			`)
			time.Sleep(time.Second * 5)
			log.Printf("loop: %v", i)
		}

		productItems := page.MustElements("a.gCNFxQ")

		for _, v := range productItems {
			x, _ := v.Property("href")
			productItemLinks = append(productItemLinks, x.String())
		}

		log.Println(productItemLinks)

		// log.Println(len(productItems))

		page.MustElement(`button[data-test="next"]`).MustClick()

		page.MustWaitLoad()

		time.Sleep(time.Second * 10)
	}

	log.Println("Full Product Listing Downloaded with URLs.")
}

func GetProductItemDetails(browser *rod.Browser) {

}
