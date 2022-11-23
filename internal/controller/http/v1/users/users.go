package users

import (
	"github.com/flaiers/flallet/internal/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func New(config Config) func(router fiber.Router) {
	controller := &Controller{service: config.Service}

	return func(router fiber.Router) {
		router.Get("/:userId", controller.readUserBalance)
		router.Post("/:userId/accrue", controller.accrueUserBalance)
		router.Post("/:userId/transfer", controller.transferUserBalance)
		router.Post("/:userId/withdraw", controller.withdrawUserBalance)
	}
}

// @summary  Read User Balance
// @tags     users
// @id       readUserBalance
// @accept   json
// @produce  json
// @param    userId path     string true " " format(uuid)
// @response 200    {object} dto.UserRead
// @response 400    {object} utils.JSONError
// @response 500    {object} utils.JSONError
// @router   /users/{userId} [get]
func (c *Controller) readUserBalance(ctx *fiber.Ctx) error {
	if ctx.Params("userId") == "" {
		return fiber.NewError(fiber.StatusBadRequest, "param userId is required")
	}
	userID, err := uuid.Parse(ctx.Params("userId"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "param userId is invalid")
	}
	user, err := c.service.FindOneOrCreate(userID)
	if err != nil {
		return err
	}

	return ctx.JSON(user.ToUserRead())
}

// @summary  Accrue User Balance
// @tags     users
// @id       accrueUserBalance
// @accept   json
// @produce  json
// @param    userId      path     string          true " " format(uuid)
// @param    userAccrual body     dto.UserAccrual true " "
// @response 200         {object} dto.UserRead
// @response 400         {object} utils.JSONError
// @response 422         {object} utils.JSONError
// @response 500         {object} utils.JSONError
// @router   /users/{userId}/accrue [post]
func (c *Controller) accrueUserBalance(ctx *fiber.Ctx) error {
	if ctx.Params("userId") == "" {
		return fiber.NewError(fiber.StatusBadRequest, "param userId is required")
	}
	userID, err := uuid.Parse(ctx.Params("userId"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "param userId is invalid")
	}
	userAccrual := &dto.UserAccrual{}
	if err := ctx.BodyParser(userAccrual); err != nil {
		return err
	}
	user, err := c.service.AccrueUserBalance(userID, userAccrual)
	if err != nil {
		return err
	}

	return ctx.JSON(user.ToUserRead())
}

// @summary  Transfer User Balance
// @tags     users
// @id       transferUserBalance
// @accept   json
// @produce  json
// @param    userId       path     string           true " " format(uuid)
// @param    UserTransfer body     dto.UserTransfer true " "
// @response 200          {object} dto.UserRead
// @response 400          {object} utils.JSONError
// @response 402          {object} utils.JSONError
// @response 422          {object} utils.JSONError
// @response 500          {object} utils.JSONError
// @router   /users/{userId}/transfer [post]
func (c *Controller) transferUserBalance(ctx *fiber.Ctx) error {
	if ctx.Params("userId") == "" {
		return fiber.NewError(fiber.StatusBadRequest, "param userId is required")
	}
	userID, err := uuid.Parse(ctx.Params("userId"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "param userId is invalid")
	}
	userTransfer := &dto.UserTransfer{}
	if err := ctx.BodyParser(userTransfer); err != nil {
		return err
	}
	user, err := c.service.TransferUserBalance(userID, userTransfer)
	if err != nil {
		return err
	}

	return ctx.JSON(user.ToUserRead())
}

// @summary  Withdraw User Balance
// @tags     users
// @id       withdrawUserBalance
// @accept   json
// @produce  json
// @param    userId         path     string             true " " format(uuid)
// @param    userWithdrawal body     dto.UserWithdrawal true " "
// @response 200            {object} dto.UserRead
// @response 400            {object} utils.JSONError
// @response 402            {object} utils.JSONError
// @response 422            {object} utils.JSONError
// @response 500            {object} utils.JSONError
// @router   /users/{userId}/withdraw [post]
func (c *Controller) withdrawUserBalance(ctx *fiber.Ctx) error {
	if ctx.Params("userId") == "" {
		return fiber.NewError(fiber.StatusBadRequest, "param userId is required")
	}
	userID, err := uuid.Parse(ctx.Params("userId"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "param userId is invalid")
	}
	userWithdrawal := &dto.UserWithdrawal{}
	if err := ctx.BodyParser(userWithdrawal); err != nil {
		return err
	}
	user, err := c.service.WithdrawUserBalance(userID, userWithdrawal, "approved")
	if err != nil {
		return err
	}

	return ctx.JSON(user.ToUserRead())
}
