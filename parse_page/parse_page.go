package parse_page

import (
	"fmt"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
)

func Parse(urlstr string) {

	resp, err := http.Get(urlstr)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)

	}

	articles := scrape.FindAll(root, scrape.ByClass("article__body"))

	for _, article := range articles {

		prs := scrape.FindAll(article, scrape.ByTag(atom.H2))

		for _, pr := range prs {


			next, ok := scrape.Find(pr.NextSibling.NextSibling.NextSibling.NextSibling, scrape.ByTag(atom.P))

			if ok {
				fmt.Println(scrape.Text(next))

			}

		}

	}

}
