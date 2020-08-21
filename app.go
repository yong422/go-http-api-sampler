package main

import (
	"context"
	"fmt"
	"github.com/ip2location/ip2location-go"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sampler/api"
	"sampler/api/model"
	"sampler/config"
	"time"
)

type App struct {
	config *config.SamplerConfiguration
	server *http.Server
}

func (a *App) Initialize() {
	var err error
	a.config, err = config.RoadConfiugurationFile("config/sampler_config.json")
	if err != nil {
		panic(err)
	}

	for index, node := range a.config.GetRedisClusterAddresses() {
		fmt.Println(index, " > ", node)
	}

	apiRouterHandler := api.CreateApiRouter()

	a.server = &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      apiRouterHandler, // Pass our instance of gorilla/mux in.
	}

	// model setter
	model.SetGeoIpWebServiceAuthorizationInfo(a.config.GeoIpWebServiceAccountId,
		a.config.GeoIpWebServiceLicenseKey)

	geoipReader, err := geoip2.Open(a.config.GeoIp2CityDatabase)
	if err != nil {
		// geoip file read error
		fmt.Println(err)
	}
	model.SetGeoIp(geoipReader)

	db, dbErr := ip2location.OpenDB(a.config.IpToLocationDatabase)
	if dbErr != nil {
		// geoip file read error
		fmt.Println(dbErr)
	} else {
		fmt.Println("open success > ip2location file > ", a.config.IpToLocationDatabase)
	}
	model.SetIpToLocation(db)
}

func (a *App) Run() int32 {
	var wait time.Duration
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err := a.server.Shutdown(ctx)
	if err != nil {
		panic(err)
	}

	return 1
}
