package dao

import "../models"

//VerseDao ....
type VerseDao interface {
	CreateVerse(bookID int, pathigamID int, verseID int, templeName string, pann string, verse string, explanation string)
	ReadVerse(bookID int, pathigamID int, verseID int) models.Verse
	ReadVerses(bookID int, pathigamID int, verseID int, everseID int) models.VerseNodes
	UpdateVerse(bookID int, pathigamID int, verseID int, templeName string, pann string, verse string, explanation string)
	DeleteVerse(bookID int, pathigamID int, verseID int)
}
