package main

import (
	"github.com/Yuichi-Kadota/sql_mock_trial/infra"
	"gorm.io/gorm"
)

type Product struct {
	View uint `gorm:"column:view"`
}

type ProductViewers struct {
	UserID    string `gorm:"user_id"`
	ProductID string `gorm:"product_id"`
}

func recordStats(db *gorm.DB, userID, productID int64) (err error) {
	tx := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	updSql := `UPDATE products SET view = views +1`
	updResult := db.Raw(updSql)
	if updResult.Error != nil {
		return updResult.Error
	}

	insertSql := `INSERT INTO product_viewers (user_id,product_id) VALUES(?,?)`
	insertResult := db.Raw(insertSql)
	if insertResult.Error != nil {
		return insertResult.Error
	}

	return
}

func main() {
	db := infra.NewDB()
	if err := recordStats(db, 1 /*some user id*/, 5 /*some product id*/); err != nil {
		panic(err)
	}
}
