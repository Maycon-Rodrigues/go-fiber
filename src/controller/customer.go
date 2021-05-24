package controller

import (
	"go-fiber/src/db"
	"go-fiber/src/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetCustomers(c *fiber.Ctx) error {
	offset := c.Query("page", "0")

	db, err := db.Connection()
	if err != nil {
		log.Print(err)
		return c.Status(500).JSON(fiber.Map{"error": "Error to connect on data base"})
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM customers LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		log.Print(err)
		return c.Status(500).JSON(fiber.Map{"error": "Error on get customers"})
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer

		if err := rows.Scan(&customer.ID, &customer.Name,
			&customer.Email, &customer.BirthDate, &customer.Phone,
			&customer.IsActive, &customer.CreatedAt, &customer.UpdatedAt); err != nil {
			log.Print(err)
			return c.Status(500).JSON(fiber.Map{"error": "Error on scan user"})
		}

		customers = append(customers, customer)
	}

	return c.JSON(customers)
}

func StoreCustomer(c *fiber.Ctx) error {
	return c.SendString("StoreCustomer")
}

func GetCustomer(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id")
	if err != nil {
		log.Print(err)
		return c.Status(400).JSON(fiber.Map{"error": "ID is not valid"})
	}

	db, err := db.Connection()
	if err != nil {
		log.Print(err)
		return c.Status(500).JSON(fiber.Map{"error": "Error to connect on data base"})
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM customers WHERE id = ?;", ID)
	if err != nil {
		log.Print(err)
		return c.Status(500).JSON(fiber.Map{"error": "Error on get customer"})
	}
	defer row.Close()

	var customer models.Customer
	if row.Next() {
		if err = row.Scan(&customer.ID, &customer.Name,
			&customer.Email, &customer.BirthDate, &customer.Phone,
			&customer.IsActive, &customer.CreatedAt, &customer.UpdatedAt); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error on scan user"})
		}
	}

	if customer.ID == 0 {
		return c.Status(400).JSON(fiber.Map{"message": "Can not get customer by id"})
	}

	return c.JSON(customer)
}

func UpdateCustomer(c *fiber.Ctx) error {
	return c.SendString("UpdateCustomer")
}

func RemoveCustomer(c *fiber.Ctx) error {
	return c.SendString("RemoveCustomer")
}
