package core

import (
	"backend-learning/hw2/global"
	"backend-learning/hw2/initialize"
	"fmt"
	"net/http"
	"time"
)

func Run() {

	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Debug("server run success on ", address)

	global.LOG.Info(s.ListenAndServe())
}
