// Package utils 提供应用程序的通用工具函数
// 包含日志记录、错误处理等基础功能
package utils

import (
	"fmt"
	"log"
)

// Logger 全局日志实例
// 提供统一的日志记录接口
var Logger = &logger{}

// logger 日志记录器结构体
// 实现统一的日志格式化输出
type logger struct{}

// WrapError 包装错误信息
// 参数:
//   - message: 附加的错误描述
//   - err: 原始错误
//
// 返回:
//   - error: 包装后的错误，包含原始错误信息
func WrapError(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}

// Info 输出信息级别日志
// 参数:
//   - module: 模块名称，用于标识日志来源
//   - message: 日志消息内容
func (l *logger) Info(module string, message string) {
	log.Printf("[NectarPin][%s]: %s", module, message)
}

// Infof 输出格式化信息级别日志
// 参数:
//   - module: 模块名称
//   - format: 格式化字符串
//   - v: 格式化参数
func (l *logger) Infof(module string, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	log.Printf("[NectarPin][%s]: %s", module, message)
}

// Error 输出错误级别日志
// 参数:
//   - module: 模块名称
//   - message: 错误消息内容
func (l *logger) Error(module string, message string) {
	log.Printf("[NectarPin][%s]: 错误 - %s", module, message)
}

// Errorf 输出格式化错误级别日志
// 参数:
//   - module: 模块名称
//   - format: 格式化字符串
//   - v: 格式化参数
func (l *logger) Errorf(module string, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	log.Printf("[NectarPin][%s]: 错误 - %s", module, message)
}

// Fatal 输出致命错误日志并终止程序
// 参数:
//   - module: 模块名称
//   - message: 致命错误消息
func (l *logger) Fatal(module string, message string) {
	log.Fatalf("[NectarPin][%s]: 致命错误 - %s", module, message)
}

// Fatalf 输出格式化致命错误日志并终止程序
// 参数:
//   - module: 模块名称
//   - format: 格式化字符串
//   - v: 格式化参数
func (l *logger) Fatalf(module string, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	log.Fatalf("[NectarPin][%s]: 致命错误 - %s", module, message)
}

// Success 输出成功级别日志
// 参数:
//   - module: 模块名称
//   - message: 成功消息内容
func (l *logger) Success(module string, message string) {
	log.Printf("[NectarPin][%s]: 成功 - %s", module, message)
}

// Successf 输出格式化成功级别日志
// 参数:
//   - module: 模块名称
//   - format: 格式化字符串
//   - v: 格式化参数
func (l *logger) Successf(module string, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	log.Printf("[NectarPin][%s]: 成功 - %s", module, message)
}
