package types

import (
	"fmt"
	"github.com/linksoft-dev/single/comps/go/number"
	"strings"
)

type Currency string

const (
	CurrencyBRL Currency = "BRL"
	CurrencyUSD Currency = "USD"
)

// Money typo utilizado para representar valores monetários no sistema
type Money float64

// GetFloat64 retorna o valor de money no formato float64
func (m Money) GetFloat64() float64 {
	v := number.ToFloat64(m, 2)
	return v
}

// GetFloat32 retorna o valor de money no formato float32
func (m Money) GetFloat32() float32 {
	return number.ToFloat32(m, 2)
}

// GetValueWithDecimals retorna o valor de money com o numero de casas decimais indicado
func (m Money) GetValueWithDecimals(decimals int) float64 {
	v := number.ToFloat64(m, decimals)
	return v
}

// GetFormattedBRL retorna o valor monetário em string no formato BRL utilizado no Brasil
func (m Money) GetFormattedBRL() string {
	return m.getFormatted(CurrencyBRL)
}

// GetFormattedUSD retorna o valor monetário em string no formato USD utilizado nos Estados Unidos da America
func (m Money) GetFormattedUSD() string {
	return m.getFormatted(CurrencyUSD)
}

// getFormatted retorna o valor monetário em string de acordo com o tipo de moeda utilizada
func (m Money) getFormatted(currency Currency) string {
	intValue := number.AddThousandToInt(int(m))
	floatValue := number.GetDecimalPart(float64(m))
	switch currency {
	case CurrencyUSD:
		intValue = strings.ReplaceAll(intValue, ".", ",")
		return fmt.Sprintf("$ %s.%s", intValue, floatValue)
	default:
		return fmt.Sprintf("R$ %s,%s", intValue, floatValue)
	}
}
