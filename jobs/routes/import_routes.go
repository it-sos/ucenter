package main

import (
	"log"
	"strings"
	"ucenter/web/bootstrap"
	"ucenter/web/routes"
)

func main() {
	app := bootstrap.New("punch-in", "peng.yu@yibuyiyin.com")
	app.Bootstrap()
	app.Configure(routes.Configure)
	for _, v := range app.APIBuilder.GetRoutesReadOnly() {
		if strings.Index(v.MainHandlerName(), "controllers") > -1 {
			log.Print(v.Method())
			log.Print(v.Path())
			log.Print("===========")
		}
	}
}
