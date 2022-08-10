package router

import (
	"api-address-golang/entities_core"
	"api-address-golang/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CreateAddress(response http.ResponseWriter, request *http.Request) {

	var address entities_core.Address

	err := json.NewDecoder(request.Body).Decode(&address)
	if err != nil {
		log.Printf("An error's happened in decoding json %v", err)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := service.InsertIntoDB(address)

	var result map[string]any
	if err != nil {
		result = map[string]any{
			"Status":   true,
			"Message": fmt.Sprintf("An error has occurred: %v", err),
		}
	} else {
		result = map[string]any{
			"Status":   http.StatusOK,
			"Message": fmt.Sprintf("Insert has been successful: %v", id),
		}
	}

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(result)
}

func GetAddressById(response http.ResponseWriter, request *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Error in parsing id %v", err)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	address, erro := service.GetAddressById(int64(id))
	if erro != nil {
		log.Printf("An error has occured %v", erro)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(address)

}

func GetAllAddress(response http.ResponseWriter, request *http.Request) {
	
	address, err := service.GetAllAddress()
	if err != nil {
		log.Printf("An error has occured %v", err)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(address)

}

func UpdateAddress(response http.ResponseWriter, request *http.Request) {

	var address entities_core.Address

	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Error in parsing id %v", err)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&address)
	if err != nil {
		log.Printf("Error in decoding json %v", err)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, erro := service.UpdateAddress(int64(id), address)
	if erro != nil {
		log.Printf("Error in parsing id %v", erro)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result := map[string]any{
		"Status":   http.StatusOK,
		"Message": fmt.Sprintf("Updated has been succesful!"),
	}

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(result)

}

func DeleteAddress(response http.ResponseWriter, request *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("An error has occured %v", err)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, erro := service.DeleteAddress(int64(id))

	if erro != nil {
		log.Printf("Error in removing address by id %v", erro)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result := map[string]any{
		"Status":   http.StatusOK,
		"Message": fmt.Sprintf("Delete has been succesful"),
	}

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(result)

}
