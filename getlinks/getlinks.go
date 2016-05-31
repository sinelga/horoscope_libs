package getlinks

import (
	"fmt"
	"github.com/sinelga/horoscope_libs/domains"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"time"
)

func GetLinks(urlstr string) []domains.Linkinfo {

	fmt.Println(urlstr)
	resp, err := http.Get(urlstr)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)

	}

	matcher := func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom == atom.A {

			return scrape.Attr(n.Parent, "class") == "grid-item"
		}
		return false
	}
	//	fortuneresors := &domains.Fortuneresors{}

	grid, ok := scrape.Find(root, scrape.ByClass("grid"))

	var arrLinkinfo []domains.Linkinfo

	if ok {

		gridItems := scrape.FindAll(grid, matcher)

		var now = time.Now()

		for _, itemA := range gridItems {

			linkinfo := domains.Linkinfo{
				Created_at: now,
				Type:       "daily_horoscope",
				Link:       scrape.Attr(itemA, "href"),
			}

			arrLinkinfo = append(arrLinkinfo, linkinfo)

//			fmt.Println(scrape.Attr(itemA, "href"))

		}
		//		fortuneresors.Links = arrLinkinfo

	}


	return arrLinkinfo
}
