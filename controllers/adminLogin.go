package controllers

import (
	"jwt_token/database"
	"jwt_token/helpers"
	"jwt_token/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ===================================================================== REGISTER=====================================================

func AdminRegister(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Admin := models.Admin{}

	if contentType == appJSON{
		c.ShouldBindJSON(&Admin)
	
	}else {
		c.ShouldBind(&Admin)

	}
	err := db.Debug().Create(&Admin).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": 	Admin.ID,
		"email": Admin.Email,
		"Full_name": Admin.FirstName,

	})


}
//==================================================================END ============================================



func AdminLogin(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Admin := models.Admin{}
	password := ""

	if contentType == appJSON{
		c.ShouldBindJSON(&Admin)
	
	}else {
		c.ShouldBind(&Admin)

	}
	password = Admin.Password

	err := db.Debug().Where("email = ?", Admin.Email).Take(&Admin).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" :"Unathorized",
			"message": "Invalid Email / Password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(Admin.Password),[]byte(password))

	if !comparePass{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error" :"Unathorized",
			"message": "Invalid Email / Password",

		})
		return
	}
	token := helpers.GenerateToken(Admin.ID, Admin.Email)

	c.JSON(http.StatusOK,gin.H{
		"token" : token,
	})

}