package main

import (
	"aid.dev/goody/db"
	"aid.dev/goody/router"
)

func init() {
	db.InitDb()
}

func main() {
	defer db.MysqlDB.Close()
	router.Router()
}
