package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// smsServiceFaker is a faker for messageService
type messageServiceFaker struct {
	mock.Mock
}

// Our mocked smsService method
func (m *messageServiceFaker) SendChargeNotification(value int) error {
	args := m.Called(value)
	return args.Error(0)
}

func TestChargeCustomer(t *testing.T) {
	smsService := new(messageServiceFaker)
	smsService.On("SendChargeNotification", 100).
		Return(errors.New("connection error"))
	bankService := BankService{smsService}
	assert.Nil(t, bankService.ChargeCustomer(100))
	smsService.AssertExpectations(t)
}

//--- FAIL: TestChargeCustomer (0.00s)
//    bank_test.go:28:
//        	Error Trace:	bank_test.go:28
//        	Error:      	Expected nil, but got: &errors.errorString{s:"could not charge money: connection error"}
//        	Test:       	TestChargeCustomer
//    bank_test.go:29: PASS:	SendChargeNotification(int)
//FAIL
