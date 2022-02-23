package config

import (
    "fmt"
    "os"
    "project-layout-go/internal/infrastructure/bootstrap"
    "project-layout-go/pkg/configx"
    "project-layout-go/pkg/utils/debugutil"
)

var (
    // DevelopEnv 开发环境
    DevelopEnv = "develop"

    // TestEnv 测试环境
    TestEnv = "test"

    // GrayEnv 灰度环境
    GrayEnv = "gray"

    // ProductEnv 生产环境
    ProductEnv = "product"
)

var (
    // AppCfg app配置实例
    AppCfg *configx.Cfg

    // Env 环境变量
    Env string

    // Addr http服务地址
    Addr string

    // ApplicationRoot 项目跟目录
    ApplicationRoot string

    // AppName 项目名称
    AppName string

    serverCfg *srvCfg
)

type srvCfg struct {
    Name string `toml:"name"`
    Env  string `mapstructure:"environment"`
    Addr string `toml:"addr"`
}

func SetUp() bootstrap.BeforeServerFunc {
    return func() error {
        //// 设置项目根目录
        //ApplicationRoot = strings.TrimRight(*flagutil.GetConfigPrefix(), "/") + "/"

        // load config
        ApplicationRoot, _ = os.Getwd()
        debugutil.DebugPrint(ApplicationRoot, 0)
        tmpCfg, err := configx.NewCfgFromFile(ApplicationRoot + "/conf/app.toml")
        if err != nil {
            return err
        }
        AppCfg = tmpCfg
        if err := AppCfg.UnmarshalKey("server", &serverCfg); err != nil {
            return err
        }
        AppName = serverCfg.Name
        Env = serverCfg.Env
        Addr = serverCfg.Addr
        if Env != DevelopEnv && Env != TestEnv && Env != GrayEnv && Env != ProductEnv {
            return fmt.Errorf("environment配置的值错误，必须为下面几个值之一：develop, test, gray, product")
        }
        return nil
    }
}
