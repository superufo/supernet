package viper

type ProxyCfgMap map[string]ProxyCfg

type ProxyCfg struct {
	Name          string   `yaml:"name"`
	StartProtocal int      `yaml:"startProtocal"`
	EndProtocal   int      `yaml:"endProtocal"`
	Strategy      string   `yaml:"strategy"`
	Addr          []string `yaml:"addr"`
	Maxgrpc       int      `yaml:"maxgrpc"`
	Mingrpc       int      `yaml:"mingrpc"`
}
