package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Models"
	"net/http"
	"strconv"
)

func ListAddress(c *gin.Context) {
	var Address []Models.Address
	address_type := c.DefaultPostForm("address_type", "0")
	int_address_type, _ := strconv.Atoi(address_type)
	if int_address_type > 1 || int_address_type < 0 {
		ApiHelpers.RespondJSON(c, 0, "", "address_type非法参数")
		return
	}
	err := Models.GetAllAddress(&Address)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, Address, err.Error())
	} else {

		ApiHelpers.RespondJSON(c, 200, Address, "success")
	}
}

func AddNewAddress(c *gin.Context) {
	var Address Models.Address

	if err := c.ShouldBind(&Address); nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code": 0,
			"Msg":  err.Error(),
		})
		return
	}
	address_type := c.DefaultPostForm("address_type", "0")
	uid := c.PostForm("uid")
	int_address_type, _ := strconv.Atoi(address_type)

	if int_address_type > 0 {
		int_address_type = 1
	} else {
		int_address_type = 0
	}
	Address.AddressType = int_address_type
	Intuid, _ := strconv.Atoi(uid)
	Address.CreatorId = Intuid

	err := Models.AddNewAddress(&Address)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, Address, "success")
	}
}

func GetOneAddress(c *gin.Context) {
	id := c.PostForm("address_id")
	var Address Models.Address
	err := Models.GetOneAddress(&Address, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, Address, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, Address, "success")
	}
}

func PutOneAddress(c *gin.Context) {
	var Address Models.Address
	id := c.PostForm("address_id")
	err := Models.GetOneAddress(&Address, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, Address, err.Error())
		return
	}
	c.ShouldBind(&Address)
	err = Models.PutOneAddress(&Address, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, Address, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, Address, "success")
	}
}

func DeleteAddress(c *gin.Context) {
	var Address Models.Address
	id := c.PostForm("address_id")
	err := Models.DeleteAddress(&Address, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, Address, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, Address, "success")
	}
}
