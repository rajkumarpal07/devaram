package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"./dao"
	sl "./dao/sqlite"
	util "./util"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func connect() *sql.DB {
	config, err := util.GetConfiguration()
	db, err := sql.Open(config.Engine, config.Database)
	if err != nil {
		log.Println(err)
	}
	return db
}

func closedb(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}

//VerseHandler ....
func VerseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db := connect()
	bid1, _ := strconv.Atoi(vars["bid"])
	pid1, _ := strconv.Atoi(vars["pid"])
	vid1, _ := strconv.Atoi(vars["vid"])

	var intf dao.VerseDao
	intf = sl.VerseImplSqlite{}

	data := intf.ReadVerse(bid1, pid1, vid1)

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	js, _ := json.Marshal(data)
	w.Write(js)
	defer closedb(db)
}

//VersesHandler ....
func VersesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db := connect()

	bid1, _ := strconv.Atoi(vars["bid"])
	pid1, _ := strconv.Atoi(vars["pid"])
	vid1, _ := strconv.Atoi(vars["vid"])
	eid1, _ := strconv.Atoi(vars["eid"])

	var intf dao.VerseDao
	intf = sl.VerseImplSqlite{}

	data1 := intf.ReadVerses(bid1, pid1, vid1, eid1)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	js, _ := json.Marshal(data1)
	w.Write(js)
	defer closedb(db)
}

func main() {

	//sl.LoadIt()
	r := mux.NewRouter()
	r.HandleFunc("/{bid:[0-9]+}/{pid:[0-9]+}/{vid:[0-9]+}", VerseHandler)

	r.HandleFunc("/{bid:[0-9]+}/{pid:[0-9]+}/{vid:[0-9]+}-{eid:[0-9]+}", VersesHandler)

	srv := &http.Server{
		Handler: r, Addr: "127.0.0.1:4567", WriteTimeout: 15 * time.Second, ReadTimeout: 15 * time.Second}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
