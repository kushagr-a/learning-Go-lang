package notes

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Repo  --> data access layer --> mongodb se interact karega
type Repo struct {
	coll *mongo.Collection
}

// NewRepo --> constructor and mongoDb mei notes collection ko point karega
func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		coll: db.Collection("notes"),
	}
}

// CreateNote --> note create karega
func (r *Repo) CreateNote(ctx context.Context, note Note) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.coll.InsertOne(opCtx, note)
	if err != nil {
		return Note{}, err
	}

	return note, nil
}
