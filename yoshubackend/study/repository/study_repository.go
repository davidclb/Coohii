package repository

import (
	"context"
	"log"

	postgres "yoshubackend/db/sqlc"
)


type StudyRepository interface {
/* 	FilteredCharactersByKeyset(ctx context.Context, Limit int32, ID int32, Carac string, list_carac []*postgres.Character ) (error)
	ListCharactersByKeyset(ctx context.Context, Limit int32, ID int32, list_carac []*postgres.Character ) ( error) */

	FilteredCharactersByKeyset(ctx context.Context, Limit int32, ID int32, Carac string) ([]*postgres.Character, error)    
	ListCharactersByKeyset(ctx context.Context, Limit int32, ID int32) ([]*postgres.Character, error)
}

type StudyRepo struct {
	Querie *postgres.Queries
}

func NewStudyRepository(r *StudyRepo) StudyRepository {
	return &StudyRepo{Querie: r.Querie}
}



func (r *StudyRepo) ListCharactersByKeyset(ctx context.Context, Limit int32, ID int32) ([]*postgres.Character, error) {

	listCharactersByKeyset := postgres.ListCharactersByKeysetParams{
		ID: ID,
		Pagesize: Limit,
	}

    log.Print(listCharactersByKeyset)
    result, err := r.Querie.ListCharactersByKeyset(ctx,  listCharactersByKeyset)
	if err != nil {
		return  nil, err
	}
	
	var characters []*postgres.Character
    for i := range result {
        characters = append(characters, &result[i])
    }

    log.Printf("ListCharacterbyKeyset repository: %v items retrived", len(characters))

    return characters, nil
}

func (r *StudyRepo) FilteredCharactersByKeyset(ctx context.Context, Limit int32, ID int32, Carac string) ([]*postgres.Character, error) {
    filteredCharactersByKeyset := postgres.FilteredCharactersByKeysetParams{
        Pagesize: Limit,
        ID:       ID,
        Carac:    Carac,
    }

    result, err := r.Querie.FilteredCharactersByKeyset(ctx, filteredCharactersByKeyset)

    if err != nil {
        return nil, err
    }

    var characters []*postgres.Character
    for i := range result {
        characters = append(characters, &result[i])
    }


    return characters, nil
}