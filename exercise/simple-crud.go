package exercise

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Database
const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = ""
	dbName   = "gocrud_test"
)

type User struct {
	ID    int
	Name  string
	Email string
}

/*		CREATE USER SECTION		*/
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// masuk ke database
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// jika error hentikan program dan shutdown
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// parsing data JSON dari request Body
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	CreateUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Gagal untuk membuat user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Data user Berhasil dibuat!")
}

// exec query untuk membuat user
func CreateUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"

	// exec query dan pengecekan data
	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}
	return nil
}

/*		READ USER SECTION		*/
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// ambil parameter "id" dari URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// convert menjadi integer
	userID, err := strconv.Atoi(idStr)

	user, err := GetUser(db, userID)
	if err != nil {
		http.Error(w, "User tidak ditemukan", http.StatusNotFound)
		return
	}

	// Convert user object ke JSON dan mengirimkannya ke response
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUser(db *sql.DB, id int) (*User, error) {
	// get data & row
	query := "SELECT * FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)

	// checking
	if err != nil {
		return nil, err
	}

	return user, nil
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// ambil parameter "id" dari URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// convert menjadi integer
	userID, err := strconv.Atoi(idStr)

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)

	UpdateUser(db, userID, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Data user tidak dapat ditemukan", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Data user berhasil dipebaharui")
}

func UpdateUser(db *sql.DB, id int, name, email string) error {
	query := "UPDATE users SET name = ?, email = ?, WHERE id = ?"
	_, err := db.Exec(query, name, email, id)

	if err != nil {
		return err
	}

	return nil
}

/*		DELETE SECTION		*/
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// ambil parameter "id" dari URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// convert menjadi integer
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Parameter 'id' Tidak sesuai!", http.StatusBadRequest)
		return
	}

	user := DeleteUser(db, userID)
	if err != nil {
		http.Error(w, "Data user tidak dapat ditemukan", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Data user berhasil dihapus")

	// convert user object menjadi JSON dan mengirimnya ke response
	w.Header().Set("Contet-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func ActivateApp() {
	// membuat router baru (mux)
	r := mux.NewRouter()

	// define rute HTTP dengan menggunakan router
	r.HandleFunc("/user", createUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

	// start HTTP server di port 8080 (localhost:8090)
	log.Println("Server Listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
