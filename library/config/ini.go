package config

var (
	conf_path string = "/home/ldap/lifuxing/fxcodes/gowebchat/conf"
)

func Init() {
	Parse_log_config()
}
