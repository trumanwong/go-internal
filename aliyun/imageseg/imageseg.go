package imageseg

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	imageseg20191230 "github.com/alibabacloud-go/imageseg-20191230/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	client *imageseg20191230.Client
}

func NewClient(accessKeyId, accessKeySecret, endpoint string) (*Client, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Endpoint:        tea.String(endpoint),
	}
	client, err := imageseg20191230.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}

// SegmentHDCommonImage 通用高清分割
func (this Client) SegmentHDCommonImage(imageUrl string, async bool) (*imageseg20191230.SegmentHDCommonImageResponse, error) {
	req := &imageseg20191230.SegmentHDCommonImageRequest{
		ImageUrl: tea.String(imageUrl),
		Async:    tea.Bool(async),
	}
	resp, err := this.client.SegmentHDCommonImage(req)
	return resp, err
}

// SegmentCommonImage 通用分割
func (this Client) SegmentCommonImage(imageUrl string, returnForm string) (*imageseg20191230.SegmentCommonImageResponse, error) {
	req := &imageseg20191230.SegmentCommonImageRequest{
		ImageURL:   tea.String(imageUrl),
		ReturnForm: tea.String(returnForm),
	}
	resp, err := this.client.SegmentCommonImage(req)
	return resp, err
}

// SegmentBody 人体分割
func (this Client) SegmentBody(imageUrl, returnForm string, async bool) (*imageseg20191230.SegmentBodyResponse, error) {
	req := &imageseg20191230.SegmentBodyRequest{
		ImageURL:   tea.String(imageUrl),
		ReturnForm: tea.String(returnForm),
		Async:      tea.Bool(async),
	}
	resp, err := this.client.SegmentBody(req)
	return resp, err
}

// SegmentHDBody 高清人体分割
func (this Client) SegmentHDBody(imageUrl string) (*imageseg20191230.SegmentHDBodyResponse, error) {
	req := &imageseg20191230.SegmentHDBodyRequest{
		ImageURL: tea.String(imageUrl),
	}
	resp, err := this.client.SegmentHDBody(req)
	return resp, err
}
