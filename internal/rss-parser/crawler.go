package rss_parser

import (
	"fmt"
	"time"

	"github.com/samvimes01/go-rss/internal/config"
	"github.com/samvimes01/go-rss/internal/env"
)

func CrawlFeeds(e *env.Env, cfg config.APPConfiger) {
	ctx := cfg.GetCtx()
	var interval time.Duration

	if e.FeedFetchInterval > 0 {
		interval = time.Duration(e.FeedFetchInterval) * time.Second
	} else {
		interval = time.Hour
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	fetchData(cfg) // since ticker will fire after interval

	go func() {
		for {
			select {
			case <-(*ctx).Done():
				fmt.Println("Crawler finished: ", time.Now())
				return
			// interval task
			case tm := <-ticker.C:
				fmt.Println("The Current time is: ", tm)
				fetchData(cfg)
			}
		}
	}()
	fmt.Println("Crawler started: ", time.Now())
}
