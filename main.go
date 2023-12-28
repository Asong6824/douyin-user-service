package main

import (
	user "github.com/Asong6824/douyin-user-service/kitex_gen/user/userservice"
	"net"
	"os"
	"io"
	"github.com/nacos-group/nacos-sdk-go/clients"
    "github.com/nacos-group/nacos-sdk-go/common/constant"
    "github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/Asong6824/douyin-user-service/global"
	"github.com/Asong6824/douyin-user-service/package/setting"
	"log"
)

func Init() {
	err := setupSetting()
	if err != nil {
		panic(err)
	}
	err = setupLogger()
	if err != nil {
		panic(err)
	}
}

func main() {
	Init()
	sc := []constant.ServerConfig{
        *constant.NewServerConfig("172.17.0.2", 8848),
    }
    
    cc := constant.ClientConfig{
        NamespaceId:         "public",
        TimeoutMs:           5000,
        NotLoadCacheAtStart: true,
        LogDir:              "/tmp/nacos/log",
        CacheDir:            "/tmp/nacos/cache",
        LogLevel:            "info",
    }
    
    cli, err := clients.NewNamingClient(
        vo.NacosClientParam{
            ClientConfig:  &cc,
            ServerConfigs: sc,
        },
    )
    if err != nil {
        panic(err)
    }
	addr, err := net.ResolveTCPAddr("tcp", global.EngineSetting.HostPort)
    if err != nil {
        panic(err)
    }
    svr := user.NewServer(
		new(UserServiceImpl),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: global.EngineSetting.ServiceName}),
		server.WithServiceAddr(addr),
    )
    if err := svr.Run(); err != nil {
        log.Println("server stopped with error:", err)
    } else {
        log.Println("server stopped")
    }

}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Engine", &global.EngineSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    fileWriter := io.MultiWriter(f,os.Stdout)
    klog.SetOutput(fileWriter)
	return nil
}
