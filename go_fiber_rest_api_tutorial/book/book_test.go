package book_test

import (
	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/suite"
)

type BookTestSuit struct {
	suite.Suite
	app *fiber.App
}



func (book *BookTestSuit) SetupSuit() {
	suite.app = mai

	
}
