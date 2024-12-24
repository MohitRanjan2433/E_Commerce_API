package order

import (
	"github.com/gofiber/fiber/v2"
	"mohit.com/ecom-api/models"
)


func CreateOrderHandler(c *fiber.Ctx) error {
	var req models.CreateOrderRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	// Validate input fields for items
	if len(req.Items) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No items in the order",
		})
	}

	// Validate each item
	var totalPrice float64
	var orderItems []models.OrderItem
	for _, item := range req.Items {
		if item.Quantity <= 0 || item.Price <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid item: Price and Quantity must be greater than 0",
			})
		}

		// Calculate total price and prepare the order item
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID, // Use the ObjectID directly here
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
		totalPrice += item.Price * float64(item.Quantity)
	}

	// Validate shipping info
	if req.ShippingInfo.Address == "" || req.ShippingInfo.City == "" ||
		req.ShippingInfo.State == "" || req.ShippingInfo.ZipCode == "" ||
		req.ShippingInfo.Country == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Shipping information is incomplete",
		})
	}

	// Ensure userID exists in the context (you must be authenticated)
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UserID is missing in the context. Please authenticate.",
		})
	}

	// Call the CreateOrder function from the models package
	order, err := models.CreateOrder(userID, orderItems, totalPrice, req.ShippingInfo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create order: " + err.Error(),
		})
	}

	// Return the created order
	return c.Status(fiber.StatusCreated).JSON(order)
}