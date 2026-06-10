package controller

import (
	"blog/database"
	"blog/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const maxUploadSize = int64(5 * 1024 * 1024)

var allowedUploadExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

func saveUpload(c *fiber.Ctx, file *multipart.FileHeader) (string, error) {
	if file == nil || file.Size == 0 {
		return "", nil
	}

	if file.Size > maxUploadSize {
		return "", fmt.Errorf("file exceeds %d bytes", maxUploadSize)
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedUploadExtensions[ext] {
		return "", fmt.Errorf("unsupported file extension")
	}

	uploadDir := filepath.Join(".", "static", "uploads")
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", err
	}

	safeFilename := fmt.Sprintf("%s%s", uuid.NewString(), ext)
	filename := filepath.Join(uploadDir, safeFilename)

	if err := c.SaveFile(file, filename); err != nil {
		return "", err
	}

	return filepath.ToSlash(filename), nil
}

// BlogList list blogs
func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"status": "Ok",
		"msg":    "blog list success",
	}

	time.Sleep(time.Millisecond * 1500)
	db := database.DBConn

	var records []model.Blog

	db.Find(&records)

	context["blog_records"] = records
	c.Status(200)
	return c.JSON(context)
}

// BlogDetail GET blog
func BlogDetail(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"status": "",
		"msg":    "",
	}
	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record is not exists.")
		context["msg"] = "Record is not exists."

		c.Status(404)
		return c.JSON(context)
	}
	context["record"] = record
	context["status"] = "Ok"
	context["msg"] = "Blog detail success"
	c.Status(200)
	return c.JSON(context)
}

// BlogCreate Add a blog in database
func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"status": "Ok",
		"msg":    "blog create success",
	}
	record := new(model.Blog)

	if err := c.BodyParser(record); err != nil {
		log.Println("blog create err: ", err)
		context["status"] = ""
		context["msg"] = "Something went wrong."
	}

	// File upload
	file, err := c.FormFile("file")
	if err != nil && err != fiber.ErrBadRequest {
		log.Println("Error in file upload:", err)
		return c.Status(400).SendString("File upload error")
	}

	if file != nil {
		filename, err := saveUpload(c, file)
		if err != nil {
			log.Println("Error in file uploading:", err)
			return c.Status(400).SendString("Invalid file upload")
		}
		record.Image = filename
	}

	result := database.DBConn.Create(record)
	if result.Error != nil {
		log.Println("Error in saving data:", result.Error)
		return c.Status(500).SendString("Error saving data")
	}

	context["msg"] = "Record is saved successfully."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

// BlogUpdate  Update a blog
func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"status": "Ok",
		"msg":    "Update Blog",
	}
	//http://localhost:8000/2
	id := c.Params("id")
	var record model.Blog
	database.DBConn.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not Found.")
		context["status"] = "error"
		context["msg"] = "Record not Found."
		c.Status(400)
		return c.JSON(context)
	}
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
		context["status"] = "error"
		context["msg"] = "Something went wrong."
		c.Status(400)
		return c.JSON(context)
	}
	// File upload
	file, err := c.FormFile("file")

	if err != nil && err != fiber.ErrBadRequest {
		// Handle error other than bad request
		log.Println("Error in file upload.", err)
		context["status"] = "error"
		context["msg"] = "Error in file upload."
		c.Status(400)
		return c.JSON(context)
	}

	if file != nil && file.Size > 0 {
		filename, err := saveUpload(c, file)
		if err != nil {
			log.Println("Error in file uploading...", err)
			context["status"] = "error"
			context["msg"] = "Invalid file upload."
			c.Status(400)
			return c.JSON(context)
		}

		// Set image path to the struct
		record.Image = filename
	}

	result := database.DBConn.Save(&record)

	if result.Error != nil {
		log.Println("Error in saving data.")
		context["status"] = "error"
		context["msg"] = "Error in saving data."
		c.Status(400)
		return c.JSON(context)
	}
	context["msg"] = "Record updated successfully."
	context["data"] = record
	c.Status(200)
	return c.JSON(context)
}

// BlogDelete Delete a blog
func BlogDelete(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"status": "",
		"msg":    "",
	}
	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record is not exists.")
		context["msg"] = "Record is not exists."

		c.Status(400)
		return c.JSON(context)
	}
	// Remove image
	filename := record.Image

	err := os.Remove(filename)
	if err != nil {
		log.Println("Error in deleting file.", err)
	}

	result := database.DBConn.Delete(&record)
	if result.Error != nil {
		context["msg"] = "Something went wrong."
	}

	context["status"] = "Ok"
	context["msg"] = "Record is deleted successfully."
	c.Status(200)
	return c.JSON(context)
}
