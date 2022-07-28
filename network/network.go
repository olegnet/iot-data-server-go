package network

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"postgres"
	"strconv"
)

func Start(bind string, database postgres.Database) {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/sensor/:sensor_id", func(c *fiber.Ctx) error {
		sensorId, err := strconv.ParseInt(c.Params("sensor_id"), 0, 32)
		if err != nil {
			return BadRequest(c, "Error: sensor_id="+c.Params("sensor_id"))
		}

		sensor, errDb := database.GetLatestTemperature(sensorId)
		if errDb != nil {
			log.Println("GetLatestTemperature ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendString(fmt.Sprintf("sensor_id=%d temperature=%.2f time=%s",
			sensorId, sensor.Temperature, sensor.Time))
	})

	app.Post("/sensor/:sensor_id/:temperature", func(c *fiber.Ctx) error {
		sensorId, errSens := strconv.ParseInt(c.Params("sensor_id"), 0, 32)
		if errSens != nil {
			return BadRequest(c, "Error: sensor_id="+c.Params("sensor_id"))
		}

		temperature, errTemp := strconv.ParseFloat(c.Params("temperature"), 32)
		if errTemp != nil {
			return BadRequest(c, "Error: temperature="+c.Params("temperature"))
		}

		result, errDb := database.InsertTemperature(sensorId, temperature)
		if errDb != nil {
			log.Println("InsertTemperature ", errDb)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendString(fmt.Sprintf("result=%d", result))
	})

	log.Fatal(app.Listen(bind))
}

func BadRequest(c *fiber.Ctx, str string) error {
	log.Print(str)
	return c.Status(fiber.StatusBadRequest).SendString(str)
}
