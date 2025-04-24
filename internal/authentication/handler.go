package authentication

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/victorsvart/go-ecommerce/internal/core/domain"
	"github.com/victorsvart/go-ecommerce/pkg/appcontext"
	"github.com/victorsvart/go-ecommerce/pkg/middleware"
	"github.com/victorsvart/go-ecommerce/pkg/token"
	"github.com/victorsvart/go-ecommerce/pkg/utils"
)

const (
	tokenName = "auth_token"
)

var (
	ErrWrongCredentials = errors.New("wrong credentials")
	ErrAuthenticating   = errors.New("internal authentication error")
)

type AuthHandler struct {
	usecases domain.UserUseCases
}

func NewAuthHandler(api chi.Router, usecases domain.UserUseCases) {
	handler := &AuthHandler{usecases}
	api.Route("/auth", func(r chi.Router) {
		r.With(middleware.Auth).Get("/me", handler.Me)
		r.Post("/login", handler.Login)
		r.Post("/register", handler.Register)
		r.Post("/logout", handler.Logout)
	})
}

func (a *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	// Log incoming request headers
	log.Println("Request Headers:", r.Header)

	// Log all cookies
	log.Println("Cookies:", r.Cookies())

	// If you're looking for a specific cookie (like auth_token)
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		log.Println("Error retrieving auth_token cookie:", err)
	} else {
		log.Println("auth_token cookie:", cookie.Value)
	}
	authCtx, err := appcontext.GetAuthContext(r.Context())
	if err != nil {
		utils.RespondJSON(w, http.StatusUnauthorized, false, err.Error())
		return
	}

	user, err := a.usecases.GetById(r.Context(), authCtx.UserID)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, false, err.Error())
		return
	}

	if user == nil {
		utils.RespondJSON(w, http.StatusUnauthorized, false, "Invalid user or not logged in")
		return
	}

	utils.RespondJSON(w, http.StatusOK, true, "ok")
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	}
	user, err := a.usecases.GetByEmail(r.Context(), input.Email)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	if err := utils.ComparePassword(user.Password, input.Password); err != nil {
		utils.RespondJSON(w, http.StatusUnauthorized, false, ErrWrongCredentials.Error())
		return
	}

	jwt, err := token.GenerateJWT(user.ID, user.RoleID, user.Email)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	shouldSecure, err := strconv.ParseBool(os.Getenv("SECURE_TOKEN"))
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    jwt,
		HttpOnly: true,
		Secure:   shouldSecure,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
	})

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
		Name:     tokenName,
		Value:    "",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, c)
}
