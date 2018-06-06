package config

var (
	conf_path string = "/home/ldap/lifuxing/gopath/src/github.com/jcsz/gowebchat/conf"
)

func Init() {
	Parse_log_config()
	Parse_mysql_config()
}
