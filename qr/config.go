package qr

import (
	"errors"
	"fmt"
	"github.com/MuhammadSuryono/module-golang-server/http/response"
	"github.com/MuhammadSuryono/module-golang-server/server"
)

var skipIndex = []int{
	2, 9, 10, 11, 12, 13, 14, 15, 33, 34, 35,
}

var indexInteger = []int{
	17, 18, 19, 20, 21, 25,
}

var indexdate = []int{
	14,
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var firstRow = make([]string, 0)
var totalData = 534536
var totalInsert = 0

var dataHeaders = []string{
	"type",
	"mid",
	"mcc",
	"tid",
	"trx_type",
	"mrc_stl",
	"resp_code",
	"type_qr",
	"mdr_amount",
	"tip_amount",
	"gross_amount",
	"trans_amount",
	"net_amount",
	"time",
	"trx_date",
	"branding_id",
	"sales_code",
	"sales_user_code",
	"sales_user_name",
	"department_code",
	"department_user_code",
	"department_user_name",
	"division_code",
	"division_user_code",
	"division_user_name",
}

var dataHeaderCsvValidate = []string{
	"TYPE",
	"MID",
	"MERCHANT_NAME",
	"MCC",
	"TID",
	"TRX_TYPE",
	"MRC_STL",
	"RESP_CODE",
	"TYPE_QR",
	"FLAG1",
	"FLAG2",
	"FLAG3",
	"FLAG4",
	"FLAG5",
	"FLAG6",
	"FLAG7",
	"MDR_AMOUNT",
	"TIP_AMOUNT",
	"GROSS_AMOUNT",
	"TRANS_AMOUNT",
	"NET_AMOUNT",
	"TIME",
	"TRX_DATE",
	"branding_id",
	"sales_code",
	"sales_user_code",
	"sales_user_name",
	"arco_code",
	"arco_user_code",
	"arco_user_name",
	"div_code",
	"div_user_code",
	"div_user_name",
	"sales_area_id",
}

func generateQuestionsMark(n int) []string {
	s := make([]string, 0)
	for i := 0; i < n; i++ {
		if contains(indexdate, i) {
			s = append(s, fmt.Sprintf("TO_DATE($%d, 'DD/MM/YYYY')", i+1))
		} else {
			s = append(s, fmt.Sprintf("$%d", i+1))
		}
	}
	return s
}

func validateHeader(header []string) error {
	if len(dataHeaderCsvValidate) == len(header) {
		for i, _ := range dataHeaderCsvValidate {
			if dataHeaderCsvValidate[i] != header[i] {
				server.Context.JSON(response.BAD_REQUEST_CODE, response.FailureResponse(
					response.BAD_REQUEST_STATUS,
					"Header QR un-validate",
					nil))
				return errors.New("Header QR un-validate")
			}
		}
	}

	if len(dataHeaderCsvValidate) != len(header) {
		server.Context.JSON(response.BAD_REQUEST_CODE, response.FailureResponse(
			response.BAD_REQUEST_STATUS,
			fmt.Sprintf("Header QR requires %d but given %d", len(dataHeaderCsvValidate), len(header)),
			nil))
		return errors.New(fmt.Sprintf("Header QR requires %d but given %d", len(dataHeaderCsvValidate), len(header)))
	}
	return nil
}
