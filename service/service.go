package service

import (
	"fmt"

	"gitee.com/go-apiServer/libs"
	"gitee.com/go-apiServer/model"
	"github.com/kataras/iris/v12"
)

type service struct {
	Property model.Property
}

func NewService(property model.Property) *service {
	return &service{
		Property: property,
	}
}

func (s *service) handler(ctx iris.Context) {
	url := ctx.Request().URL
	fmt.Println(url)
	contentByte, err := libs.ReadJSON(fmt.Sprintf("%s/%s", s.Property.DataPath, url.Path))
	if err != nil {
		ctx.JSON(contentByte)
		return
	}
	ctx.NotFound()
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

	// handle xx.json
	app.Get("/{name:string regexp(.*.json$)}", s.handler)
	app.Put("/{name:string regexp(.*.json$)}", s.handler)
	app.Delete("/{name:string regexp(.*.json$)}", s.handler)
	// handle xx.post.json
	app.Post("/{name:string regexp(.*.post.json$)}", s.handler)

	return app
}

func (s *service) Start() {
	app := s.newApp()
	app.Logger().SetLevel(s.Property.DebugLevel)
	app.Listen(fmt.Sprintf(":%s", s.Property.Port))
}
