package storage

import (
	"os"

	"github.com/jinzhu/gorm"
	"gitlab.com/menuxd/api-rest/pkg/ad"
	"gitlab.com/menuxd/api-rest/pkg/bill"
	"gitlab.com/menuxd/api-rest/pkg/category"
	"gitlab.com/menuxd/api-rest/pkg/click"
	"gitlab.com/menuxd/api-rest/pkg/client"
	"gitlab.com/menuxd/api-rest/pkg/dish"
	"gitlab.com/menuxd/api-rest/pkg/order"
	"gitlab.com/menuxd/api-rest/pkg/promotion"
	"gitlab.com/menuxd/api-rest/pkg/question"
	"gitlab.com/menuxd/api-rest/pkg/rating"
	"gitlab.com/menuxd/api-rest/pkg/stay"
	"gitlab.com/menuxd/api-rest/pkg/table"
	"gitlab.com/menuxd/api-rest/pkg/user"
	"gitlab.com/menuxd/api-rest/pkg/waiter"
)

// DBName Database name.
const DBName = "menuxd"

var (
	conn       *gorm.DB
	connString string
)

// createDBSession Create a new connection with the database.
func createDBSession() error {
	var err error
	// pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	5432, "localhost", "postgres", "superadmin", "menuxd")
	conn, err = gorm.Open("postgres", connString)
	if err != nil {
		return err
	}
	return nil
}

// getSession Returns the gorm conn.
func getSession() *gorm.DB {
	if conn == nil {
		createDBSession()
	}
	return conn
}

func migration() error {
	return conn.AutoMigrate(
		&ad.Ad{},
		&bill.Bill{},
		&category.Category{},
		&client.Client{},
		&dish.Dish{},
		&dish.Ingredient{},
		&promotion.Promotion{},
		&table.Table{},
		&user.User{},
		&waiter.Waiter{},
		&order.Order{},
		&order.Item{},
		&order.IngredientSelected{},
		&click.Click{},
		&stay.Stay{},
		&question.Question{},
		&rating.Rating{},
	).Error
}

// InitData initialize the conn.
func InitData() error {
	var err error
	err = createDBSession()
	if err != nil {
		return err
	}

	err = migration()
	if err != nil {
		return err
	}
	return nil
}

// Close the gorm conn.
func Close() error {
	return conn.Close()
}

func init() {
	connString = os.Getenv("DATABASE_URL")
}
