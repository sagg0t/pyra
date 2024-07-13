package main

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"

	"pyra/pkg/db"
	"pyra/pkg/foodproducts"
)

func main() {
	dbConf := db.NewConfig("postgres")
	// dbConf.Attrs.Add("sslmode", fetchEnv("DB_SSLMODE", "disable"))
	dbConn, err := pgx.Connect(context.Background(), dbConf.String())
	if err != nil {
		panic(err)
	}

	err = seedFoodProducts(dbConn)
	if err != nil {
		panic(err)
	}
}

func seedFoodProducts(conn *pgx.Conn) error {
	f, err := os.Open("./database/seeds/новое питание - продукты.csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)
	r.Comma = ';'

	foodProducts := make([]foodproducts.FoodProduct, 0)

	for {
		csvRecord, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		now := time.Now()
		calories, err := strconv.ParseFloat(csvRecord[2], 32)
		if err != nil {
			log.Println(err)
			continue
		}
		proteins, err := strconv.ParseFloat(csvRecord[3], 32)
		if err != nil {
			log.Println(err)
			continue
		}
		fats, err := strconv.ParseFloat(csvRecord[4], 32)
		if err != nil {
			log.Println(err)
			continue
		}
		carbs, err := strconv.ParseFloat(csvRecord[5], 32)
		if err != nil {
			log.Println(err)
			continue
		}

		record := foodproducts.FoodProduct{
			Name:      csvRecord[0],
			Calories:  float32(calories),
			Proteins:  float32(proteins),
			Fats:      float32(fats),
			Carbs:     float32(carbs),
			CreatedAt: now,
			UpdatedAt: now,
		}

		foodProducts = append(foodProducts, record)
	}

	ctx := context.Background()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	for _, product := range foodProducts {
		_, err := tx.Exec(ctx, `INSERT INTO food_products (
			name, calories, proteins, fats, carbs, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		);`, product.Name, product.Calories, product.Proteins, product.Fats, product.Carbs, product.CreatedAt, product.UpdatedAt)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
