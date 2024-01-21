package main

import (
	"embed"
	"fmt"
	httpLib "net/http"
	"os"
	"system-point/delivery/http"
	"system-point/domain"

	"github.com/gofiber/template/html/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

//go:embed resource/*
var resourcefs embed.FS

func main() {
	domain := domain.ConstructDomain()
	engine := html.NewFileSystem(httpLib.FS(resourcefs), ".html")
	app := http.NewHttpDelivery(domain, engine)
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT_APP")))
}
