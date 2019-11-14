package es

import (
	"context"
	"github.com/weylau/myblog-api/app/configs"
	"gopkg.in/olivere/elastic.v6"
)

func NewClient() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(configs.Configs.EsHost))
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()
	_, _, err = client.Ping(configs.Configs.EsHost).Do(ctx)
	if err != nil {
		panic(err)
	}
	return client
}
