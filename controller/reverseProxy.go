package controller

import (
	"fmt"
	pb "github.com/FakJeongTeeNhoi/api-gateway/generated/space"
	"github.com/FakJeongTeeNhoi/api-gateway/grpcClient"
	"github.com/FakJeongTeeNhoi/api-gateway/model/response"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

var routesMapping = map[string]string{}

func addReverseProxyRoutes() {
	routesMapping["/user"] = os.Getenv("UMS_URL")
	routesMapping["/space"] = os.Getenv("CMS_URL")
	routesMapping["/reserve"] = os.Getenv("RMS_URL")
	routesMapping["/report"] = os.Getenv("RPS_URL")
	routesMapping["/grpc/space"] = os.Getenv("GRPC_URL")
	routesMapping["/grpc/room"] = os.Getenv("GRPC_URL")
}

func InitReverseProxyRoutes(server *gin.Engine) {
	server.Any("/*path", reverseProxy)
	addReverseProxyRoutes()
}

func stripPath(path string) (string, string) {
	// path must start with some in routesMapping
	for key, value := range routesMapping {
		// check the prefix
		if (len(path) == len(key) && path[:len(key)] == key) ||
			(len(path) > len(key) && path[:len(key)] == key && path[len(key)] == '/') {
			return path[len(key):], value
		}
	}
	return path, ""
}

func reverseProxy(c *gin.Context) {
	stripedPath, targetURL := stripPath(c.Param("path"))
	if targetURL == "" {
		response.BadRequest("Invalid path").AbortWithError(c)
		return
	}

	if c.Param("path") == "/grpc/space" {
		// grpcClient
		// if method is POST, call create
		if c.Request.Method == "POST" {
			createSpaceRequest := &pb.CreateSpaceRequest{}
			err := c.BindJSON(createSpaceRequest)
			if err != nil {
				response.BadRequest("Failed to bind request").AbortWithError(c)
				return
			}

			resp, err := grpcClient.SpaceClient.Client.CreateSpace(c, createSpaceRequest)
			if err != nil {
				response.InternalServerError(err.Error()).AbortWithError(c)
				return
			}

			c.JSON(http.StatusOK, resp)
		}

		return
	}

	url := targetURL + stripedPath + "?" + c.Request.URL.RawQuery

	// call target URL
	request, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		return
	}

	request.Header = c.Request.Header

	request.Body = c.Request.Body

	if user, ok := c.Get("user"); ok {
		for key, value := range user.(map[string]interface{}) {
			request.Header.Add(key, fmt.Sprint(value))
		}
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		response.InternalServerError("Failed to call target URL").AbortWithError(c)
		return
	}

	// copy response
	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}
	c.Status(resp.StatusCode)
	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		response.InternalServerError("Failed to copy response").AbortWithError(c)
		return
	}
}
