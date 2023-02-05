package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"thucidol/component/appctx"
	"thucidol/middleware"
	"thucidol/pubsub/localpubsub"
	"thucidol/subscriber"
)

func main() {

	// dsn := os.Getenv("DB_URL")
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	secretKey := os.Getenv("SYSTEM_SECRET")
	db, err := gorm.Open(postgres.Open("host=127.0.0.1 user=postgres password=YeRjCHvgxw22 dbname=gorm port=5432 sslmode=disable "), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db, err)
	fmt.Println("? Connected Successfully to the Database")
	ps := localpubsub.NewPubSub()
	appCtx := appctx.NewAppCtx(db, secretKey, ps)

	//setup subscriber
	// subscriber.Setup(appCtx,context.Background())
	_ = subscriber.NewEngine(appCtx).Start()

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	//POST
	v1 := r.Group("/v1")
	r.Static("/static", "./static")

	setupRoute(appCtx, v1)
	setupAdminRoute(appCtx, v1)

	r.Run()

}
