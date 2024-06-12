package endpoints

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/saddam-satria/posq-be/commons"
	"github.com/saddam-satria/posq-be/domains"
	"github.com/saddam-satria/posq-be/domains/apis"
	"github.com/saddam-satria/posq-be/models"
	"github.com/saddam-satria/posq-be/repositories"
)

func Checkout(ctx *fiber.Ctx) error {
	requestData := new(apis.CheckoutRequest)
	code := fiber.StatusBadRequest
	if err := ctx.BodyParser(requestData); err != nil {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.BAD_REQUEST[domains.En], code, nil))
		return nil
	}

	if requestData.PaymentType != "cash" && requestData.PaymentType != "debit" {
		ctx.Status(code).JSON(commons.GetResponse[any](commons.BAD_REQUEST[domains.En], code, nil))
		return nil
	}

	if requestData.CustomerId != nil {
		var customer models.Customer
		customerId, _ := uuid.Parse(*requestData.CustomerId)
		if err := repositories.FindCustomerById(customerId, &customer); err != nil {
			ctx.Status(fiber.StatusNotFound).JSON(commons.GetResponse[any](commons.NOT_FOUND[domains.En], fiber.StatusNotFound, nil))
			return nil
		}

	}

	header := new(apis.AuthHeader)
	ctx.ReqHeaderParser(header)

	token := header.Authorization

	tokenParsed := strings.Split(token, " ")

	jwtToken := tokenParsed[1]

	parsedToken, _ := commons.VerifyToken(jwtToken)

	userId := fmt.Sprintf("%v", parsedToken["id"])

	t := time.Now()

	items := []models.OrderProduct{}

	for _, item := range requestData.Items {
		items = append(
			items,
			models.OrderProduct{
				Quantity:         item.Quantity,
				Description:      item.Note,
				ProductVariantId: item.VariantId,
			},
		)
	}

	customerId := requestData.CustomerId

	if customerId != nil && *customerId == "" {
		customerId = nil
	}

	paymentType := models.CASH

	if requestData.PaymentType == "debit" {
		paymentType = models.DEBT
	}

	order := models.Order{
		ReferenceId:      fmt.Sprintf("TRX-%s", t.Format("20060102150405")),
		OrderNote:        requestData.PaymentNote,
		PaymentType:      paymentType,
		Tendered:         requestData.Tendered,
		CustomerId:       customerId,
		UserCredentialId: userId,
		Products:         items,
	}

	var response apis.CheckoutResponse

	if err := repositories.Checkout(&order, &response); err != nil {
		ctx.Status(code).JSON(commons.GetResponse[any](err.Error(), code, nil))
		return nil
	}

	ctx.Status(fiber.StatusCreated).JSON(commons.GetResponse(commons.SUCCESS[domains.En], fiber.StatusCreated, response))
	return nil

}
