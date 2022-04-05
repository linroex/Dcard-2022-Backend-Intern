package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func init() {
	getRand = func(_ int) int {
		return 1
	}
}

func TestRandString(t *testing.T) {
	assert.Equal(t, RandString(3), "bbb")
}

func TestAddUrlsWithExpired(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	router := setupRouter(&Env{db: db})

	var jsonStr = []byte(`{"url":"https://coder.tw", "expireAt": "2022-04-08T09:20:41Z"}`)

	mock.ExpectPrepare("^INSERT.*expired_at.*").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/urls", bytes.NewBuffer(jsonStr))

	router.ServeHTTP(w, req)
	body, _ := io.ReadAll(w.Body)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"id":"1bbb","shortUrl":"`+os.Getenv("baseUrl")+`1bbb"}`, string(body))
}

func TestAddUrlsWithErrorExpired(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	router := setupRouter(&Env{db: db})

	var jsonStr = []byte(`{"url":"https://coder.tw", "expireAt": "22-04-08T09:20:41Z"}`)

	mock.ExpectPrepare("^INSERT.*expired_at.*").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/urls", bytes.NewBuffer(jsonStr))

	router.ServeHTTP(w, req)
	body, _ := io.ReadAll(w.Body)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"message":"expiredAt invalid"}`, string(body))
}

func TestAddUrlsWithErrorUrl(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	router := setupRouter(&Env{db: db})

	var jsonStr = []byte(`{"url":"coder.tw", "expireAt": "2022-04-08T09:20:41Z"}`)

	mock.ExpectPrepare("^INSERT.*expired_at.*").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/urls", bytes.NewBuffer(jsonStr))

	router.ServeHTTP(w, req)
	body, _ := io.ReadAll(w.Body)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"message":"url invalid"}`, string(body))
}

func TestAddUrlsNoExpiredAt(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	router := setupRouter(&Env{db: db})

	var jsonStr = []byte(`{"url":"https://coder.tw"}`)

	mock.ExpectPrepare("^INSERT.*").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/urls", bytes.NewBuffer(jsonStr))

	router.ServeHTTP(w, req)
	body, _ := io.ReadAll(w.Body)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"id":"1bbb","shortUrl":"`+os.Getenv("baseUrl")+`1bbb"}`, string(body))
}

func TestGoUrl(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	router := setupRouter(&Env{db: db})

	rows := sqlmock.NewRows([]string{"url"}).AddRow("https://me.coder.tw")
	mock.ExpectQuery("^SELECT.*").WillReturnRows(rows)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1bbb", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 302, w.Code)
	assert.Equal(t, "https://me.coder.tw", w.Result().Header.Get("Location"))
}

func TestGoUrlNotFound(t *testing.T) {

	db, _, _ := sqlmock.New()
	defer db.Close()

	router := setupRouter(&Env{db: db})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1bbb", nil)

	router.ServeHTTP(w, req)
	body, _ := io.ReadAll(w.Body)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "Not Found", string(body))
}
