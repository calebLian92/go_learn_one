package server

import (
	"fmt"
	"learn-one/cmd/common/interfaces/user"
)

// PaymentService 结构体
type PaymentService struct {
	orderManager interface{}       // 先使用空接口
	manager      user.OrderManager // 实际接口
}

// NewPaymentService 创建支付服务
func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

// SetOrderManager 设置订单管理器
func (s *PaymentService) SetOrderManager(manager user.OrderManager) {
	s.orderManager = manager
	s.manager = manager
}

// ProcessPayment 处理支付
func (s *PaymentService) ProcessPayment(amount float64) error {
	fmt.Printf("支付服务: 开始处理支付 %.2f\n", amount)

	if s.manager == nil {
		return fmt.Errorf("订单管理器未设置")
	}

	// 模拟支付处理
	fmt.Println("支付服务: 支付处理中...")

	// 支付成功后通知订单服务
	err := s.manager.OrderPaid("order_123")
	if err != nil {
		return err
	}

	// 更新订单状态
	err = s.manager.UpdateOrderStatus("order_123", "已支付")
	if err != nil {
		return err
	}

	fmt.Println("支付服务: 支付处理完成")
	return nil
}
