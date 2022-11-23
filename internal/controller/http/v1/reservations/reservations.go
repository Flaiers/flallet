package reservations

import (
	"github.com/flaiers/flallet/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func New(config Config) func(router fiber.Router) {
	controller := &Controller{service: config.Service}

	return func(router fiber.Router) {
		router.Post("", controller.createReservation)
		router.Delete("/reject", controller.rejectReservation)
		router.Delete("/approve", controller.approveReservation)
	}
}

// @summary  Create Reservation
// @tags     reservations
// @id       createReservation
// @accept   json
// @produce  json
// @param    reservationCreate body     dto.ReservationCreate true " "
// @response 201               {object} dto.ReservationRead
// @response 402               {object} utils.JSONError
// @response 422               {object} utils.JSONError
// @response 500               {object} utils.JSONError
// @router   /reservations [post]
func (c *Controller) createReservation(ctx *fiber.Ctx) error {
	reservationCreate := &dto.ReservationCreate{}
	if err := ctx.BodyParser(reservationCreate); err != nil {
		return err
	}
	reservation, err := c.service.Create(reservationCreate)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(reservation.ToReservationRead())
}

// @summary  Reject Reservation
// @tags     reservations
// @id       rejectReservation
// @accept   json
// @produce  json
// @param    reservationFilter body     dto.ReservationFilter true " "
// @response 200               {object} dto.ReservationRead
// @response 422               {object} utils.JSONError
// @response 500               {object} utils.JSONError
// @router   /reservations/reject [delete]
func (c *Controller) rejectReservation(ctx *fiber.Ctx) error {
	reservationFilter := &dto.ReservationFilter{}
	if err := ctx.BodyParser(reservationFilter); err != nil {
		return err
	}
	reservation, err := c.service.Reject(reservationFilter)
	if err != nil {
		return err
	}

	return ctx.JSON(reservation.ToReservationRead())
}

// @summary  Approve Reservation
// @tags     reservations
// @id       approveReservation
// @accept   json
// @produce  json
// @param    reservationFilter body     dto.ReservationFilter true " "
// @response 200               {object} dto.ReservationRead
// @response 422               {object} utils.JSONError
// @response 500               {object} utils.JSONError
// @router   /reservations/approve [delete]
func (c *Controller) approveReservation(ctx *fiber.Ctx) error {
	reservationFilter := &dto.ReservationFilter{}
	if err := ctx.BodyParser(reservationFilter); err != nil {
		return err
	}
	reservation, err := c.service.Approve(reservationFilter)
	if err != nil {
		return err
	}

	return ctx.JSON(reservation.ToReservationRead())
}
