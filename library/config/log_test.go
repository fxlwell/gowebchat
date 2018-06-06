package config

import (
	"testing"
)

func Test_parse_log(t *testing.T) {
	Parse_log_config()
	t.Log(LogPath)
	if LogPath == "./var/logs" {
		t.Fail()
	}
}
