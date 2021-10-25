package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/SHoar/master_go_programming_course_notes/scratchpad/goProductApi/entity"
)

func GetProductHandler() http.HandlerFunc {
	/* GetProductHandler is used to get data inside the products defined on our product catalog */
	return func(rw http.ResponseWriter, r *http.Request) {
		// read JSON File
		data, err := ioutil.ReadFile("./mockData/data.json")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Write the body with Json data
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write(data)
	}
}

func CreateProductHandler() http.HandlerFunc {
	/* CreateProductHandler is used to create a new product and add to our product store */
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read incoming JSON from req body
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// if no body is associted, return with StatusBadRequest
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// check if data is proper JSON (data validation)
		var product entity.Product
		err = json.Unmarshal(data, &product)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid Data Format"))
			return
		}

		// Load existing products and append the data to product list

		var products []entity.Product
		data, err = ioutil.ReadFile("./mockData/data.json")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Add new product to our list
		products = append(products, product)

		// Write updated JSON file
		updatedData, err := json.Marshal(products)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = ioutil.WriteFile("./data/data.json", updatedData, os.ModePerm)

		// return after writing Body
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Added New Product"))
	}
}
