package user

// 定义接口
type Notifier interface {
	SendNotification(message string)
}

type UserService struct {
	notifier Notifier // 依赖接口，而不是具体实现
}

func NewUserService(notifier Notifier) *UserService {
	return &UserService{notifier: notifier}
}

// 支付案例
// 在 main 包中定义接口，这样所有包都能使用
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type OrderManager interface {
	UpdateOrderStatus(orderID string, status string) error
	OrderPaid(orderID string) error
}
