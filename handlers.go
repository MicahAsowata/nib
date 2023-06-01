package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/MicahAsowata/nib/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pocketbase/dbx"
	"golang.org/x/crypto/bcrypt"
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

func (r *Repo) CreateUser(c *fiber.Ctx) error {
	user := models.User{}
	//Get User
	c.Accepts("application/json")
	err := c.BodyParser(&user)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "invalid data 1",
		})
	}
	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "invalid data 2",
		})
	}
	// Insert it to the DB
	_, err = r.db.Insert("users", dbx.Params{
		"name":     user.Name,
		"email":    user.Email,
		"password": hash,
	}).Execute()

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "there was an issue 3",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "welcome to Nib",
	})
}

func (r *Repo) LoginUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	// Get Name and Email
	cred := models.LoginCredentials{}
	err := c.BodyParser(&cred)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "invalid data 11",
		})
	}
	type ValidUser struct {
		Name     string
		Email    string
		Password []byte
	}

	user := ValidUser{}
	// Check if the email is correct
	err = r.db.Select("name", "email", "password").From("users").Where(dbx.HashExp{"email": cred.Email}).One(&user)
	if err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "something is wrong with those credentials 12",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "there was a problem finding you in our systems 13",
		})
	}
	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(cred.Password))
	if err != nil {
		log.Println(err)
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "there was a problem finding you in our systems 14",
			})
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "something is wrong with those credentials 15",
		})
	}
	// Give the user the desired response
	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"success": fmt.Sprintf("Welcome %s 🎉🎉", user.Name),
	})
}
