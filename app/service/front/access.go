package front

import "myblog-api/app/db/es"
import "context"

type AccessLog struct {
}

type AccessLogParams struct {
	Ip        string `json:"ip"`
	Timestamp int64  `json:"timestamp"`
	Path      string `json:"path"`
}

func (this *AccessLog) Add(log AccessLogParams) error {
	esconn := es.Default().GetConn()
	ctx := context.Background()
	_, err := esconn.Index().
		Index("myblog_access_log").
		Type("access_log").
		BodyJson(log).
		Do(ctx)
	return err
}
