package data

// import (
//  "github.com/cc14514/go-geoip2-db"
//  "github.com/labstack/echo/v4"
//  "github.com/labstack/echo/v4/middleware"
//  "github.com/thinkeridea/go-extend/exnet"
//  "io/ioutil"
//  "log"
//  "net"
//  "net/http"
//  "os"
// )

// // ReDirectConfig defines the config for Redirect middleware.
// type ReDirectConfig struct {
//  // Skipper defines a function to skip middleware.
//  middleware.Skipper

//  // Status code to be used when redirecting the request.
//  // Optional. Default value http.StatusMovedPermanently.
//  Code int `yaml:"code"`
// }

// // DefaultRedirectConfig is the default Redirect middleware config.
// var DefaultRedirectConfig = ReDirectConfig{
//  Skipper: middleware.DefaultSkipper,
//  Code:    http.StatusFound,
// }

// // NoticeRedirect redirects non www requests to www.
// // For example, http://labstack.com will be redirect to http://www.labstack.com/notice.html.
// // Usage `Echo#Pre(WWWRedirect())`
// func NoticeRedirect() echo.MiddlewareFunc {
//  return NoticeRedirectWithConfig(DefaultRedirectConfig)
// }

// // NoticeRedirectWithConfig returns an HTTPSRedirect middleware with config.
// // See `WWWRedirect()`.
// func NoticeRedirectWithConfig(config ReDirectConfig) echo.MiddlewareFunc {
//  return redirect(config)
// }

// func redirect(config ReDirectConfig) echo.MiddlewareFunc {
//  if config.Skipper == nil {
//   config.Skipper = DefaultRedirectConfig.Skipper
//  }
//  if config.Code == 0 {
//   config.Code = DefaultRedirectConfig.Code
//  }
//  db, err := geoip2db.NewGeoipDbByStatik()
//  if err != nil {
//   log.Println(err.Error())
//  }
//  return func(next echo.HandlerFunc) echo.HandlerFunc {
//   return func(c echo.Context) error {
//    if config.Skipper(c) {
//     return next(c)
//    }
//    req := c.Request()
//    publicIP := exnet.ClientPublicIP(req)
//    if publicIP != "" && db != nil {
//     record, _ := db.City(net.ParseIP(publicIP))
//     if record.Country.IsoCode == "CN" {
//      log.Printf("IP(%s) from China is forbidden to access!", publicIP)
//      f, err := os.Open("./static/notice.html")
//      defer func() {
//       _ = f.Close()
//      }()
//      if err != nil {
//       log.Println(err.Error())
//       return next(c)
//      }
//      b, err := ioutil.ReadAll(f)
//      if err != nil {
//       log.Println(err.Error())
//       return next(c)
//      }
//      c.Response().Header().Set("cache-control", "private, max-age=0, no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
//      return c.Blob(http.StatusForbidden, "text/html; charset=utf-8", b)
//     }
//    }
//    return next(c)
//   }
//  }
// }
