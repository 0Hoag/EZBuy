package main

import (
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

// Handler function để forward request đến service tương ứng
func createProxyHandler(targetURL string) gin.HandlerFunc {
	target, err := url.Parse(targetURL)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	return func(c *gin.Context) {
		// Set host của request về service đích
		c.Request.URL.Host = target.Host
		c.Request.URL.Scheme = target.Scheme
		c.Request.Header.Set("X-Forwarded-Host", c.Request.Header.Get("Host"))

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()

	// Service URLs
	authServiceURL := "http://localhost:8081"
	productServiceURL := "http://localhost:8082"

	authPaths := []string{
		"/v1/api/user/*path",
		"/v1/api/role/*path",
		"/v1/api/permission/*path",
		"/v1/api/auth/*path",
	}

	for _, path := range authPaths {
		router.Any(path, createProxyHandler(authServiceURL))
	}

	productPaths := []string{
		"/v1/api/categories/*path",
		"/v1/api/products/*path",
	}

	for _, path := range productPaths {
		router.Any(path, createProxyHandler(productServiceURL))
	}

	router.Run(":8080")
}
