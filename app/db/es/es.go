package es

import (
	"context"
	"github.com/olivere/elastic"
	"myblog-api/app/config"
	"myblog-api/app/loger"
)

type Es struct {
	conn *elastic.Client
}

func Default() *Es {
	es := &Es{}
	conn, err := elastic.NewClient(elastic.SetURL(config.Configs.EsHost))
	if err != nil {
		loger.Default().Error("es connect error:", err.Error())
		panic(err.Error())
	}
	ctx := context.Background()
	_, _, err = conn.Ping(config.Configs.EsHost).Do(ctx)
	if err != nil {
		loger.Default().Error("es ping error:", err.Error())
		panic(err.Error())
	}
	es.conn = conn
	return es
}

func (this *Es) GetConn() *elastic.Client {
	return this.conn
}
