// Package server ...
package server

import (
	"net/http"
	_ "net/http/pprof" // pprof debug

	"github.com/dup2X/crsdk/api/controller/cr"
	//	"github.com/dup2X/crsdk/api/controller/project"
	//	"github.com/dup2X/crsdk/api/controller/user"
	"github.com/dup2X/gopkg/config"
	"github.com/dup2X/gopkg/httpsvr"
)

// Serve 启动服务
func Serve(conf config.Configer) error {
	sec, err := conf.GetSection("http")
	if err != nil {
		return err
	}
	go func() {
		http.ListenAndServe(":6061", nil)
	}()
	return httpServe(sec)
}

func httpServe(sec config.Sectioner) error {
	addr := sec.GetStringMust("bind_addr", "127.0.0.1:10025")
	slatimeout := sec.GetIntMust("sla_timeout", 3000)
	svr := httpsvr.New(addr, httpsvr.WithDumpAccess(true), httpsvr.SetHandleDefaultTimeout(slatimeout))
	svr.AddMiddleware(httpsvr.RecoveryWithMetric)
	//	svr.AddRoute("POST", "/api/project/query", &project.QueryController{})                // 查询项目
	//	svr.AddRoute("POST", "/api/project/branch", &project.BranchController{})              // 查询分支
	//	svr.AddRoute("POST", "/api/project/branch/pull", &project.BranchPullController{})     // 更新分支
	//	svr.AddRoute("POST", "/api/project/branch/create", &project.BranchCreateController{}) // 创建分支
	//	svr.AddRoute("POST", "/api/project/branch/merge", &project.BranchMergeController{})   // Merge分支
	// svr.AddRoute("POST", "/api/cr/create", &cr.CreateController{}) // 创建CR
	svr.AddRoute("POST", "/api/cr/query", &cr.QueryController{}) // 查询CR状态

	return svr.Serve()
}
