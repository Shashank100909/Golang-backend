package handlers

import (
	"net/http"
	"strconv"

	"github.com/Shashank100909/STUDENTS-API/dtos"
	"github.com/Shashank100909/STUDENTS-API/models"
	services "github.com/Shashank100909/STUDENTS-API/service"
	"github.com/Shashank100909/STUDENTS-API/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var input dtos.RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *UserHandler) CreateStudent(c *gin.Context) {
	var req models.Student
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.userService.CreateStudent(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Student registered successfully"})
}

func (h *UserHandler) GetStudent(c *gin.Context) {
	StudentIDstr := c.Query("id")

	StudentID := 0
	var err error
	if StudentIDstr != "" {
		StudentID, err = strconv.Atoi(StudentIDstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid student ID"})
			return
		}
	}

	Students, err := h.userService.GetStudent(StudentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK,
				gin.H{"Error": "Students not found"})
			return
		}
		c.JSON(http.StatusInternalServerError,
			gin.H{"Error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{"Students": Students})
}

func (h *UserHandler) Login(c *gin.Context) {
	var input dtos.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.FindUserByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

func (h *UserHandler) DeleteStudent(c *gin.Context) {
	studentIDstr := c.Param("id")

	StudentID, err := strconv.Atoi(studentIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect student Id"})
		return
	}
	err = h.userService.DeleteStudent(StudentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest,
				gin.H{"Error": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError,
			gin.H{"Error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{"Success": "Student Deleted successfully"})
}
