package http

import (
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type ApplicationInfo struct {
	Application string `json:"application"`
	Author      string `json:"author"`
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PageHandler interface {
	RegistrationPageRoute(g *echo.Group)
}

type ApiHandler interface {
	RegistrationApiRoute(g *echo.Group)
}

type FilesHandler interface {
	RegistrationFilesRoute(g *echo.Group)
}
type ApiRouteGroup interface {
	RegistrationRouteApi(g *echo.Group)
}

type Server struct {
	info ApplicationInfo

	e *echo.Echo

	apiGroup  *echo.Group
	pageGroup *echo.Group
	fileGroup *echo.Group
}

func NewServer() *Server {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover(), middleware.CORS())
	e.File("/favicon.ico", "pkg/server/http/static/favicon.png")
	s := &Server{
		e:         e,
		apiGroup:  e.Group("/api"),
		pageGroup: e.Group(""),
		fileGroup: e.Group("/files"),
	}
	e.GET("/", s.mainPage)
	s.apiGroup.Use(errorApiMiddleware())
	s.pageGroup.Use(errorPageMiddleware())
	return s
}

func (s *Server) WithAuthor(author string) {
	s.info.Author = author
}

func (s *Server) WithApplicationName(name string) {
	s.info.Application = name
}

func (s *Server) WithFavicon(path string) {
	s.e.File("/favicon.ico", path)
}

func (s *Server) WithIndexPage(path string) {
	s.e.File("/", path)
}

func (s *Server) RegistrationPage(h PageHandler) {
	h.RegistrationPageRoute(s.pageGroup)
}

func (s *Server) RegistrationApi(h ApiHandler) {
	h.RegistrationApiRoute(s.apiGroup)
}

func (s *Server) RegistrationFilesHandler(h FilesHandler) {
	h.RegistrationFilesRoute(s.fileGroup)
}

func (s *Server) Run(host, port string) error {
	return s.e.Start(net.JoinHostPort(host, port))
}

func (s *Server) mainPage(c echo.Context) error {
	contentType := c.Request().Header.Get(echo.HeaderContentType)
	if contentType == echo.MIMEApplicationJSON {
		return c.JSON(http.StatusOK, s.info)
	}
	return c.String(http.StatusOK, "Index page")
}

func errorApiMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}
			httpErr, ok := err.(*echo.HTTPError)
			if ok && httpErr.Code == http.StatusNotFound {
				return c.JSON(http.StatusNotFound, nil)
			}
			zap.L().Error("error api http execute", zap.Error(err), zap.String("url", c.Request().URL.String()))
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
		}
	}
}

func errorPageMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}
			httpErr, ok := err.(*echo.HTTPError)
			if !ok {
				zap.L().Error("error http execute", zap.Error(err), zap.String("url", c.Request().URL.String()))
				return err
			}
			if httpErr.Code == http.StatusNotFound {
				return c.Redirect(http.StatusMovedPermanently, "/")
			}

			return err
		}
	}
}

func (s *Server) Close() {
	_ = s.e.Close()
}
