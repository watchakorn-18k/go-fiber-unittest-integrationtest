package configuration

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		AppName:     ")϶ go-fiber-unittest ϵ(",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
}
