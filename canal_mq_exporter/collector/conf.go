package collector

import "github.com/spf13/viper"

type ServerConf struct {
	Host  string `json:"host"`
	Uri   string `json:"uri"`
	MqUrl string `json:"mqUrl"`
	Type  string `json:"type"`
}

var Conf = new(ServerConf)

func InitConf() error {
	var v *viper.Viper
	v = viper.New()
	v.SetConfigName("config.yaml")
	v.AddConfigPath("./")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	err := v.UnmarshalKey("serverConf", &Conf)
	if err != nil {
		return err
	}
	return nil
}
