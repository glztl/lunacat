package api

import (
	"github.com/glztl/lunacat/backend/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化 API 路由
func InitRouter() *gin.Engine {	// 定义初始化函数，返回一个Gin路由引擎对象
	// 创建 Gin 路由实例, gin.Engine = 区块链API的总开关、总调度器
	r := gin.Default()	// 创建路由

	// 1. 获取整条区块链
	// func(c *gin.Context) {} 是一个匿名函数，作为路由处理器，当用户访问 /blocks 时会执行这个函数. 请求上下文，这次请求的所有信息都在这里
	r.GET("/blocks", func(c *gin.Context) {
		// 当用户访问127.0.0.1:8080/blocks时，执行这个函数
		bc := core.GetBlockchainInstance()
		c.JSON(http.StatusOK, gin.H{
			"blocks": bc.Blocks,
		})
	})

	// 2. 挖矿
	r.POST("/mine", func(c *gin.Context) {
		// 接收前端传来的数据
		var req struct {
			Data string `json:"data" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
			return 
		}

		// 执行挖矿
		bc := core.GetBlockchainInstance()
		newBlock := bc.AddBlock(req.Data)

		c.JSON(http.StatusOK, gin.H {
			"message": "挖矿成功",
			"block": newBlock,
		})
	})

	return r	// 返回路由引擎对象
}