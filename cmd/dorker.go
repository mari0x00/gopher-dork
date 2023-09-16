package cmd

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

type Result struct {
	Id          int
	Name        string
	Url         string
	Description string
}

const MAX_LIMIT = 1000

func Dork(query string, limit int) ([]Result, error) {
	c := colly.NewCollector()
	start := 0
	var results []Result
	if limit < 1 {
		limit = MAX_LIMIT
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("[-] Visiting: ", r.URL)
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
		r.Headers.Set("cookie", "CONSENT=YES+")
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(r.Request.URL.String(), "consent") {
			fmt.Println("Need to give consent")
		}
	})

	c.OnHTML("div#center_col", func(e *colly.HTMLElement) {
		id := len(results)
		e.ForEach("h3[class]:not([role])", func(i int, e *colly.HTMLElement) {
			el_id := id + i
			if el_id < limit {
				result := Result{
					Id:   el_id,
					Name: e.Text,
				}
				results = append(results, result)
			}
		})
		e.ForEach("a[jsname][data-ved][ping]", func(i int, e *colly.HTMLElement) {
			el_id := id + i
			if el_id < len(results) {
				results[el_id].Url = e.Attr("href")
			}

		})
		e.ForEach("div[style=\"-webkit-line-clamp:2\"] > span:not([class])", func(i int, e *colly.HTMLElement) {
			el_id := id + i
			if el_id < len(results) {
				stripped := strings.ReplaceAll(e.Text, "... ", "")
				results[el_id].Description = stripped
			}
		})
	})

	for len(results) < limit {
		before := len(results)
		urlToVisit := fmt.Sprintf("https://www.google.com/search?q=%s&start=%d&sa=N", url.QueryEscape(query), start)
		err := c.Visit(urlToVisit)
		if err != nil {
			return nil, fmt.Errorf("dork: %w", err)
		}
		if len(results) <= before {
			break
		}
		start += 10
	}
	return results, nil
}
