package logging

import (
	"github.com/mitra-gh/CarBid/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"fatal":  zapcore.FatalLevel,
}

type zapLogger struct {
	cfg    *configs.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *configs.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zapLogger) getZapLevel(level string) zapcore.Level {
	lvl, ok := zapLevelMap[level]
	if !ok {
		return zapcore.InfoLevel
	}
	return lvl
}


func (l *zapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: l.cfg.Logger.FilePath,
		MaxSize:  1,
		MaxBackups: 10,
		MaxAge:     5,
		Compress:   true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		l.getZapLevel(l.cfg.Logger.Level),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	l.logger = logger.Sugar()
}



func (l *zapLogger) Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(extra, category, subCategory)
	l.logger.Debugw(message, params...)
}

func (l *zapLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l *zapLogger) Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(extra, category, subCategory)
	l.logger.Infow(message, params...)
}

func (l *zapLogger) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (l *zapLogger) Warn(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(extra, category, subCategory)
	l.logger.Warnw(message, params...)
}

func (l *zapLogger) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

func (l *zapLogger) Error(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(extra, category, subCategory)
	l.logger.Errorw(message, params...)
}

func (l *zapLogger) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

func (l *zapLogger) Fatal(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(extra, category, subCategory)
	l.logger.Fatalw(message, params...)
}

func (l *zapLogger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

