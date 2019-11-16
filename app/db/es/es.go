package es

import (
	"context"
	"github.com/weylau/myblog-api/app/config"
	"gopkg.in/olivere/elastic.v6"
)

type Es struct {
	conn *elastic.Client
}

func Default() *Es {
	es := &Es{}
	conn, err := elastic.NewClient(elastic.SetURL(config.Configs.EsHost))
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()
	_, _, err = conn.Ping(config.Configs.EsHost).Do(ctx)
	if err != nil {
		panic(err)
	}
	es.conn = conn
	return es
}

func (this *Es) GetConn() *elastic.Client {
	return this.conn
}
