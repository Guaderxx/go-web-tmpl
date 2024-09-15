package web

import (
	"fmt"
	"net/http"

	"github.com/Guaderxx/gowebtmpl/pkg/alog"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Routes(r *gin.Engine, logger alog.ALogger) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response[string]{
			Code: 200,
			Data: "pong",
		})
	})

	r.GET("/test", func(c *gin.Context) {
		var p Person
		if c.ShouldBind(&p) == nil {
			logger.WithGroup("person").
				Info("", "name", p.Name,
					"address", p.Address,
					"birthday", p.Birthday)
		}

		c.JSON(http.StatusOK, Response[Person]{
			Code: 200,
			Data: p,
		})
	})

	r.GET("/:name/:id", func(c *gin.Context) {
		var s SUri
		if err := c.ShouldBindUri(&s); err != nil {
			c.JSON(http.StatusOK, Response[any]{
				Code:  400,
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, Response[SUri]{
			Code: 200,
			Data: s,
		})
	})

	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBindWith(&b, binding.Query); err == nil {
			c.JSON(http.StatusOK, Response[gin.H]{
				Code: 200,
				Data: gin.H{
					"msg": "Booking date are valid.",
				},
			})
		} else {
			c.JSON(http.StatusOK, Response[any]{
				Code:  400,
				Error: err.Error(),
			})
		}
	})

	r.GET("/png", func(c *gin.Context) {
		res, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || res.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := res.Body
		contentLength := res.ContentLength
		contentType := res.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NoSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		logger.Info("cookie value", "key", "gin_cookie", "value", cookie)
	})

	r.POST("/single-file", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		logger.Info("upload file succeed", "filename", file.Filename)
		c.SaveUploadedFile(file, "dst-path")
		c.JSON(http.StatusOK, Response[string]{
			Code: 200,
			Data: fmt.Sprintf("`%s` uploaded.", file.Filename),
		})
	})

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/multifile", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			logger.Info("file upload succeed", "filename", file.Filename)

			c.SaveUploadedFile(file, "dst-path")
		}

		c.JSON(http.StatusOK, Response[string]{
			Code: 200,
			Data: fmt.Sprintf("`%d` files uploaded.", len(files)),
		})
	})
}
