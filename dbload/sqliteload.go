package dbload

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//VerseNodes ....
type VerseNodes struct {
	Verses []Verse `json:"verses"`
}

// Verse ....
type Verse struct {
	BookID      int    `json:"book_id"`
	PathigamID  int    `json:"pathigam_id"`
	VerseID     int    `json:"verse_id"`
	TempleName  string `json:"temple_name"`
	Pann        string `json:"pann"`
	Verse       string `json:"verse"`
	Explanation string `json:"explanation"`
}

func connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./devaram.db")
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

//CreateVerse ... Inserts a Verse in to the DB
func CreateVerse(db *sql.DB, bookID int, pathigamID int, verseID int, templeName string, pann string, verse string, explanation string) {

	statement, _ := db.Prepare("INSERT INTO verses (book_id, pathigam_id, verse_id, temple_name, pann, verse, explanation) VALUES (?, ?, ?, ?, ?, ?, ?)")
	statement.Exec(bookID, pathigamID, verseID, templeName, pann, verse, explanation)
	fmt.Println("Inserted the verse into devaramDB!")

}

//ReadVerse ....Retrieves only one verse from the DB and wraps in the struct Verse.
func ReadVerse(db *sql.DB, bookID int, pathigamID int, verseID int) Verse {

	rows, _ := db.Query("SELECT book_id, pathigam_id, verse_id, temple_name, pann, verse, explanation FROM verses where book_id=? and pathigam_id=? and verse_id=?", bookID, pathigamID, verseID)
	var oneverse Verse
	for rows.Next() {
		rows.Scan(&oneverse.BookID, &oneverse.PathigamID, &oneverse.VerseID, &oneverse.TempleName, &oneverse.Pann, &oneverse.Verse, &oneverse.Explanation)
	}

	return oneverse
}

//ReadVerses .... Retrieves one or more verses from the DB and wraps in the struct VerseNodes.
func ReadVerses(db *sql.DB, bookID int, pathigamID int, verseID int, everseID int) VerseNodes {

	rows, _ := db.Query("SELECT book_id, pathigam_id, verse_id, temple_name, pann, verse, explanation FROM verses where verse_id between ? and ?", verseID, everseID)
	var varray VerseNodes
	var tempverse Verse

	for rows.Next() {
		rows.Scan(&tempverse.BookID, &tempverse.PathigamID, &tempverse.VerseID, &tempverse.TempleName, &tempverse.Pann, &tempverse.Verse, &tempverse.Explanation)

		varray.Verses = append(varray.Verses, tempverse)
	}
	return varray
}

//UpdateVerse ... Update one Verse in the DB
func UpdateVerse(db *sql.DB, bookID int, pathigamID int, verseID int, templeName string, pann string, verse string, explanation string) {
	statement1, _ := db.Prepare("update verses set verse=?, explanation=?, temple_name=?, pann=? where book_id=? and pathigam_id=? and verse_id=?")
	statement1.Exec(verse, explanation, templeName, pann, bookID, pathigamID, verseID)
	fmt.Println("Successfully updated the verse and explanation in devaramDB!")

}

//DeleteVerse ... delete a Verse in the DB
func DeleteVerse(db *sql.DB, bookID int, pathigamID int, verseID int) {

	statement, _ := db.Prepare("delete from verses where book_id=? and pathigam_id=? and verse_id=?")
	statement.Exec(bookID, pathigamID, verseID)
	fmt.Println("Successfully deleted the verse in database!")
}

//LoadIt .... Loads the JSON FILE DATA in to the SQLITE DB named devaram.db
func LoadIt() {
	db := connect()

	file, _ := ioutil.ReadFile("data.json")
	data := VerseNodes{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Verses); i++ {
		CreateVerse(db, data.Verses[i].BookID, data.Verses[i].PathigamID,
			data.Verses[i].VerseID, data.Verses[i].TempleName,
			data.Verses[i].Pann, data.Verses[i].Verse, data.Verses[i].Explanation)
	}

	defer closedb(db)
}
