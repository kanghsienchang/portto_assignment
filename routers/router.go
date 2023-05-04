package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kanghsienchang/portto_assignment/controllers"
)

func InitRouter(r *gin.Engine) {

	api := r.Group("/api")
	{
		// Blocks
		api.GET("/blocks", controllers.GetBlocks)
		api.GET("/blocks/:id", controllers.GetBlockById)

		// Transactions
		api.GET("/transactions/:txHash", controllers.GetTransactionByTxHash)
	}
}
