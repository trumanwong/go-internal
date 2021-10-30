package es

import (
	"github.com/elastic/go-elasticsearch/v7"
)

func NewElasticClient(username, password string, addresses []string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Username:  username,
		Password:  password,
		Addresses: addresses,
	}
	return elasticsearch.NewClient(cfg)
}
