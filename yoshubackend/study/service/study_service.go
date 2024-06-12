package service

import (
	"context"
	"log"

	"yoshubackend/db/sqlc"
	"yoshubackend/study/repository"

	"github.com/gofiber/fiber/v2"
)

type studyService struct {
	StudyRepository repository.StudyRepository
}



type StudyConfig struct{
	Repository repository.StudyRepository
 }

func NewStudyService(c *StudyConfig) StudyService {
	return &studyService{
		StudyRepository:  c.Repository,
		}
}


func (US *studyService ) GetHanzi(ctx *fiber.Ctx, filter string, limit int32,  lastID int32 ) ([]*sqlc.Character, error){

var characters []*sqlc.Character
c := context.Background()

var err error

log.Print(filter)
log.Print(limit)

if filter != "" {
	characters, err = US.StudyRepository.FilteredCharactersByKeyset(c,  int32(limit), int32(lastID), filter)
	
} else {

	log.Print("pas filtré")	
	characters, err = US.StudyRepository.ListCharactersByKeyset(c, int32(limit), int32(lastID))
}
if err != nil {
	return nil, err
}


log.Printf("ListCharacterbyKeyset service: %v items retrived", len(characters))
return characters, err

}

