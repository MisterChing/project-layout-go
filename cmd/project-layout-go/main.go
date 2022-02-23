package main

import (
    "context"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "os"
    "os/signal"
    "project-layout-go/internal/infrastructure/bootstrap"
    "project-layout-go/internal/infrastructure/common/cache"
    "project-layout-go/internal/infrastructure/common/config"
    "project-layout-go/internal/infrastructure/common/db"
    "project-layout-go/internal/router"
    "project-layout-go/pkg/osx"
    "project-layout-go/pkg/utils/debugutil"
    "syscall"
    "time"
)

func main() {
    // bootloader
    bootloader := bootstrap.NewBootStrap()

    bootloader.AddBeforeServerFunc(
        config.SetUp(),
        //logger.SetUp(),
        db.SetUp(),
        cache.SetUp(),
    )
    bootloader.AddAfterServerFunc(
        db.Destroy(),
        cache.Destroy(),
    )
    if err := bootloader.SetUp(); err != nil {
        log.Fatal(err)
    }
    defer bootloader.Destroy()

    // 设置进程最大打开文件描述符个数
    osx.SetRLimit()

    // set gin mode
    if config.Env == config.GrayEnv || config.Env == config.ProductEnv {
        gin.SetMode(gin.ReleaseMode)
    } else {
        gin.SetMode(gin.DebugMode)
    }
    engine := gin.Default()

    router.Init(engine)

    // Listen and Server in 0.0.0.0:8080
    serverAddr := config.Addr
    srv := &http.Server{
        Addr:    serverAddr,
        Handler: engine,
    }
    debugutil.DebugPrint(srv, 0)

    // Start Server
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()

    // hold the main goroutine
    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Server is shutting down...")

    // wait active connections to get done.
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown: ", err)
    }
    log.Println("Server gracefully shutdown.")
}
