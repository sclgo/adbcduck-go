package types

import (
	"fmt"

	"github.com/apache/arrow-go/v18/arrow/decimal"
)

// Decimal extends the arrow.Decimal[T] with a record of the source scale value
// which adds the ability to print the number
type Decimal[T decimal.DecimalTypes] struct {
	decimal.Num[T]
	// The original scale of the decimal
	Scale int32
}

// String implements fmt.Stringer
func (d Decimal[T]) String() string {
	return d.ToString(d.Scale)
}

var _ fmt.Stringer = Decimal[decimal.Decimal128]{}
