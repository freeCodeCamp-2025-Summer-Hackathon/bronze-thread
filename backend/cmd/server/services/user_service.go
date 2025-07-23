package services

import (
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/db"
)

type UserService struct{}

func (s *UserService) GetById(id string) (*db.User, error) {
	var user db.User
	result := db.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// func (s *UserService) UpdateById(currentUser db.User, data any) (*db.User, error) {
// 	var user db.User
// 	updateResult := db.DB.Where("ID = ?", currentUser).Model(&user).Updates(data)

// 	if updateResult.Error != nil {
// 		return nil, updateResult.Error
// 	}

// 	fmt.Print("USER UPDSTED:::", updateResult.Config)

// 	return &user, updateResult.Error
// }

var User = &UserService{}
