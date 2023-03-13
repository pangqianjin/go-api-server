package service

import (
	"fmt"
	"regexp"

	"gitee.com/go-apiServer/conf"
	"gitee.com/go-apiServer/libs"
	"gitee.com/go-apiServer/logger"
	"gitee.com/go-apiServer/model"
	"github.com/kataras/iris/v12"
)

type service struct {
	Property model.Property
}

func NewService() *service {
	property := model.Property{}
	libs.ReadFileToModel(conf.PropertyFileName, &property)

	fmt.Printf("%#v\n", property)
	return &service{
		Property: property,
	}
}

func (s *service) handler(ctx iris.Context) {
	url := ctx.Request().URL
	fmt.Println(url)
	method := ctx.Method()
	path := fmt.Sprintf("%s/%s", s.Property.DataPath, url.Path)
	if method == "POST" {
		jsonReg := regexp.MustCompile(`.json`)
		path = jsonReg.ReplaceAllString(path, ".post.json") // replace .json to .post.json
	}

	contentByte, err := libs.ReadJSON(path)
	if err != nil {
		fmt.Println("err:", err)
		ctx.NotFound()
		return
	}
	ctx.JSON(contentByte)
}

func (s *service) newApp() *iris.Application {
	app := iris.New()

	// handle statics resources
	statics := app.Party("/")
	statics.HandleDir("/statics", iris.Dir(s.Property.PublicPath), iris.DirOptions{
		Compress:   false,
		ShowList:   false,
		ShowHidden: false,
		Cache: iris.DirCacheOptions{
			// enable in-memory cache and pre-compress the files.
			Enable: true,
			// do not compress files smaller than size.
			CompressMinSize: 300,
			// available encodings that will be negotiated with client's needs.
			Encodings: []string{"gzip", "br" /* you can also add: deflate, snappy */},
		},
		DirList: iris.DirListRich(),
	})

	return app
}

func (s *service) AddRoutes(app *iris.Application) {
	app.Any("/{request:string}", s.handler)
	app.Any("/{group:string}/{request:string}", s.handler)
	app.Any("/{prefix:string}/{group:string}/{request:string}", s.handler)

	return
}

func (s *service) Start() {
	logger := logger.CustomLogger()
	app := s.newApp()
	app.Use(logger)
	s.AddRoutes(app)
	app.Logger().SetLevel(s.Property.DebugLevel)
	app.Listen(fmt.Sprintf(":%s", s.Property.Port))
}
