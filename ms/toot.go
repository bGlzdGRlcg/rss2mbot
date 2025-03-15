package ms

import (
	"context"
	"fmt"

	"github.com/mattn/go-mastodon"
)

func PostdToot(ms_c *mastodon.Client, content string) {
	toot := mastodon.Toot{
		Status:     content,
		Visibility: "direct",
	}
	_, err := ms_c.PostStatus(context.Background(), &toot)
	if err != nil {
		fmt.Println(err)
	}
}
