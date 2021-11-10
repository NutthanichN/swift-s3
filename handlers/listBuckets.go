package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ncw/swift/v2"
)

func ListBuckets(ctx *fiber.Ctx) error {
	swiftConnection := ctx.UserContext().Value("swiftConnection").(*swift.Connection)
	containers, err := swiftConnection.ContainerNames(context.TODO(), nil)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot list container names",
		})
	}
	fmt.Println(fmt.Sprintf("ContainerNames: %v\n", containers))
	allContainers, err := swiftConnection.ContainerNamesAll(context.TODO(), nil)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot list all container names",
		})
	}
	fmt.Println(fmt.Sprintf("ContainerNamesAll: %v\n", allContainers))
	c3, err := swiftConnection.Containers(context.TODO(), nil)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot list containers",
		})
	}
	fmt.Println(fmt.Sprintf("Containers: %v\n", c3))
	c4, err := swiftConnection.ContainersAll(context.TODO(), nil)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot list all containers",
		})
	}
	fmt.Println(fmt.Sprintf("ContainersAll: %v\n", c4))
	return ctx.SendStatus(fiber.StatusOK)
}
