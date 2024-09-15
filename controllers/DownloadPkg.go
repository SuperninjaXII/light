package controllers

import (
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func InstallPackage(c *fiber.Ctx) error {
	packageName := c.Query("name")

	cmd := exec.Command("apt", "install", "-y", packageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\nOutput: %s\n", err, output)
		return c.Status(fiber.StatusInternalServerError).SendString("Error installing package. Please check the package name and try again.")
	}

	return c.SendString(string(output))
}
