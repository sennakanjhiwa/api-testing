package controllers

import (
	"github.com/sev-2/raiden"
)

type UserRequest struct {
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"Age"`
}

type UserController struct {
	raiden.ControllerBase
	Http    string `path:"/users" type:"custom"`
	Payload *UserRequest
	Result  []UserResponse
}

func (c *UserController) Get(ctx raiden.Context) error {
	c.Result = []UserResponse{
		{ID: 1, Name: "Jhon", Age: 35},
	}
	return ctx.SendJson(c.Result)
}
