package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kanghsienchang/portto_assignment/routers"
)

func main() {
	r := gin.Default()

	routers.InitRouter(r)

	r.Run()
}
