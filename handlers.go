package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/MicahAsowata/nib/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pocketbase/dbx"
)

func (r *Repo) Index(c *fiber.Ctx) error {
	books := &[]models.Book{}
	err := r.db.Select("id", "title").From("book").All(books)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "no books were found",
			})
			return nil
		}
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "All books",
		"books":   books,
	})
}

func (r *Repo) Create(c *fiber.Ctx) error {
	c.Accepts("application/json")
	book := models.Book{}
	err := c.BodyParser(&book)
	if err != nil {
		log.Print(err.Error())
		c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "the data is invalid",
		})
		return nil
	}

	_, err = r.db.Insert("book", dbx.Params{
		"title": book.Title,
	}).Execute()
	if err != nil {
		log.Println(err.Error())
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "book could not be created",
		})
		return nil
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "book created successfully",
	})
}
