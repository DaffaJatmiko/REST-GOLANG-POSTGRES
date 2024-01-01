package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DaffaJatmiko/rest-golang-postgres/models"
	"github.com/DaffaJatmiko/rest-golang-postgres/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Mahasiswa struct {
	Name 						string 				`json:"name"`
	Birth_date  		time.Time			`json:"birth_date"`
	Gender 					string 				`json:"gender"`
	IDJurusan    		uint          `json:"id_jurusan"`
	IDHobi       		uint          `json:"id_hobi"`
	MahasiswaJurusan string    		`json:"mahasiswa_jurusan"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_mahasiswa", r.CreateMahasiswa)
	api.Get("/mahasiswa", r.GetAllMahasiswa)
	api.Get("/get_mahasiswa/:id", r.GetMahasiswaByID)
	api.Delete("/delete_mahasiswa/:id", r.DeleteMahasiswa)

	// Fungsi-fungsi untuk Hobi
	api.Post("/create_hobi", r.CreateHobi)
	api.Get("/hobi", r.GetAllHobi)
	api.Delete("/delete_hobi/:id", r.DeleteHobi)

	// Fungsi-fungsi untuk Jurusan
	api.Post("/create_jurusan", r.CreateJurusan)
	api.Get("/jurusan", r.GetAllJurusan)
	api.Delete("/delete_jurusan/:id", r.DeleteJurusan)
}

func (r *Repository) CreateHobi(context *fiber.Ctx) error {
	hobi := models.Hobi{}
	err := context.BodyParser(&hobi)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&hobi).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create hobi"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "successfully created Hobi",
		"data":    hobi,
	})
	return nil
}


func (r *Repository) GetAllHobi(context *fiber.Ctx) error {
	hobiModels := &[]models.Hobi{}

	err := r.DB.Find(hobiModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "failed to get Hobi"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "successfully fetch hobi",
		"data":    hobiModels,
	})
	return nil
}

func (r *Repository) DeleteHobi(context *fiber.Ctx) error {
	hobiModel := models.Hobi{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(hobiModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete hobi",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Delete successfully",
	})
	return nil
}

func (r *Repository) CreateJurusan(context *fiber.Ctx) error {
	jurusan := models.Jurusan{}
	err := context.BodyParser(&jurusan)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&jurusan).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create jurusan"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "successfully created Jurusan",
		"data":    jurusan,
	})
	return nil
}


func (r *Repository) GetAllJurusan(context *fiber.Ctx) error {
	jurusanModels := &[]models.Jurusan{}

	err := r.DB.Find(jurusanModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "failed to get Jurusan"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "successfully fetch jurusan",
		"data":    jurusanModels,
	})
	return nil
}

func (r *Repository) DeleteJurusan(context *fiber.Ctx) error {
	jurusanModel := models.Jurusan{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(jurusanModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete jurusan",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Delete successfully",
	})
	return nil
}

func (r *Repository) CreateMahasiswa(context *fiber.Ctx) error {
	mahasiswa := Mahasiswa{}
	
	err := context.BodyParser(&mahasiswa)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message" : "request failed"})
		return err
	}
	
	err = r.DB.Create(&mahasiswa).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message" : "could not create mahasiswa"})
			return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message" : "successfully created Mahasiswa"})
	return nil
}

func (r *Repository) GetAllMahasiswa(context *fiber.Ctx) error {
	mahasiswaModels := &[]models.Mahasiswa{}

	err := r.DB.Preload("Jurusan").Preload("Hobi").Find(mahasiswaModels).Error
    if err != nil {
        context.Status(http.StatusBadRequest).JSON(
            &fiber.Map{"message": "failed to get Mahasiswa"})
        return err
    }

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message" : "successfully fetch mahasiswa",
		"data" : mahasiswaModels,
	})
	return nil
}

func (r *Repository) GetMahasiswaByID(context *fiber.Ctx) error {
	id := context.Params("id")
	mahasiswaModel := &models.Mahasiswa{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(mahasiswaModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the id"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Mahasiswa id fetched successfully",
		"data":    mahasiswaModel,
	})
	return nil
}

func (r *Repository) DeleteMahasiswa(context *fiber.Ctx) error {
	mahasiswaModel := models.Mahasiswa{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(mahasiswaModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete mahasiswa",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Delete successfully",
	})
	return nil
}


func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	err = models.MigrateMahasiswa(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}


	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}