package logging

import (
	"os"

	"github.com/mitra-gh/CarBid/configs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var zerologLevelMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
	"panic": zerolog.PanicLevel,
}


type zerologLogger struct {
	cfg    *configs.Config
	logger *zerolog.Logger
}

func NewZerologLogger(cfg *configs.Config) *zerologLogger {
	logger:= &zerologLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zerologLogger) getLevel() zerolog.Level {
	level, ok := zerologLevelMap[l.cfg.Logger.Level]
	if !ok {
		return zerolog.InfoLevel
	}
	return level
}


func (l *zerologLogger) Init(){
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	
	file, err := os.OpenFile(l.cfg.Logger.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("could not open log file: " + err.Error())
	}

	var logger = zerolog.New(file).
	With().
	Timestamp().
	Str("AppName", "CarBid").
	Str("LoggerName", "ZerologLogger").
	Logger()

	zerolog.SetGlobalLevel(l.getLevel())
	l.logger = &logger
}

func (l *zerologLogger) Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	l.logger.Debug().
	Str("Category", string(category)).
	Str("SubCategory", string(subCategory)).
	Interface("Params", mapToZeroParams(extra)).
	Msg(message)
}

func (l *zerologLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debug().
	Msgf(format, v...)
}

func (l *zerologLogger) Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	l.logger.Info().
	Str("Category", string(category)).
	Str("SubCategory", string(subCategory)).
	Interface("Params", mapToZeroParams(extra)).
	Msg(message)
}	

func (l *zerologLogger) Infof(format string, v ...interface{}) {
	l.logger.Info().
	Msgf(format, v...)
}

func (l *zerologLogger) Warn(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	l.logger.Warn().
	Str("Category", string(category)).
	Str("SubCategory", string(subCategory)).
	Interface("Params", mapToZeroParams(extra)).
	Msg(message)
}	

func (l *zerologLogger) Warnf(format string, v ...interface{}) {
	l.logger.Warn().
	Msgf(format, v...)
}	

func (l *zerologLogger) Error(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	l.logger.Error().
	Str("Category", string(category)).
	Str("SubCategory", string(subCategory)).
	Interface("Params", mapToZeroParams(extra)).
	Msg(message)
}		

func (l *zerologLogger) Errorf(format string, v ...interface{}) {
	l.logger.Error().
	Msgf(format, v...)
}		

func (l *zerologLogger) Fatal(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	l.logger.Fatal().
	Str("Category", string(category)).
	Str("SubCategory", string(subCategory)).
	Interface("Params", mapToZeroParams(extra)).
	Msg(message)
}		

func (l *zerologLogger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatal().
	Msgf(format, v...)
}	



