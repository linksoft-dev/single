package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney(t *testing.T) {
	money1 := Money(0.91)
	assert.Equal(t, 0.91, money1.GetFloat64())
}
