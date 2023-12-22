package paghiper

// EXEMPLOS DE USO DA BIBLIOTECA
// nao pode ser testes porque nao se pode deixar os testes dependente de comunicacao externa

import (
	"comps/payment"
	// Paghiper "comps/payment/paghiper"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	InvalidPaghiperApikey        = ""
	InvalidPaghiperToken         = ""
	InvalidPaghiperTransactionId = ""
	ValidPaghiperApikey          = ""
	ValidPaghiperToken           = ""
	ValidPaghiperTransactionId   = ""
)

func TestPagHiperCreate(t *testing.T) {
	testName := "PagHiper Create test: "

	// create a invalid billet
	billet := payment.Billet{}
	_, err := Create(billet)
	assert.NotEqual(t, nil, err, "Error: error is expected for invalid billet")

	// create invalid fields
	billet.Id = "323303351111111111111111111111111323303351111111111111111111111111"
	billet.ClientEmail = "323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303323303351111111111111111111111111323303351111111111111111111111111351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111noreply@sigeflex.com"
	billet.ClientName = "323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303323303351111111111111111111111111323303351111111111111111111111111351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111noreply@sigeflex.com"
	billet.ClientCpfCnpj = "111111111111111@sigeflex.com"

	_, err = Create(billet)
	assert.NotEqual(t, nil, err, "%sError: error is expected for invalid billet", testName)

	//Create a valid billet
	billet.ApiKey = InvalidPaghiperApikey
	billet.Token = InvalidPaghiperToken
	billet.Id = "32330335"
	billet.ClientName = "Jose Inventado"
	billet.ClientCpfCnpj = "13032513000126"
	billet.ClientPhone = "1140638785"
	billet.ClientEmail = "noreply@sigeflex.com"
	billet.ClientAddress = "Av Brigadeiro Faria Lima"
	billet.ClientNumber = "1461"
	billet.ClientDistrict = "Jardim Paulistano"
	billet.ClientCity = "Sao Paulo"
	billet.ClientState = "SP"
	billet.ClientZipCode = "01452002"
	billet.Value = 10.00
	billet.DueDate = time.Now().Add(time.Hour * 24 * 3) //after 3 days

	//item 1
	billetItem := payment.BilletItem{}
	billetItem.Id = "1"
	billetItem.Description = "piscina de bolinha"
	billetItem.Quantity = 1
	billetItem.Value = 10.12
	billet.Items = append(billet.Items, billetItem)

	//item 2
	billetItem = payment.BilletItem{}
	billetItem.Id = "2"
	billetItem.Description = "pula pula"
	billetItem.Quantity = 2
	billetItem.Value = 20.00
	billet.Items = append(billet.Items, billetItem)

	//item 3
	billetItem = payment.BilletItem{}
	billetItem.Id = "3"
	billetItem.Description = "mala de viagem"
	billetItem.Quantity = 3.5
	billetItem.Value = 40.00
	billet.Items = append(billet.Items, billetItem)

	// isso está comentado pois nao deve se comunicar e fazer operacao com API externa dentro dos testes
	// mas deve ficar comentando pra ver como consumir o servico
	//billet, err = Create(billet)

	// os dados abaixo estao sendo mockados simunando um retorno da API
	billet.TransactionId = "123"
	billet.PdfUrl = "test"
	assert.NotEqual(t, "", billet.TransactionId, "%sbillet creation did not return  TransactionId", testName)
	assert.NotEqual(t, "", billet.PdfUrl, "%sbillet creation did not return PdfUrl", testName)
}

func TestPagHiperStatus(t *testing.T) {
	testName := "PagHiper Status test: "

	//finding with invalid apiKey
	_, err := GetStatus(ValidPaghiperToken, InvalidPaghiperApikey, InvalidPaghiperTransactionId)
	assert.NotEqual(t, nil, err, "%sError: error is expected for non-existent transaction", testName)

	//finding with invalid token
	_, err = GetStatus(InvalidPaghiperToken, ValidPaghiperApikey, InvalidPaghiperTransactionId)
	assert.NotEqual(t, nil, err, "%sError: error is expected for non-existent transaction", testName)

	//finding inexistent transaction
	_, err = GetStatus(ValidPaghiperToken, ValidPaghiperApikey, InvalidPaghiperTransactionId)
	assert.NotEqual(t, nil, err, "%sError: error is expected for non-existent transaction", testName)

	//finding valid transaction
	responseMap, err := GetStatus(ValidPaghiperToken, ValidPaghiperApikey, ValidPaghiperTransactionId)
	assert.Nil(t, err, "Error: %v", err)
	assert.Equal(t, ValidPaghiperTransactionId, responseMap.PagHiperTransactionId, "%squery dont return expected Id", testName)
	assert.Equal(t, "success", responseMap.Result, "%sstatus returned by the query other than expected", testName)
	assert.Equal(t, "transacao encontrada", responseMap.ResponseMessage, "%sstatus returned by the query other than expected", testName)
	assert.Equal(t, "canceled", responseMap.Status, "%sstatus returned by the query other than expected", testName)
}

// TestPagHiperList esse teste está pegando uma transacao(id da transacao PT96FVD2A9YWRWJQ) que ja existe e checando se
// retorna corretamente os dados
func TestPagHiperList(t *testing.T) {
	testName := "PagHiper List test: "

	resp, err := GetList(ValidPaghiperToken,
		ValidPaghiperApikey, "2019-04-12", "2019-04-12")
	assert.Nil(t, err, "Error: %v", err)
	assert.NotEqual(t, 0, len(resp), "%snot returns list", testName)
	assert.Equal(t, "PT96FVD2A9YWRWJQ", resp[0].TransactionId, "%stransaction_id different than expected", testName)
	assert.Equal(t, "57016", resp[0].OrderId, "%sorder_id different than expected", testName)
	assert.Equal(t, "completed", resp[0].Status, "%sstatus different than expected", testName)
	assert.Equal(t, "2019-04-12 08:10:43", resp[0].StatusDate, "%sstatus_date different than expected", testName)
	assert.Equal(t, "2019-04-12", resp[0].DueDate, "%sdue_date different than expected", testName)
	assert.Equal(t, "16373", resp[0].ValueCents, "%svalue_cents different than expected", testName)
	assert.Equal(t, "0", resp[0].DiscountCents, "%sdiscount_cents different than expected", testName)
	assert.Equal(t, "199", resp[0].ValueFeeCents, "%svalue_fee_cents different than expected", testName)
	assert.Equal(t, "noreply@sigeflex.com", resp[0].PayerEmail, "%spayer_email different than expected", testName)
	assert.Equal(t, "POTIGUAR DIESEL LTDA ME", resp[0].PayerName, "%spayer_name different than expected", testName)
	assert.Equal(t, "12165870000108", resp[0].PayerCpfCnpj, "%spayer_cpf_cnpj different than expected", testName)
	assert.Equal(t, "33167340", resp[0].PayerPhone, "%spayer_phone different than expected", testName)
	assert.Equal(t, "2019-04-11 09:06:27", resp[0].CreateDate, "%screate_date different than expected", testName)
	assert.Equal(t, "2019-04-12 05:06:43", resp[0].PaidDate, "%spaid_date different than expected", testName)
}

func TestPagHiperCancel(t *testing.T) {
	testName := "PagHiper Cancel test: "

	phToken := ""
	phTransactionId := ""
	phApiKey := ValidPaghiperApikey

	//testing using a invalid token and invalid transaction Id
	phToken = ""

	phTransactionId = "BPV661O7AVLORCN5"

	_, err := Cancel(phToken, phApiKey, phTransactionId)
	assert.NotEqual(t, nil, err, "%sError: error is espected for (token ou apiKey inválidos)", testName)

	//testing using a valid token and invalid transaction Id
	phToken = ValidPaghiperToken

	_, err = Cancel(phToken, phApiKey, phTransactionId)
	assert.NotEqual(t, nil, err, "%sError: error is expected for invalid transaction_id", testName)

	//test with valid Id with canceled operation
	// isso está comentado pois nao deve se comunicar e fazer operacao com API externa dentro dos testes
	// mas deve ficar comentando pra ver como consumir o servico
	//respose, err := Cancel(phToken, phApiKey, phTransactionId)
}

func TestPagHiperCreatePix(t *testing.T) {
	testName := "PagHiper Create test: "

	// create a invalid billet
	billet := payment.Billet{}
	_, err := CreatePix(billet)
	assert.NotEqual(t, nil, err, "Error: error is expected for invalid billet")

	// create invalid fields
	billet.Id = "323303351111111111111111111111111323303351111111111111111111111111"
	billet.ClientEmail = "323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303323303351111111111111111111111111323303351111111111111111111111111351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111noreply@sigeflex.com"
	billet.ClientName = "323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303323303351111111111111111111111111323303351111111111111111111111111351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111323303351111111111111111111111111noreply@sigeflex.com"
	billet.ClientCpfCnpj = "111111111111111@sigeflex.com"

	_, err = CreatePix(billet)
	assert.NotEqual(t, nil, err, "%sError: error is expected for invalid billet", testName)

	//Create a valid billet
	billet.ApiKey = InvalidPaghiperApikey
	billet.Token = InvalidPaghiperToken
	billet.Id = "32330335"
	billet.ClientName = "Jose Inventado"
	billet.ClientCpfCnpj = "13032513000126"
	billet.ClientPhone = "1140638785"
	billet.ClientEmail = "noreply@sigeflex.com"
	billet.ClientAddress = "Av Brigadeiro Faria Lima"
	billet.ClientNumber = "1461"
	billet.ClientDistrict = "Jardim Paulistano"
	billet.ClientCity = "Sao Paulo"
	billet.ClientState = "SP"
	billet.ClientZipCode = "01452002"
	billet.Value = 10.00
	billet.DueDate = time.Now().Add(time.Hour * 24 * 3) //after 3 days

	//item 1
	billetItem := payment.BilletItem{}
	billetItem.Id = "1"
	billetItem.Description = "piscina de bolinha"
	billetItem.Quantity = 1
	billetItem.Value = 10.12
	billet.Items = append(billet.Items, billetItem)

	//item 2
	billetItem = payment.BilletItem{}
	billetItem.Id = "2"
	billetItem.Description = "pula pula"
	billetItem.Quantity = 2
	billetItem.Value = 20.00
	billet.Items = append(billet.Items, billetItem)

	//item 3
	billetItem = payment.BilletItem{}
	billetItem.Id = "3"
	billetItem.Description = "mala de viagem"
	billetItem.Quantity = 3.5
	billetItem.Value = 40.00
	billet.Items = append(billet.Items, billetItem)

	// isso está comentado pois nao deve se comunicar e fazer operacao com API externa dentro dos testes
	// mas deve ficar comentando pra ver como consumir o servico
	//billet, err = CreatePix(billet)

	// os dados abaixo estao sendo mockados simunando um retorno da API
	billet.TransactionId = "123"
	billet.PdfUrl = "test"
	assert.NotEqual(t, "", billet.TransactionId, "%sbillet creation did not return  TransactionId", testName)
	assert.NotEqual(t, "", billet.PdfUrl, "%sbillet creation did not return PdfUrl", testName)
}
