package test_integration

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/daisuke0925m/sample-saas/domain"
	"github.com/daisuke0925m/sample-saas/infrastructure"
	"github.com/daisuke0925m/sample-saas/interfaces/controllers"
	"github.com/gin-gonic/gin"
	"github.com/sebdah/goldie/v2"
)

func setUp(t *testing.T) *gin.Engine {
	// TODO change by env
	conn, err := sql.Open("mysql", "root:@tcp("+os.Getenv("DB_HOST")+":3306)/"+"sample_test")
	if err != nil {
		t.Fatalf(err.Error())
	}
	tx, err := conn.Begin()
	if err != nil {
		t.Fatalf(err.Error())
	}

	infraTx := infrastructure.Tx{
		Tx: tx,
	}

	userController := controllers.NewUserController(infraTx)
	_, err = userController.Interactor.UserRepository.Store(domain.User{
		ID:        0,
		FirstName: "firstName",
		LastName:  "lastName",
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	r := infrastructure.NewRouter(userController)
	t.Cleanup(func() {
		if err := infraTx.RollBack(); err != nil {
			t.Fatalf("failed RollBack")
		}
	})
	return r
}

func TestUsers(t *testing.T) {
	r := setUp(t)
	api := "/users"
	req := httptest.NewRequest(http.MethodGet, api, nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatal(rec)
	}
	g := goldie.New(t)
	ret := regexp.MustCompile(`\"id\":[0-9]+`).ReplaceAllString(rec.Body.String(), `"id":0`)
	g.AssertJson(t, t.Name(), ret)
}
