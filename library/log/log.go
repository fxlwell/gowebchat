package log

import (
	"fmt"
	lib_tool "github.com/jcsz/gowebchat/library/tool"
	"log"
	"os"
	"time"
)

type Logger struct {
	log     *log.Logger
	file    string
	logfile string
	level   int
	suffix  string
}

const (
	//指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
	LEVEL_OFF   = 0
	LEVEL_PANIC = 1
	LEVEL_FATAL = 2
	LEVEL_ERROR = 3
	LEVEL_WARN  = 4
	LEVEL_INFO  = 5
	LEVEL_DEBUG = 6
	LEVEL_ALL   = 7
)

const (
	//日志时间分表格式
	SUFFIX_YM       = "200601"
	SUFFIX_YM_TYPE  = "2006-01"
	SUFFIX_YMD      = "20060102"
	SUFFIX_YMD_TYPE = "2006-01-02"
	SUFFIX_YMDH     = "2006-01-02-15"
	SUFFIX_YMDHI    = "2006-01-02-15-04"
	SUFFIX_YMDHIS   = "2006-01-02-15-04-05"
)

func NewLogger(logFile string, level int, suffix string) *Logger {

	l := &Logger{
		file:   logFile,
		level:  level,
		suffix: suffix,
	}

	logf := l.get_log_file()
	l.logfile = logf

	if err := lib_tool.File_not_exists_to_create(logf); err != nil {
		panic(err)
	}

	logFd, err := os.OpenFile(logf, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		panic(err)
	}

	logger := log.New(logFd, "", log.Ldate|log.Ltime|log.Lshortfile)

	l.log = logger

	go l.check_and_update_file_name()

	return l
}

func (l *Logger) get_log_file() string {
	suffix := time.Now().Format(l.suffix)
	new_file := l.file + "." + suffix
	return new_file
}

func (l *Logger) check_and_update_file_name() {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			newf := l.get_log_file()
			if newf != l.logfile {
				logFd, err := os.OpenFile(newf, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
				if err == nil {
					l.logfile = newf
					l.log.SetOutput(logFd)
				}
			}
		}
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.level >= LEVEL_DEBUG {
		l.log.SetPrefix("debug ")
		l.log.Output(2, fmt.Sprintf(format, args...))
	}
}
func (l *Logger) Infof(format string, args ...interface{}) {
	if l.level >= LEVEL_INFO {
		l.log.SetPrefix("info ")
		l.log.Output(2, fmt.Sprintf(format, args...))
	}
}
func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.level >= LEVEL_WARN {
		l.log.SetPrefix("warn ")
		l.log.Output(2, fmt.Sprintf(format, args...))
	}
}
func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.level >= LEVEL_ERROR {
		l.log.SetPrefix("error ")
		l.log.Output(2, fmt.Sprintf(format, args...))
	}
}
func (l *Logger) Fatalf(format string, args ...interface{}) {
	if l.level >= LEVEL_FATAL {
		l.log.SetPrefix("fatal ")
		l.log.Output(2, fmt.Sprintf(format, args...))
	}
}
func (l *Logger) Printf(format string, args ...interface{}) {
	l.log.Output(2, fmt.Sprintf(format, args...))
}
