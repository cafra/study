package unit

import "fmt"

// MessageService 通知客户被收取的费用
type MessageService interface {
	SendChargeNotification(int) error
}

// SMSService 是 MessageService 的实现
type SMSService struct{}

// SendChargeNotification 通过 SMS 来告知客户他们被收取费用
// 这就是我们将要模拟的方法
func (sms SMSService) SendChargeNotification(value int) error {
	fmt.Println("Sending Production Charge Notification")
	return nil
}

// MyService 使用 MessageService 来通知客户
type MyService struct {
	messageService MessageService
}

// ChargeCustomer 向客户收取费用
// 在真实系统中，我们会模拟这个
// 但是在这里，我想在每次运行测试时都赚点钱
func (a MyService) ChargeCustomer(value int) (err error) {
	err = a.messageService.SendChargeNotification(value)
	if err != nil {
		fmt.Printf("Charging SendChargeNotification err %v\n", err)
		return err
	}
	fmt.Printf("Charging Customer For the value of %d\n", value)
	return nil
}
