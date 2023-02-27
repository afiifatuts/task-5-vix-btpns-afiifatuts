package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afiifatuts/go-authentication/helpers"
	"github.com/afiifatuts/go-authentication/initializer"
	"github.com/afiifatuts/go-authentication/models"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func Signup(c *gin.Context) {
	db := initializer.GetDB()
	contentType := helpers.GetContextType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
	})
}

func Login(c *gin.Context) {
	db := initializer.GetDB()
	contentType := helpers.GetContextType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email / password",
		})

		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email / password",
		})

		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(c *gin.Context) {
	db := initializer.GetDB()
	contentType := helpers.GetContextType(c)
	_, _ = db, contentType

	User := models.User{}
	NewUser := models.User{}

	id := c.Param("userId")

	err := db.First(&User, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&NewUser)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = db.Model(&User).Updates(NewUser).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"updated_at": User.UpdatedAt,
	})
}

// package controllers

// import (
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/afiifatuts/go-authentication/initializer"
// 	"github.com/afiifatuts/go-authentication/models"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// 	"golang.org/x/crypto/bcrypt"
// )

// func Signup(c *gin.Context) {
// 	//Get the emial/pass of req body
// 	var body struct {
// 		Email    string
// 		Password string
// 	}

// 	if c.Bind(&body) != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Failed to read body",
// 		})
// 		return
// 	}

// 	// Hash the password
// 	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Failed to hash password",
// 		})
// 		return
// 	}
// 	//Create the user
// 	user := models.User{Email: body.Email, Password: string(hash)}

// 	result := initializer.DB.Create(&user)

// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Failed to create user",
// 		})
// 		return
// 	}

// 	//Respon
// 	c.JSON(http.StatusOK, gin.H{})
// }

// func Login(c *gin.Context) {
// 	//Get the email and pass off req body
// 	var body struct {
// 		Email    string
// 		Password string
// 	}

// 	if c.Bind(&body) != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Failed to read body",
// 		})
// 		return
// 	}

// 	// Look up requested user
// 	var user models.User
// 	initializer.DB.First(&user, "email = ?", body.Email)

// 	if user.ID == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid email or password",
// 		})

// 		return
// 	}

// 	//Compare sent in pass with saved user pass hash

// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid email or password",
// 		})
// 		return
// 	}
// 	//Generate a jwt token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub": user.ID,
// 		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
// 	})

// 	// Sign and get the complete encoded token as a string using the secret
// 	tokenString, err := token.SignedString([]byte(os.Getenv("JTW")))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Failed to create token",
// 		})
// 		return
// 	}

// 	//Send it back
// 	c.SetSameSite(http.SameSiteLaxMode)
// 	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

// 	c.JSON(http.StatusOK, gin.H{})
// }

// func Validate(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "I'm Logged in",
// 	})
// }
