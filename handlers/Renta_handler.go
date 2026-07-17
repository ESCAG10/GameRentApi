package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type RentaHandler struct {
    RentaService *services.RentaService
}

func NewRentaHandler(service *services.RentaService) *RentaHandler {
    return &RentaHandler{RentaService: service}
}

func (h *RentaHandler) CreateRenta(c *fiber.Ctx) error {
    var renta models.Renta
    if err := c.BodyParser(&renta); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
    }
    if err := h.RentaService.Create(&renta); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(fiber.StatusCreated).JSON(renta)
}

func (h *RentaHandler) GetRenta(c *fiber.Ctx) error {
    renta, err := h.RentaService.FindAll()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(renta)
}

func (h *RentaHandler) GetRentaByID(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    renta, err := h.RentaService.FindByID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Renta no encontrada"})
    }
    return c.JSON(renta)
}

func (h *RentaHandler) UpdateRenta(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    var renta models.Renta
    if err := c.BodyParser(&renta); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
    }
    if err := h.RentaService.Update(id, &renta); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Renta actualizada correctamente"})
}

func (h *RentaHandler) DeleteRenta(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    if err := h.RentaService.Delete(id); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Renta eliminada correctamente"})
}
