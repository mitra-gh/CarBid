package logging

import "github.com/mitra-gh/CarBid/configs"

type Logger interface {
	Init()

	Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Warn(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Error(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Fatal(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})

	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

func NewLogger(cfg *configs.Config) Logger {
	if cfg.Logger.Type == "zap" {
		return NewZapLogger(cfg)
	}
	if cfg.Logger.Type == "zerolog" {
		return NewZerologLogger(cfg)
	}
	return nil
}
