package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/config"
	"xyz.xeonds/nano-oj/lib"
	"xyz.xeonds/nano-oj/model"
)

var db *gorm.DB

func init() {
	config := lib.LoadConfig[config.Config]()
	db = lib.NewDB(&config.DatabaseConfig, func(db *gorm.DB) error {
		return db.AutoMigrate(&model.Submission{}, &model.Problem{}, &model.User{}, &model.Contest{}, &model.Notification{})
	})
}

func TestUserRegister(t *testing.T) {
	var err error
	user := model.User{
		Username: "testuser",
		Password: "testpassword",
		Email:    "test@email.com",
	}
	var count int64
	if err = db.Where("username = ?", user.Username).Find(new([]model.User)).Count(&count).Error; count != 0 {
		t.Fatal("User already exists: ", err)
		return
	}
	if user.Password, err = lib.HashedPassword(user.Password), db.Create(&user).Error; err != nil {
		t.Fatal("Failed to create user: ", err)
		return
	}
	if user.ID == 1 { // if it is the first user, set it as admin
		user.AccountInfo.Permission = 0
		if err := db.Save(&user).Error; err != nil {
			t.Fatal("Failed to set user as admin: ", err)
		}
	}

}

func TestUserLogin(t *testing.T) {
	input, user := model.User{
		Username: "testuser",
		Password: "testpassword",
		Email:    "test@email.com",
	}, new(model.User)

	if err := db.Where("email = ?", input.Email).Find(user).Error; err != nil {
		t.Fatal("Find user by email failed: ", err)
	}
	t.Log("input pass: ", input.Password, "user pass: ", user.Password)
	if err := lib.CheckPasswordHash(input.Password, user.Password); err != nil {
		t.Fatal("Incorrect password: ", err)
	}
	claim := lib.UserClaim{
		Name:       user.Username,
		Expire:     time.Now().Add(time.Hour * 24),
		Permission: int(user.AccountInfo.Permission),
	}
	token, err := lib.GenerateToken(&claim)
	if err != nil {
		t.Fatal("Failed to generate token: ", err)
	}
	t.Log("Token: ", token)
}

func TestUserPermission(t *testing.T) {
	input, user := model.User{
		Username: "testuser",
		Password: "testpassword",
		Email:    "test@email.com",
	}, new(model.User)

	if err := db.Where("email = ?", input.Email).Find(user).Error; err != nil {
		t.Fatal("Find user by email failed: ", err)
	}
	t.Log("input pass: ", input.Password, "user pass: ", user.Password)
	if err := lib.CheckPasswordHash(input.Password, user.Password); err != nil {
		t.Fatal("Incorrect password: ", err)
	}
	claim := lib.UserClaim{
		Name:       user.Username,
		Expire:     time.Now().Add(time.Hour * 24),
		Permission: int(user.AccountInfo.Permission),
	}
	token, err := lib.GenerateToken(&claim)
	if err != nil {
		t.Fatal("Failed to generate token: ", err)
	}
	parsed, err := lib.ParseToken(token)
	if err != nil {
		t.Fatal("Failed to parse token:", err)
	}
	if func(permLo, permHi int) error {
		if parsed.Permission < permLo || parsed.Permission > permHi {
			return errors.New("permission not valid")
		}
		return nil
	}(0, 1) != nil {
		t.Error("Priviledge not included in allowance area")
	}
}

func TestLoginAPI(t *testing.T) {
	testCases := map[string]model.User{
		"testuser1": {
			Username: "testuser1",
			Email:    "test1@test",
			Password: "testpass1",
		},
		"testuser2": {
			Username: "testuser2",
			Email:    "test2@test",
			Password: "testpass2",
		},
		"testuser3": {
			Username: "testuser3",
			Email:    "test3@test",
			Password: "testpass3",
		},
	}
	r := gin.Default()
	lib.AddLoginAPI(r, "", db)
	w := httptest.NewRecorder()
	for _, testCase := range testCases {
		// Register API test
		registerPayload, err := json.Marshal(&testCase)
		if err != nil {
			t.Fatal("Failed to marshal register payload:", err)
		}
		registerReq, err := http.NewRequest("POST", "/register", bytes.NewReader(registerPayload))
		if err != nil {
			t.Fatal("Failed to create register request:", err)
		}
		registerReq.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, registerReq)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		}

		// Login API test
		loginPayload, err := json.Marshal(&testCase)
		if err != nil {
			t.Fatal("Failed to marshal login payload:", err)
		}
		loginReq, err := http.NewRequest("POST", "/login", bytes.NewReader(loginPayload))
		if err != nil {
			t.Fatal("Failed to create login request:", err)
		}
		loginReq.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, loginReq)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		}
	}
}

func TestUserLoginWithWrongPassword(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	lib.AddLoginAPI(router, "", db)

	// Register a user
	registerData := model.User{
		Username: "testuser",
		Email:    "user@email",
		Password: "correctpassword",
	}
	registerBody, _ := json.Marshal(registerData)
	registerRequest, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(registerBody))
	registerResponse := httptest.NewRecorder()
	router.ServeHTTP(registerResponse, registerRequest)

	// Check the response status code for registration
	if registerResponse.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, registerResponse.Code)
	}

	// Create a test request with a wrong password
	loginData := model.User{
		Email:    "user@email",
		Password: "wrongpassword",
	}
	loginBody, _ := json.Marshal(loginData)
	loginRequest, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginBody))
	loginResponse := httptest.NewRecorder()

	// Perform the login request
	router.ServeHTTP(loginResponse, loginRequest)

	// Check the response status code for login
	if loginResponse.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, but got %d", http.StatusUnauthorized, loginResponse.Code)
	}

	// Check the response body
	expectedResponse := map[string]interface{}{
		"error": "Incorrect password",
	}
	var responseBody map[string]interface{}
	json.Unmarshal(loginResponse.Body.Bytes(), &responseBody)
	if !reflect.DeepEqual(responseBody, expectedResponse) {
		t.Errorf("Expected response body %+v, but got %+v", expectedResponse, responseBody)
	}
}

func TestHandleFind(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	lib.AddLoginAPI(router, "", db)
	router.GET("/find",
		lib.HandleFind[model.User](
			func(c *gin.Context) *gorm.DB {
				return db.Where("username = ?", "testuser")
			}))

	// Register a user
	registerBody, _ := json.Marshal(model.User{
		Username: "testuser",
		Email:    "user@email",
		Password: "correctpassword",
	})
	registerRequest, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(registerBody))
	registerResponse := httptest.NewRecorder()
	router.ServeHTTP(registerResponse, registerRequest)
	if registerResponse.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, registerResponse.Code)
	}

	findRequest, _ := http.NewRequest("GET", "/find", bytes.NewBuffer([]byte("testuser")))
	findResponse := httptest.NewRecorder()
	router.ServeHTTP(findResponse, findRequest)
	expectedResponse := map[string]interface{}{
		"username": "testuser",
		"email":    "user@email",
		"password": "correctpassword",
	}
	var responseBody map[string]interface{}
	json.Unmarshal(findResponse.Body.Bytes(), &responseBody)
	if responseBody["email"] != "user@email" {
		t.Errorf("Expected response body %+v, but got %+v", expectedResponse, responseBody)
	}
}
