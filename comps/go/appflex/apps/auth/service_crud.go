package auth

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/waykiss/wkcomps/str"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	dao Dao
}

// NewService factory to Create mew userService
func NewService() (s Service, err error) {
	dao, err := NewDao()
	if err != nil {
		return
	}
	s.dao = dao
	return
}

// register a user with name, email and password
func (s *Service) register(name, email, password string, age uint) (err error) {
	u := NewModel(name, email)
	u.Password = password
	u.Age = age
	err = s.Create(&u)
	if err != nil {
		return
	}
	return
}

// Create a new user, perform data parser, validations and persist
func (s *Service) Create(m *Model) (err error) {
	// set default fields
	m.CreatedAt = time.Now()
	m.UpdatedAt = m.CreatedAt
	m.Id = str.Uuid()
	if m.Status == "" {
		m.Status = StatusUnconfirmed
	}
	m.ConfirmationCode = str.RandString(4, str.RandStringCharsOnlyNumbers.String())
	inputParser(m)

	// validation
	if err = checkPasswordPolicy(m.Password); err != nil {
		return
	}
	// generate hashed password
	m.Password = hashPassword(m.Password)

	// validate all the model with hashed password
	if err = validate(m); err != nil {
		return
	}
	userDb, err := s.dao.FindByEmail(m.Email)
	if err != nil {
		return
	}
	if userDb.Id != "" {
		err = fmt.Errorf("user with email %s already exists", m.Email)
		return
	}
	//Create using DAO
	err = s.dao.Create(*m)
	return
}

// Update a user, perform parser and validation
func (s *Service) Update(m *Model) (err error) {
	m.UpdatedAt = time.Now()
	inputParser(m)
	if err = validate(m); err != nil {
		return
	}
	err = s.dao.Update(*m)
	return
}

// Delete Delete e organization in the database and in the cloudStorage
func (s *Service) Delete(id string) (err error) {
	if id == "" {
		return fmt.Errorf("id not passed by parameter")
	}
	query := Query{}
	query.Id = id
	records, err := s.Find(query)
	if err != nil {
		return
	}
	if len(records) == 0 {
		return fmt.Errorf("user not found")
	}
	err = s.dao.Delete(records[0])
	return
}

// Find General function to query the models
func (s *Service) Find(query Query) (r []Model, err error) {
	r, err = s.dao.Find(query)
	return
}

// login perform login to the system, returns loginInfoModel as a result of the login
func (s *Service) login(email, password string) (l loginInfoModel, err error) {
	email = str.LowerNoSpaceNoAccent(email)
	user, err := s.dao.FindByEmail(email)
	if err != nil {
		return
	}

	if user.Id == "" {
		err = fmt.Errorf("%s", invalidUserAndPassword)
		return
	}

	// check if password is correct
	if !comparePassword(user.Password, password) {
		err = fmt.Errorf("%s", invalidUserAndPassword)
		return
	}
	l.Name = user.Name
	l.Email = user.Email
	l.UserId = user.Id
	l.Token = generateJwtToken(user.Id, 30)
	return
}

// generateJwtToken returns the signed string(token)
func generateJwtToken(userId string, minutes float64) string {
	mySigningKey := []byte(os.Getenv("JWT_SECRET"))

	type Claims struct {
		UserId string `json:"userId"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := Claims{
		UserId: userId,
	}
	claims.IssuedAt = time.Now().Unix()

	if minutes > 0 {
		claims.ExpiresAt = time.Now().UTC().Add(time.Duration(minutes) * time.Minute).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Error(err)
	}
	return tokenStr
}

// comparePassword check if the hashed password corresponds to the plain password
func comparePassword(hashedPwd, plainPwd string) bool {
	// will be a string so we'll need to convert it to a byte slice
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)); err != nil {
		return false
	}
	return true
}

// hashPassword function to encrypt the user password
func hashPassword(pwd string) string {
	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
