package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	// Inisialisasi database SQLMock untuk pengujian
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	// Konfigurasi GORM dengan database SQLMock
	gormDB, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

// Struktur Repository harus didefinisikan pada file main.go
type Repository struct {
	DB *gorm.DB
}

func TestCreateHobi(t *testing.T) {
	// Setup database dan Fiber App untuk testing
	db, _, err := setupTestDB()
	assert.NoError(t, err)
	defer db.Close()

	app := fiber.New()
	repo := Repository{
		DB: db,
	}
	repo.SetupRoutes(app)

	hobiPayload := []byte(`{"nama": "Olahraga"}`)

	req := httptest.NewRequest(http.MethodPost, "/api/create_hobi", bytes.NewBuffer(hobiPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateJurusan(t *testing.T) {
	// Setup database dan Fiber App untuk testing
	db, _, err := setupTestDB()
	assert.NoError(t, err)
	defer db.Close()

	app := fiber.New()
	repo := Repository{
		DB: db,
	}
	repo.SetupRoutes(app)

	jurusanPayload := []byte(`{"nama": "Teknik Informatika"}`)

	req := httptest.NewRequest(http.MethodPost, "/api/create_jurusan", bytes.NewBuffer(jurusanPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateMahasiswa(t *testing.T) {
	// Setup database dan Fiber App untuk testing
	db, _, err := setupTestDB()
	assert.NoError(t, err)
	defer db.Close()

	app := fiber.New()
	repo := Repository{
		DB: db,
	}
	repo.SetupRoutes(app)

	mahasiswaPayload := []byte(`{
		"name": "John Doe",
		"birth_date": "1990-01-01T00:00:00Z",
		"gender": "Male",
		"id_jurusan": 1,
		"id_hobi": 1,
		"mahasiswa_jurusan": "Teknik Informatika"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/create_mahasiswa", bytes.NewBuffer(mahasiswaPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
