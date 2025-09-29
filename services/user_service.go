package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"example.com/net-http-class/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// custom struct for JWT
type CustomClaims struct {
	Email   string    `json:"email"`
	User_ID uuid.UUID `json:"userid"`
	jwt.RegisteredClaims
}

type Userservicedependencies struct {
	db *gorm.DB
}

func NewUserservicedependencies(DB *gorm.DB) *Userservicedependencies {
	return &Userservicedependencies{db: DB}
}

func (s *Userservicedependencies) Createuser(user *models.Users) error {
	// Hashpassword using Bcrypt
	hashedpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return fmt.Errorf("hashing Password failed: %w", err)
	}

	user.Password = string([]byte(hashedpass)) // setting incoming datapassword to the hashed password

	result := s.db.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("user creation failed: %w", result.Error)
	}
	return nil
}

func (s *Userservicedependencies) Loginuser(user *models.Users) (string, error) {
	Plainpassword := user.Password
	result := s.db.Where("email = ?", user.Email).First(user)
	if result.Error != nil {
		return "", fmt.Errorf("couldn't find user: %w", result.Error)
	}

	// compare plain password to hashedpassword
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Plainpassword))
	if err != nil {
		return "", fmt.Errorf("password do not match: %w", result.Error)
	}

	// secretKey for signing JWT token
	secretKey := os.Getenv("JWT_SECRET")

	claims := CustomClaims{
		Email:   user.Email,
		User_ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, err
}
