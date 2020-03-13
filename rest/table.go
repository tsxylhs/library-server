package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type table int

var Table table

func (table) list(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (table) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (table) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (table) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (table) save(c *gin.Context) {

}
func (table) Register(r *gin.RouterGroup) {
	r.GET("/v1/table/:id/list", Table.list)
	r.GET("/v1/table/:id", Table.get)
	r.PUT("/v1/table/:id", Table.put)
	r.DELETE("/v1/table/:id", Table.delete)
	r.POST("/v1/table", Table.save)
}
