package users

import "bGlzdGRlcg/rss2mbot/rss"

type User struct {
	Userid   int64
	Binduser string
	IsBind   bool
	Subs     []string
	RSSFeeds map[string]*rss.RSSWatcher
}

func (u *User) AddSub(sub string) {
	u.Subs = append(u.Subs, sub)
}

func (u *User) Bind(user string) {
	u.Binduser = user
	u.IsBind = true
}

func (u *User) AddRSSFeed(url string) {
	if u.RSSFeeds == nil {
		u.RSSFeeds = make(map[string]*rss.RSSWatcher)
	}
	u.RSSFeeds[url] = rss.NewRSSWatcher(url)
}
