package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type notes int

var Notes notes

func (notes) list(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (notes) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (notes) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (notes) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (notes) save(c *gin.Context) {

}
func (notes) Register(r *gin.RouterGroup) {
	r.GET("/v1/notes/:id/list", Notes.list)
	r.GET("/v1/notes/:id", Notes.get)
	r.PUT("/v1/notes/:id", Notes.put)
	r.DELETE("/v1/notes/:id", Notes.delete)
	r.POST("/v1/notes", Notes.save)
}
