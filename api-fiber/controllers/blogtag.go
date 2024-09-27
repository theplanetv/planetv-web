package controllers

import (
	"api-fiber/libs"
	"api-fiber/models"
	"api-fiber/services"

	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BlogTagController struct {
	service services.BlogTagService
}

func (controller *BlogTagController) New() {
	controller.service.Open()
}

func (controller *BlogTagController) Count(c *fiber.Ctx) error {
	// Open and close database after ending
	controller.service.Open()
	defer controller.service.Close()

	// Execute command
	data, err := controller.service.Count()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": libs.GET_DATA_UNSUCCESS,
		})
	}

	// Return success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data":    data,
		"message": libs.GET_DATA_SUCCESS,
	})
}

func (controller *BlogTagController) GetAll(c *fiber.Ctx) error {
	// Get limit query
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	if limit < 10 {
		limit = 10
	} else if limit > 50 {
		limit = 50
	}

	// Get page query
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	if page < 1 {
		page = 1
	}

	// Open and close database after ending
	controller.service.Open()
	defer controller.service.Close()

	// Execute command
	data, err := controller.service.GetAll(&limit, &page)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": libs.GET_DATA_UNSUCCESS,
		})
	}

	// Return success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data":    data,
		"message": libs.GET_DATA_SUCCESS,
	})
}

func (controller *BlogTagController) Create(c *fiber.Ctx) error {
	// Get input body
	inputdata := models.BlogTag{}
	if err := c.BodyParser(&inputdata); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": libs.INPUT_ERROR,
		})
	}

	// Open and close database after ending
	controller.service.Open()
	defer controller.service.Close()

	// Execute command
	err := controller.service.Create(&inputdata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": libs.CREATE_UNSUCCESS,
		})
	}

	// Return success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": libs.CREATE_SUCCESS,
	})
}

func (controller *BlogTagController) Update(c *fiber.Ctx) error {
	// Get input body
	inputdata := models.BlogTag{}
	if err := c.BodyParser(&inputdata); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": libs.INPUT_ERROR,
		})
	}

	// Open and close database after ending
	controller.service.Open()
	defer controller.service.Close()

	// Execute command
	err := controller.service.Update(&inputdata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": libs.UPDATE_UNSUCCESS,
		})
	}

	// Return success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": libs.UPDATE_SUCCESS,
	})
}

func (controller *BlogTagController) Remove(c *fiber.Ctx) error {
	// Get id param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": libs.INPUT_ERROR,
		})
	}

	// Open and close database after ending
	controller.service.Open()
	defer controller.service.Close()

	// Execute command
	err = controller.service.Remove(&id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": libs.REMOVE_UNSUCCESS,
		})
	}

	// Return success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": libs.REMOVE_SUCCESS,
	})
}
