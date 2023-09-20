package config

type Config struct {
	RootPath string
	Mysqls   []Mysql `mapstructure:"mysqls" json:"mysql" yaml:"mysqls"`
	Log      Log     `mapstructure:"log" json:"log" yaml:"log"`
	Jaeger   Jaeger  `mapstructure:"jaeger" json:"jaeger" yaml:"jaeger"`
	Minio    Minio   `mapstructure:"minio" json:"minio" yaml:"minio"`
}

type Mysqls struct {
	Mysqls []*Mysql `mapstructure:"mysqls" json:"mysqls" yaml:"mysqls"`
}

type Mysql struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`             // 服务器地址:端口
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             //:端口
	Config   string `mapstructure:"config" json:"config" yaml:"config"`       // 高级配置
	Dbname   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // 数据库名
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	Conn     string `mapstructure:"conn" json:"conn" yaml:"conn"`
}

type Log struct {
	LogLevel          string `mapstructure:"log_level" json:"log_level" yaml:"log_level"`                                  // 日志打印级别 debug  info  warning  error
	LogFormat         string `mapstructure:"log_format" json:"log_format" yaml:"log_format"`                               // 输出日志格式	logfmt, json
	LogPath           string `mapstructure:"log_path" json:"log_path" yaml:"log_path"`                                     // 输出日志文件路径
	LogFileName       string `mapstructure:"log_file_name" json:"log_file_name" yaml:"log_file_name"`                      // 输出日志文件名称
	LogFileMaxSize    int    `mapstructure:"log_file_max_size" json:"log_file_max_size" yaml:"log_file_max_size"`          // 【日志分割】单个日志文件最多存储量 单位(mb)
	LogFileMaxBackups int    `mapstructure:"log_file_max_backups" json:"log_file_max_backups" yaml:"log_file_max_backups"` // 【日志分割】日志备份文件最多数量
	LogMaxAge         int    `mapstructure:"log_max_age" json:"log_max_age" yaml:"log_max_age"`                            // 日志保留时间，单位: 天 (day)
	LogCompress       bool   `mapstructure:"log_compress" json:"log_compress" yaml:"log_compress"`                         // 是否压缩日志
	LogStdout         bool   `mapstructure:"log_stdout" json:"log_stdout" yaml:"log_stdout"`                               // 是否输出到控制台
}

type Minio struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyID     string `mapstructure:"accessKeyID" json:"accessKeyID" yaml:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey" json:"secretAccessKey" yaml:"secretAccessKey"`
}

type Jaeger struct {
	CollectorEndpoint  string `mapstructure:"collector_endpoint" json:"collector_endpoint" yaml:"collector_endpoint"`
	LocalAgentHostPort string `mapstructure:"localAgentHostPort" json:"localAgentHostPort" yaml:"localAgentHostPort"`
	ServiceName        string `mapstructure:"serviceName" json:"serviceName" yaml:"serviceName"`
}
