package auth

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"splitmoney/pkg/domain/user"
)

type App struct {
	app         *fiber.App
	authService user.UserManager
}

func NewApp(authService user.UserManager) (*App, error) {
	app := &App{
		app:         fiber.New(),
		authService: authService,
	}
	app.SetupRoutes()
	return app, nil
}

func (a *App) SetupRoutes() {
	a.app.Post("/signup", a.SignUp)
	a.app.Post("/signout", a.SignOut)
}

func (a *App) Start() {
	log.Fatal(a.app.Listen(":3000"))
}

func (a *App) SignUp(c *fiber.Ctx) error {
	data := c.Body()
	var req SignupRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return user.ErrorFailedToSignup
	}
	accId, err := a.authService.SignUp(c.Context(), user.User{
		Email: req.Email,
		Phone: req.Phone,
		Profile: user.Profile{
			Name: req.Name,
		},
	})
	if err != nil && err == user.ErrorEmailAlreadyInUse {
		c.Status(http.StatusBadRequest)
		return user.ErrorFailedToSignup
	}
	resp := SingupResponse{
		AccountID: accId,
	}
	respData, err := json.Marshal(&resp)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return user.ErrorFailedToSignup
	}
	c.Write(respData)
	slog.Info("User singup successful")
	c.Status(http.StatusOK)
	return nil
}

func (a *App) SignOut(c *fiber.Ctx) error {
	return nil
}
