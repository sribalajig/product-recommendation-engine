package main

import (
	_ "home24/recommendation.api/routers"

	"github.com/astaxie/beego"
	"home24/core/repository"
	"home24/persistence"
	"home24/recommendation/engine"
)

func main() {
	repository.DataProvider = persistence.NewElasticProvider()
	engine.DataProvider = persistence.NewElasticProvider()

	beego.Run()
}
