package bots

import (
	"bGlzdGRlcg/rss2mbot/users"
	"os"
	"sync"
)

const (
	Start_Text = `
	This is listder's test bot.
	`
	sublim = 5
)

var (
	User_map     = make(map[int64]*users.User)
	userMapMutex sync.RWMutex
	Token        = os.Getenv("BOT_TOKEN")
)
