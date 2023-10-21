package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/waykiss/wkgo/adapters/rest"
)

type adapter struct {
	apps []rest.AppInterface
}

var webserver *gin.Engine

var Adapter adapter

//Run method to start fiber webserver
func (f adapter) Run() {
	webserver = gin.Default()
	panic("implement-me")
}
