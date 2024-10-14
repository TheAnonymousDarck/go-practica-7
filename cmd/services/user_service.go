package services

import (
	"github.com/TheAnonymousDarck/go-practica-7/cmd/database"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

// GetAllUsers Obtener todos los usuarios
func (s *UserService) GetAllUsers() ([]database.User, error) {
	var users []database.User
	err := s.DB.Find(&users).Error
	return users, err
}

// GetUserByID Obtener un usuario por ID
func (s *UserService) GetUserByID(id int) (*database.User, error) {
	var user database.User
	err := s.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser Crear un nuevo usuario
func (s *UserService) CreateUser(user *database.User) error {
	return s.DB.Create(user).Error
}

// UpdateUser Actualizar un usuario
func (s *UserService) UpdateUser(id int, updatedUser *database.User) error {
	var user database.User
	err := s.DB.First(&user, id).Error
	if err != nil {
		return err
	}
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	return s.DB.Save(&user).Error
}

// DeleteUser Eliminar un usuario
func (s *UserService) DeleteUser(id int) error {
	return s.DB.Delete(&database.User{}, id).Error
}
