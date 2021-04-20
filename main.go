package main

import (
	"fmt"

	"net/http"
	"os"
	"sync"
	"time"

	route "github.com/adefemi171/goHelm/pkg/route"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
)

var (
	g            errgroup.Group
	reqProcessed = promauto.NewCounter(prometheus.CounterOpts{
		//name of number of requests metric
		Name: "app_number_of_requests",
		Help: "The total number of processed requests",
	})
)

func webRoute() http.Handler {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		//increase number of processed request
		reqProcessed.Inc()
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	route.UserRoutes(r)

	return r
}

func prometheusRoute() http.Handler {
	srv := http.DefaultServeMux
	srv.Handle("/metrics", promhttp.Handler())
	return srv
}

// main function
func main() {
	webServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", GetEnv("APP_PORT", "8080")),
		Handler:      webRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	prometheusServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", GetEnv("PROMETHEUS_PORT", "9110")),
		Handler:      prometheusRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		webServer.ListenAndServe()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		prometheusServer.ListenAndServe()
		wg.Done()
	}()

	wg.Wait()
}

//GetEnv Utility to retrieve env and provide default
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
