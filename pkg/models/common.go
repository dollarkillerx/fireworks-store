package models

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type CommonModel struct {
	ID        string    `gorm:"primarykey;type:varchar(360)"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type OrderType string

const (
	OrderTypeOrderTypeToBePaid          OrderType = "OrderTypeToBePaid"
	OrderTypeOrderTypeCancelled         OrderType = "OrderTypeCancelled"
	OrderTypeOrderTypePaymentSuccessful OrderType = "OrderTypePaymentSuccessful"
)

var AllOrderType = []OrderType{
	OrderTypeOrderTypeToBePaid,
	OrderTypeOrderTypeCancelled,
	OrderTypeOrderTypePaymentSuccessful,
}

func (e OrderType) IsValid() bool {
	switch e {
	case OrderTypeOrderTypeToBePaid, OrderTypeOrderTypeCancelled, OrderTypeOrderTypePaymentSuccessful:
		return true
	}
	return false
}

func (e OrderType) String() string {
	return string(e)
}

func (e *OrderType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderType", str)
	}
	return nil
}

func (e OrderType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WithdrawType string

const (
	WithdrawTypeWithdrawTypeWait    WithdrawType = "WithdrawTypeWait"
	WithdrawTypeWithdrawTypeSuccess WithdrawType = "WithdrawTypeSuccess"
	WithdrawTypeWithdrawTypeFail    WithdrawType = "WithdrawTypeFail"
)

var AllWithdrawType = []WithdrawType{
	WithdrawTypeWithdrawTypeWait,
	WithdrawTypeWithdrawTypeSuccess,
	WithdrawTypeWithdrawTypeFail,
}

func (e WithdrawType) IsValid() bool {
	switch e {
	case WithdrawTypeWithdrawTypeWait, WithdrawTypeWithdrawTypeSuccess, WithdrawTypeWithdrawTypeFail:
		return true
	}
	return false
}

func (e WithdrawType) String() string {
	return string(e)
}

func (e *WithdrawType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WithdrawType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WithdrawType", str)
	}
	return nil
}

func (e WithdrawType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type IBStateType string

const (
	IBStateTypeIBStateTypeActivation IBStateType = "IBStateTypeActivation"
	IBStateTypeIBStateTypeFreeze     IBStateType = "IBStateTypeFreeze"
)

var AllIBStateType = []IBStateType{
	IBStateTypeIBStateTypeActivation,
	IBStateTypeIBStateTypeFreeze,
}

func (e IBStateType) IsValid() bool {
	switch e {
	case IBStateTypeIBStateTypeActivation, IBStateTypeIBStateTypeFreeze:
		return true
	}
	return false
}

func (e IBStateType) String() string {
	return string(e)
}

func (e *IBStateType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IBStateType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IBStateType", str)
	}
	return nil
}

func (e IBStateType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
