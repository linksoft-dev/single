package paghiper

const (
	typeBoleto                pagHiperType = "BOLETO"
	typePix                   pagHiperType = "PIX"
	host                      string       = "https://api.paghiper.com"
	hostPìx                   string       = "https://pix.paghiper.com"
	pagHiperEndPointCreate    string       = "/transaction/create/"
	pagHiperEndPointStatus    string       = "/transaction/status/"
	pagHiperEndPointList      string       = "/transaction/list/"
	pagHiperEndPointCancel    string       = "/transaction/cancel/"
	pagHiperPixEndPointCreate string       = "/invoice/create/"
	pagHiperPixEndPointStatus string       = "/invoice/status/"
	pagHiperPixEndPointList   string       = "/invoice/list/"
	pagHiperPixEndPointCancel string       = "/invoice/cancel/"
	StatusCancelado           string       = "Cancelado"
	StatusCanceled            string       = "canceled"
	StatusPaid                string       = "paid"
	StatusCompleted           string       = "completed"
	StatusAprovado            string       = "Aprovado"
)
