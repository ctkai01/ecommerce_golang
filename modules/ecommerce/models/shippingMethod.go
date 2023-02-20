package models

import (
	"errors"
	"time"
)

// type PaymentTypeValue int

// const (
// 	CODPayment PaymentTypeValue = iota
// 	CREDIT_CARD
// )

// var valuePaymentTypes = [2]string{"COD", "CREDIT_CARD"}

// func (item *PaymentTypeValue) String() string {
// 	 return valuePaymentTypes[*item]
// }

// func parseStr2PaymentStatus(s string) (PaymentTypeValue, error) {
// 	for i := range valuePaymentTypes {
// 		if valuePaymentTypes[i] == s {
// 			return PaymentTypeValue(i), nil
// 		}
// 	}
// 	return PaymentTypeValue(0), errors.New("Invalid payment type string")
// }

// // Read from DB to Server
// func (pt *PaymentTypeValue) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)

// 	if !ok {
// 		return errors.New(fmt.Sprintf("Fail to scan data from sql: %s", value))
// 	}

// 	v, err := parseStr2PaymentStatus(string(bytes))

// 	if err != nil {
// 		return errors.New(fmt.Sprintf("Fail to scan data from sql: %s", value))
// 	}
// 	*pt = v
// 	return nil
// }

// // Write PaymentTypeValue to JSON
// func (pt *PaymentTypeValue) MarshalJSON() ([]byte, error) {
// 	if pt == nil {
// 		return nil, nil
// 	}
// 	return []byte(fmt.Sprintf("\"%s\"", pt.String())), nil
// }

// // Write from Server to DB
// func (pt *PaymentTypeValue) Value() (driver.Value, error) {
// 	if pt == nil {
// 		return nil, nil
// 	}
// 	return  pt.String(), nil
// }

// // Read JSON to PaymentTypeValue
// func (pt *PaymentTypeValue) UnmarshalJSON(data []byte) error {
// 	str := strings.ReplaceAll(string(data), "\"", "")

// 	ptValue, err := parseStr2PaymentStatus(str)

// 	if err != nil {
// 		return err
// 	}

// 	*pt = ptValue

//		return nil
//	}
var (
	ErrorNotFoundShippingMethod = errors.New("Not found shipping method")
)

type ShippingMethod struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ShippingMethod) TableName() string { return "shipping_methods" }

type CreateShippingMethod struct {
	ID    int    `json:"-"`
	Name  string `json:"name" validate:"required,gte=1,lte=50"`
	Price int    `json:"price" validate:"required,numeric"`
}

func (CreateShippingMethod) TableName() string { return ShippingMethod{}.TableName() }

type UpdateShippingMethod struct {
	Name  string `json:"name" validate:"required,gte=1,lte=50"`
	Price int    `json:"price" validate:"required,numeric"`
}

func (UpdateShippingMethod) TableName() string { return ShippingMethod{}.TableName() }
