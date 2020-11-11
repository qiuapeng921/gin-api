package mysql

type Config struct {
	Host     string
	Database string
	Username string
	Password string
	Charset  string
	MaxIdle  int
	MaxOpen  int
}
