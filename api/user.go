package api

import (
	"middleearth/eateries/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// requests
type registerUserRequest struct {
	Name         string `json:"name" binding:"required"`
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
}

type loginUserRequest struct {
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
}

type userResponse struct {
	ID           string `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	EmailAddress string `json:"emailAddress" binding:"required,email"`
}

type loggedInUserResponse struct {
	ID           string `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Token        string `json:"token" binding:"required"`
}

type UserAPI struct {
	Service *service.UserService
}

func NewUserAPI(Service *service.UserService) *UserAPI {
	return &UserAPI{Service: Service}
}

// RegisterUser godoc
// @Summary      Register a user
// @Description  register a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param		 register body    api.registerUserRequest   true  "Register user"
// @Success      200  {array}   api.userResponse
// @Router       /users/register [post]
func (userAPI *UserAPI) RegisterUser(ginContext *gin.Context) {
	var request registerUserRequest
	if err := ginContext.BindJSON(&request); err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	// map to user action
	action := service.RegisterUserAction{
		Name:         request.Name,
		EmailAddress: request.EmailAddress,
		Password:     request.Password,
	}

	user, err := userAPI.Service.RegisterUser(action)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	// map to user response
	response := userResponse{
		ID:           user.ID.String(),
		Name:         user.Name,
		EmailAddress: user.EmailAddress,
	}

	ginContext.IndentedJSON(http.StatusOK, response)
}

// RegisterUser godoc
// @Summary      Login user
// @Description  Login user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param		 register body    api.loginUserRequest   true  "Login user"
// @Success      200  {array}   api.loggedInUserResponse
// @Router       /users/login [post]
func (userAPI *UserAPI) LoginUser(ginContext *gin.Context) {
	var request loginUserRequest
	if err := ginContext.BindJSON(&request); err != nil {
		return
	}

	// map to login user action
	action := service.LoginUserAction{
		EmailAddress: request.EmailAddress,
		Password:     request.Password,
	}

	loggedInUser, err := userAPI.Service.LoginUser(action)
	if err != nil {
		ErrorResponse(err, ginContext)
		return
	}

	// map to login user response
	response := loggedInUserResponse{
		ID:           loggedInUser.ID.String(),
		Name:         loggedInUser.Name,
		EmailAddress: loggedInUser.EmailAddress,
		Token:        loggedInUser.Token,
	}

	ginContext.IndentedJSON(http.StatusOK, response)
}
