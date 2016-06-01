package parse_page

import (
	"fmt"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"github.com/sinelga/horoscope_libs/domains"
)

func Parse(urlstr string) []domains.Zodiac {

	resp, err := http.Get(urlstr)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)

	}

	articles := scrape.FindAll(root, scrape.ByClass("article__body"))
	
	var zodiacs []domains.Zodiac

	for _, article := range articles {

		prs := scrape.FindAll(article, scrape.ByTag(atom.H2))

		for _, pr := range prs {

			var zodiac domains.Zodiac
			fmt.Println(scrape.Text(pr))
			zodiac.Name=scrape.Text(pr)			
			next, ok := scrape.Find(pr.NextSibling.NextSibling.NextSibling.NextSibling, scrape.ByTag(atom.P))

			if ok {
				fmt.Println(scrape.Text(next))
				zodiac.Contents=scrape.Text(next)
				
				zodiacs=append(zodiacs,zodiac)				

			}

		}

	}

return zodiacs
}
