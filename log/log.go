package log

import (
	"os"
	"time"

	"git.bgenius.cn/universe/base/stack"
	"github.com/rs/zerolog"
)

const (
	FieldXid = "_xid"
)

var root zerolog.Logger

func init() {
	zerolog.LevelFieldName = "L"
	zerolog.MessageFieldName = "M"
	zerolog.TimestampFieldName = "T"

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	root = zerolog.New(output).With().Timestamp().Logger() // debug by default
}

func SetLevel(level string) {
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		lvl = zerolog.NoLevel
	}
	root.Level(lvl)
}

func SetXid(xid string) {
	root.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str(FieldXid, xid)
	})
}

func Debug(msg string) {
	root.Debug().Msg(msg)
}

func Debugf(fmt string, v ...interface{}) {
	root.Debug().Msgf(fmt, v...)
}

func DebugKv(msg string, kv ...interface{}) {
	root.Debug().Fields(convertKv2Fx(kv)).Msg(msg)
}

func DebugFx(msg string, fx ...stack.Fx) {
	m := stack.Fx{}
	for _, f := range fx {
		for k, v := range f {
			m[k] = v
		}
	}
	root.Debug().Fields(m).Msg(msg)
}

func Info(msg string) {
	root.Info().Msg(msg)
}

func Infof(fmt string, v ...interface{}) {
	root.Info().Msgf(fmt, v...)
}

func InfoKv(msg string, kv ...interface{}) {
	root.Info().Fields(convertKv2Fx(kv)).Msg(msg)
}

func InfoFx(msg string, fx ...stack.Fx) {
	m := stack.Fx{}
	for _, f := range fx {
		for k, v := range f {
			m[k] = v
		}
	}
	root.Info().Fields(m).Msg(msg)
}

func Err(msg string) {
	root.Error().Msg(msg)
}

func Errf(fmt string, v ...interface{}) {
	root.Error().Msgf(fmt, v...)
}

func ErrKv(msg string, kv ...interface{}) {
	root.Error().Fields(convertKv2Fx(kv)).Msg(msg)
}

func ErrFx(msg string, fx ...stack.Fx) {
	m := stack.Fx{}
	for _, f := range fx {
		for k, v := range f {
			m[k] = v
		}
	}
	root.Error().Fields(m).Msg(msg)
}

func Warn(msg string) {
	root.Warn().Msg(msg)
}

func Warnf(fmt string, v ...interface{}) {
	root.Warn().Msgf(fmt, v...)
}

func WarnKv(msg string, kv ...interface{}) {
	root.Warn().Fields(convertKv2Fx(kv)).Msg(msg)
}

func WarnFx(msg string, fx ...stack.Fx) {
	m := stack.Fx{}
	for _, f := range fx {
		for k, v := range f {
			m[k] = v
		}
	}
	root.Warn().Fields(m).Msg(msg)
}

func Fatal(msg string) {
	root.Fatal().Msg(msg)
}

func Fatalf(fmt string, v ...interface{}) {
	root.Fatal().Msgf(fmt, v...)
}

func FatalKv(msg string, kv ...interface{}) {
	root.Fatal().Fields(convertKv2Fx(kv)).Msg(msg)
}

func FatalFx(msg string, fx ...stack.Fx) {
	m := stack.Fx{}
	for _, f := range fx {
		for k, v := range f {
			m[k] = v
		}
	}
	root.Fatal().Fields(m).Msg(msg)
}

func convertKv2Fx(kv []interface{}) stack.Fx {
	size := len(kv)
	if size == 0 || size%2 != 0 {
		return stack.Fx{}
	}

	fx := stack.Fx{}
	for i := 0; i < size; i += 2 {
		if n, ok := kv[i].(string); ok {
			fx[n] = kv[i+1]
		}
	}
	return fx
}
