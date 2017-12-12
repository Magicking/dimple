package internal

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/ethereum/go-ethereum/common"
)

type Wrinkle struct {
	gorm.Model
	To     common.Address
	TxHash common.Hash
	Amount big.Int
}

func InsertWrinkle(ctx context.Context, wkl *Wrinkle) error {
	db := DBFromContext(ctx)

	if err := db.Create(wkl).Error; err != nil {
		return err
	}

	return nil
}

func GetLimitedDB(ctx context.Context, limit int) (wrinkles []Wrinkle, err error) {
	db := DBFromContext(ctx)

	cursor := db.Limit(limit).Order("id desc").Find(&wrinkles)
	if cursor.RecordNotFound() {
		log.Println("RecordNotFound: No entry found in database")
		return nil, nil
	}
	if err = cursor.Error; err != nil {
		return nil, fmt.Errorf("Error getting entry: %v", err)
	}
	return wrinkles, nil
}

func InitDatabase(dbDsn string) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	for i := 1; i < 10; i++ {
		db, err = gorm.Open("postgres", dbDsn)
		if err == nil || i == 10 {
			break
		}
		sleep := (2 << uint(i)) * time.Second
		log.Printf("Could not connect to DB: %v", err)
		log.Printf("Waiting %v before retry", sleep)
		time.Sleep(sleep)
	}
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&Wrinkle{}).Error; err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
