package config

import "time"

var SignKey = []byte("asd@#lskd2!aw32k34242WSASdsk32")

const (
	AccessExpireTime  = time.Minute * 2
	RefreshExpireTime = time.Hour * 24
)
