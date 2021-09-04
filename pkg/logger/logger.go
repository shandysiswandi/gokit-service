package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	log "github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

var (
	logInstance   log.Logger
	once          sync.Once
	timeLayout    log.Valuer
	logrusPackage string
)

func init() {
	once.Do(func() {
		// prepare
		timeLayout = log.DefaultTimestamp

		// instanciate
		logInstance = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
		logInstance = log.With(logInstance, "time", timeLayout, "caller", getCaller())
	})
}

// getPackageName reduces a fully qualified function name to the package name
// There really thought to be to be a better way
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

// getCaller is function to retrieves the name of the first non-this calling function
func getCaller() log.Valuer {
	return func() interface{} {
		// initial variable
		pcs := make([]uintptr, 10)
		runtime.Callers(0, pcs)

		// search runtime caller file and line for this package
		for i := 0; i < 10; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "getCaller") {
				logrusPackage = getPackageName(funcName)
				break
			}
		}

		// search runtime caller file and line for value
		frames := runtime.CallersFrames(pcs[:runtime.Callers(4, pcs)])
		for f, again := frames.Next(); again; f, again = frames.Next() {
			if getPackageName(f.Function) != logrusPackage {
				return fmt.Sprintf("%s:%d", f.File, f.Line)
			}
		}

		// the default caller file and line
		return fmt.Sprintf("%s:%d", "-", 0)
	}
}

// SetOutput set output for logger
func SetOutput(out io.Writer) {
	if out == nil {
		out = os.Stderr
	}

	logInstance = log.NewJSONLogger(log.NewSyncWriter(out))
	logInstance = log.With(logInstance, "time", timeLayout, "caller", getCaller())
}

// SetTimeFormat set time format for logger
//
// Todo: not implement
func SetTimeFormat(time time.Time, layout string) {
	timeLayout = log.DefaultTimestampUTC
	logInstance = log.With(logInstance, "time", timeLayout, "caller", getCaller())
}

func Info(args ...interface{}) {
	args = append([]interface{}{"msg"}, args...)
	level.Info(logInstance).Log(args...)
}

func Debug(args ...interface{}) {
	args = append([]interface{}{"msg"}, args...)
	level.Debug(logInstance).Log(args...)
}

func Warn(args ...interface{}) {
	args = append([]interface{}{"msg"}, args...)
	level.Warn(logInstance).Log(args...)
}

func Error(args ...interface{}) {
	args = append([]interface{}{"msg"}, args...)
	level.Error(logInstance).Log(args...)
}
