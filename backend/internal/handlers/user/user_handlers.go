package handlers

import (
	"net/http"

	"github.com/shaikhjunaidx/pennywise-backend/internal/handlers"
	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
)

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// SignUpHandler handles user registration requests.
// @Summary User Registration
// @Description Registers a new user with the given username, email, and password.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   username  body  string  true  "Username"
// @Param   email     body  string  true  "Email"
// @Param   password  body  string  true  "Password"
// @Success 201 {object} UserResponse "Created User"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/signup [post]
func SignUpHandler(s *user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := handlers.ParseJSONRequest(w, r, &req); err != nil {
			return
		}

		user, err := s.SignUp(req.Username, req.Email, req.Password)
		if err != nil {
			handlers.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		handlers.SendJSONResponse(w, user, http.StatusCreated)
	}
}

// LoginHandler handles user login requests.
// @Summary User Login
// @Description Authenticates a user and returns a JWT token.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   username  body  string  true  "Username"
// @Param   password  body  string  true  "Password"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /api/login [post]
func LoginHandler(s *user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := handlers.ParseJSONRequest(w, r, &req); err != nil {
			return
		}

		token, err := s.Login(req.Username, req.Password)
		if err != nil {
			handlers.SendErrorResponse(w, err.Error(), http.StatusUnauthorized)
			return
		}

		handlers.SendJSONResponse(w, map[string]string{"token": token}, http.StatusOK)
	}
}
