package main

import (
	"context"
	"testing"
	"time"
)

func TestClassicContextTimeout(t *testing.T) {
	// 创建一个带有超时时间的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 调用ClassicContextTimeout函数，传入带有超时时间的上下文
	err := ClassicContextTimeout(ctx)

	// 判断函数返回值，如果返回值不为nil，则输出错误信息
	if err != nil {
		t.Errorf("got %v; want nil", err)
	}
}
