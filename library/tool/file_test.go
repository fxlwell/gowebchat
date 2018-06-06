package tool

import (
	"testing"
)

func Test_file_exist(t *testing.T) {
	re := File_exists("/home/ldap/lifuxing/gopath/src/github.com/jcsz/gowebchat/library/tool/filex.go")
	re1 := File_exists("/home/ldap/lifuxing/gopath/src/github.com/jcsz/gowebchat/library/tool/file.go")
	re2 := File_exists("/home/ldap/lifuxing/gopath/src/github.com/jcsz/gowebchat/library/tool")
	re3 := File_exists("/home/ldap/lifuxing/gopath/src/github.com/jcsz/gowebchat/library/tool/")
	re4 := File_exists("/home/ldap/lifuxing/gopath/src/github.com/jcsz/gowebchat/library/toolx")
	if re == false && re1 == true && re2 == true && re3 == true && re4 == false {
	} else {
		t.Fail()
	}
}
