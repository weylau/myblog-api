package front

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"myblog-api/app/db/mongo"
	"time"
)

type AccessLog struct {
}

type AccessLogParams struct {
	Ip        string `json:"ip"`
	Timestamp int64  `json:"timestamp"`
	Path      string `json:"path"`
	Date      string `json:"date"`
}

func (this *AccessLog) Add(log *AccessLogParams) error {
	mongoconn := mongo.Default().GetConn()
	collection := mongoconn.Database("myblog").Collection("access_log")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	access_log, err := bson.Marshal(log)
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(ctx, access_log)
	return err
}
