package main

import (
    // "database/sql"
    // _ "github.com/go-sql-driver/mysql"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "time"
)

type App struct {
    Router *mux.Router
    // DB     *sql.DB
}

var notes = []Note{Note{ID: "1", Timestamp: 1, Text: "1"},
                   Note{ID: "2", Timestamp: 2, Text: "2"},
                   Note{ID: "3", Timestamp: 3, Text: "3"}}

func (a *App) Initialize() {//(user, password, dbname string) {
    // connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
    // var err error
    // a.DB, err = sql.Open("mysql", connectionString)
    // if err != nil {
    //     log.Fatal(err)
    // }

    a.Router = mux.NewRouter()
    a.initializeRoutes()
}

func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/notes", GetNotes).Methods("GET")
    a.Router.HandleFunc("/notes/{id}", GetNote).Methods("GET")
    a.Router.HandleFunc("/notes/{id}", CreateNote).Methods("POST")
    a.Router.HandleFunc("/notes/{id}", DeleteNote).Methods("DELETE")
}


// --- handlers ---


func GetNotes(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(notes)
}

// it works slow
func GetNote(w http.ResponseWriter, r *http.Request) {
    time.Sleep(10000 * time.Millisecond)
    params := mux.Vars(r)
    for _, item := range notes {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Note{})
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var note Note
    _ = json.NewDecoder(r.Body).Decode(&note)
    note.ID = params["id"]
    notes = append(notes, note)
    json.NewEncoder(w).Encode(notes)
}

// it returns 500 error code
func DeleteNote(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 - Something bad happened!"))

    // params := mux.Vars(r)
    // for index, item := range notes {
    //     if item.ID == params["id"] {
    //         notes = append(notes[:index], notes[index+1:]...)
    //         break
    //     }
    //     json.NewEncoder(w).Encode(notes)
    // }
}