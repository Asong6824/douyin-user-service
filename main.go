package main

import (
	"log"
	user "github.com/Asong6824/douyin-user-service/kitex_gen/user/userhandler"
)

func main() {
	sc := []constant.ServerConfig{
        *constant.NewServerConfig("127.0.0.1", 8848),
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
    
    svr := user.NewServer(
		new(UserHandlerImpl),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: global.EngineSetting.Registry.ServiceName}),
		server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}),
    )
    if err := svr.Run(); err != nil {
        log.Println("server stopped with error:", err)
    } else {
        log.Println("server stopped")
    }

}
