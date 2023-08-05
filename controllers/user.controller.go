package controllers

import (
	"fmt"

	"github.com/Izzy499/crud_api/initializers"
	"github.com/Izzy499/crud_api/models"
	"github.com/Izzy499/crud_api/structs"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	v := validator.New()
	var userModel *models.User

	user := &structs.Register{}

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	err = v.Struct(*user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("%v %v %v", err.Field(), err.Tag(), err.Param()),
			})
		}
	}

	initializers.DB.Where("email = ?", user.Email).First(&userModel)
	if userModel.Id != 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"error":   "user already exists",
		})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	userDetails := models.User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Password:    string(hashPassword),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}

	result := initializers.DB.Create(&userDetails)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "created successfully",
		"data": models.User{
			FirstName:   userDetails.FirstName,
			LastName:    userDetails.LastName,
			PhoneNumber: userDetails.PhoneNumber,
			Email:       userDetails.Email,
		},
	})
}

func Login(c *fiber.Ctx) error {
	v := validator.New()
	var userModel models.User
	user := &structs.Login{}

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	err = v.Struct(*user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("%v %v %v", err.Field(), err.Tag(), err.Param()),
			})
		}
	}

	initializers.DB.Where("email = ?", user.Email).First(&userModel)
	if userModel.Id == 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"error":   "incorrect email or password",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "incorrect email or password",
		})
	}

	sess, err := initializers.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"err":     err.Error(),
		})
	}

	sess.Set("userId", userModel.Id)

	err = sess.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"err":     err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}

func GetUserDetails(c *fiber.Ctx) error {
	var user models.User

	userId := c.Locals("userId")
	if userId == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "not authenticated",
		})
	}

	fmt.Println(userId)
	initializers.DB.Select("id,first_name, last_name, email, phone_number").Find(&user, userId)

	if user.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func GetUserById(c *fiber.Ctx) error {
	var user models.User
	userId := c.Params("userId")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "no id passed",
		})
	}

	initializers.DB.Select("id, first_name, last_name,phone_number").Find(&user, userId)
	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}
