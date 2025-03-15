package ms

import (
	"os"
)

var (
	HOST   = os.Getenv("MS_HOST")
	Cid    = os.Getenv("MS_CID")
	Secret = os.Getenv("MS_SECRET")
	Token  = os.Getenv("MS_TOKEN")
)
