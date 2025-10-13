package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"example.com/net-http-class/models"
	"example.com/net-http-class/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// custom struct for JWT
type CustomClaims struct {
	Email   string    `json:"email"`
	User_ID uuid.UUID `json:"userid"`
	jwt.RegisteredClaims
}

func (s *UserService) Createuser(user *models.Users) error {
	// Hashpassword using Bcrypt
	hashedpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return fmt.Errorf("hashing Password failed: %w", err)
	}

	user.Password = string(hashedpass) // setting incoming datapassword to the hashed password

	err = s.userRepo.Create(user)
	if err != nil {
		return fmt.Errorf("user creation failed: %w", err)
	}
	return nil
}

func (s *UserService) Loginuser(user *models.Users) (string, error) {
	Plainpassword := user.Password
	dbUser, err := s.userRepo.FindUserByEmail(user.Email)
	if err != nil {
		return "", fmt.Errorf("couldn't find user: %w", err)
	}

	// compare plain password to hashedpassword
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(Plainpassword))
	if err != nil {
		return "", fmt.Errorf("password do not match: %w", err)
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
