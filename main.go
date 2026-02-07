package main

import (
	"encoding/json"
	"fendi/modul-03-task/config"
	"fendi/modul-03-task/database"
	"fendi/modul-03-task/handler"
	"fendi/modul-03-task/repository"
	"fendi/modul-03-task/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type StatusResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

func main() {
	var err error

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_, err = os.Stat(".env")
	if err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	conf := config.Config{
		AppPort: viper.GetString("APP_PORT"),
		DBConn:  viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(conf.DBConn)
	if err != nil {
		log.Fatalf("Error: Unable to connect to database, %v", err.Error())
	}
	defer db.Close()

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, categoryRepo)
	productHandler := handler.NewProductHandler(productService)

	checkoutRepo := repository.NewCheckoutRepository(db)
	checkoutService := service.NewCheckoutService(checkoutRepo, productRepo)
	checkoutHandler := handler.NewCheckoutHandler(checkoutService)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if r.URL.Path == "/" {
				w.Header().Set("Content-Type", "application/json")

				json.NewEncoder(w).Encode(StatusResponse{
					Code:   200,
					Status: "OK",
				})
				return
			}
		}

		http.Error(w, "Not Found", http.StatusNotFound)
	})

	http.HandleFunc("/products", productHandler.HandleProduct)

	http.HandleFunc("/products/", productHandler.HandleProductItem)

	http.HandleFunc("/categories", categoryHandler.HandleCategory)

	http.HandleFunc("/categories/", categoryHandler.HandleCategoryItem)

	http.HandleFunc("/checkouts", checkoutHandler.HandleCheckout)

	fmt.Println("Server is up and running")
	fmt.Printf("http://localhost:%s\n", conf.AppPort)

	addr := fmt.Sprintf(":%s", conf.AppPort)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Error: Unable to start server, %v", err.Error())
		return
	}
}
