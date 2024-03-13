package lib_test

import (
	"testing"
	"time"

	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/config"
	"xyz.xeonds/nano-oj/lib"
	"xyz.xeonds/nano-oj/model"
)

var db *gorm.DB

func init() {
	config, _ := lib.LoadConfig[config.Config]()
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
	if err = db.Where("username = ?", user.Username).Find(new(model.User)).Error; err == nil {
		t.Log("User already exists: ", err)
		return
	}
	if user.Password, err = lib.HashedPassword(user.Password), db.Create(&user).Error; err != nil {
		t.Fatal("Failed to create user: ", err)
		return
	}
	if user.ID == 1 { // if it is the first user, set it as admin
		user.AccountInfo.Permission = 0
		if err := db.Save(&user); err != nil {
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
	if err := lib.CheckPasswordHash(input.Password, user.Password); err != nil {
		t.Log("input pass: ", input.Password, "user pass: ", user.Password)
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
