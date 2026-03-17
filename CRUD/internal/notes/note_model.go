package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	// bson tag mongodb ke liye
	// json tag API response ke liye

	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Title string `bson:"title" json:"title"`

	Content string `bson:"content" json:"content"`

	Pinned bool `bson:"pinned" json:"pinned"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`

	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// Request body for creating a note
type CreateNoteRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Pinned  bool   `json:"pinned"`
}
