package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/papahmsolo/go-coding-dojo/handlers/internal/service"
)

var ErrInvalidParameter = errors.New("invalid parameter")

type SignupRequest struct {
	Name     string `json:"user_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignupResponse struct {
	ID string `json:"user_id"`
}

type ErrorResponse struct {
	Message string `json:"msg"`
}

type Server struct {
	userService service.UserService
}

func NewServer(us service.UserService) *Server {
	return &Server{userService: us}
}

func (s SignupRequest) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("%v: %s", ErrInvalidParameter, "name")
	}
	if s.Login == "" {
		return fmt.Errorf("%v: %s", ErrInvalidParameter, "login")
	}
	if s.Password == "" {
		return fmt.Errorf("%v: %s", ErrInvalidParameter, "password")
	}
	return nil
}

func (s *Server) HandleSignup(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	err = req.Validate()
	if err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	err = s.userService.CreateUser(req.Name, req.Login, req.Password)
	if errors.Is(err, service.ErrUserAlreadyExists) {
		writeErr(w, http.StatusConflict, err)
		return
	}
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func writeErr(w http.ResponseWriter, statusCode int, errMsg error) {
	w.WriteHeader(statusCode)
	errRes := ErrorResponse{Message: errMsg.Error()}
	err := json.NewEncoder(w).Encode(&errRes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
