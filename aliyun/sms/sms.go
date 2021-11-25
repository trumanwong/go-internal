package sms

import (
	"github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsclient "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type SmsClient struct {
	dysmsClient *dysmsclient.Client
}

func NewSmsClient(accessKeyId, accessKeySecret, endPoint string) (*SmsClient, error) {
	config := &client.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String(endPoint)
	result, err := dysmsclient.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &SmsClient{dysmsClient: result}, nil
}

// SendSms 发送短信
func (this *SmsClient) SendSms(phone, signName, templateCode, param string) error {
	sendSmsRequest := &dysmsclient.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(templateCode),
		TemplateParam: tea.String(param),
	}
	_, err := this.dysmsClient.SendSms(sendSmsRequest)
	return err
}
