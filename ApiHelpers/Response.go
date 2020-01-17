package ApiHelpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code   int
	Msg    interface{}
	Result interface{}
}

func RespondJSON(w *gin.Context, status int, payload interface{}, msg string) {
	fmt.Println("code ", status)
	var res ResponseData

	res.Code = status
	res.Msg = msg
	res.Result = payload

	w.JSON(200, res)
}
