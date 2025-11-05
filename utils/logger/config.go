package logger

type Config struct {
	Enable    bool   `mapstructure:"enable"`
	Level     string `mapstructure:"level"`
	FileName  string `mapstructure:"filename"`
	DirPath   string `mapstructure:"dir-path"`
	MaxSize   int    `mapstructure:"max-size"`
	MaxBackup int    `mapstructure:"max-backup"`
	MaxAge    int    `mapstructure:"max-age"`
	LocalTime bool   `mapstructure:"local-time"`
	Compress  bool   `mapstructure:"compress"`
}

func DefaultConfig() *Config {
	return &Config{
		Enable:    false,
		Level:     "DEBUG",
		FileName:  "",
		DirPath:   "logs",
		MaxSize:   64,
		MaxBackup: 3,
		MaxAge:    28,
		LocalTime: true,
		Compress:  false,
	}
}
