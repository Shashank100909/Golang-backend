package dao

import (
	"github.com/Shashank100909/STUDENTS-API/configs"
	"github.com/Shashank100909/STUDENTS-API/models"
	"gorm.io/gorm"
)

type UserDAO interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	CreateStudent(student *models.Student) error
	GetStudent(StudentID int) ([]models.Student, error)
	DeleteStudent(StudentID int) error
}

type userDAO struct {
	db *gorm.DB
}

func NewUserDAO() UserDAO {
	return &userDAO{
		db: configs.DB,
	}
}

func (dao *userDAO) CreateUser(user *models.User) error {
	return dao.db.Create(user).Error
}

func (dao *userDAO) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := dao.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *userDAO) CreateStudent(student *models.Student) error {
	return dao.db.Create(student).Error
}

func (dao *userDAO) GetStudent(StudentID int) ([]models.Student, error) {

	var students []models.Student
	query := dao.db.Model(&models.Student{}).
		Select(`
            id,
            name,
            class,
            section
        `)

	if StudentID != 0 {
		query = query.Where("id = ?", StudentID)
	}

	err := query.Find(&students).Error
	if err != nil {
		return nil, err
	}

	if len(students) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return students, nil
}

func (dao *userDAO) DeleteStudent(StudentID int) error {
	var student models.Student
	err := dao.db.Where("id = ?", StudentID).First(&student).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
	}
	return dao.db.Delete(&student).Error
}
