package rest

import (
	"github.com/gin-gonic/gin"
)

func NewResource(router *gin.Engine, controller Controller) {
	resource := controller.GetResource()
	router.GET(getResourcePath(resource), controller.GetAll)
	router.POST(getResourcePath(resource), controller.Create)
}

func getResourcePath(category string, params ...string) string {
	resource := "/" + category
	for _, p := range params {
		resource += "/" + p
	}
	return resource
}