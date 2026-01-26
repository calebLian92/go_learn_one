package server

import (
	"fmt"
	"learn-one/cmd/common/interfaces/user"
)

// OrderService 结构体
type OrderService struct {
	paymentProcessor interface{}           // 先使用空接口
	processor        user.PaymentProcessor // 实际接口，需要从main包导入
}

// 注意：这里 PaymentProcessor 是从main包导入的接口
// 在真实项目中，需要导入 main 包或专门的接口包

// NewOrderService 创建订单服务
func NewOrderService() *OrderService {
	return &OrderService{}
}

// SetPaymentProcessor 设置支付处理器
func (s *OrderService) SetPaymentProcessor(processor user.PaymentProcessor) {
	s.paymentProcessor = processor
	s.processor = processor
}

// CreateOrder 创建订单
func (s *OrderService) CreateOrder(amount float64) error {
	fmt.Println("订单服务: 开始创建订单")

	if s.processor == nil {
		return fmt.Errorf("支付处理器未设置")
	}

	// 调用支付处理器
	err := s.processor.ProcessPayment(amount)
	if err != nil {
		return fmt.Errorf("支付失败: %v", err)
	}

	fmt.Println("订单服务: 订单创建成功")
	return nil
}

// OrderPaid 订单支付完成
func (s *OrderService) OrderPaid(orderID string) error {
	fmt.Printf("订单服务: 订单 %s 已支付\n", orderID)
	return nil
}

// UpdateOrderStatus 更新订单状态
func (s *OrderService) UpdateOrderStatus(orderID string, status string) error {
	fmt.Printf("订单服务: 更新订单 %s 状态为 %s\n", orderID, status)
	return nil
}
