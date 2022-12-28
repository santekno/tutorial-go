package cache

import "time"

const (
	albumsKey      = "albums"
	albumDetailKey = "albums:%d"
	expiration     = time.Hour * 1
)
