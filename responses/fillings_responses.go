package responses

import "github.com/gofiber/fiber/v2"

type FillingResponse struct {
	Km     int        `json:"km"`
	Volume int        `json:"volume"`
	Date   string     `json:"date"`
	Data   *fiber.Map `json:"data"`
}
