package jsondb

import (
	"deni1688/exerciseTracker/config"
	"deni1688/exerciseTracker/domain"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	scribble "github.com/nanobox-io/golang-scribble"
)

type exerciseRepository struct {
	db *scribble.Driver
}

func NewExerciseRepository() (domain.Repository, error) {
	dir := config.GetString("storage.path")
	if dir == "" {
		return nil, errors.New("error getting the JSON_DB_DIR env variable")
	}

	db, err := scribble.New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	return &exerciseRepository{db}, nil
}

func (r *exerciseRepository) FindAll() (*[]domain.Exercise, error) {
	records, err := r.db.ReadAll(domain.Collection)
	if err != nil {
		fmt.Println("Error", err)
	}

	results := make([]domain.Exercise, 0)
	for _, record := range records {
		ex := domain.Exercise{}
		_ = json.Unmarshal([]byte(record), &ex)
		results = append(results, ex)
	}
	return &results, nil
}

func (r *exerciseRepository) Create(ex *domain.Exercise) (string, error) {
	ex.ID = uuid.New().String()
	ex.Created = time.Now().Unix()

	if err := r.db.Write(domain.Collection, ex.ID, ex); err != nil {
		return "", err
	}

	return ex.ID, nil
}
