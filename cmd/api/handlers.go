package main

import (
	"errors"
	"net/http"
)

func (app *Config) Login(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	//Validate the user again the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.writeError(w, errors.New("Invalid Credentials"), http.StatusBadRequest)
		return
	}

	//Check if the password is correct
	if user.Password != requestPayload.Password {
		app.writeError(w, errors.New("Invalid Credentials"), http.StatusBadRequest)
		return
	}

	payload := JsonResponse{
		Error:   false,
		Message: "Login Successful",
		Data:    user,
	}

	app.writeJSON(w, http.StatusOK, payload)

}
