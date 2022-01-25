package darabonba

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dypnsapi20170525 "github.com/alibabacloud-go/dypnsapi-20170525/client"
	"github.com/alibabacloud-go/tea/tea"
)

func NewClient(accessKeyId *string, accessKeySecret *string) (result *dypnsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dypnsapi.aliyuncs.com")
	result = &dypnsapi20170525.Client{}
	result, _err = dypnsapi20170525.NewClient(config)
	return result, _err
}
