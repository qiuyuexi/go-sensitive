package server

import (
	"context"
	"fmt"
	"go-sensitive/internal/app/api"
	"go-sensitive/internal/pkg/ahocorasick"
	"go-sensitive/internal/pkg/config"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var signalChan chan os.Signal
var httpServer http.Server

//服务启动初始化
func Start(port int, configPath string) {

	//配置目录设置
	config.SetWorkPath(configPath)

	//构建ac自动机
	ahocorasick.BuildAhocorasickDict()

	//监听敏感词是否更新
	wordsWatch()

	//信号注册
	registerSignal()

	//记录进程id
	logPid()

	//开启http服务
	httpServerStart(port)

	//信号监听
	watchSignal()
}

//信号注册
func registerSignal() {
	signalChan = make(chan os.Signal, 10)
	var signalList = []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
	}
	signal.Notify(signalChan, signalList...)
}

//记录进程id
func logPid() {
	pid := os.Getpid()
	msg := fmt.Sprintf("服务启动,pid:%d", pid)
	fmt.Println(msg)
}

//开启http服务
func httpServerStart(port int) {
	go func() {
		http.HandleFunc("/filter", api.FilterHandel);
		httpServer = http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: http.DefaultServeMux,
		}
		httpServer.ListenAndServe()
	}()
	return
}

//监听信号
func watchSignal() {
	signalVal := <-signalChan
	fmt.Println("信号内容:", signalVal)
	httpServerShutdown()
}

//收到退出信号，http服务退出
func httpServerShutdown() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	httpServer.Shutdown(ctx)
	fmt.Println("服务关闭")
	return
}
