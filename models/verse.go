package models

//Verse ...
type Verse struct {
	BookID      int    `json:"book_id"`
	PathigamID  int    `json:"pathigam_id"`
	VerseID     int    `json:"verse_id"`
	TempleName  string `json:"temple_name"`
	Pann        string `json:"pann"`
	Verse       string `json:"verse"`
	Explanation string `json:"explanation"`
	Translation string `json:"translation"`
}
