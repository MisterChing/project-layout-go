package config

import (
	"github.com/spf13/cast"
	"io/ioutil"
	"project-layout-go/internal/infrastructure/bootstrap"
	"project-layout-go/pkg/configx"
)

var (
	EndpointPool = make(map[string]*ServiceEndPoint)
)

type endpointConf map[string]interface{}

type ServiceEndPoint struct {
	name string
	conf endpointConf
}

func (endpoint *ServiceEndPoint) GetString(key string) string {
	if v, ok := endpoint.conf[key]; ok {
		return cast.ToString(v)
	}
	return ""
}

func (endpoint *ServiceEndPoint) GetBool(key string) bool {
	if v, ok := endpoint.conf[key]; ok {
		return cast.ToBool(v)
	}
	return false
}

func (endpoint *ServiceEndPoint) GetInt(key string) int {
	if v, ok := endpoint.conf[key]; ok {
		return cast.ToInt(v)
	}
	return 0
}

func GetServiceEndpoint(name string) *ServiceEndPoint {
	if v, ok := EndpointPool[name]; ok {
		return v
	}
	return nil
}

func InitServicesConfig() bootstrap.BeforeServerFunc {
	return func() error {
		servicesConfPath := ApplicationRoot + "/conf/services"
		fileList, err := ioutil.ReadDir(servicesConfPath)
		if err != nil {
			return err
		}
		if len(fileList) > 0 {
			for _, file := range fileList {
				tmpCfg, err := configx.NewCfgFromFile(servicesConfPath + "/" + file.Name())
				if err != nil {
					return err
				}
				tmpEndpoint := new(ServiceEndPoint)
				tmpEndpoint.name = tmpCfg.GetString("service.name")
				tmpEndpoint.conf = tmpCfg.GetStringMap("endpoint." + Env)
				EndpointPool[tmpEndpoint.name] = tmpEndpoint
			}
		}
		return nil
	}
}
