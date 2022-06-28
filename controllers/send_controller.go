package controllers

import (
	"github.com/pfthink/go-whatsappproxy/services"
	"github.com/pfthink/go-whatsappproxy/structs"
	"github.com/pfthink/go-whatsappproxy/utils"
	"github.com/pfthink/go-whatsappproxy/validations"
)

type SendController struct {
	Service services.SendService
}

func NewSendController(service services.SendService) SendController {
	return SendController{Service: service}
}

func (controller *SendController) Route(app *fiber.App) {
	app.Post("/send/message", controller.SendText)
	app.Post("/send/image", controller.SendImage)
	app.Post("/send/file", controller.SendFile)
	app.Post("/send/video", controller.SendVideo)
}

func (controller *SendController) SendText(c *fiber.Ctx) error {
	var request structs.SendMessageRequest
	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)

	// add validation send message
	validations.ValidateSendMessage(request)

	if request.Type == structs.TypeGroup {
		request.Phone = request.Phone + "@g.us"
	} else {
		request.Phone = request.Phone + "@s.whatsapp.net"
	}

	response, err := controller.Service.SendText(c, request)
	utils.PanicIfNeeded(err)

	return c.JSON(utils.ResponseData{
		Code:    200,
		Message: response.Status,
		Results: response,
	})
}

func (controller *SendController) SendImage(c *fiber.Ctx) error {
	var request structs.SendImageRequest
	request.Compress = true

	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)

	file, err := c.FormFile("image")
	utils.PanicIfNeeded(err)

	request.Image = file

	//add validation send image
	validations.ValidateSendImage(request)

	if request.Type == structs.TypeGroup {
		request.Phone = request.Phone + "@g.us"
	} else {
		request.Phone = request.Phone + "@s.whatsapp.net"
	}

	response, err := controller.Service.SendImage(c, request)
	utils.PanicIfNeeded(err)

	return c.JSON(utils.ResponseData{
		Code:    200,
		Message: response.Status,
		Results: response,
	})
}

func (controller *SendController) SendFile(c *fiber.Ctx) error {
	var request structs.SendFileRequest
	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)

	file, err := c.FormFile("file")
	utils.PanicIfNeeded(err)

	request.File = file

	//add validation send image
	validations.ValidateSendFile(request)

	if request.Type == structs.TypeGroup {
		request.Phone = request.Phone + "@g.us"
	} else {
		request.Phone = request.Phone + "@s.whatsapp.net"
	}

	response, err := controller.Service.SendFile(c, request)
	utils.PanicIfNeeded(err)

	return c.JSON(utils.ResponseData{
		Code:    200,
		Message: response.Status,
		Results: response,
	})
}

func (controller *SendController) SendVideo(c *fiber.Ctx) error {
	var request structs.SendVideoRequest
	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)

	video, err := c.FormFile("video")
	utils.PanicIfNeeded(err)

	request.Video = video

	//add validation send image
	validations.ValidateSendVideo(request)

	if request.Type == structs.TypeGroup {
		request.Phone = request.Phone + "@g.us"
	} else {
		request.Phone = request.Phone + "@s.whatsapp.net"
	}

	response, err := controller.Service.SendVideo(c, request)
	utils.PanicIfNeeded(err)

	return c.JSON(utils.ResponseData{
		Code:    200,
		Message: response.Status,
		Results: response,
	})
}
