package service_test

import (
	"testing"

	"github.com/Delaram-Gholampoor-Sagha/testing_in_go/service"
	"github.com/Delaram-Gholampoor-Sagha/testing_in_go/service/mocks"
	
)

// type smsServiceMock struct {
// 	mock.Mock
// }

// func (m *smsServiceMock) SendChargeNotification(value int) bool {
// 	fmt.Println("Mocked charge notification")
// 	arg := m.Called(value)
// 	return arg.Bool(0)
// }

func TestChargeCustomer(t *testing.T) {
	smsService := new(mocks.MessageService)
	smsService.On("SendChargeNotification", 100).Return(true)

	myService := service.MyService{smsService}

	myService.ChargeCustomer(100)
	smsService.AssertExpectations(t)

}
