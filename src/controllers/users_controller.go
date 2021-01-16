package controllers

var (
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
	CreateUser()
}

type usersController struct{}

func (c *usersController) CreateUser() {

}
