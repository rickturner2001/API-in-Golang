package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/rickturner2001/crm-test/database"
	"github.com/rickturner2001/crm-test/models"
	"golang.org/x/crypto/bcrypt"
)




func Home(c *fiber.Ctx) error{
	return c.SendString("Hello World!")
}

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	db := database.DBConection
	var data map[string]string
	if err :=c.BodyParser(&data); err != nil{
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	
	user := models.User{
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	db.Create(&user)
	return c.JSON(user)
}


func Login(c *fiber.Ctx) error{
	var data map[string]string
	if err :=c.BodyParser(&data); err != nil{
		return err
	}

	var user models.User

	database.DBConection.Where("email = ?", data["email"]).First(&user)
	
	if user.ID == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix() ,
	})

	token, err := claims.SignedString([] byte(SecretKey))

	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not log in",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func User(c *fiber.Ctx) error{
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated user",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	
	database.DBConection.Where("id = ?", claims.Issuer).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error{
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}