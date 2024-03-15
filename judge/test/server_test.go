package test

import (
	"errors"
	"testing"
	"time"

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
		"testuser": {
			Username: "testuser",
			Email:    "test@test",
			Password: "testpass",
		},
	}
	r := gin.Default()
	lib.AddLoginAPI(r, "", db)
	w := httptest.NewRecorder()
	for name, testCase := range testCases {
		payload, _ := json.Marshal(&testCase)
		req, _ := http.NewRequest("POST", "/login", payload)
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		}
	}
}