package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"../../dao"
	"../../models"
	util "../../util"
)

//VerseImplSqlite ...
type VerseImplSqlite struct {
}

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

//CreateVerse ... Inserts a Verse in to the DB
func (dao VerseImplSqlite) CreateVerse(bookID int, pathigamID int, verseID int, templeName string, pann string, verse string, explanation string) {
	db := connect()
	statement, _ := db.Prepare("INSERT INTO verses (book_id, pathigam_id, verse_id, temple_name, pann, verse, explanation) VALUES (?, ?, ?, ?, ?, ?, ?)")
	statement.Exec(bookID, pathigamID, verseID, templeName, pann, verse, explanation)
	fmt.Println("Inserted the verse into devaramDB!")

	defer closedb(db)

}

//ReadVerse ....Retrieves only one verse from the DB and wraps in the struct Verse.
func (dao VerseImplSqlite) ReadVerse(bookID int, pathigamID int, verseID int) models.Verse {
	db := connect()
	rows, _ := db.Query("SELECT book_id, pathigam_id, verse_id, temple_name, pann, verse, explanation FROM verses where book_id=? and pathigam_id=? and verse_id=?", bookID, pathigamID, verseID)
	var oneverse models.Verse
	for rows.Next() {
		rows.Scan(&oneverse.BookID, &oneverse.PathigamID, &oneverse.VerseID, &oneverse.TempleName, &oneverse.Pann, &oneverse.Verse, &oneverse.Explanation)
	}
	defer closedb(db)
	return oneverse
}

//ReadVerses .... Retrieves one or more verses from the DB and wraps in the struct VerseNodes.
func (dao VerseImplSqlite) ReadVerses(bookID int, pathigamID int, verseID int, everseID int) models.VerseNodes {
	db := connect()
	rows, _ := db.Query("SELECT book_id, pathigam_id, verse_id, temple_name, pann, verse, explanation FROM verses where verse_id between ? and ?", verseID, everseID)
	var varray models.VerseNodes
	var tempverse models.Verse

	for rows.Next() {
		rows.Scan(&tempverse.BookID, &tempverse.PathigamID, &tempverse.VerseID, &tempverse.TempleName, &tempverse.Pann, &tempverse.Verse, &tempverse.Explanation)

		varray.Verses = append(varray.Verses, tempverse)
	}
	defer closedb(db)
	return varray
}

//UpdateVerse ... Update one Verse in the DB
func (dao VerseImplSqlite) UpdateVerse(bookID int, pathigamID int, verseID int, templeName string, pann string, verse string, explanation string) {
	db := connect()
	statement1, _ := db.Prepare("update verses set verse=?, explanation=?, temple_name=?, pann=? where book_id=? and pathigam_id=? and verse_id=?")
	statement1.Exec(verse, explanation, templeName, pann, bookID, pathigamID, verseID)
	fmt.Println("Successfully updated the verse and explanation in devaramDB!")
	defer closedb(db)
}

//DeleteVerse ... delete a Verse in the DB
func (dao VerseImplSqlite)  DeleteVerse(bookID int, pathigamID int, verseID int) {
	db := connect()
	statement, _ := db.Prepare("delete from verses where book_id=? and pathigam_id=? and verse_id=?")
	statement.Exec(bookID, pathigamID, verseID)
	fmt.Println("Successfully deleted the verse in database!")
	defer closedb(db)
}

//LoadIt .... Loads the JSON FILE DATA in to the SQLITE DB named devaram.db
func LoadIt() {
	db := connect()

	file, _ := ioutil.ReadFile("data.json")
	data := models.VerseNodes{}

	_ = json.Unmarshal([]byte(file), &data)
	var intf dao.VerseDao
	intf = VerseImplSqlite{}

	for i := 0; i < len(data.Verses); i++ {

		intf.CreateVerse(data.Verses[i].BookID, data.Verses[i].PathigamID,
			data.Verses[i].VerseID, data.Verses[i].TempleName,
			data.Verses[i].Pann, data.Verses[i].Verse, data.Verses[i].Explanation)
	}

	defer closedb(db)
}
