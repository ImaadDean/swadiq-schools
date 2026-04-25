package fees

import (
	"strconv"
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"

	"github.com/gofiber/fiber/v2"
)

// GetClassFeesAPI handles fetching default fees for a class
func GetClassFeesAPI(c *fiber.Ctx) error {
	classID := c.Query("class_id")
	termID := c.Query("term_id")

	if classID == "" || termID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "class_id and term_id are required",
		})
	}

	fees, err := database.GetClassFees(config.GetDB(), classID, termID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch class fees: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    fees,
	})
}

// UpsertClassFeeAPI handles setting or updating a default fee for a class
func UpsertClassFeeAPI(c *fiber.Ctx) error {
	type Request struct {
		ClassID    string `json:"class_id"`
		FeeTypeID  string `json:"fee_type_id"`
		TermID     string `json:"term_id"`
		Amount     string `json:"amount"` // String to handle various form inputs
		IsRequired bool   `json:"is_required"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid amount",
		})
	}

	if req.ClassID == "" || req.FeeTypeID == "" || req.TermID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "class_id, fee_type_id, and term_id are required",
		})
	}

	err = database.UpsertClassFee(config.GetDB(), req.ClassID, req.FeeTypeID, req.TermID, amount, req.IsRequired)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to save class fee: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Class fee saved successfully",
	})
}
