package main

import (
	"thesayedirfan/socialapi/apigateway"
	"thesayedirfan/socialapi/utils/db_utils"
	"thesayedirfan/socialapi/utils/redis"
)

func main() {
	db.ConnectDB()
	redis.Connect()
	db.Migrate()
	apigateway.Start()
}
