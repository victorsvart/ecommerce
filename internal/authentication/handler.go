package authentication

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/victorsvart/go-ecommerce/internal/user/domain"
	"github.com/victorsvart/go-ecommerce/pkg/token"
	"github.com/victorsvart/go-ecommerce/pkg/utils"
)

type AuthHandler struct {
	usecases domain.UserUseCases
}

var (
	ErrWrongCredentials = errors.New("wrong credentials")
	ErrAuthenticating   = errors.New("internal authentication error")
)

func NewAuthHandler(api chi.Router, usecases domain.UserUseCases) {
	handler := &AuthHandler{usecases}
	api.Route("/auth", func(r chi.Router) {
		r.Post("/login", handler.Login)
		r.Post("/register", handler.Register)
		r.Post("/logout", handler.Logout)
	})
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	}
	user, err := a.usecases.GetByEmail(r.Context(), input.Email)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, ErrWrongCredentials)
		return
	}

	if err := utils.ComparePassword(user.Password, input.Password); err != nil {
		utils.RespondJSON(w, http.StatusUnauthorized, false, ErrWrongCredentials.Error())
		return
	}

	jwt, err := token.GenerateJWT(user.ID, user.Email)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	shouldSecure, err := strconv.ParseBool(os.Getenv("SECURE_TOKEN"))
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	cookie := http.Cookie{
		Name:     "auth_token",
		Value:    jwt,
		HttpOnly: true,
		Secure:   shouldSecure,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, &cookie)
	utils.RespondJSON(w, http.StatusOK, true, "Logged in!")
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	}

	user := input.ToUser()
	if err := a.usecases.Create(r.Context(), &user); err != nil {
		utils.RespondJSON(w, http.StatusConflict, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, true, "Registered")
}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	}

	http.SetCookie(w, c)
}
