package paghiper

import (
	"bytes"
	"comps/payment"
	"encoding/json"
	"errors"
	"github.com/linksoft-dev/single/comps/go/validation"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/linksoft-dev/single/comps/go/date"
	"github.com/linksoft-dev/single/comps/go/number"

	"golang.org/x/text/language"
)

// validatePaghiperFormat validates the received data and returns an error if any invalid data is found
func validatePaghiperFormat(billet BilletPaghiperFormat) error {
	v := validation.NewValidation(language.BrazilianPortuguese)

	//required fields
	v.IsFilled("TypeBankSlip", billet.TypeBankSlip, 1, 11)
	validatePaghiperBasicFormat(v, billet)
	if !v.Validated() {
		return nil
	}
	return nil
}

// validatePaghiperPixFormat validates the received data and returns an error if any invalid data is found
func validatePaghiperPixFormat(billet BilletPaghiperFormat) error {
	v := validation.NewValidation(language.BrazilianPortuguese)
	//required fields
	validatePaghiperBasicFormat(v, billet)
	if !v.Validated() {
		return nil
	}
	return nil
}

// validatePaghiperBasicFormat validates the received basic data and returns an error if any invalid data is found
func validatePaghiperBasicFormat(v *validation.Validation, billet BilletPaghiperFormat) {
	//required fields
	v.IsFilled("ApiKey", billet.ApiKey, 1, 50)
	v.IsFilled("OrderId", billet.OrderId, 1, 64)
	v.IsFilled("DaysDueDate", billet.DaysDueDate, 1, 3)
	v.IsFilled("PayerEmail", billet.PayerEmail, 1, 255)
	v.IsFilled("PayerName", billet.PayerName, 1, 255)
	v.IsCpfCnpjValid("PayerCpfCnpj", billet.PayerCpfCnpj)

	for _, item := range billet.Items {
		v.IsFilled("item Id", item.Id, 1, 32)
		v.IsFilled("item Description", item.Description, 1, 255)
		v.IsFilled("item Quantity", strconv.Itoa(item.Quantity), 1, 9999)
		v.IsFilled("item Value", strconv.Itoa(item.Value), 1, 99999)

		//fields range
		v.IsByteLength("item Description", item.Description, 1, 255)
	}
}

// convertToPaghiperFormat converts the billet received into the pagHiper format used to generate the data in JSON format.
func convertToPaghiperFormat(phType pagHiperType, billet payment.Billet) (b BilletPaghiperFormat, err error) {
	//Calculates how many days are left until the ticket expires
	daysDueDate := date.GetDaysBetweenDates(time.Now(), billet.DueDate)
	if daysDueDate < 1 {
		daysDueDate = 1
	}

	// Checking of fines and jutes
	percentageFine := billet.PercentageFine
	latePaymentFine := 0
	if percentageFine > 2 {
		latePaymentFine = 2
	} else if percentageFine < 2 && percentageFine > 0 {
		latePaymentFine = 1
	} else {
		latePaymentFine = 0
	}

	perDayInterest := billet.PerDayInterest != 0

	billetPagHiper := BilletPaghiperFormat{}
	billetPagHiper.OrderId = billet.Id
	billetPagHiper.ApiKey = billet.ApiKey
	billetPagHiper.PayerEmail = billet.ClientEmail
	billetPagHiper.PayerName = billet.ClientName
	billetPagHiper.PayerCpfCnpj = billet.ClientCpfCnpj
	billetPagHiper.PayerPhone = billet.ClientPhone
	billetPagHiper.PayerStreet = billet.ClientAddress
	billetPagHiper.PayerNumber = billet.ClientNumber
	billetPagHiper.PayerComplement = ""
	billetPagHiper.PayerDistrict = billet.ClientDistrict
	billetPagHiper.PayerCity = billet.ClientCity
	billetPagHiper.PayerState = billet.ClientState
	billetPagHiper.PayerZipCode = billet.ClientZipCode
	billetPagHiper.NotificationUrl = billet.NotificationUrl
	if billet.Desconto > 0 {
		billetPagHiper.DiscountCents = int(number.RoundFloat(billet.Desconto*100, 0))
	}
	billetPagHiper.ShippingPriceCents = ""
	billetPagHiper.ShippingMethods = ""
	billetPagHiper.FixedDescription = false
	billetPagHiper.TypeBankSlip = "boletoA4"
	billetPagHiper.PartnersId = billet.PartnersId
	billetPagHiper.DaysDueDate = strconv.Itoa(daysDueDate)
	billetPagHiper.LatePaymentFine = strconv.Itoa(latePaymentFine)
	billetPagHiper.PerDayInterest = perDayInterest
	billetPagHiper.OpenAfterDayDue = 30

	for _, item := range billet.Items {
		itemPagHiper := BilletPaghiperItemFormat{}
		itemPagHiper.Id = item.Id
		itemPagHiper.Description = item.Description
		itemPagHiper.Quantity = int(item.Quantity)
		itemPagHiper.Value = int(number.RoundFloat(item.Value*100, 0))
		billetPagHiper.Items = append(billetPagHiper.Items, itemPagHiper)
	}

	switch phType {
	case typePix:
		err = validatePaghiperPixFormat(billetPagHiper)
	default:
		err = validatePaghiperFormat(billetPagHiper)
	}
	if err != nil {
		return BilletPaghiperFormat{}, err
	}

	return billetPagHiper, nil
}

// Validates the apikey paghiper in order to verify that it starts with the character apk_ and contains the correct number of characters.
func validateApikey(paghiperApiKey string) error {
	apiKeyLength := len(paghiperApiKey)
	if apiKeyLength < 3 || apiKeyLength > 50 {
		return errors.New("Invalid PAGHIPER apiKey, review apiKey and inform a valid value")
	}

	if !strings.HasPrefix(paghiperApiKey, "apk_") {
		return errors.New("Invalid PAGHIPER format apiKey, review apiKey and inform a valid value")
	}

	return nil
}

// pagHiperRequest converts the data into the format used in requests (json), communicates with PagHiper, processes and returns the result.
func pagHiperRequest(phType pagHiperType, endPoint string, data interface{}) (interface{}, error) {
	hostPagHiper := host
	if phType == typePix {
		hostPagHiper = hostPìx
	}
	urlRequest := hostPagHiper + endPoint
	requestBody, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return "Error", err
	}
	timeout := time.Duration(20 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest("POST", urlRequest, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Println(err)
		return "Error", err
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return "Error", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "Error", err
	}

	if resp.StatusCode != 201 {
		log.Println("PAGHIPER - Reject: ", endPoint, string(body))
		return "PAGHIPER - Reject: " + endPoint, errors.New(string(body))
	}
	rec := jsonRequest
	json.Unmarshal(body, &rec)

	return rec, nil
}

// Convert the received payment slip into PagHiper format and send a request to be created at the payment gateway PAGHIPER
func Create(billet payment.Billet) (payment.Billet, error) {
	return CreateByType(typeBoleto, billet)
}

// CreateByType the received payment slip into PagHiper format and send a request by type to be created at the payment gateway PAGHIPER
func CreateByType(phType pagHiperType, billet payment.Billet) (payment.Billet, error) {
	billetPh, err := convertToPaghiperFormat(phType, billet)
	if err != nil {
		log.Println("error", err)
		return payment.Billet{}, err
	}

	err = validateApikey(billetPh.ApiKey)

	if err != nil {
		log.Println(err)
		return payment.Billet{}, err
	}

	//IF YOU HAVE TRANSACTONID YOU MUST PERFORM THE CONSULTATION
	//IF THE CONSULTATION RETURNS MATURITY BEYOND 30 DAYS, THE BILLET MUST BE RECREATED (CANCEL THE FIRST AND CREATE ANOTHER WITH THE VALUE AND NEW VALIDITY DATE)
	//IF YOU ARE STILL WITHIN THE PAYMENT LIMIT, RETURN THE EXISTING BILLET THAT MUST ALREADY HAVE THE ACCESS URL
	if billet.TransactionId != "" {
		response, err := GetStatusByType(phType, billet.Token, billet.ApiKey, billet.TransactionId)
		if err != nil {
			log.Println("error", err)
			return payment.Billet{}, err
		}

		if response.Success {
			dueDate := date.FromString(response.DueDate)
			// if you have not exceeded the 30-day limit after expiration, return the existing billet
			if date.GetDaysBetweenDates(time.Now(), dueDate) > -30 {
				return billet, nil
			} else { // if you exceeded the 30-day limit after expiration, cancel the billet already created and create a new billet with an updated expiration
				statusCancel, err := CancelByType(phType, billet.Token, billet.ApiKey, billet.TransactionId)
				if err != nil {
					log.Println("error", err)
					return payment.Billet{}, err
				} else if statusCancel != StatusCanceled {
					return payment.Billet{}, errors.New("The bank ticket cannot be canceled, please contact your system administrator!")
				}
			}
		}
	}
	switch phType {
	case typePix:
		response, err := pagHiperRequest(typePix, pagHiperPixEndPointCreate, billetPh)
		if err != nil {
			return payment.Billet{}, err
		}
		responseMap := response.(map[string]interface{})["pix_create_request"].(map[string]interface{})
		billet.TransactionId = responseMap["transaction_id"].(string)
		billet.PixUrl = responseMap["pix_code"].(map[string]interface{})["pix_url"].(string)
		billet.QrcodeBase64 = responseMap["pix_code"].(map[string]interface{})["qrcode_base64"].(string)
		billet.QrcodeImageUrl = responseMap["pix_code"].(map[string]interface{})["qrcode_image_url"].(string)
		billet.Emv = responseMap["pix_code"].(map[string]interface{})["emv"].(string)
		billet.BacenUrl = responseMap["pix_code"].(map[string]interface{})["bacen_url"].(string)
	default:
		response, err := pagHiperRequest(phType, pagHiperEndPointCreate, billetPh)
		if err != nil {
			return payment.Billet{}, err
		}
		responseMap := response.(map[string]interface{})["create_request"].(map[string]interface{})
		billet.TransactionId = responseMap["transaction_id"].(string)
		billet.PdfUrl = responseMap["bank_slip"].(map[string]interface{})["url_slip_pdf"].(string)
	}

	return billet, nil
}

// GetStatus checks and returns the billet status.
func GetStatus(paghiperToken, paghiperApikey, paghiperTransactionId string) (Response, error) {
	return GetStatusByType(typeBoleto, paghiperToken, paghiperApikey, paghiperTransactionId)
}

// GetStatusByType checks and returns the data status by type.
func GetStatusByType(phType pagHiperType, paghiperToken, paghiperApikey, paghiperTransactionId string) (Response, error) {
	err := validateApikey(paghiperApikey)
	if err != nil {
		log.Println(err)
		return Response{Success: false}, err
	}

	type data struct {
		Token         string `json:"token"`
		ApiKey        string `json:"apiKey"`
		TransactionId string `json:"transaction_id"`
	}

	dados := data{paghiperToken, paghiperApikey, paghiperTransactionId}

	endPointStatus := pagHiperEndPointStatus
	if phType == typePix {
		endPointStatus = pagHiperPixEndPointStatus
	}
	response, err := pagHiperRequest(phType, endPointStatus, dados)
	if err != nil {
		log.Println(err)
		return Response{Success: false}, err
	}
	responseMap := response.(map[string]interface{})["status_request"].(map[string]interface{})
	responseStatus := Response{
		Success:               true,
		Result:                responseMap["result"].(string),
		ResponseMessage:       responseMap["response_message"].(string),
		PagHiperTransactionId: paghiperTransactionId,
		Status:                responseMap["status"].(string),
		StatusDate:            responseMap["status_date"].(string),
		DueDate:               responseMap["due_date"].(string),
		ValueCents:            strconv.FormatFloat(responseMap["value_cents"].(float64), 'f', 0, 64),
	}
	if phType == typePix {
		responseStatus.PixCode.PixUrl = responseMap["pix_code"].(map[string]interface{})["pix_url"].(string)
		responseStatus.PixCode.QrcodeBase64 = responseMap["pix_code"].(map[string]interface{})["qrcode_base64"].(string)
		responseStatus.PixCode.QrcodeImageUrl = responseMap["pix_code"].(map[string]interface{})["qrcode_image_url"].(string)
		responseStatus.PixCode.Emv = responseMap["pix_code"].(map[string]interface{})["emv"].(string)
		responseStatus.PixCode.BacenUrl = responseMap["pix_code"].(map[string]interface{})["bacen_url"].(string)
	}
	return responseStatus, nil
}

// GetList returns the list of paid slips filtered by payment date.
func GetList(paghiperToken, paghiperApikey, fromDateString, toDateString string) ([]Response, error) {
	return GetListByType(typeBoleto, paghiperToken, paghiperApikey, fromDateString, toDateString)
}

// GetListByType returns the list of paid slips by type filtered by payment date.
func GetListByType(phType pagHiperType, paghiperToken, paghiperApikey, fromDateString, toDateString string) ([]Response, error) {
	fromDate := time.Now()
	toDate := time.Now()
	err := validateApikey(paghiperApikey)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if fromDateString != "" {
		fromDate = date.FromString(fromDateString)
		if fromDate.IsZero() {
			log.Println(err)
			return []Response{}, err
		}
	}
	if toDateString != "" {
		toDate = date.FromString(toDateString)
		if toDate.IsZero() {
			log.Println(err)
			return []Response{}, err
		}
	}

	type Data struct {
		Token       string `json:"token"`
		ApiKey      string `json:"apiKey"`
		InitialDate string `json:"initial_date"`
		FinalDate   string `json:"final_date"`
		// find by payment date
		FilterDate string `json:"filter_date"`
		Status     string `json:"status"`
	}
	data := Data{
		Token:       paghiperToken,
		ApiKey:      paghiperApikey,
		InitialDate: fromDate.Format("2006-01-02"),
		FinalDate:   toDate.Format("2006-01-02"),
		FilterDate:  "paid_date",
		Status:      "paid",
	}

	endPointList := pagHiperEndPointList
	if phType == typePix {
		endPointList = pagHiperPixEndPointList
	}

	response, err := pagHiperRequest(phType, endPointList, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	responsePaghiperList := ResponsePaghiperList{}

	responseMap := response.(map[string]interface{})["transaction_list_request"]
	responsePaghiperList.Result = responseMap.(map[string]interface{})["result"].(string)
	responsePaghiperList.ResponseMessage = responseMap.(map[string]interface{})["response_message"].(string)

	for _, item := range responseMap.(map[string]interface{})["transaction_list"].([]interface{}) {
		returnItem := Response{}
		itemMap := item.(map[string]interface{})
		returnItem.TransactionId = itemMap["transaction_id"].(string)
		returnItem.OrderId = itemMap["order_id"].(string)
		returnItem.Status = itemMap["status"].(string)
		returnItem.StatusDate = itemMap["status_date"].(string)
		returnItem.DueDate = itemMap["due_date"].(string)
		returnItem.ValueCents = strconv.FormatFloat(itemMap["value_cents"].(float64), 'f', 0, 64)
		returnItem.DiscountCents = strconv.FormatFloat(itemMap["discount_cents"].(float64), 'f', 0, 64)
		returnItem.ValueFeeCents = strconv.FormatFloat(itemMap["value_fee_cents"].(float64), 'f', 0, 64)
		returnItem.PayerEmail = itemMap["payer_email"].(string)
		returnItem.PayerName = itemMap["payer_name"].(string)
		returnItem.PayerCpfCnpj = itemMap["payer_cpf_cnpj"].(string)
		returnItem.PayerPhone = itemMap["payer_phone"].(string)
		returnItem.CreateDate = itemMap["create_date"].(string)
		returnItem.PaidDate = itemMap["paid_date"].(string)
		if phType == typePix {
			returnItem.PixCode.PixUrl = itemMap["pix_code"].(map[string]interface{})["pix_url"].(string)
			returnItem.PixCode.QrcodeBase64 = itemMap["pix_code"].(map[string]interface{})["qrcode_base64"].(string)
			returnItem.PixCode.QrcodeImageUrl = itemMap["pix_code"].(map[string]interface{})["qrcode_image_url"].(string)
			returnItem.PixCode.Emv = itemMap["pix_code"].(map[string]interface{})["emv"].(string)
			returnItem.PixCode.BacenUrl = itemMap["pix_code"].(map[string]interface{})["bacen_url"].(string)
		}

		responsePaghiperList.TransactionList = append(responsePaghiperList.TransactionList, returnItem)
	}

	return responsePaghiperList.TransactionList, nil
}

// Cancel requests the cancellation of a billet and returns the response of the requested transaction
func Cancel(pagHiperToken, pagHiperApikey, pagHiperTransacaoId string) (string, error) {
	return CancelByType(typeBoleto, pagHiperToken, pagHiperApikey, pagHiperTransacaoId)
}

// CancelByType requests the cancellation of a data by type and returns the response of the requested transaction
func CancelByType(phType pagHiperType, pagHiperToken, pagHiperApikey, pagHiperTransacaoId string) (string, error) {
	err := validateApikey(pagHiperApikey)
	if err != nil {
		log.Println(err)
		return "Error", err
	}

	type Data struct {
		Token         string `json:"token"`
		ApiKey        string `json:"apiKey"`
		TransactionId string `json:"transaction_id"`
		Status        string `json:"status"`
	}

	data := Data{
		Token:         pagHiperToken,
		ApiKey:        pagHiperApikey,
		TransactionId: pagHiperTransacaoId,
		Status:        "canceled",
	}
	endPointCancel := pagHiperEndPointCancel
	if phType == typePix {
		endPointCancel = pagHiperPixEndPointCancel
	}
	_, err = pagHiperRequest(phType, endPointCancel, data)
	responseString := StatusCanceled
	if err != nil {
		if !strings.Contains(err.Error(), "status atual do pedido é canceled") {
			log.Println(err)
			return "Error", err
		}
	}

	return responseString, nil
}

// CreatePix the received payment slip into PagHiper pix format and send a request to be created at the payment gateway PAGHIPER
func CreatePix(billet payment.Billet) (payment.Billet, error) {
	return CreateByType(typePix, billet)
}

// GetStatusPix checks and returns the pix status.
func GetStatusPix(paghiperToken, paghiperApikey, paghiperTransactionId string) (Response, error) {
	return GetStatusByType(typePix, paghiperToken, paghiperApikey, paghiperTransactionId)
}

// GetListPix returns the list of pix slips filtered by payment date.
func GetListPix(paghiperToken, paghiperApikey, fromDateString, toDateString string) ([]Response, error) {
	return GetListByType(typePix, paghiperToken, paghiperApikey, fromDateString, toDateString)
}

// CancelPix requests the cancellation of a pix and returns the response of the requested transaction
func CancelPix(pagHiperToken, pagHiperApikey, pagHiperTransacaoId string) (string, error) {
	return CancelByType(typePix, pagHiperToken, pagHiperApikey, pagHiperTransacaoId)
}
