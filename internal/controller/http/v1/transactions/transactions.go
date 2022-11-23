package transactions

import (
	"github.com/flaiers/flallet/internal/dto"
	"github.com/flaiers/flallet/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func New(config Config) func(router fiber.Router) {
	controller := &Controller{service: config.Service}

	return func(router fiber.Router) {
		router.Get("", controller.readTransactions)
	}
}

// @summary  Read Transactions
// @tags     transactions
// @id       readTransactions
// @accept   json
// @produce  json
// @param    userId           query    string               true  " " format(uuid)
// @param    page             query    int                  false " " default(1)
// @param    size             query    int                  false " " default(50)
// @param    transactionOrder query    dto.TransactionOrder false " "
// @response 200              {object} pkg.Pagination{items=[]dto.TransactionRead}
// @response 400              {object} utils.JSONError
// @response 500              {object} utils.JSONError
// @router   /transactions [get]
func (c *Controller) readTransactions(ctx *fiber.Ctx) error {
	if ctx.Query("userId") == "" {
		return fiber.NewError(fiber.StatusBadRequest, "param userId is required")
	}
	userID, err := uuid.Parse(ctx.Query("userId"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "param userId is invalid")
	}
	pagination := &pkg.Pagination[dto.TransactionRead]{}
	if err := ctx.QueryParser(pagination); err != nil {
		return err
	}
	transactionOrder := &dto.TransactionOrder{}
	if err := ctx.QueryParser(transactionOrder); err != nil {
		return err
	}
	transactions, err := c.service.FindByUserID(userID, pagination, transactionOrder)
	if err != nil {
		return err
	}
	var transactionsRead []*dto.TransactionRead
	for _, transaction := range transactions {
		transactionsRead = append(transactionsRead, transaction.ToTransactionRead())
	}

	return ctx.JSON(pagination.Paginate(transactionsRead, len(transactionsRead)))
}
