package services

import (
	"errors"

	dao "github.com/Shashank100909/STUDENTS-API/daos"
	"github.com/Shashank100909/STUDENTS-API/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(input models.User) error
	CreateStudent(req models.Student) error
	GetStudent(StudentID int) ([]models.Student, error)
	FindUserByUsername(username string) (*models.User, error)
	DeleteStudent(StudentID int) error
}

type userService struct {
	userDAO dao.UserDAO
}

func NewUserService(userDAO dao.UserDAO) UserService {
	return &userService{userDAO}
}

func (s *userService) Register(input models.User) error {
	_, err := s.userDAO.GetUserByUsername(input.Username)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		FirstName: input.FirstName,
		LastName: input.LastName,
		Age: input.Age,
		Gender: input.Gender,
		Email: input.Email,
		MobileNumber: input.MobileNumber,
		Username: input.Username,
		Password: string(hashedPassword),
	}
	return s.userDAO.CreateUser(&user)
}

func (s *userService) CreateStudent(req models.Student) error {
	student := models.Student{
		Name:    req.Name,
		Class:   req.Class,
		Section: req.Section,
	}

	return s.userDAO.CreateStudent(&student)
}

func (s *userService) GetStudent(StudentID int) ([]models.Student, error) {
	return s.userDAO.GetStudent(StudentID)
}

func (s *userService) FindUserByUsername(username string) (*models.User, error) {
	return s.userDAO.GetUserByUsername(username)
}

func (s *userService) DeleteStudent(StudentID int) error {
	return s.userDAO.DeleteStudent(StudentID)
}
