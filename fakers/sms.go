package main

// SMSService is our implementation of MessageService.
type SMSService struct{}

// SendChargeNotification notifies clients they have been
// charged via SMS.
// This is the method we are going to mock.
func (sms SMSService) SendChargeNotification(value int) error {
	// <send message>()
	return nil
}
