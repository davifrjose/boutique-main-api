package http

import (
	"github.com/davifrjose/boutique-main-api/internal/core/domain"
	"github.com/davifrjose/boutique-main-api/internal/core/port"
	"github.com/gin-gonic/gin"
)

// UserHandler represents the HTTP handler for user-related requests
type UserHandler struct {
	svc port.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

// registerRequest represents the request body for creating a user
type registerRequest struct {
	FirstName     string `json:"first_name" binding:"required" example:"John"`
	LastName     string `json:"last_name" binding:"required" example:"Doe"`
	Email    string `json:"email" binding:"required,email" example:"test@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"12345678"`
	PhoneNumber string `json:"phone_number" binding:"required,min=8" example:"12345678"`
	Address 	string `json:"address"  example:"Rua Dom Afonso Henriques 655 2D"`
	City      string `json:"city"  example:"Luanda"`
	Province  string `json:"province"  example:"Luanda"`
	Zip 			string	`json:"zip"  example:"433-007"`
	Country   string	`json:"country" example:"Portugal"`
}

// Register godoc
//
//	@Summary		Register a new user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		registerRequest	true	"Register request"
//	@Success		200				{object}	userResponse	"User created"
//	@Failure		400				{object}	errorResponse	"Validation error"
//	@Failure		401				{object}	errorResponse	"Unauthorized error"
//	@Failure		404				{object}	errorResponse	"Data not found error"
//	@Failure		409				{object}	errorResponse	"Data conflict error"
//	@Failure		500				{object}	errorResponse	"Internal server error"
//	@Router			/users [post]

func (uh *UserHandler) Register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	userParams := domain.CreateUserParams{
		FirstName:     req.FirstName,
		LastName: req.LastName,
		Email:    req.Email,
		Password: req.Password,
		PhoneNumber: req.PhoneNumber,
		Address: req.Address,
		City: req.City,
		Province: req.Province,
		Zip: req.Zip,
		Country: req.Country,
	}

	user, err := uh.svc.Register(ctx, &userParams)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newUserResponse(user)

	handleSuccess(ctx, rsp)
}
