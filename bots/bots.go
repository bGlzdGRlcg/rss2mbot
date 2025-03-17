package bots

import (
	"bGlzdGRlcg/rss2mbot/ms"
	"bGlzdGRlcg/rss2mbot/users"
	"os"
	"strconv"

	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v4"

	"github.com/joho/godotenv"
	"github.com/mattn/go-mastodon"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	ms.HOST = os.Getenv("MS_HOST")
	ms.Cid = os.Getenv("MS_CID")
	ms.Secret = os.Getenv("MS_SECRET")
	ms.Token = os.Getenv("MS_TOKEN")
	Token = os.Getenv("BOT_TOKEN")
}

func Start_bots() {
	config := &mastodon.Config{
		Server:       ms.HOST,
		ClientID:     ms.Cid,
		ClientSecret: ms.Secret,
		AccessToken:  ms.Token,
	}

	ms_c := mastodon.NewClient(config)

	if err := loadUserMap(); err != nil {
		fmt.Println(err)
	}

	pref := tele.Settings{
		Token:  Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("bot is running...")

	b.Handle("/start", func(c tele.Context) error {
		if User_map[c.Sender().ID] == nil {
			User_map[c.Sender().ID] = &users.User{
				Userid: c.Sender().ID,
				IsBind: false,
			}
		}
		if err := saveUserMap(); err != nil {
			fmt.Println(err)
		}
		return c.Send(Start_Text)
	})

	b.Handle("/ping", func(c tele.Context) error {
		if User_map[c.Sender().ID] == nil {
			User_map[c.Sender().ID] = &users.User{
				Userid: c.Sender().ID,
				IsBind: false,
			}
		}
		if err := saveUserMap(); err != nil {
			fmt.Println(err)
		}
		return c.Send("200 ok")
	})

	b.Handle("/bind", func(c tele.Context) error {
		if User_map[c.Sender().ID] == nil {
			User_map[c.Sender().ID] = &users.User{
				Userid: c.Sender().ID,
				IsBind: false,
			}
		}
		if c.Message().Payload == "" {
			return c.Send("Please provide UserID")
		}
		acc, err := ms.GetAcc(ms_c, c.Message().Payload)
		if err != nil {
			return c.Send("User not found")
		}
		User_map[c.Sender().ID].Bind(string(acc.ID))
		User_map[c.Sender().ID].IsBind = true
		if err := saveUserMap(); err != nil {
			fmt.Println(err)
		}
		return c.Send("Bind success")
	})

	b.Handle("/getid", func(c tele.Context) error {
		if User_map[c.Sender().ID] == nil {
			User_map[c.Sender().ID] = &users.User{
				Userid: c.Sender().ID,
				IsBind: false,
			}
		}
		if err := saveUserMap(); err != nil {
			fmt.Println(err)
		}
		if c.Message().Payload == "" {
			return c.Send("Please provide UserName")
		}
		acc, err := ms.GetAccList(ms_c, c.Message().Payload)
		if err != nil {
			return c.Send("User not found")
		}
		var send_text string
		for _, v := range acc.Accounts {
			send_text += fmt.Sprintf("ID: %s --- Name: %s\n%s\n\n", v.ID, v.Username, v.URL)
		}
		return c.Send(send_text)
	})

	b.Handle("/getinfo", func(c tele.Context) error {
		if User_map[c.Sender().ID] == nil {
			User_map[c.Sender().ID] = &users.User{
				Userid: c.Sender().ID,
				IsBind: false,
			}
		}
		if err := saveUserMap(); err != nil {
			fmt.Println(err)
		}
		send_text := fmt.Sprintf("Userid: %d\nBinduser: %s\nIsBind: %t\nSubs: %v\n", User_map[c.Sender().ID].Userid, User_map[c.Sender().ID].Binduser, User_map[c.Sender().ID].IsBind, User_map[c.Sender().ID].Subs)
		return c.Send(send_text)
	})

	b.Handle("/getsublist", func(c tele.Context) error {
		if User_map[c.Sender().ID] == nil {
			User_map[c.Sender().ID] = &users.User{
				Userid: c.Sender().ID,
				IsBind: false,
			}
		}
		if err := saveUserMap(); err != nil {
			fmt.Println(err)
		}
		var send_text string
		for index, v := range User_map[c.Sender().ID].Subs {
			send_text += fmt.Sprintf("%d: %s\n", index, v)
		}
		return c.Send(send_text)
	})

	b.Handle("/sub", func(c tele.Context) error {
		if User_map[c.Sender().ID] == nil {
			User_map[c.Sender().ID] = &users.User{
				Userid: c.Sender().ID,
				IsBind: false,
			}
		}
		if len(User_map[c.Sender().ID].Subs) > sublim {
			return c.Send("Sublist is full")
		}
		if User_map[c.Sender().ID] == nil {
			return c.Send("Please /start first")
		}

		if !User_map[c.Sender().ID].IsBind {
			return c.Send("Please /bind your Mastodon account first")
		}

		url := c.Message().Payload
		if url == "" {
			return c.Send("Please provide subscription link")
		}

		User_map[c.Sender().ID].AddRSSFeed(url)

		User_map[c.Sender().ID].AddSub(url)

		if err := saveUserMap(); err != nil {
			fmt.Println(err)
			return c.Send("Failed to save subscription")
		}

		return c.Send("Sub success")
	})

	b.Handle("/unsub", func(c tele.Context) error {
		if User_map[c.Sender().ID] == nil {
			User_map[c.Sender().ID] = &users.User{
				Userid: c.Sender().ID,
				IsBind: false,
			}
		}
		if User_map[c.Sender().ID] == nil {
			return c.Send("Please /start first")
		}
		if !User_map[c.Sender().ID].IsBind {
			return c.Send("Please /bind your Mastodon account first")
		}
		if User_map[c.Sender().ID].Subs == nil {
			return c.Send("You have no subscription")
		}
		index := c.Message().Payload
		if index == "" {
			return c.Send("Please provide index")
		}
		i, err := strconv.Atoi(index)
		if err != nil {
			return c.Send("Index must be a number")
		}
		if i >= len(User_map[c.Sender().ID].Subs) {
			return c.Send("Index out of range")
		}
		User_map[c.Sender().ID].RSSFeeds[User_map[c.Sender().ID].Subs[i]].Close()
		delete(User_map[c.Sender().ID].RSSFeeds, User_map[c.Sender().ID].Subs[i])
		User_map[c.Sender().ID].Subs = append(User_map[c.Sender().ID].Subs[:i], User_map[c.Sender().ID].Subs[i+1:]...)
		if err := saveUserMap(); err != nil {
			fmt.Println(err)
		}
		send_text := fmt.Sprintf("Unsub success\nThe subscription list is now:\n%s", User_map[c.Sender().ID].Subs)
		return c.Send(send_text)
	})

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for range ticker.C {
			for _, user := range User_map {
				if user.IsBind && user.RSSFeeds != nil {
					for _, watcher := range user.RSSFeeds {
						items, err := watcher.CheckNew()
						if err != nil {
							fmt.Printf("Error checking RSS feed: %v\n", err)
							continue
						}
						m_user, err := ms.GetAcc(ms_c, user.Binduser)
						if err != nil {
							fmt.Printf("Error getting Mastodon user: %v\n", err)
							continue
						}
						for _, item := range items {
							content := fmt.Sprintf("@%s \n📰 %s\n\nLink: %s",
								m_user.Username,
								item.Title,
								item.Link)
							ms.PostdToot(ms_c, content)
						}
					}
				}
			}
		}
	}()

	b.Start()
}
