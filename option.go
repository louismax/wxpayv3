package wxpayv3

import "github.com/louismax/wxpayv3/core"

// ClientOption Client初始化参数
type ClientOption interface {
	Apply(settings *core.DialSettings) error
}

// ErrorOption 错误初始化参数，用于返回错误
type ErrorOption struct{ Error error }

// Apply 返回初始化错误
func (w ErrorOption) Apply(_ *core.DialSettings) error {
	return w.Error
}
