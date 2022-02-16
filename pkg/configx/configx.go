package configx

import (
	"github.com/spf13/viper"
)

type Cfg struct {
	*viper.Viper
}

func NewCfgFromFile(file string) (*Cfg, error) {
	v := viper.New()
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	cfg := &Cfg{
		v,
	}
	return cfg, nil
}
