package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func validateEnvVariables() {
	envVars := [5]string{"PORT", "IS_SPA", "SPA_ENTRYPOINT", "STATIC_FOLDER", "PUBLIC_PATH"}

    for _, value := range envVars {
		_, defined := os.LookupEnv(value)
		if !defined {
			panic("ENV Variable is not defined: " + value);
		}
    }
}

func convertEnvVariableToBoolean(value string) bool {
	convertedValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return convertedValue
}

func main() {
	validateEnvVariables()

	PORT := os.Getenv("PORT")
	IS_SPA := convertEnvVariableToBoolean(os.Getenv("IS_SPA"))
	SPA_ENTRYPOINT := os.Getenv("SPA_ENTRYPOINT")
	STATIC_FOLDER := os.Getenv("STATIC_FOLDER")
	PUBLIC_PATH := os.Getenv("PUBLIC_PATH")

	app := fiber.New()

	app.Static(PUBLIC_PATH, STATIC_FOLDER)

	if IS_SPA {
		app.Get("/*", func(ctx fiber.Ctx) error {
			return ctx.SendFile(STATIC_FOLDER + "/" + SPA_ENTRYPOINT)
		})
	}

	log.Fatal(app.Listen(":" + PORT))
}
