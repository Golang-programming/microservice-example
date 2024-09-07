package routes

import (
	"api-gateway/config"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.Any("/user/*action", proxyToUserService)
	router.Any("/product/*action", proxyToProductService)
	router.Any("/order/*action", proxyToOrderService)
}

func proxyToUserService(c *gin.Context) {
	proxyRequest(c, config.Cfg.UserServiceURL)
}

func proxyToProductService(c *gin.Context) {
	proxyRequest(c, config.Cfg.ProductServiceURL)
}

func proxyToOrderService(c *gin.Context) {
	proxyRequest(c, config.Cfg.OrderServiceURL)
}

func proxyRequest(c *gin.Context, target string) {
	targetURL, err := url.Parse(target)
	fmt.Println("targetURL", targetURL)
	if err != nil {
		log.Printf("Error parsing target URL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	originalWriter := c.Writer
	c.Writer = &responseCaptureWriter{ResponseWriter: originalWriter, body: &bytes.Buffer{}}

	c.Request.URL.Path = c.Param("action")
	fmt.Println("action", c.Param("action"))

	proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
		log.Printf("Error when proxying to %s: %v", target, e)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Bad Gateway", "message": "Upstream server is unreachable"})
	}

	proxy.ServeHTTP(c.Writer, c.Request)

	responseBody := c.Writer.(*responseCaptureWriter).body.String()
	fmt.Println("responseBody", responseBody)
	log.Printf("Response from %s: %s", target, responseBody)
}

type responseCaptureWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseCaptureWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
