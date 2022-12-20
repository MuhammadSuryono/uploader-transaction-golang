package sales

import "time"

type SalesRaw struct {
	KDCODE         int         `json:"KD_CODE"`
	MID            int         `json:"MID"`
	MERCHANTNAME   string      `json:"MERCHANT_NAME"`
	BATCHNOEDC     int         `json:"BATCH_NO_EDC"`
	VCARDTYPE      string      `json:"V_CARD_TYPE"`
	BATCHNO        int         `json:"BATCH_NO"`
	APPRVCODE      interface{} `json:"APPRV_CODE"`
	TRXDATE        string      `json:"TRX_DATE"`
	SETTELMENTDATE string      `json:"SETTELMENT_DATE"`
	DAYBATCH       string      `json:"DAYBATCH"`
	COUNT          int         `json:"COUNT"`
	CARDTYPE       string      `json:"CARD_TYPE"`
	TRXAMOUNT      int         `json:"TRXAMOUNT"`
	DISCAMOUNT     string      `json:"DISCAMOUNT"`
	NETTAMOUNT     string      `json:"NETTAMOUNT"`
	DISCRATE       int         `json:"DISC_RATE"`
	CARDNUMBER     string      `json:"CARDNUMBER"`
}

type SalesVolume struct {
	Id                  string    `json:"id"`
	ApproveCode         string    `json:"approve_code"`
	BatchNo             string    `json:"batch_no"`
	BatchNoEdc          string    `json:"batch_no_edc"`
	CardNumber          string    `json:"card_number"`
	CardType            string    `json:"card_type"`
	ConvertType         string    `json:"convert_type"`
	ConverterCardNumber string    `json:"converter_card_number"`
	Count               int       `json:"count"`
	DayBatch            time.Time `json:"day_batch"`
	DiscAmount          float64   `json:"disc_amount"`
	DiscRate            int       `json:"disc_rate"`
	KdCode              string    `json:"kd_code"`
	MerchantName        string    `json:"merchant_name"`
	Mid                 string    `json:"mid"`
	NetAmount           float64   `json:"net_amount"`
	SettlementDate      time.Time `json:"settlement_date"`
	Terminal            string    `json:"terminal"`
	TrxAmount           int       `json:"trx_amount"`
	TrxDate             time.Time `json:"trx_date"`
	VCardType           string    `json:"v_card_type"`
	CardBinTypeId       int       `json:"card_bin_type_id"`
	BrandingId          int       `json:"branding_id"`
	DepartmentCode      string    `json:"department_code"`
	DepartmentUserCode  string    `json:"department_user_code"`
	DepartmentUserName  string    `json:"department_user_name"`
	DivisionCode        string    `json:"division_code"`
	DivisionUserCode    string    `json:"division_user_code"`
	DivisionUserName    string    `json:"division_user_name"`
	SalesCode           string    `json:"sales_code"`
	SalesUserCode       string    `json:"sales_user_code"`
	SalesUserName       string    `json:"sales_user_name"`
	SalesAreaId         int       `json:"sales_area_id"`
}

type MerchantInfo struct {
	Mid                string      `json:"mid"`
	Acc                interface{} `json:"acc"`
	AccNumber          interface{} `json:"acc_number"`
	Cif                interface{} `json:"cif"`
	CifNo              interface{} `json:"cif_no"`
	CoverageArea       string      `json:"coverage_area"`
	CoverageAreaCode   interface{} `json:"coverage_area_code"`
	Department         interface{} `json:"department"`
	DepartmentName     string      `json:"department_name"`
	Division           interface{} `json:"division"`
	DivisionName       string      `json:"division_name"`
	FlagType           string      `json:"flag_type"`
	MerchantName       string      `json:"merchant_name"`
	MidType            string      `json:"mid_type"`
	Sales              interface{} `json:"sales"`
	SalesArea          int         `json:"sales_area"`
	SalesName          string      `json:"sales_name"`
	BrandingId         int         `json:"branding_id"`
	MccDescriptionId   int         `json:"mcc_description_id"`
	MccGroupId         int         `json:"mcc_group_id"`
	DepartmentCode     string      `json:"department_code"`
	DepartmentUserCode string      `json:"department_user_code"`
	DivisionCode       string      `json:"division_code"`
	DivisionUserCode   string      `json:"division_user_code"`
	SalesCode          string      `json:"sales_code"`
	SalesUserCode      string      `json:"sales_user_code"`
}

type CardBin struct {
	Bin               string `json:"bin"`
	BankId            int    `json:"bank_id"`
	Brand             string `json:"brand"`
	CardBinTypeId     int    `json:"card_bin_type_id"`
	CardProductId     int    `json:"card_product_id"`
	Description       string `json:"description"`
	ConvertCardNumber string `json:"convert_card_number"`
}
