package setting

type Config struct {
	Mysql  MySQLSetting  `mapstucture:"mysql"`
	Logger LoggerSetting `mapstucture:"logger"`
}

type MySQLSetting struct {
	Host            string `mapstucture:"host"`
	Port            int    `mapstucture:"port"`
	Username        string `mapstucture:"username"`
	Password        string `mapstucture:"password"`
	Dbname          string `mapstucture:"dbname"`
	MaxIdleConns    int    `mapstucture:"maxIdleConns"`
	MaxOpenConns    int    `mapstucture:"maxOpenConns"`
	ConnMaxLifetime int    `mapstucture:"connMaxLifetime"`
}

type LoggerSetting struct {
	LogLevel    string `mapstucture:"logLevel"`
	FileLogName string `mapstucture:"fileLogName"`
	MaxSize     int    `mapstucture:"maxSize"`
	MaxBackups  int    `mapstucture:"maxBackups"`
	MaxAge      int    `mapstucture:"maxAge"`
	Compress    bool   `mapstucture:"compress"`
}
