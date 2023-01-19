package controller

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/thanh-vt/splash-inventory-service/internal/middleware"
	"github.com/thanh-vt/splash-inventory-service/internal/model"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

func GetAllSupplier(w http.ResponseWriter, r *http.Request) {
	suppliers := &[]model.Supplier{}
	DB.Find(&suppliers)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, &suppliers)
}
func GetSupplier(w http.ResponseWriter, r *http.Request) {
	existedSupplier := &model.Supplier{}
	//r.URL.Query().Get()
	code := chi.URLParam(r, "code")
	result := DB.First(&existedSupplier, "code = ?", code)
	if result.Error != nil {
		middleware.NotFound(w, r, errors.New("supplier not found"))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, &existedSupplier)
}

func CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier model.Supplier
	if err := render.DecodeJSON(r.Body, &supplier); err != nil {
		middleware.BadRequest(w, r, errors.New("supplier request is invalid"))
		return
	}
	err := DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Create(&supplier).Error; err != nil {
			// return any error will roll back
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, &supplier)
}

func UpdateCartProduct(w http.ResponseWriter, r *http.Request) {
	var supplier model.Supplier
	if err := render.DecodeJSON(r.Body, &supplier); err != nil {
		middleware.BadRequest(w, r, errors.New("supplier request is invalid"))
		return
	}
	existedSupplier := &model.Supplier{}
	code := chi.URLParam(r, "code")
	result := DB.First(&existedSupplier, "code = ?", code)
	if result.Error != nil {
		middleware.NotFound(w, r, errors.New("supplier not found"))
		return
	}
	existedSupplier.Name = supplier.Name
	err := DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Save(&existedSupplier).Error; err != nil {
			// return any error will roll back
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, &existedSupplier)
}

func DeleteCartProduct(w http.ResponseWriter, r *http.Request) {
	existedSupplier := &model.Supplier{}
	code := chi.URLParam(r, "code")
	result := DB.First(&existedSupplier, "code = ?", code)
	if result.Error != nil {
		middleware.NotFound(w, r, errors.New("supplier not found"))
		return
	}
	err := DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Delete(&existedSupplier).Error; err != nil {
			// return any error will roll back
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
	render.Status(r, http.StatusOK)
}
