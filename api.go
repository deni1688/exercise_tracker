package main

import "github.com/gin-gonic/gin"

type api struct {
	Engine *gin.Engine
}

func newApi(engine *gin.Engine) *api {
	return &api{engine}
}

func (a *api) Run() {
	a.Engine.Run()
}

func (a *api) InitExerciseApi() *api {
	a.Engine.GET("/exercises", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "notImplmented",
		})
	})

	a.Engine.POST("/exercises", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "notImplmented",
		})
	})

	return a
}

func (a *api) InitWeightApi() *api {
	a.Engine.GET("/weight", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "notImplmented",
		})
	})

	a.Engine.POST("/weight", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "notImplmented",
		})
	})

	return a
}

func (a *api) InitNutritionApi() *api {
	a.Engine.GET("/nutrition", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "notImplmented",
		})
	})

	a.Engine.POST("/nutrition", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "notImplmented",
		})
	})

	return a
}
