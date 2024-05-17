package logger

type Config struct {
	Level      int    `mapstructure:"level" json:"level"`
	Identifier string `mapstructure:"identifier" json:"identifier"`
	Elastic    struct {
		Addresses     []string `mapstructure:"addresses" json:"addresses"`
		Username      string   `mapstructure:"username" json:"username"`
		Password      string   `mapstructure:"password" json:"password"`
		CAFingerprint string   `mapstructure:"ca" json:"ca"`
		Index         string   `mapstructure:"index" json:"index"`
	} `mapstructure:"elastic" json:"elastic"`
	Methods struct {
		LvFatal string `mapstructure:"fatal" json:"fatal"`
		LvError string `mapstructure:"error" json:"error"`
		LvWarn  string `mapstructure:"warn" json:"warn"`
		LvInfo  string `mapstructure:"info" json:"info"`
		LvDebug string `mapstructure:"debug" json:"debug"`
	} `mapstructure:"methods" json:"methods"`
}

func Init(cfg *Config) {
	Level = cfg.Level
	Identifier = cfg.Identifier

	if cfg.Elastic.Addresses != nil {
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
