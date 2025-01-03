package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"renebizelli/api/internal/dtos"
	"renebizelli/api/internal/entities"
	"renebizelli/api/internal/infra/database"
	"time"

	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB     database.UserInterface
	JWT        *jwtauth.JWTAuth
	JWTExpires int64
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, expiresIn int64) *UserHandler {
	return &UserHandler{
		UserDB:     db,
		JWT:        jwt,
		JWTExpires: expiresIn,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param request body dtos.CreateUserInput true "User request"
// @Success 200 {object} entities.User
// @Router /users [post]
// @Failure 500 {object} Error
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var dto dtos.CreateUserInput

	e := json.NewDecoder(r.Body).Decode(&dto)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(dto)
	fmt.Println(dto)

	user, e := entities.NewUser(dto.Name, dto.Email, dto.Password)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e = u.UserDB.Create(user)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: e.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Login godoc
// @Summary User login
// @Description User login
// @Tags login
// @Accept  json
// @Produce  json
// @Param request body dtos.LoginInput true "Login request"
// @Success 200 {object} string
// @Router /login [post]
// @Failure 401
// @Failure 404
// @Failure 400
func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	var dto dtos.LoginInput

	e := json.NewDecoder(r.Body).Decode(&dto)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, e := u.UserDB.FindByEmail(dto.Email)
	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !user.ValidatePassword(dto.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := map[string]interface{}{
		"sub": user.Id,
		"exp": time.Now().Add(time.Second * time.Duration(u.JWTExpires)).Unix(),
	}

	_, stringToken, ejwt := u.JWT.Encode(claims)

	if ejwt != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	json.NewEncoder(w).Encode(stringToken)
}
