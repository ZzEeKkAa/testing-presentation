package main

import (
	"fmt"
	"log"
)

// MessageService handles notifying clients they have
// been charged.
type MessageService interface {
	SendChargeNotification(int) error
}

// BankService uses the MessageService to notify clients.
type BankService struct {
	messageService MessageService
}

// ChargeCustomer performs the charge to the customer.
func (a BankService) ChargeCustomer(value int) error {
	// <charging money>()
	if err := a.messageService.SendChargeNotification(value); err != nil {
		return fmt.Errorf("could not charge money: %v", err)
	}
	return nil
}

// A "Production" Example.
func main() {
	smsService := SMSService{}
	bank := BankService{smsService}
	if err := bank.ChargeCustomer(100); err != nil {
		log.Fatal(err)
	}
}
