package controllers

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SearchPackage(c *fiber.Ctx) error {
	packageName := c.Query("name")

	// Check if the package name is provided
	if packageName == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Package name query is required")
	}

	// Execute the appstreamcli search command
	cmd := exec.Command("appstreamcli", "search", packageName)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing appstreamcli search:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching package list")
	}

	// Split the output into lines
	packages := strings.Split(string(output), "\n")

	var parsedPackages []map[string]string

	// Regular expression to capture relevant fields from the appstreamcli output
	identifierRegex := regexp.MustCompile(`^Identifier:\s+([^\s]+)`)
	nameRegex := regexp.MustCompile(`^Name:\s+(.+)$`)
	summaryRegex := regexp.MustCompile(`^Summary:\s+(.+)$`)
	packageRegex := regexp.MustCompile(`^Package:\s+(.+)$`)
	homepageRegex := regexp.MustCompile(`^Homepage:\s+(.+)$`)
	iconRegex := regexp.MustCompile(`^Icon:\s+(.+)$`)

	var currentPackage map[string]string

	// Iterate over each line and extract information
	for _, line := range packages {
		if identifierRegex.MatchString(line) {
			// Start a new package map when we encounter an Identifier
			if currentPackage != nil {
				parsedPackages = append(parsedPackages, currentPackage)
			}
			currentPackage = make(map[string]string)
			currentPackage["identifier"] = identifierRegex.FindStringSubmatch(line)[1]
		} else if nameRegex.MatchString(line) && currentPackage != nil {
			currentPackage["name"] = nameRegex.FindStringSubmatch(line)[1]
		} else if summaryRegex.MatchString(line) && currentPackage != nil {
			currentPackage["summary"] = summaryRegex.FindStringSubmatch(line)[1]
		} else if packageRegex.MatchString(line) && currentPackage != nil {
			currentPackage["package"] = packageRegex.FindStringSubmatch(line)[1]
		} else if homepageRegex.MatchString(line) && currentPackage != nil {
			currentPackage["homepage"] = homepageRegex.FindStringSubmatch(line)[1]
		} else if iconRegex.MatchString(line) && currentPackage != nil {
			currentPackage["icon"] = iconRegex.FindStringSubmatch(line)[1]
		}
	}

	// Add the last package if it exists
	if currentPackage != nil {
		parsedPackages = append(parsedPackages, currentPackage)
	}

	// Return the parsed package data as JSON
	if len(parsedPackages) == 0 {
		return c.Status(fiber.StatusNotFound).SendString("No packages found")
	}

	return c.JSON(parsedPackages)
}

// get one package
func ShowPackage(c *fiber.Ctx) error {
	packageName := c.Query("name")

	// Check if the package name is provided
	if packageName == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Package name query is required")
	}

	// Execute the appstreamcli search command
	cmd := exec.Command("appstreamcli", "get", packageName)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing appstreamcli search:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching package list")
	}

	// Split the output into lines
	packages := strings.Split(string(output), "\n")

	var parsedPackages []map[string]string

	// Regular expression to capture relevant fields from the appstreamcli output
	identifierRegex := regexp.MustCompile(`^Identifier:\s+([^\s]+)`)
	nameRegex := regexp.MustCompile(`^Name:\s+(.+)$`)
	summaryRegex := regexp.MustCompile(`^Summary:\s+(.+)$`)
	packageRegex := regexp.MustCompile(`^Package:\s+(.+)$`)
	homepageRegex := regexp.MustCompile(`^Homepage:\s+(.+)$`)
	iconRegex := regexp.MustCompile(`^Icon:\s+(.+)$`)

	var currentPackage map[string]string

	// Iterate over each line and extract information
	for _, line := range packages {
		if identifierRegex.MatchString(line) {
			// Start a new package map when we encounter an Identifier
			if currentPackage != nil {
				parsedPackages = append(parsedPackages, currentPackage)
			}
			currentPackage = make(map[string]string)
			currentPackage["identifier"] = identifierRegex.FindStringSubmatch(line)[1]
		} else if nameRegex.MatchString(line) && currentPackage != nil {
			currentPackage["name"] = nameRegex.FindStringSubmatch(line)[1]
		} else if summaryRegex.MatchString(line) && currentPackage != nil {
			currentPackage["summary"] = summaryRegex.FindStringSubmatch(line)[1]
		} else if packageRegex.MatchString(line) && currentPackage != nil {
			currentPackage["package"] = packageRegex.FindStringSubmatch(line)[1]
		} else if homepageRegex.MatchString(line) && currentPackage != nil {
			currentPackage["homepage"] = homepageRegex.FindStringSubmatch(line)[1]
		} else if iconRegex.MatchString(line) && currentPackage != nil {
			currentPackage["icon"] = iconRegex.FindStringSubmatch(line)[1]
		}
	}

	// Add the last package if it exists
	if currentPackage != nil {
		parsedPackages = append(parsedPackages, currentPackage)
	}

	// Return the parsed package data as JSON
	if len(parsedPackages) == 0 {
		return c.Status(fiber.StatusNotFound).SendString("No packages found")
	}

	return c.JSON(parsedPackages)
}
