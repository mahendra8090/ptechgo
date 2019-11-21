package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type profiles struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

func getprofiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	//	json.NewEncoder(w).Encode(books)
}

func addprofile(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "root:securepassword@tcp(localhost:3306)/userlogin?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Print(err.Error)

	}
	if err == nil {
		fmt.Print("open")

	}

	defer db.Close()
	w.Header().Set("Content-type", "application/json")
	var P profiles

	_ = json.NewDecoder(r.Body).Decode(&P)
	json.NewEncoder(w).Encode(P)
	//fmt.Print(User.Username)

	//db.AutoMigrate(&Book{})
	db.Create(&profiles{Name: P.Name, Email: P.Email, Mobile: P.Mobile})

}
func updateprofile(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "root:securepassword@tcp(localhost:3306)/userlogin?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Print(err.Error)

	}
	if err == nil {
		fmt.Print("open")

	}

	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	var pro profiles
	db.Where("id=?", id).Find(&pro)
	pro.Mobile = "8090007284"
	db.Save(&pro)
}
func deleteprofile(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "root:securepassword@tcp(localhost:3306)/userlogin?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Print(err.Error)

	}
	if err == nil {
		fmt.Print("open")

	}

	defer db.Close()
	vars := mux.Vars(r)
	id := vars["id"]
	var pro profiles
	db.Where("id=?", id).Find(&pro)
	db.Delete(&pro)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/profiles", getprofiles).Methods("GET")
	r.HandleFunc("/api/profile", addprofile).Methods("POST")
	r.HandleFunc("/profile/{id}", updateprofile).Methods("PUT")
	r.HandleFunc("/profile/{id}", deleteprofile).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}
