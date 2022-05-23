package rodhandlers

import (
	"log"
	"time"

	"github.com/go-rod/rod"
)

func PageLoader() {
	browser := rod.New().MustConnect().NoDefaultDevice()
	page := browser.MustPage("https://www.target.com/s?searchTerm=enfamil+baby+formula").MustWindowFullscreen()

	page.MustWaitLoad()

	time.Sleep(time.Second * 10)

	for i := 0; i <= 15; i++ {
		page.Eval(`	
			window.scrollBy(0,600)
		`)
		time.Sleep(time.Second * 5)
		log.Printf("loop: %v", i)
	}

	// 	document.getElementById("email-address").scrollIntoView({behavior: "smooth"});

	x := page.MustElements("a.gCNFxQ") //.MustInput("earth")

	for _, y := range x {
		log.Println(y.Property("href")) // Attribute("href"))
	}
	// page.MustElement("#search-form > fieldset > button").MustClick()

	log.Println(len(x))
	// log.Println(x)
	// page.MustWaitLoad() //.MustScreenshot("a.png")
	// time.Sleep(time.Second * 30)
}
