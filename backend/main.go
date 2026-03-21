package main


import (
	"github.com/glztl/lunacat/backend/api"
)

func main() {
	// 初始化路由
	r := api.InitRouter()

	// 启动服务，端口8080
	println(" 🚀 Lunacat 区块链 API 服务启动: http://localhost:8080")
	r.Run(":8080")
}