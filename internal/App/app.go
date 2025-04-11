package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewAppCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "流量采集器",
		Long:  "一个go开发的流量采集器",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runApp()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return errors.New("参数不能为空")
				}
			}
			return nil
		},
	}

	cobra.OnInitialize(initConfig)
	return cmd
}

func runApp() error {
	// 在这里实现你的应用程序逻辑
	g := gin.New()
	if err := initStore(); err != nil {
		return err
	}

	gin.SetMode(viper.GetString("runmode"))

	if err := InitRouter(g); err != nil {
		return err
	}

	httpSrv := runHttpSecureService(g)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func runHttpSecureService(g *gin.Engine) *http.Server {
	// 在这里实现你的应用程序逻辑
	httpSrv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: g,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return httpSrv
}
