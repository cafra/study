package unit

import (
	"fmt"
	"testing"

	"github.com/cafra/study/unit/mocks"
)

func TestMyService_ChargeCustomer(t *testing.T) {
	type fields struct {
		messageService MessageService
	}
	type args struct {
		value int
	}

	smsService := new(mocks.MessageService)
	smsService.On("SendChargeNotification", 100).Return(nil)
	smsService.On("SendChargeNotification", 200).Return(fmt.Errorf("Internal err"))

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "正常测试",
			fields:  struct{ messageService MessageService }{messageService: SMSService{}},
			args:    struct{ value int }{value: 50},
			wantErr: false,
		},
		{
			name:    "mock测试正常",
			fields:  struct{ messageService MessageService }{messageService: smsService},
			args:    struct{ value int }{value: 100},
			wantErr: false,
		},
		{
			name:    "mock测试异常",
			fields:  struct{ messageService MessageService }{messageService: smsService},
			args:    struct{ value int }{value: 200},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := MyService{
				messageService: tt.fields.messageService,
			}
			if err := a.ChargeCustomer(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ChargeCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	// 最后，我们验证 myService.ChargeCustomer 调用了我们模拟的 SendChargeNotification 方法
	smsService.AssertExpectations(t)
}
