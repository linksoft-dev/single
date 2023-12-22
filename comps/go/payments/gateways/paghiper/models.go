package paghiper

type pagHiperType string

type BilletPaghiperFormat struct {
	ApiKey             string                     `json:"apiKey"`
	OrderId            string                     `json:"order_id"`
	PayerEmail         string                     `json:"payer_email"`
	PayerName          string                     `json:"payer_name"`
	PayerCpfCnpj       string                     `json:"payer_cpf_cnpj"`
	PayerPhone         string                     `json:"payer_phone"`
	PayerStreet        string                     `json:"payer_street"`
	PayerNumber        string                     `json:"payer_number"`
	PayerComplement    string                     `json:"payer_complement"`
	PayerDistrict      string                     `json:"payer_district"`
	PayerCity          string                     `json:"payer_city"`
	PayerState         string                     `json:"payer_state"`
	PayerZipCode       string                     `json:"payer_zip_code"`
	NotificationUrl    string                     `json:"notification_url"`
	DiscountCents      int                        `json:"discount_cents"`
	ShippingPriceCents string                     `json:"shipping_price_cents"`
	ShippingMethods    string                     `json:"shipping_methods"`
	FixedDescription   bool                       `json:"fixed_description"`
	TypeBankSlip       string                     `json:"type_bank_slip"`
	PartnersId         string                     `json:"partners_id"`
	DaysDueDate        string                     `json:"days_due_date"`
	LatePaymentFine    string                     `json:"late_payment_fine"`
	PerDayInterest     bool                       `json:"per_day_interest"`
	OpenAfterDayDue    int                        `json:"open_after_day_due"`
	Items              []BilletPaghiperItemFormat `json:"items"`
}

type BilletPaghiperItemFormat struct {
	Id          string `json:"item_id"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Value       int    `json:"price_cents"`
}

var jsonRequest map[string]interface{}

type ResponsePaghiperList struct {
	Result          string     `json:"result"`
	ResponseMessage string     `json:"response_message"`
	TransactionList []Response `json:"transaction_list"`
}

type Response struct {
	TransactionId         string `json:"transaction_id"`
	OrderId               string `json:"order_id"`
	Status                string `json:"status"`
	StatusDate            string `json:"status_date"`
	PagHiperTransactionId string
	DueDate               string `json:"due_date"`
	ValueCents            string `json:"value_cents"`
	DiscountCents         string `json:"discount_cents"`
	ValueFeeCents         string `json:"value_fee_cents"`
	PayerEmail            string `json:"payer_email"`
	PayerName             string `json:"payer_name"`
	PayerCpfCnpj          string `json:"payer_cpf_cnpj"`
	PayerPhone            string `json:"payer_phone"`
	CreateDate            string `json:"create_date"`
	PaidDate              string `json:"paid_date"`
	Success               bool
	Result                string
	ResponseMessage       string
	//campos pix
	PixCode struct {
		QrcodeBase64   string `json:"qrcode_base64"`
		QrcodeImageUrl string `json:"qrcode_image_url"`
		Emv            string `json:"emv"`
		PixUrl         string `json:"pix_url"`
		BacenUrl       string `json:"bacen_url"`
	} `json:"pix_code"`
	HttpCode string `json:"http_code"`
}
