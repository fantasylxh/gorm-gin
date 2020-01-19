package ApiHelpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
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
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func Geturl(r *http.Request) string {
	return r.Host
}
