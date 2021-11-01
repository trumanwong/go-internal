package util

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/cdn"
)

type Qiniu struct {
	mac *qbox.Mac
	cdnManager *cdn.CdnManager
}

func NewQiniu(accessKey string, secretKey string) *Qiniu {
	mac := qbox.NewMac(accessKey, secretKey)
	cdnManager := cdn.NewCdnManager(mac)
	return &Qiniu{mac: mac, cdnManager: cdnManager}
}