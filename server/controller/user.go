package controller

//
//import (
//	"blog/database"
//	"golang.org/x/crypto/bcrypt"
//
//	//"blog/helper"
//	"blog/model"
//	"github.com/gofiber/fiber/v2"
//	//"golang.org/x/crypto/bcrypt"
//	"log"
//)
//
//// FormData represents the payload for login and register requests
//type FormData struct {
//	Email    string `json:"email"`
//	Password string `json:"password"`
//}
//
//// Login handles user login and returns a JWT token if authentication is successful
//func Login(c *fiber.Ctx) error {
//	returnObject := fiber.Map{
//		"status": "",
//		"msg":    "Something went wrong.",
//	}
//
//	// Check user credentials
//	var formData FormData
//	if err := c.BodyParser(&formData); err != nil {
//		log.Println("Form parsing error:", err)
//		returnObject["msg"] = "Invalid request"
//		return c.Status(fiber.StatusBadRequest).JSON(returnObject)
//	}
//
//	var user model.User
//	result := database.DBConn.Where("email = ?", formData.Email).First(&user)
//	if result.Error != nil {
//		returnObject["msg"] = "User not found."
//		return c.Status(fiber.StatusBadRequest).JSON(returnObject)
//	}
//
//	// Validate password
//	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))
//	if err != nil {
//		log.Println("Invalid password:", err)
//		returnObject["msg"] = "Password doesn't match"
//		return c.Status(fiber.StatusUnauthorized).JSON(returnObject)
//	}
//
//	// Create token
//	token, err := helper.GenerateToken(user)
//	if err != nil {
//		returnObject["msg"] = "Token creation error."
//		return c.Status(fiber.StatusUnauthorized).JSON(returnObject)
//	}
//
//	returnObject["token"] = token
//	returnObject["user"] = user
//	returnObject["status"] = "OK"
//	returnObject["msg"] = "User authenticated"
//	return c.Status(fiber.StatusOK).JSON(returnObject)
//}
//
//// Register handles user registration
//func Register(c *fiber.Ctx) error {
//	returnObject := fiber.Map{
//		"status": "OK",
//		"msg":    "Register route",
//	}
//
//	// Collect form data
//	var formData FormData
//	if err := c.BodyParser(&formData); err != nil {
//		log.Println("Error parsing request:", err)
//		returnObject["msg"] = "Error occurred."
//		return c.Status(fiber.StatusBadRequest).JSON(returnObject)
//	}
//
//	// Add form data to model
//	var user model.User
//	user.Email = formData.Email
//	user.Password = helper.HashPassword(formData.Password)
//
//	result := database.DBConn.Create(&user)
//	if result.Error != nil {
//		log.Println("Database error:", result.Error)
//		returnObject["msg"] = "User already exists."
//		return c.Status(fiber.StatusBadRequest).JSON(returnObject)
//	}
//
//	returnObject["msg"] = "User added successfully."
//	return c.Status(fiber.StatusCreated).JSON(returnObject)
//}
//
//// Logout handles user logout (not implemented)
//func Logout(c *fiber.Ctx) error {
//	return c.SendStatus(fiber.StatusNotImplemented)
//}
//
//// RefreshToken handles token refreshing
//func RefreshToken(c *fiber.Ctx) error {
//	returnObject := fiber.Map{
//		"status": "OK",
//		"msg":    "Refresh Token route",
//	}
//
//	email := c.Locals("email")
//	if email == nil {
//		log.Println("Email key not found.")
//		returnObject["msg"] = "Email not found."
//		return c.Status(fiber.StatusUnauthorized).JSON(returnObject)
//	}
//
//	var user model.User
//	result := database.DBConn.Where("email = ?", email).First(&user)
//	if result.Error != nil {
//		returnObject["msg"] = "User not found."
//		return c.Status(fiber.StatusBadRequest).JSON(returnObject)
//	}
//
//	token, err := helper.GenerateToken(user)
//	if err != nil {
//		returnObject["msg"] = "Token creation error."
//		return c.Status(fiber.StatusUnauthorized).JSON(returnObject)
//	}
//
//	returnObject["token"] = token
//	returnObject["user"] = user
//
//	return c.Status(fiber.StatusOK).JSON(returnObject)
//}
//
//// Profile handles user profile requests (protected route)
//func Profile(c *fiber.Ctx) error {
//	returnObject := fiber.Map{
//		"status": "OK",
//		"msg":    "User profile route. This is a protected route accessible to authenticated users only.",
//	}
//
//	return c.Status(fiber.StatusOK).JSON(returnObject)
//}
