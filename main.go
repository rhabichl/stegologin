package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"

	"github.com/auyer/steganography"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/profile/change_profile_pict", func(c *fiber.Ctx) error {
		// get the file from the body
		file, err := c.FormFile("profile_picture")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("couldn't get file from form")
		}

		// open the file
		f, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("couldn't open the file")
		}
		defer f.Close()

		var image image.Image
		// get the filetype and decode accordingly, only jpeg and png are supported
		switch file.Header.Get("Content-Type") {
		case "image/jpeg":
			image, err = jpeg.Decode(f)
		case "image/png":
			image, err = png.Decode(f)
		default:
			return c.Status(fiber.StatusBadRequest).SendString("file wasn't jpeg or png")
		}
		// handle the error for the file
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("invalid file: " + err.Error())
		}
		// get the message size
		msgSize := steganography.GetMessageSizeFromImage(image)
		// if there isn't a msg in the picture the size will be the max size for a unit32
		if msgSize == 4294967264 {
			// file is normal there, this isn't a login request
			return returnDefault(c)
		}
		// get the hidden credentials
		msg := steganography.Decode(msgSize, image)
		// check if the login is correct and if not just pretend like nothing happend
		if checkLogin(string(msg)) {
			return c.SendString("you are in :)")
		}
		// if the login isn't valid mimik default behavior
		return returnDefault(c)
	})

	log.Fatal(app.Listen(":3000"))
}

// a function to check if the encoded msg from the picture is valid
func checkLogin(msg string) bool {
	// TODO: Implement me
	return true
}

// a function to mimik the default behavior of the app
func returnDefault(c *fiber.Ctx) error {
	// TODO: Implement me
	return c.SendString("nice picture")
}
