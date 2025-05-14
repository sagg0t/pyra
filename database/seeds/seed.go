package main

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"pyra/pkg/db"
	"pyra/pkg/log"
	"pyra/pkg/nutrition"
)

var l = log.NewLogger()

func main() {
	dbConf := db.NewConfig("pgx")
	// dbConf.Attrs.Add("sslmode", fetchEnv("DB_SSLMODE", "disable"))
	dbConn, err := db.New(context.Background(), dbConf, l)
	if err != nil {
		panic(err)
	}

	err = seedFoodProducts(dbConn)
	if err != nil {
		panic(err)
	}
}

func seedFoodProducts(conn db.DBTX) error {
	f, err := os.Open("./database/seeds/products_sample.csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)
	r.Comma = ';'

	foodProducts := make([]nutrition.Product, 0)

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
			l.Debug(err.Error())
			continue
		}
		proteins, err := strconv.ParseFloat(csvRecord[3], 32)
		if err != nil {
			l.Debug(err.Error())
			continue
		}
		fats, err := strconv.ParseFloat(csvRecord[4], 32)
		if err != nil {
			l.Debug(err.Error())
			continue
		}
		carbs, err := strconv.ParseFloat(csvRecord[5], 32)
		if err != nil {
			l.Debug(err.Error())
			continue
		}

		record := nutrition.Product{
			Name: nutrition.ProductName(csvRecord[0]),
			Macro: nutrition.Macro{
				Calories: float32(calories),
				Proteins: float32(proteins),
				Fats:     float32(fats),
				Carbs:    float32(carbs),
			},
			CreatedAt: now,
			UpdatedAt: now,
		}

		foodProducts = append(foodProducts, record)
	}

	ctx := context.Background()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, product := range foodProducts {
		_, err := tx.ExecContext(ctx, `INSERT INTO food_products (
			name, calories, proteins, fats, carbs, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		);`, product.Name, product.Calories, product.Proteins, product.Fats, product.Carbs, product.CreatedAt, product.UpdatedAt)
		if err != nil {
			l.Debug(err.Error())
			continue
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
