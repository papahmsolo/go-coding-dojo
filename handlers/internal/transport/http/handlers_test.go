package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mock_http "github.com/papahmsolo/go-coding-dojo/handlers/internal/mock"
	"github.com/papahmsolo/go-coding-dojo/handlers/internal/service"
	transport_http "github.com/papahmsolo/go-coding-dojo/handlers/internal/transport/http"
)

func TestSignupHandler(t *testing.T) {
	const (
		someName     = "Name"
		someLogin    = "Login"
		somePassword = "Password"
	)

	var testError = errors.New("some error")

	cases := []struct {
		name                string
		request             transport_http.SignupRequest
		expectedResposeBody transport_http.ErrorResponse
		expectedStatusCode  int
		serviceError        error
		isMockCalled        bool
	}{
		{
			name: "should return 400 when name is empty",
			request: transport_http.SignupRequest{
				Name:     "",
				Login:    someLogin,
				Password: somePassword,
			},
			expectedResposeBody: transport_http.ErrorResponse{"invalid parameter: name"},
			expectedStatusCode:  http.StatusBadRequest,
		},
		{
			name: "should return 400 when login is empty",
			request: transport_http.SignupRequest{
				Name:     someName,
				Login:    "",
				Password: somePassword,
			},
			expectedResposeBody: transport_http.ErrorResponse{"invalid parameter: login"},
			expectedStatusCode:  http.StatusBadRequest,
		},
		{
			name: "should return 400 when password is empty",
			request: transport_http.SignupRequest{
				Name:     someName,
				Login:    someLogin,
				Password: "",
			},
			expectedResposeBody: transport_http.ErrorResponse{"invalid parameter: password"},
			expectedStatusCode:  http.StatusBadRequest,
		},
		// {
		// 	name: "should return 200 when request is valid and successful",
		// 	request: transport_http.SignupRequest{
		// 		Name:     someName,
		// 		Login:    someLogin,
		// 		Password: somePassword,
		// 	},
		// 	expectedStatusCode: http.StatusOK,
		// 	isMockCalled:       true,
		// 	serviceError:       nil,
		// },
		{
			name: "should return 409 when user already exists",
			request: transport_http.SignupRequest{
				Name:     someName,
				Login:    someLogin,
				Password: somePassword,
			},
			expectedResposeBody: transport_http.ErrorResponse{service.ErrUserAlreadyExists.Error()},
			expectedStatusCode:  http.StatusConflict,
			isMockCalled:        true,
			serviceError:        service.ErrUserAlreadyExists,
		},
		{
			name: "should return 500 when service returns an error",
			request: transport_http.SignupRequest{
				Name:     someName,
				Login:    someLogin,
				Password: somePassword,
			},
			expectedResposeBody: transport_http.ErrorResponse{testError.Error()},
			expectedStatusCode:  http.StatusInternalServerError,
			isMockCalled:        true,
			serviceError:        testError,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockUserService := mock_http.NewMockUserService(ctrl)
			s := transport_http.NewServer(mockUserService)

			reqBody, err := json.Marshal(tt.request)
			if err != nil {
				t.Errorf("cannot marshal request body: %v", err)
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(reqBody))

			if tt.isMockCalled {
				mockUserService.EXPECT().
					CreateUser(someName, someLogin, somePassword).
					Return(tt.serviceError)
			}

			s.HandleSignup(w, r)

			code := w.Result().StatusCode
			if code != tt.expectedStatusCode {
				t.Errorf("expected status code: %d, but got: %d", tt.expectedStatusCode, code)
			}

			var resBody transport_http.ErrorResponse
			err = json.NewDecoder(w.Result().Body).Decode(&resBody)
			if err != nil {
				t.Errorf("cannot decode response body: %v", err)
			}
			defer w.Result().Body.Close()

			if tt.expectedResposeBody != resBody {
				t.Errorf("expected response body: %+v, but got: %+v", tt.expectedResposeBody, resBody)
			}
		})
	}
}
