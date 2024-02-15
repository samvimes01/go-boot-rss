package rss_parser

import (
	"fmt"
	"time"

	"github.com/samvimes01/go-rss/internal/config"
	"github.com/samvimes01/go-rss/internal/env"
)

func CrawlFeeds(e *env.Env, cfg config.APPConfiger) {
	var interval time.Duration
	if e.FeedFetchInterval > 0 {
		interval = time.Duration(e.FeedFetchInterval) * time.Second
	} else {
		interval = time.Hour
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	fetchData(cfg)
	ctx := cfg.GetCtx()

	go func() {
		for {
			select {
			case <-(*ctx).Done():
				fmt.Println("Finished: ", time.Now())
				return
			// interval task
			case tm := <-ticker.C:
				fmt.Println("The Current time is: ", tm)
				fetchData(cfg)
			}
		}
	}()
	fmt.Println("Started: ", time.Now())

	<-(*ctx).Done()
}
