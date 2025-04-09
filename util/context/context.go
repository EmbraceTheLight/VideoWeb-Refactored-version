package context

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// DefaultTimeout 默认超时时间
const DefaultTimeout = 5 * time.Second

// ContextKey 常量 key，用于存储和提取值
type ContextKey string

const (
	TraceIDKey ContextKey = "trace_id"
	UserIDKey  ContextKey = "user_id"
)

// NewBaseContext 初始化基础 context
func NewBaseContext() context.Context {
	return context.Background()
}

// WithTimeout 派生带超时的 context
func WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}
	if parent == nil {
		return context.WithTimeout(context.Background(), timeout)
	}
	return context.WithTimeout(parent, timeout)
}

// WithValue 向 context 中注入值
func WithValue(parent context.Context, key any, value interface{}) context.Context {
	return context.WithValue(parent, key, value)
}

// GetValue 从 context 中获取值
func GetValue(ctx context.Context, key any) (interface{}, bool) {
	value := ctx.Value(key)
	return value, value != nil
}

// MustGetValue 从 context 中获取值，如果不存在则 panic
func MustGetValue(ctx context.Context, key any) any {
	value := ctx.Value(key)

	if value == nil {
		errString := fmt.Sprintf("key %v not found in context", key)
		panic(errors.New(errString))
	}
	return value
}
