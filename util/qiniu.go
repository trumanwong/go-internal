package util

import (
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/cdn"
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