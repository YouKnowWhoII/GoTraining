package handlers

import (
	"GO/Fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

func SetAuthDB(database *gorm.DB) {
	db = database
}

// to be moved to env
const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	// parse body
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// hash password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	// create user
	user := models.User{
		Name:        data["name"],
		Email:       data["email"],
		Password:    password,
		Permissions: data["permissions"],
	}

	db.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	// parse body
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	db.Where("email = ?", data["email"]).First(&user)

	// user not found
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	// create token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	// sign token
	token, err := claims.SignedString([]byte(SecretKey))

	// token creation failed
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	// set cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	// send cookie
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	// get cookie
	cookie := c.Cookies("jwt")

	// parse token
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// token parsing failed
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// get user from token
	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	db.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// remove cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func CheckPermissions(c *fiber.Ctx) error {
	// get cookie
	cookie := c.Cookies("jwt")

	// parse token
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	// token parsing failed
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// get user from token
	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	db.Where("id = ?", claims.Issuer).First(&user)

	// check permissions, if not admin, return forbidden
	if user.Permissions != "admin" {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "insufficient permissions",
		})
	}

	return c.Next()
}
