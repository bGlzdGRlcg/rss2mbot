package bots

import (
	"bGlzdGRlcg/rss2mbot/users"
	"os"
	"sync"
)

const (
	Start_Text = `
/getid username (获取Mastodon ID)
/getinfo (获取账号信息)
/bind Mastodon_ID (绑定Mastodon账号)
/sub rss_url (订阅rss)
/unsub index (取消订阅rss)
/getsublist (获取订阅列表)
/ping (200 ok)
	`
	sublim = 5
)

var (
	User_map     = make(map[int64]*users.User)
	userMapMutex sync.RWMutex
	Token        = os.Getenv("BOT_TOKEN")
)
