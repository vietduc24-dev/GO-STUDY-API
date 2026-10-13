package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	MySQL  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"log"`
	Redis  RedisSetting  `mapstructure:"redis"`
	Kafka  KafkaSetting  `mapstructure:"kafka"`
}

type ServerSetting struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DataBase int    `mapstructure:"database"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbName"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_Log_Name string `mapstructure:"file_log_name"`
	Max_Size      int    `mapstructure:"max_size"`
	Max_Backups   int    `mapstructure:"max_backups"`
	Max_Age       int    `mapstructure:"max_age"`
	Compress      bool   `mapstructure:"compress"`
}

type KafkaSetting struct {
	Broker string `mapstructure:"broker"`
	Topic  string `mapstructure:"topic"`
}
