package main

import (
	"context"
	"encoding/xml"
	"html"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rssFeed := RSSFeed{}
	decoder := xml.NewDecoder(resp.Body)
	if err := decoder.Decode(&rssFeed); err != nil {
		return nil, err
	}

	cleanFeed := cleanHTML(&rssFeed)
	return cleanFeed, nil
}

func cleanHTML(r *RSSFeed) *RSSFeed {
	r.Channel.Title = html.UnescapeString(r.Channel.Title)
	r.Channel.Description = html.UnescapeString(r.Channel.Description)

	for i := range r.Channel.Item {
		r.Channel.Item[i].Title = html.UnescapeString(r.Channel.Item[i].Title)
		r.Channel.Item[i].Description = html.UnescapeString(r.Channel.Item[i].Description)
	}

	return r
}
