package bootstrap

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/spf13/viper"
	"time"
	"ucenter/s/core"
	"ucenter/s/db"
	"ucenter/s/errors"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

// New returns a new Bootstrapper.
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// set config
func SetupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/" + core.GetEnviron())
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func SetupInitDb() {
	db.Init()
}

// SetupViews loads the templates.
func (b *Bootstrapper) SetupViews(viewsDir string) {
	b.RegisterView(iris.HTML(viewsDir, ".html").Layout("shared/layout.html").Reload(!core.IsProudctEnv()))
}

// `(context.StatusCodeNotSuccessful`,  which defaults to >=400 (but you can change it).
func (b *Bootstrapper) SetupErrorHandlers() {
	b.APIBuilder.ConfigureContainer().Container.GetErrorHandler = func(*context.Context) hero.ErrorHandler {
		return hero.ErrorHandlerFunc(func(ctx *context.Context, err error) {
			if err != hero.ErrStopExecution {
				if status := ctx.GetStatusCode(); status == 0 || !context.StatusCodeNotSuccessful(status) {
					ctx.StatusCode(hero.DefaultErrStatusCode)
				}

				if isOutJson(ctx) {
					ctx.ContentType(context.ContentJSONHeaderValue)
				}
				_, _ = ctx.WriteString(err.Error())
			}

			ctx.StopExecution()
		})
	}

	b.OnAnyErrorCode(func(ctx iris.Context) {
		res := errors.Errors{}
		res.SetErrCode(ctx.GetStatusCode())
		res.SetMessage(iris.StatusText(ctx.GetStatusCode()))

		if isOutJson(ctx) {
			ctx.JSON(res)
			return
		}

		err := iris.Map{
			"errCode": res.GetErrCode(),
			"message": res.GetMessage(),
		}
		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

func isOutJson(ctx iris.Context) bool {
	return ctx.IsAjax() ||
		ctx.URLParamExists("json") ||
		ctx.GetHeader("Accept") == context.ContentJSONHeaderValue
}

const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./web/public/"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
)

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./web/views")
	b.SetupErrorHandlers()

	// static files
	b.Favicon(StaticAssets + Favicon)
	b.HandleDir("/public", iris.Dir(StaticAssets))

	// middleware, after static files
	b.Use(recover.New())
	b.Use(logger.New())

	SetupConfig()
	SetupInitDb()
	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
