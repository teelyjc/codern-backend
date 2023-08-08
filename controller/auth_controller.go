package controller

import (
	"github.com/codern-org/codern/domain"
	"github.com/codern-org/codern/internal/payload"
	"github.com/codern-org/codern/internal/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthController struct {
	logger    *zap.Logger
	validator domain.PayloadValidator

	authUsecase   domain.AuthUsecase
	googleUsecase domain.GoogleUsecase
	userUsecase   domain.UserUsecase
}

func NewAuthContoller(
	logger *zap.Logger,
	validator domain.PayloadValidator,
	authUsecase domain.AuthUsecase,
	googleUsecase domain.GoogleUsecase,
	userUsecase domain.UserUsecase,
) *AuthController {
	return &AuthController{
		logger:        logger,
		validator:     validator,
		authUsecase:   authUsecase,
		googleUsecase: googleUsecase,
		userUsecase:   userUsecase,
	}
}

// Me godoc
//
// @Summary 		Get an user data
// @Description	Get an authenticated user data
// @Tags 				auth
// @Accept 			json
// @Produce 		json
// @Security 		ApiKeyAuth
// @param 			sid header string true "Session ID"
// @Success 		200	{object}	domain.User
// @Failure			400 {object}	response.GenericErrorResponse "If `sid` header is missing"
// @Failure			401 {object}	response.GenericErrorResponse "If something wrong on authentication"
// @Router 			/api/auth/me [get]
func (c *AuthController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(domain.User)
	return ctx.JSON(user)
}

// SignIn godoc
//
// @Summary 		Sign in with self provider
// @Description Sign in with email & password provided by the user
// @Tags 				auth
// @Accept 			json
// @Produce 		json
// @Param				credentials	body	payload.AuthSignIn true "Email and password for authentication"
// @Success			200
// @Router 			/api/auth/signin [post]
func (c *AuthController) SignIn(ctx *fiber.Ctx) error {
	var payload payload.AuthSignIn
	if ok, err := c.validator.ValidateBody(&payload, ctx); !ok {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "signin",
	})
}

// SignOut godoc
//
// @Summary 		Sign out
// @Description Sign out and remove a sid cookie header
// @Tags 				auth
// @Produce 		json
// @Security 		ApiKeyAuths
// @param 			sid header string true "Session ID"
// @Success			200
// @Router 			/api/auth/signout [post]
func (c *AuthController) SignOut(ctx *fiber.Ctx) error {
	sid := ctx.Get("sid")

	if err := c.authUsecase.SignOut(sid); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.GenericErrorResponse{
			Code:    response.ErrUnauthorized,
			Message: err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "signout",
	})
}

func (c *AuthController) GetGoogleAuthUrl(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"url": c.googleUsecase.GetOAuthUrl(),
	})
}

func (c *AuthController) SignInWithGoogle(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "google",
	})
}
