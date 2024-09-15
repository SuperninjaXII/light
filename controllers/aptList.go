package controllers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ListPackages(c *fiber.Ctx) error {
	cmd := exec.Command("apt", "list")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching package list")
	}

	packages := strings.Split(string(output), "\n")
	var parsedPackages []map[string]string

	for _, pkg := range packages {
		if pkg != "" {
			parts := strings.Fields(pkg)
			if len(parts) >= 2 {
				packageName := strings.Split(parts[0], "/")[0]
				version := parts[1]
				parsedPackages = append(parsedPackages, map[string]string{
					"name":    packageName,
					"version": version,
				})
			}
		}
	}

	return c.JSON(parsedPackages)
}
