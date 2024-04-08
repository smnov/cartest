package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/smnov/cartest/docs"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

// @Summary      GetCarsHandler
// @Description  Get a list of cars with pagination support
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        page query int true "Page number"
// @Param        page_size query int true "Number of items per page"
// @Param make query string false "Car make"
// @Param model query string false "Car model"
// @Param year query int false "Car year"
// @Success      200 {array} Car "Successful response with an array of cars"
// @Failure      400 {object} APIError "Bad request"
// @Failure      404 {object} APIError "Resource not found"
// @Failure      500 {object} APIError "Internal server error"
// @Router       /cars/get [get]
func (s *Server) GetCarsHandler(w http.ResponseWriter, r *http.Request) error {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return err
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return err
	}

	make := r.URL.Query().Get("make")
	model := r.URL.Query().Get("model")
	yearStr := r.URL.Query().Get("year")
	var year int
	if yearStr != "" {
		year, err = strconv.Atoi(yearStr)
		if err != nil {
			return err
		}
	}

	s.logger.Info("Handling GetCars request")

	cars, err := s.db.GetCars(page, pageSize, make, model, year)
	if err != nil {
		s.logger.Debug("error while getting cars", "error", err.Error())
		return err
	}

	return WriteJSON(w, 200, cars)
}

// @Summary      DeleteCarHandler
// @Description  Delete a car by ID
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id path int true "Car ID"
// @Success      200 {integer} integer "ID of the deleted car"
// @Failure      400 {object} APIError "Bad request"
// @Failure      404 {object} APIError "Resource not found"
// @Failure      500 {object} APIError "Internal server error"
// @Router       /cars/delete/{id} [delete]
func (s *Server) DeleteCarHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return err
	}
	s.logger.Debug(fmt.Sprintf("Handling DeleteCar request for ID: %v", id))
	err = s.db.DeleteCarByID(id)
	if err != nil {
		s.logger.Debug("car deletion error", "error", err.Error())
		return err
	}

	return WriteJSON(w, 200, id)
}

// @Summary      UpdateCarHandler
// @Description  Update a car by ID
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id path int true "Car ID"
// @Success      200 {integer} integer "ID of the updated car"
// @Failure      400 {object} APIError "Bad request"
// @Failure      404 {object} APIError "Resource not found"
// @Failure      500 {object} APIError "Internal server error"
// @Router       /cars/update/{id} [put]
func (s *Server) UpdateCarHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return err
	}
	s.logger.Debug(fmt.Sprintf("Handling UpdateCar request for ID: %v", id))
	updatedCar := Car{}
	err = s.db.UpdateCarByID(id, &updatedCar)
	if err != nil {
		s.logger.Debug("update car error", "error", err.Error())
		return err
	}
	return WriteJSON(w, 200, id)
}

// @Summary      AddCarHandler
// @Description  Add one or more cars
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        regNums body string true "Registration numbers of cars (comma-separated)"
// @Success      201 {array} Car "Successful response with an array of added cars"
// @Failure      400 {object} APIError "Bad request"
// @Failure      500 {object} APIError "Internal server error"
// @Router       /cars/add [post]
func (s *Server) AddCarHandler(w http.ResponseWriter, r *http.Request) error {
	var cars []*Car
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var requestData struct {
		RegNums []string `json:"regNums"`
	}
	if err := json.Unmarshal(body, &requestData); err != nil {
		return err
	}

	s.logger.Info("Handling AddCar request")
	if err := s.db.AddCars(requestData.RegNums); err != nil {
		return err
	}

	for _, regNum := range requestData.RegNums {
		carInfo, err := getCarInfo(regNum)
		if err != nil {
			s.logger.Debug("Error getting car info", "error", err.Error())
			return err
		}
		s.logger.Info("Received car info from external API", "car info", carInfo)
		cars = append(cars, carInfo)
	}

	return WriteJSON(w, http.StatusCreated, cars)
}

// External API call
func getCarInfo(regNum string) (*Car, error) {
	apiUrl := "http://external-api.com/info?regNum=" + regNum

	response, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var car Car
	if err := json.Unmarshal(body, &car); err != nil {
		return nil, err
	}

	return &car, nil
}
