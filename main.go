package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type HomePageResponse struct {
	Description string `json:"description"`
	Path        string `json:"path"`
	Method      string `json:"method"`
}

var customerList []Customer

var homePageResponse []HomePageResponse

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(homePageResponse)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(customerList) == 0 {
		w.WriteHeader(404)
	}
	json.NewEncoder(w).Encode(customerList)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(customerList) == 0 {
		w.WriteHeader(404)
	}
	params := mux.Vars(r)
	for index, customer := range customerList {
		if customer.ID == params["id"] {
			customerList = append(customerList[:index], customerList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(customerList)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(customerList) == 0 {
		w.WriteHeader(404)
	}
	params := mux.Vars(r)
	for _, customer := range customerList {
		if customer.ID == params["id"] {
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.ContentLength == 0 {
		w.WriteHeader(404)
		fmt.Print("empty")
	}
	if r.ContentLength != 0 {
		var customer Customer
		json.NewDecoder(r.Body).Decode((&customer))
		customer.ID = strconv.Itoa((rand.Intn(10000)))
		customerList = append(customerList, customer)
		json.NewEncoder(w).Encode(customer)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	if len(customerList) == 0 {
		w.WriteHeader(404)
	}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, customer := range customerList {
		if customer.ID == params["id"] {
			customerList = append(customerList[:index], customerList[index+1:]...)
			var customer Customer
			json.NewDecoder(r.Body).Decode(&customer)
			customer.ID = params["id"]
			customerList = append(customerList, customer)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()
	customerList = append(customerList, Customer{ID: "1", Name: "Ben", Role: "Customer", Email: "ben.tran@jungtalens.com", Phone: "0123456789", Contacted: true})
	customerList = append(customerList, Customer{ID: "2", Name: "Lock", Role: "Seller", Email: "lock.huynh@jungtalens.com", Phone: "0987654321", Contacted: false})
	customerList = append(customerList, Customer{ID: "3", Name: "James", Role: "Manager", Email: "james.vo@jungtalens.com", Phone: "1234509876", Contacted: false})
	customerList = append(customerList, Customer{ID: "4", Name: "Han", Role: "Customer", Email: "han.nguyen@jungtalens.com", Phone: "0192837465", Contacted: true})

	homePageResponse = append(homePageResponse, HomePageResponse{Description: "Getting a single customer", Path: "/customers/{id}", Method: "GET"})
	homePageResponse = append(homePageResponse, HomePageResponse{Description: "Getting all customers", Path: "/customers", Method: "GET"})
	homePageResponse = append(homePageResponse, HomePageResponse{Description: "Creating a customer", Path: "/customers", Method: "POST"})
	homePageResponse = append(homePageResponse, HomePageResponse{Description: " Updating a customer", Path: "/customers/{id}", Method: "PUT"})
	homePageResponse = append(homePageResponse, HomePageResponse{Description: " Deleting a customer", Path: "/customers/{id}", Method: "DELETE"})

	r.HandleFunc("/", homePage)
	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/customers", addCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
