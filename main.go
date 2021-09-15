package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	//"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book Struct (Model)
type Book struct {
	ID 		string `json:"id"`
	Isbn 	string `json:"isbn"`
	Title 	string `json:"title"`
	Author 	*Author `json:"id"`
}

// Init books var as a slice Book struct
 var books []Book

//****************************************************************************************************
// Get All Books 																					 *
//**************************************************************************************************>>
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//****************************************************************************************************
// Get single Book																					 *
//**************************************************************************************************>>
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//****************************************************************************************************
// Create a new Book 																				 *
//**************************************************************************************************>>
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// Create an ID
	// We use the stringConverter to convert the Int id
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe - just for example
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
//****************************************************************************************************
// Update a Book 																					 *
//**************************************************************************************************>>
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		// check for the id
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			// Create an ID
			// We use the stringConverter to convert the Int id
			book.ID =  params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//****************************************************************************************************
// Delete a Book 																					 *
//**************************************************************************************************>>
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

func main() {
	//fmt.Println("Hello World")

	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	books = append(books, Book{ID: "2", Isbn: "448744", Title: "Book Two", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "3", Isbn: "448745", Title: "Book Three", Author: &Author{Firstname: "Casa", Lastname: "Blanca"}})
	books = append(books, Book{ID: "4", Isbn: "448746", Title: "Book Four", Author: &Author{Firstname: "Za", Lastname: "Gora"}})
	books = append(books, Book{ID: "5", Isbn: "448747", Title: "Book Five", Author: &Author{Firstname: "Ouar", Lastname: "Zazate"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}