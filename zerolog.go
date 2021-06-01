package zerolog

import (
	"embed"
	"os"

	"github.com/go-gluon/gluon"
	"github.com/go-gluon/gluon/log"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type ZerologConfig struct {
	Log struct {
		Debug bool `config:"debug"`
		Json  bool `config:"json"`
	} `config:"log"`
}

type ExtensionProvider struct {
}

func (e ExtensionProvider) NewExtesion() gluon.Extension {
	return gluon.Extension{
		Name:     "zerolog",
		Priority: 10,
		Init:     Init,
		Config:   &ZerologConfig{},
	}
}

func Init(resources embed.FS, config interface{}) error {

	tmp := config.(*ZerologConfig)
	if !tmp.Log.Json {
		zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: zerolog.TimeFieldFormat}).With().Logger()
	} else {
		zlog.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}

	if !tmp.Log.Debug {
		zlog.Logger = zlog.Level(zerolog.DebugLevel)
	} else {
		zlog.Logger = zlog.Level(zerolog.InfoLevel)
	}
	log.Log = New(zlog.Logger)
	return nil
}

type Logger struct {
	logger zerolog.Logger
}

func New(logger zerolog.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	sendEvent(l.logger.Trace(), msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	sendEvent(l.logger.Debug(), msg, fields...)
}

func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	sendEvent(l.logger.Info(), msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	sendEvent(l.logger.Warn(), msg, fields...)
}

func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	sendEvent(l.logger.Error(), msg, fields...)
}

func sendEvent(event *zerolog.Event, msg string, fields ...map[string]interface{}) {
	if len(fields) > 0 {
		event.Fields(fields[0])
	}

	event.Msg(msg)
}
