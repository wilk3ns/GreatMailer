package api

import (
	"GreatMailer/models"
	"GreatMailer/verification"
	"encoding/json"
	"io"
	"net/http"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Handle preflight request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		// Decode the request body
		var emailReq models.EmailRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&emailReq); err != nil {
			response := models.Response{
				Message: "Invalid request body",
				Status:  false,
			}
			mResp, _ := json.Marshal(response)
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write(mResp)
			if err != nil {
				return
			}
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(r.Body)

		// Validate required fields
		if emailReq.Email == "" || emailReq.Message == "" {
			response := models.Response{
				Message: "Email and message are required fields",
				Status:  false,
			}
			mResp, _ := json.Marshal(response)
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write(mResp)
			if err != nil {
				return
			}
			return
		}

		// Send email with the provided data
		res, err := verification.SendEmail(emailReq.Name, emailReq.Email, emailReq.Message)

		var response models.Response
		if err == nil {
			response = models.Response{
				Message: res,
				Status:  true,
			}
			mResp, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(mResp)
			if err != nil {
				return
			}
		} else {
			response = models.Response{
				Message: res,
				Status:  false,
			}
			mResp, _ := json.Marshal(response)
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write(mResp)
			if err != nil {
				return
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
