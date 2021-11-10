package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/ncw/swift/v2"
	"github.com/rs/zerolog/log"
)

func StartServer(serverAddress string) {
	serverApp := fiber.New()
	serverApp.Use(func(c *fiber.Ctx) error {
		swiftConnection := InitiateSwiftConnection()
		c.SetUserContext(context.WithValue(c.UserContext(), "swiftConnection", swiftConnection))
		return c.Next()
	})
	registerRoutes(serverApp)
	log.Fatal().Err(serverApp.Listen(serverAddress))
}

func InitiateSwiftConnection() *swift.Connection {
	config := GetConfig()
	//fmt.Println(config.SwiftUser, config.SwiftAccessKey, config.SwiftDomain, config.SwiftAuthUrl)
	swiftConn := &swift.Connection{
		UserName: config.SwiftUser,
		ApiKey:   config.SwiftAccessKey,
		Domain:   config.SwiftDomain,
		AuthUrl:  config.SwiftAuthUrl,
	}
	err := swiftConn.Authenticate(context.Background())
	if err != nil {
		panic(err)
	}
	return swiftConn
}
