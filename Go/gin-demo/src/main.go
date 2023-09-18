package main

import (
	"fmt"
	"log"
	"os"

	"evision.com/go-demo/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Father struct {
	X string
}

func (f *Father) Say() {
	fmt.Printf("hello boy")
}

type Son struct {
	Father
	money int64
}

type Logger interface {
	Info()
	Warn()
	Error()
}

type StdLogger struct {
}

func (l *StdLogger) Info() {
	fmt.Print("info")
}

func (l *StdLogger) Warn() {
	fmt.Print("warn")
}

func (l *StdLogger) Error() {
	fmt.Print("error")
}

func init() {
	fmt.Print("我是第一个被执行的东西")
}

func main() {

	logrus.SetOutput(os.Stdout)
	//logrus.SetFormatter(&logrus.JSONFormatter{
	//	TimestampFormat:  "2006-01-02 15:04:05",
	//	DisableTimestamp: false,
	//})
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(true)

	xson := Son{Father{
		X: "",
	}, 11}

	/*son := Son{
		money: 1000,
	}*/

	xson.Say()
	fmt.Print(xson.money)

	engine := gin.Default()

	/// 拦截器
	// 1。 exceptionHandler
	engine.Use(func(context *gin.Context) {
		defer func() {
			e := recover()
			if e != nil {
				context.JSON(200, "系统异常")
				log.Print("222\n")
				context.Abort()
				return
			}
		}()
		context.Next()
	})

	/// 2. Token filter
	/*engine.Use(func(context *gin.Context) {
		token := context.GetHeader("Token")
		if token == "" {
			context.JSON(200, "token异常")
			context.Abort()
			log.Print("111\n")
			return
		}
		context.Next()
	})*/

	/*engine.GET("/user", func(context *gin.Context) {
		//context.JSON(200, "hello world")
		panic(errors.New("fuck"))
	})*/

	//userEngine := engine.Group("/user")
	//router.NewUserRouter(userEngine)
	userRouter := router.NewUserRouter(engine)
	userRouter.Handle()

	err := engine.Run(":7004")
	//err := http.ListenAndServe(":7004", engine)
	if err != nil {
		log.Fatalln("listen server with port 7004 error", err)
	}
}
