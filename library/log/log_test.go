package log

import (
	"testing"
	"time"
)

func Test_new_logger(t *testing.T) {
	logger := NewLogger("/home/ldap/lifuxing/gopath/src/github.com/jcsz/gowebchat/library/log/test.log", LEVEL_FATAL, SUFFIX_YMDHIS)
	for i := 0; i < 10000; i++ {
		logger.Debugf("1111111111111111111111")
		logger.Infof("222222222222222222222")
		logger.Warnf("333333333333333333333")
		logger.Errorf("444444444444444444444")
		logger.Fatalf("555555555555555555555")
		time.Sleep(time.Second)
	}
}
