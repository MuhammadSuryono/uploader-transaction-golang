package sales

import (
	"encoding/json"
	"fmt"
	database2 "github.com/MuhammadSuryono/module-golang-database/config/database"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"time"
)

func InvokeDataSales(filename string) {
	salesRawFile, _ := os.Open(filename)

	jsonParser := json.NewDecoder(salesRawFile)
	var salesRawData []SalesRaw
	if err := jsonParser.Decode(&salesRawData); err != nil {
		log.Fatalln("Error decode: ", err.Error())
	}
	var salesData []SalesVolume
	log.Println("Loop data", len(salesRawData))
	for _, sales := range salesRawData {

		var mf MerchantInfo
		q := database2.Connection.Table("merchant_info").Where("mid = ?", fmt.Sprintf("%d", sales.MID)).Debug().First(&mf)

		if q.Error == nil {
			log.Println("Validate data sales volume: Start")

			log.Println("Validate data sales volume: Data Merchant", mf)
			dayBatch, _ := time.Parse("13/12/22", sales.DAYBATCH)
			settlementDate, _ := time.Parse("13/12/22", sales.SETTELMENTDATE)
			trxDate, _ := time.Parse("13/12/22", sales.TRXDATE)
			discAmount, _ := strconv.ParseFloat(sales.DISCAMOUNT, 32)
			nettAmount, _ := strconv.ParseFloat(sales.NETTAMOUNT, 32)
			log.Println("Parse data to sales volume")
			var cardBin CardBin
			database2.Connection.Table("card_bin").Where("bin = ?", sales.CARDNUMBER[0:6]).First(&cardBin)
			salesVolume := SalesVolume{
				Id:                  uuid.New().String(),
				ApproveCode:         fmt.Sprintf("%v", sales.APPRVCODE),
				BatchNo:             fmt.Sprintf("%d", sales.BATCHNO),
				BatchNoEdc:          fmt.Sprintf("%d", sales.BATCHNOEDC),
				CardNumber:          sales.CARDNUMBER,
				CardType:            sales.CARDTYPE,
				ConvertType:         "",
				ConverterCardNumber: sales.CARDNUMBER[0:6],
				Count:               sales.COUNT,
				DayBatch:            dayBatch,
				DiscAmount:          discAmount,
				DiscRate:            sales.DISCRATE,
				KdCode:              fmt.Sprintf("%d", sales.KDCODE),
				MerchantName:        mf.MerchantName,
				Mid:                 mf.Mid,
				NetAmount:           nettAmount,
				SettlementDate:      settlementDate,
				Terminal:            "",
				TrxAmount:           sales.TRXAMOUNT,
				TrxDate:             trxDate,
				VCardType:           "",
				CardBinTypeId:       cardBin.CardBinTypeId,
				BrandingId:          mf.BrandingId,
				DepartmentCode:      mf.DepartmentCode,
				DepartmentUserCode:  mf.DepartmentUserCode,
				DepartmentUserName:  mf.DepartmentName,
				DivisionCode:        mf.DivisionCode,
				DivisionUserCode:    mf.DivisionUserCode,
				DivisionUserName:    mf.DivisionName,
				SalesCode:           mf.SalesCode,
				SalesUserCode:       mf.SalesUserCode,
				SalesUserName:       mf.SalesName,
				SalesAreaId:         mf.SalesArea,
			}
			salesData = append(salesData, salesVolume)

			database2.Connection.Table("transaction_sales").Create(&salesVolume)
		}
	}
}

func generateQuestionsMark(n int) []string {
	s := make([]string, 0)
	for i := 0; i < n; i++ {
		s = append(s, fmt.Sprintf("$%d", i+1))
	}
	return s
}
