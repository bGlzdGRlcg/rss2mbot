# Rss2Mastodon bot

## 开始之前

从`BotFather`那里创建一个 Bot,获取到`telegram`的 bot token

在 `MS_HOST/settings/applications` (MS_HOST 为你的`Mastodon`地址) 创建一个应用

之后你可以得到

应用 ID "xxx"
应用密钥 "xxx"
你的访问令牌 "xxx"

## 配置

在项目目录新建`.env`文件

```.env
BOT_TOKEN="xxx:xxx"
MS_HOST="https://xxx.xxx"
MS_CID="xxx"
MS_SECRET="xxx"
MS_TOKEN="xxx"
```

-   BOT_TOKEN 为`telegram`的 bot token
-   MS_HOST 为你的`Mastodon`地址
-   MS_CID 为你获取到的`Mastodon`应用 ID
-   MS_SECRET 为你获取到的`Mastodon`应用密钥
-   MS_TOKEN 为你获取到的`Mastodon`你的访问令牌

## 运行

终端下输入

```shell
go mod tidy && go build
./rss2mbot
```
