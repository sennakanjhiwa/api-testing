package controllers

import (
	"github.com/sev-2/raiden"
)

type BookRequest struct {
}

type BookResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BooksController struct {
	raiden.ControllerBase
	Http    string `path:"/books" type:"custom"`
	Payload *BookRequest
	Result  []BookResponse
}

func (c *BooksController) Get(ctx raiden.Context) error {
	c.Result = []BookResponse{
		{ID: 1, Name: "Book 1", Description: "Desc Book 1"},
		{ID: 2, Name: "Book 2", Description: "Desc Book 2"},
		{ID: 3, Name: "Book 3", Description: "Desc Book 3"},
		{ID: 3, Name: "Book 4", Description: "Desc Book 4"},
		{ID: 3, Name: "Book 5", Description: "Desc Book 5"},
		{ID: 3, Name: "Book 6", Description: "Desc Book 6"},
		{ID: 3, Name: "Book 7", Description: "Desc Book 7"},
	}
	return ctx.SendJson(c.Result)
}
