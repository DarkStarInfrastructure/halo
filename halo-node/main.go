package main

import (
	"halo/middlewares"
	"halo/routers"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func setupLogging(logPath *string) {
	// Logging to a file.
	f, _ := os.Create(*logPath)
	// Logging to a file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	var (
		listenAddress = kingpin.Flag(
			"halo-node-listen-address",
			"The address to listen on for HTTP requests.",
		).Default("0.0.0.0:8888").String()
		logPath = kingpin.Flag(
			"halo-node-log-path",
			"The path to the log file.",
		).Default("/var/log/halo-node.log").String()
		enableBasicAuth = kingpin.Flag(
			"halo-node-enable-basic-auth",
			"Enable basic authentication.",
		).Default("false").Bool()
		basicAuthUser = kingpin.Flag(
			"halo-node-basic-auth-user",
			"The username for basic authentication.",
		).Default("admin").String()
		basicAuthPassword = kingpin.Flag(
			"halo-node-basic-auth-password",
			"The password for basic authentication.",
		).Default("admin").String()
	)
	kingpin.Parse()

	setupLogging(logPath)

	//
	router := gin.New()

	// load middlewares
	router.Use(gin.Recovery())
	router.Use(middlewares.Logger())
	if *enableBasicAuth {
		router.Use(gin.BasicAuth(gin.Accounts{*basicAuthUser: *basicAuthPassword}))
	}

	routers.DiskRouters(router)

	router.Run(*listenAddress)

}
