package news

import (
	"fmt"

	"github.com/anaskhan96/soup"
)

// To get The news from AnimeNewsNetwork
// fmt.Println(Get_News())
func Get_News() string {
	var url string = "https://animenewsnetwork.com"
	res, err := soup.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	doc := soup.HTMLParse(res)
	news := doc.FindStrict("div", "class", "herald box news")
	data := news.Find("div", "class", "wrap")
	title := data.Find("h3").Find("a").FullText()
	newsx := data.Find("div", "class", "preview").Find("span", "class", "full")

	return fmt.Sprintf("%v\n\n%v", title, newsx.FullText())

}
