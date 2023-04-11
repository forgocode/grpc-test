package typed

type DbConfig struct {
	Database string `json:"database" yaml:"database"`
	URL      string `json:"url" yaml:"url"`

	Port   uint16 `json:"port" yaml:"port"`
	DbName string `json:"dbName" yaml:"dbName"`
	User   string `json:"user" yaml:"user"`
	Passwd string `json:"passwd" yaml:"passwd"`
}

type LogConfig struct {
	FileName string `json:"fileName" yaml:"fileName"`
	Level    string `json:"level" yaml:"level"`
}

type ConfigYaml struct {
	Db   DbConfig   `json:"db" yaml:"db"`
	Log  LogConfig  `json:"log" yaml:"log"`
	Etcd EtcdConfig `json:"etcd" yaml:"etcd"`
}

type EtcdConfig struct {
	EndPoints []string `json:"endPoints" yaml:"endPoints"`
}
