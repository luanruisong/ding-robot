package ding_robot

import "fmt"

type (
	Logger interface {
		Info(s string, i ...interface{})
		Warn(s string, i ...interface{})
		Error(s string, i ...interface{})
	}

	DefaultLogger struct {
	}
)

func (DefaultLogger) Info(s string, i ...interface{}) {
	fmt.Println(fmt.Sprintf(s, i...))
}
func (DefaultLogger) Warn(s string, i ...interface{}) {
	fmt.Println(fmt.Sprintf(s, i...))
}
func (DefaultLogger) Error(s string, i ...interface{}) {
	fmt.Println(fmt.Sprintf(s, i...))
}

var logger Logger

func init() {
	logger = &DefaultLogger{}
}

func SetLogger(l Logger) {
	logger = l
}
