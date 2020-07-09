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
	VerseNodes []Verse `json:"verses"`
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

func createVerse(db *sql.DB, bookID int, pathigamID int, verseID int, templeName string, pann string, verse string, explanation string) {

	statement, _ := db.Prepare("INSERT INTO verses (book_id, pathigam_id, verse_id, temple_name, pann, verse, explanation) VALUES (?, ?, ?, ?, ?, ?, ?)")
	statement.Exec(bookID, pathigamID, verseID, templeName, pann, verse, explanation)
	fmt.Println("Inserted the verse into devaramDB!")

}

//ReadVerse ....
func ReadVerse(db *sql.DB, bookID int, pathigamID int, verseID int) Verse {

	rows, _ := db.Query("SELECT book_id, pathigam_id, verse_id, temple_name, pann, verse, explanation FROM verses where book_id=? and pathigam_id=? and verse_id=?", bookID, pathigamID, verseID)
	var oneverse Verse
	for rows.Next() {
		rows.Scan(&oneverse.BookID, &oneverse.PathigamID, &oneverse.VerseID, &oneverse.TempleName, &oneverse.Pann, &oneverse.Verse, &oneverse.Explanation)
		//fmt.Printf("bookID:%d, pathigamID:%d, verseID:%d, %v, %v, %v, %v\n", oneverse.BookID, oneverse.PathigamID, oneverse.VerseID, oneverse.TempleName, oneverse.Pann, oneverse.Verse, oneverse.Explanation)
	}

	return oneverse
}

func readVerses(db *sql.DB, bookID int, pathigamID int, verseID int) {

	rows, _ := db.Query("SELECT book_id, pathigam_id, verse_id, temple_name, pann, verse, explanation FROM verses")
	var tempverse Verse
	for rows.Next() {
		rows.Scan(&tempverse.BookID, &tempverse.PathigamID, &tempverse.VerseID, &tempverse.TempleName, &tempverse.Pann, &tempverse.Verse, &tempverse.Explanation)
		fmt.Printf("bookID:%d, pathigamID:%d, verseID:%d, %v, %v, %v, %v\n", tempverse.BookID, tempverse.PathigamID, tempverse.VerseID, tempverse.TempleName, tempverse.Pann, tempverse.Verse, tempverse.Explanation)
	}
}

func updateVerse(db *sql.DB, bookID int, pathigamID int, verseID int, templeName string, pann string, verse string, explanation string) {
	statement1, _ := db.Prepare("update verses set verse=?, explanation=?, temple_name=?, pann=? where book_id=? and pathigam_id=? and verse_id=?")
	statement1.Exec(verse, explanation, templeName, pann, bookID, pathigamID, verseID)
	fmt.Println("Successfully updated the verse and explanation in devaramDB!")

}

func deleteVerse(db *sql.DB, bookID int, pathigamID int, verseID int) {

	statement, _ := db.Prepare("delete from verses where book_id=? and pathigam_id=? and verse_id=?")
	statement.Exec(bookID, pathigamID, verseID)
	fmt.Println("Successfully deleted the verse in database!")
}

//LoadIt ....
func LoadIt() {
	db := connect()

	file, _ := ioutil.ReadFile("data.json")
	data := VerseNodes{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.VerseNodes); i++ {
		createVerse(db, data.VerseNodes[i].BookID, data.VerseNodes[i].PathigamID,
			data.VerseNodes[i].VerseID, data.VerseNodes[i].TempleName,
			data.VerseNodes[i].Pann, data.VerseNodes[i].Verse, data.VerseNodes[i].Explanation)
	}

	defer closedb(db)
}
