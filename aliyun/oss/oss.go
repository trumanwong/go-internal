package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type Client struct {
	client *oss.Client
	bucket *oss.Bucket
}

func NewClient(endpoint, accessKeyId, accessKeySecret, bucketName string) (*Client, error) {
	// 获取OSSClient实例
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return &Client{client: client, bucket: bucket}, nil
}

// PutObject put文件至OSS
func (this *Client) PutObject(objectName string, reader io.Reader, options ...oss.Option) error {
	err := this.bucket.PutObject(objectName, reader)
	return err
}

// ListObjects 列举文件
func (this *Client) ListObjects(prefix string) ([]oss.ObjectProperties, error) {
	fileList := make([]oss.ObjectProperties, 0)
	marker := ""
	for {
		lsRes, err := this.bucket.ListObjects(oss.Marker(marker), oss.Prefix(prefix))
		if err != nil {
			return fileList, err
		}

		for _, object := range lsRes.Objects {
			fileList = append(fileList, object)
		}

		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	return fileList, nil
}

// GetSignUrl 获取签名链接
func (this *Client) GetSignUrl(objectName string) (string, error) {
	signedUrl, err := this.bucket.SignURL(objectName, oss.HTTPGet, 3600 * 8)
	if err != nil {
		return "", err
	}
	return signedUrl, nil
}

// DeleteObjects 删除对象
func (this *Client) DeleteObjects(objects []string) error {
	_, err := this.bucket.DeleteObjects(objects)
	if err != nil {
		return err
	}
	return nil
}