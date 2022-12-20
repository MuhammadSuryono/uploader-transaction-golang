package sales

import (
	"fmt"
	db "github.com/MuhammadSuryono/module-golang-database/config/database"
	"github.com/joho/godotenv"
	"testing"
	"time"
)

func TestInvokeDataSales(t *testing.T) {
	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		panic(fmt.Sprintf("Env file notfound on directory %v", errLoadEnv))
	}
	db.InitConnectionFromEnvironment().CreateNewConnection()
	go InvokeDataSales("SALES_14122022.json")
	go InvokeDataSales("SALES_16122022.json")
	go InvokeDataSales("SALES_17122022.json")
	go InvokeDataSales("SALES_18122022.json")
	go InvokeDataSales("SALES_19122022.json")

	time.Sleep(40 * time.Minute)
}
