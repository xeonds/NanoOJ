package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/lib"
	"xyz.xeonds/nano-oj/model"
)

func TestAPIPermission(t *testing.T) {
	tokens := make([]string, 3)
	users := make([]model.User, 3)
	for i := 0; i < 3; i++ {
		users[i] = model.User{
			Username: "user" + string(rune(i)),
			Password: "password",
			AccountInfo: model.AccountInfo{
				Permission: uint8(i),
			},
		}
		tokens[i], _ = lib.GenerateToken(&lib.UserClaim{
			Name:       users[i].Username,
			Permission: int(users[i].AccountInfo.Permission),
		})
	}
	db.Create(&users)
	router := gin.Default()
	lib.APIBuilder(router,
		func(group *gin.RouterGroup) *gin.RouterGroup {
			group.GET("/a", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "pong"})
			})
			return group
		},
		func(group *gin.RouterGroup) *gin.RouterGroup {
			group.Use(lib.JWTMiddleware(lib.AuthPermission(0, 1)))
			group.GET("/b", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "pong"})
			})
			return group
		})(router, "")
	// Test permission 0
	req, _ := http.NewRequest("GET", "/a", nil)
	// req.Header.Set("Authorization", tokens[0])
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error("Permission 0 failed")
	}
	// Test permission 1
	req, _ = http.NewRequest("GET", "/a", nil)
	// req.Header.Set("Authorization", tokens[1])
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error("Permission 1 failed")
	}
	// Test permission 2
	req, _ = http.NewRequest("GET", "/a", nil)
	// req.Header.Set("Authorization", tokens[2])
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error("Permission 2 failed", w.Code)
	}

	// Test permission 0
	req, _ = http.NewRequest("GET", "/b", nil)
	req.Header.Set("Authorization", tokens[0])
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error("Permission 0 failed")
	}
	// Test permission 1
	req, _ = http.NewRequest("GET", "/b", nil)
	req.Header.Set("Authorization", tokens[1])
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error("Permission 1 failed")
	}
	// Test permission 2
	req, _ = http.NewRequest("GET", "/b", nil)
	req.Header.Set("Authorization", tokens[2])
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Error("Permission 2 failed", w.Code)
	}
}
