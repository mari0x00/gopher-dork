package controllers

import (
	"context"
	"fmt"
	"net/url"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func CreateConfig(cc, lc, ua, px string, limit, start int) googlesearch.SearchOptions {
	return googlesearch.SearchOptions{
		CountryCode:    "",
		LanguageCode:   "",
		Limit:          5,
		Start:          0,
		UserAgent:      "",
		OverLimit:      false,
		ProxyAddr:      "",
		FollowNextPage: true,
	}
}

func QueryDorks(query string, config googlesearch.SearchOptions) ([]googlesearch.Result, error) {
	ctx := context.Background()
	res, err := googlesearch.Search(ctx, query, config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for i := range res {
		res[i].URL, err = url.QueryUnescape(res[i].URL)
		if err != nil {
			fmt.Println(err)
		}
	}

	return res, nil
}
