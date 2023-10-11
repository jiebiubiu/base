package config

type Config struct {
	RootPath   string
	Mysqls     []Mysql `mapstructure:"mysqls" json:"mysqls" yaml:"mysqls" config:",prefix_ext=mysqls/"`
	RedisConns []Redis `mapstructure:"redis_conns" json:"redis_conns" yaml:"redis_conns" config:",prefix_ext=redis_conns/"` // redis://<user>:<pass>@localhost:6379/<db>
	Log        Log     `mapstructure:"log" json:"log" yaml:"log" config:",prefix_ext=log/"`
	Jaeger     Jaeger  `mapstructure:"jaeger" json:"jaeger" yaml:"jaeger" config:",prefix_ext=jaeger/"`
	Minio      Minio   `mapstructure:"minio" json:"minio" yaml:"minio" config:",prefix_ext=minio/"`
	Email      int
}

type Email struct {
	Account  string `mapstructure:"account" json:"account" yaml:"account" config:"account,prefix"`
	Password string `mapstructure:"password" json:"password" yaml:"password" config:"password,prefix"`
}

type Mysqls struct {
	Mysqls []*Mysql `mapstructure:"mysqls" json:"mysqls" yaml:"mysqls"`
}

type Redis struct {
	Conn   string `mapstructure:"conn" json:"conn" yaml:"conn" config:"conn,prefix"`
	DbName string `mapstructure:"db_name" json:"db_name" yaml:"db_name" config:"db_name,prefix"` // 标识哪个连接
}

type Mysql struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path" config:"path,prefix,omitempty"`                 // 服务器地址:端口
	Port     string `mapstructure:"port" json:"port" yaml:"port" config:"port,prefix,omitempty"`                 //:端口
	Config   string `mapstructure:"config" json:"config" yaml:"config" config:"config,prefix,omitempty"`         // 高级配置
	Dbname   string `mapstructure:"db-name" json:"db-name" yaml:"db-name" config:"db-name,prefix,omitempty"`     // 数据库名
	Username string `mapstructure:"username" json:"username" yaml:"username" config:"username,prefix,omitempty"` // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password" config:"password,prefix,omitempty"` // 数据库密码
	Conn     string `mapstructure:"conn" json:"conn" yaml:"conn" config:"conn,prefix"`
}

type Log struct {
	LogLevel          string `mapstructure:"log_level" json:"log_level" yaml:"log_level" config:"log_level,prefix,omitempty"`                                             // 日志打印级别 debug  info  warning  error
	LogFormat         string `mapstructure:"log_format" json:"log_format" yaml:"log_format" config:"log_format,prefix,omitempty"`                                         // 输出日志格式	logfmt, json
	LogPath           string `mapstructure:"log_path" json:"log_path" yaml:"log_path" config:"log_path,prefix,omitempty"`                                                 // 输出日志文件路径
	LogFileName       string `mapstructure:"log_file_name" json:"log_file_name" yaml:"log_file_name" config:"log_file_name,prefix,omitempty"`                             // 输出日志文件名称
	LogFileMaxSize    int    `mapstructure:"log_file_max_size" json:"log_file_max_size" yaml:"log_file_max_size" config:"log_file_max_size,prefix,omitempty"`             // 【日志分割】单个日志文件最多存储量 单位(mb)
	LogFileMaxBackups int    `mapstructure:"log_file_max_backups" json:"log_file_max_backups" yaml:"log_file_max_backups" config:"log_file_max_backups,prefix,omitempty"` // 【日志分割】日志备份文件最多数量
	LogMaxAge         int    `mapstructure:"log_max_age" json:"log_max_age" yaml:"log_max_age" config:"log_max_age,prefix,omitempty"`                                     // 日志保留时间，单位: 天 (day)
	LogCompress       bool   `mapstructure:"log_compress" json:"log_compress" yaml:"log_compress" config:"log_compress,prefix,omitempty"`                                 // 是否压缩日志
	LogStdout         bool   `mapstructure:"log_stdout" json:"log_stdout" yaml:"log_stdout" config:"log_stdout,prefix,omitempty"`                                         // 是否输出到控制台
}

type Minio struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"  config:"endpoint,prefix,omitempty"`
	AccessKeyID     string `mapstructure:"accessKeyID" json:"accessKeyID" yaml:"accessKeyID"  config:"accessKeyID,prefix,omitempty"`
	SecretAccessKey string `mapstructure:"secretAccessKey" json:"secretAccessKey" yaml:"secretAccessKey"  config:"secretAccessKey,prefix,omitempty"`
}

type Jaeger struct {
	CollectorEndpoint  string `mapstructure:"collector_endpoint" json:"collector_endpoint" yaml:"collector_endpoint"  config:"collector_endpoint,prefix,omitempty"`
	LocalAgentHostPort string `mapstructure:"localAgentHostPort" json:"localAgentHostPort" yaml:"localAgentHostPort"  config:"localAgentHostPort,prefix,omitempty"`
	ServiceName        string `mapstructure:"serviceName" json:"serviceName" yaml:"serviceName"  config:"serviceName,prefix,omitempty"`
}
