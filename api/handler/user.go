package handler

import (
	"fiber-apis/database"
	"fiber-apis/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))
	return uid == n
}

func validUser(id string, p string) bool {
	db := database.DBConn
	var user models.User
	db.First(&user, id)
	if user.Username == "" {
		return false
	}
	return true

}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var user models.User
	db.Find(&user, id)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func Register(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't register", "data": nil})
	}

	encryptedPassword, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Password hashing error", "data": nil})
	}

	user.Password = encryptedPassword

	result := db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't register", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"register": user, "encryptedPassword": encryptedPassword})

}
