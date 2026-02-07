package setting

type Config struct {
	Mysql MySQLSetting `mapstucture:"mysql"`
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
