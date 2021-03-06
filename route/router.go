package route

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/moh-fajri/learn-jwt/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Route for mapping from json file
type Route struct {
	Path       string   `json:"path"`
	Method     string   `json:"method"`
	Module     string   `json:"module"`
	Endpoint   string   `json:"endpoint_filter"`
	Middleware []string `json:"middleware"`
}

// Init gateway router
func Init() *echo.Echo {
	routes := loadRoutes("./route/gate/")

	e := echo.New()
	// Set Bundle MiddleWare
	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength,
			echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "X-Request-Id", "device-id", "X-Summary", "X-Account-Number", "X-Business-Name", "client-secret", "client-key", "x-csrf-token", "api-key"},
		ExposeHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength,
			echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "X-Request-Id", "device-id", "X-Summary", "X-Account-Number", "X-Business-Name", "client-secret", "client-key", "x-csrf-token", "api-key"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	prefix := "/api"

	for _, route := range routes {
		e.Add(route.Method, prefix+route.Path, endpoint[route.Endpoint].Handle, chainMiddleware(route.Middleware)...)
	}

	util.CustomErrorHandler(e)

	return e
}

func loadRoutes(filePath string) []Route {
	var routes []Route
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatalf("Failed to load file: %v", err)
	}
	for _, file := range files {
		byteFile, err := ioutil.ReadFile(filePath + "/" + file.Name())
		if err != nil {
			log.Fatalf("Failed to load file: %v", err)
		}
		var tmp []Route
		if err := json.Unmarshal(byteFile, &tmp); err != nil {
			log.Fatalf("Failed to marshal file: %v", err)
		}
		routes = append(routes, tmp...)
	}

	return routes

}

func chainMiddleware(tags []string) []echo.MiddlewareFunc {
	mwHandlers := []echo.MiddlewareFunc{}
	for _, v := range tags {
		mwHandlers = append(mwHandlers, middlewareHandler[v])
	}
	return mwHandlers
}
