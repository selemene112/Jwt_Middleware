package middlewares

import (
	"jwt_token/database"
	"jwt_token/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc{
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": err.Error(),
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := models.Product{}


		err = db.Select("user_id").First(&Product, uint(productId)).Error
		if err != nil{
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Data not Found",
				"message": " Data tidak ada yang cocock",

			})
			return
		}

		if Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error ": " Unauthorized",
				"message": " tidak dapat memunculkan data",
			})

			return
		}
		c.Next()

	}
}



//===================================================================== Aut Admin =================================================
func AutAdmin() gin.HandlerFunc{
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": err.Error(),
			})
			return
		}
		AdminData := c.MustGet("userData").(jwt.MapClaims)
		AdminID := uint(AdminData["id"].(float64))
		Product := models.Admin{}


		err = db.Select("user_id").First(&Product, uint(productId)).Error
		if err != nil{
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Data not Found",
				"message": " Data tidak ada yang cocock",

			})
			return
		}

		if Product.AdminID != AdminID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error ": " Unauthorized",
				"message": " tidak dapat memunculkan data",
			})

			return
		}
		c.Next()

	}
}
// ======================================================================END Admin=========================================


func AutUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		db := database.GetDB()
		
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": err.Error(),
			})
			return
		}

		
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["userid"].(float64))
		Product := models.Product{}


		err = db.Select("user_id").First(&Product, uint(productId)).Error
		if err != nil{
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Data not Found",
				"message": " Data tidak ada yang cocock",

			})
			return
		}

		if Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error ": " Unauthorized",
				"message": " tidak dapat memunculkan data",
			})

			return
		}
		c.Next()

	}
}