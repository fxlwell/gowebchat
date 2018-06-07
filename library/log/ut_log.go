package log

import ()

type Conf struct {
	File   string
	Level  int
	Suffix string
}

var (
	_loggers map[string]*Logger = make(map[string]*Logger)
	_lognull *Logger            = NewLogger("/dev/null", LEVEL_OFF, SUFFIX_YM)
)

func Init(confs map[string]Conf) {
	for k, v := range confs {
		_loggers[k] = NewLogger(v.File, v.Level, v.Suffix)
	}
	return
}

func Logger(node string) *Logger {
	if logger, ok := _loggers[node]; ok {
		return logger
	}
	return _lognull
}
