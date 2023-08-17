package controllers

import (
	"fmt"
	"strconv"

	"github.com/Izzy499/crud_api/initializers"
	"github.com/Izzy499/crud_api/models"
	"github.com/Izzy499/crud_api/structs"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	v := validator.New()
	userId := c.Locals("userId")
	post := &structs.Post{}

	err := c.BodyParser(post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err = v.Struct(*post)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("%v %v %v", err.Field(), err.Tag(), err.Param()),
			})
		}
	}

	id, _ := userId.(int)

	postDetails := models.Post{
		Text:   post.Text,
		Images: post.Images,
		UserId: id,
	}

	result := initializers.DB.Create(&postDetails)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "created successfully",
		"data":    post,
	})
}

func GetAllPost(c *fiber.Ctx) error {
	var posts []models.Post
	var count int64
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	if limit == 0 {
		initializers.DB.Select("id,user_id,text,images").Limit(15).Find(&posts).Count(&count)
	} else {
		initializers.DB.Select("*").Limit(limit).Offset(skip).Find(&posts).Count(&count)
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Successful",
		"data":    posts,
		"count":   count,
	})
}

func AddComment(c *fiber.Ctx) error {
	
	v := validator.New()
	userId := c.Locals("userId")

	comment := &structs.Comments{}

	err := c.BodyParser(comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	err = v.Struct(*comment)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("%v %v %v", err.Field(), err.Tag(), err.Param()),
			})
		}
	}

	id, _ := userId.(int)

	postDetails := &models.Comments{
		Comment: comment.Comment,
		PostId:  comment.PostId,
		UserId:  id,
	}

	result := initializers.DB.Create(&postDetails)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   result.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}
