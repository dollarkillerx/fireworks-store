package models

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/dollarkillerx/fireworks/pkg/utils"
	"gorm.io/gorm"
)

type CommonModel struct {
	ID        string    `gorm:"primarykey;type:varchar(360)"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func DefaultCommonModel() CommonModel {
	return CommonModel{
		ID: utils.SonuFlakeStringID(),
	}
}

type IBStateType string

const (
	IBStateTypeActivation IBStateType = "Activation"
	IBStateTypeFreeze     IBStateType = "Freeze"
)

var AllIBStateType = []IBStateType{
	IBStateTypeActivation,
	IBStateTypeFreeze,
}

func (e IBStateType) IsValid() bool {
	switch e {
	case IBStateTypeActivation, IBStateTypeFreeze:
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

type WithdrawType string

const (
	WithdrawTypeWait    WithdrawType = "Wait"
	WithdrawTypeSuccess WithdrawType = "Success"
	WithdrawTypeFail    WithdrawType = "Fail"
)

var AllWithdrawType = []WithdrawType{
	WithdrawTypeWait,
	WithdrawTypeSuccess,
	WithdrawTypeFail,
}

func (e WithdrawType) IsValid() bool {
	switch e {
	case WithdrawTypeWait, WithdrawTypeSuccess, WithdrawTypeFail:
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

type OrderType string

const (
	OrderTypeToBePaid          OrderType = "ToBePaid"
	OrderTypeCancelled         OrderType = "Cancelled"
	OrderTypePaymentSuccessful OrderType = "PaymentSuccessful"
)

var AllOrderType = []OrderType{
	OrderTypeToBePaid,
	OrderTypeCancelled,
	OrderTypePaymentSuccessful,
}

func (e OrderType) IsValid() bool {
	switch e {
	case OrderTypeToBePaid, OrderTypeCancelled, OrderTypePaymentSuccessful:
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

type IBGrade string

const (
	IBGradeBronze IBGrade = "Bronze"
	IBGradeSilver IBGrade = "Silver"
	IBGradeGold   IBGrade = "Gold"
)

var AllIBGrade = []IBGrade{
	IBGradeBronze,
	IBGradeSilver,
	IBGradeGold,
}

func (e IBGrade) IsValid() bool {
	switch e {
	case IBGradeBronze, IBGradeSilver, IBGradeGold:
		return true
	}
	return false
}

func (e IBGrade) String() string {
	return string(e)
}

func (e *IBGrade) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IBGrade(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IBGrade", str)
	}
	return nil
}

func (e IBGrade) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
