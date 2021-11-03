package handler

import "github.com/poltro90/tech-panel/controller"

type UserHandler struct {
	userController controller.IUserController
}

func NewUserHandler(userController controller.IUserController) UserHandler {
	return UserHandler{
		userController: userController,
	}
}
