package util

import "github.com/charizer/snowflake"

var worker *snowflake.Worker

func init() {
	// todo 从别的什么地方配置worker-id
	worker, _ = snowflake.NewWorker(1)
}

func NextId() int64 {
	return worker.Next()
}
