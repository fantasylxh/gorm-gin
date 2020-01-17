package Controllers

import (
	"gorm-gin/ApiHelpers"
	"gorm-gin/Models"
	"github.com/gin-gonic/gin"
)

func ListBook(c *gin.Context) {
	var book []Models.Book
	err := Models.GetAllBook(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book,"api接口")
	} else {

		ApiHelpers.RespondJSON(c, 200, book,"api接口")
	}
}

func AddNewBook(c *gin.Context) {
	var book Models.Book
	c.ShouldBind(&book)
	err := Models.AddNewBook(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book,"api接口")
	} else {
		ApiHelpers.RespondJSON(c, 200, book,"api接口")
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book,"api接口")
	} else {
		ApiHelpers.RespondJSON(c, 200, book,"api接口")
	}
}

func PutOneBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book,"api接口")
	}
	c.BindJSON(&book)
	err = Models.PutOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book,"api接口")
	} else {
		ApiHelpers.RespondJSON(c, 200, book,"api接口")
	}
}

func DeleteBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.DeleteBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book,"api接口")
	} else {
		ApiHelpers.RespondJSON(c, 200, book,"api接口")
	}
}
