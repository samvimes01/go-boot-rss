package rss_parser

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/samvimes01/go-rss/internal/config"
	"github.com/samvimes01/go-rss/internal/db"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Atom    string   `xml:"atom,attr"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description   string `xml:"description"`
		Generator     string `xml:"generator"`
		Language      string `xml:"language"`
		LastBuildDate string `xml:"lastBuildDate"`
		Items         []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			PubDate     string `xml:"pubDate"` // Sat, 05 Nov 2022 00:00:00 +0000
			Guid        string `xml:"guid"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchData(cfg config.APPConfiger) {
	params := db.GetNextFeedsToFetchParams{Limit: 10, Offset: 0}
	ctx := cfg.GetCtx()
	feeds, err := cfg.GetDB().GetNextFeedsToFetch(*ctx, params)
	if err != nil {
		log.Fatalln(err)
	}
	for _, feed := range feeds {
		go fetchFeedData(cfg, &feed)
	}
}

func fetchFeedData(cfg config.APPConfiger, feed *db.Feed) {
	// bytes, err := os.ReadFile("cmd/rss-parser/blog.rss") //https://www.wagslane.dev/index.xml
	rss := Rss{}
	resp, err := http.Get(feed.Url)
	if err != nil {
		log.Println(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	if err := xml.Unmarshal(body, &rss); err != nil {
		log.Fatal(err)
		return
	}
	processFeed(cfg, &rss, feed)
}

func processFeed(cfg config.APPConfiger, rss *Rss, feed *db.Feed) {
	ctx := cfg.GetCtx()
	cfg.GetDB().MarkAsFetched(*ctx, feed.ID)
	fmt.Println("Fetched feed: ", rss.Channel.Title, feed.Url)
	for _, item := range rss.Channel.Items {
		var date time.Time
		if item.PubDate != "" {
			date, _ = time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
		}
		cfg.GetDB().CreatePost(*ctx, db.CreatePostParams{
			FeedID:      feed.ID,
			Title:       item.Title,
			Description: sql.NullString{Valid: item.Description != "", String: item.Description},
			Url:         item.Link,
			PublishedAt: sql.NullTime{Valid: item.PubDate != "", Time: date},
		})
		fmt.Println("Inserted feed item: ", item.Title, item.Link)
	}
}
