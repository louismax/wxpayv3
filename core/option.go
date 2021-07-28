package core

// ClientOption 微信支付 API v3 HTTPClient core.Client 初始化参数
type ClientOption interface {
	Join(settings *DialSettings) error
}

// ErrorOption 错误初始化参数，用于返回错误
type ErrorOption struct{ Error error }

// Join 返回初始化错误
func (w ErrorOption) Join(o *DialSettings) error {
	return w.Error
}