package models

import "time"

// we will create the model for the notes
type NotesModel struct {
	NotesTitle string
	NotesDescription string
	IsCompleted bool
	CreatedAt time.Time
	UpdatedAt time.Time

}