package logger

type Config struct {
	Level      int    `mapstructure:"level" json:"level" yaml:"level"`
	Identifier string `mapstructure:"identifier" json:"identifier" yaml:"identifier"`
	Elastic    struct {
		Addresses     []string `mapstructure:"addresses" json:"addresses" yaml:"addresses"`
		Username      string   `mapstructure:"username" json:"username" yaml:"username"`
		Password      string   `mapstructure:"password" json:"password" yaml:"password"`
		CAFingerprint string   `mapstructure:"ca" json:"ca" yaml:"ca"`
		Index         string   `mapstructure:"index" json:"index" yaml:"index"`
	} `mapstructure:"elastic" json:"elastic" yaml:"elastic"`
	Methods struct {
		LvFatal string `mapstructure:"fatal" json:"fatal" yaml:"fatal"`
		LvError string `mapstructure:"error" json:"error" yaml:"error"`
		LvWarn  string `mapstructure:"warn" json:"warn" yaml:"warn"`
		LvInfo  string `mapstructure:"info" json:"info" yaml:"info"`
		LvDebug string `mapstructure:"debug" json:"debug" yaml:"debug"`
	} `mapstructure:"methods" json:"methods" yaml:"methods"`
}

func Init(cfg *Config) {
	Level = cfg.Level
	Identifier = cfg.Identifier

	if len(cfg.Elastic.Addresses) > 0 {
		if e := InitElastic(&ElasticConfig{
			Addresses:     cfg.Elastic.Addresses,
			Username:      cfg.Elastic.Username,
			Password:      cfg.Elastic.Password,
			CAFingerprint: cfg.Elastic.CAFingerprint,
			Index:         cfg.Elastic.Index,
		}); e != nil {
			Warn("Failed to init elastic logger: %s", e.Error())
		} else {
			Debug("Elastic logger initialized")
			if cfg.Methods.LvFatal == "elastic" {
				LogMethod[FATAL] = LogElastic
			}
			if cfg.Methods.LvError == "elastic" {
				LogMethod[ERROR] = LogElastic
			}
			if cfg.Methods.LvWarn == "elastic" {
				LogMethod[WARN] = LogElastic
			}
			if cfg.Methods.LvInfo == "elastic" {
				LogMethod[INFO] = LogElastic
			}
			if cfg.Methods.LvDebug == "elastic" {
				LogMethod[DEBUG] = LogElastic
			}
		}
	}
}
