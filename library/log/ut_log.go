package log

import ()

type Conf struct {
	File   string
	Level  int
	Suffix string
}

var (
	_loggers map[string]*Logger = make(map[string]*Logger)
	_lognull *Logger
)

func Init(confs map[string]Conf) error {
	var errnull error
	_lognull, errnull = NewLogger("/dev/null", LEVEL_OFF, SUFFIX_YM)
	if errnull != nil {
		return errnull
	}

	for k, v := range confs {
		if logTmp, err := NewLogger(v.File, v.Level, v.Suffix); err != nil {
			return err
		} else {
			_loggers[k] = logTmp
		}
	}

	return nil
}

func Node(node string) *Logger {
	if logger, ok := _loggers[node]; ok {
		return logger
	}
	return _lognull
}
