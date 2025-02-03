package view

import (
	"encoding/json"
	"net/http"

	"github.com/1ef7yy/go-kafka-poc/internal/models"
)

func (v *view) ValidatePhone(w http.ResponseWriter, r *http.Request) {
	phone_number := r.URL.Query().Get("phone_number")
	if phone_number == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Phone number is required"))
		return
	}

	valid, err := v.domain.ValidatePhone(phone_number)
	if err != nil {
		v.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if !valid {
		w.WriteHeader(http.StatusOK)
		response := models.InvalidPhoneResponse{
			Status: false,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return

	}

	normalized, err := v.domain.NormalizePhone(phone_number)

	if err != nil {
		v.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	response := models.ValidPhoneResponse{
		Status:     true,
		Normalized: normalized,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
