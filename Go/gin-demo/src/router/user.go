package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	fmt.Printf("heheheheheh")
}

type UserRouter struct {
	engine *gin.Engine
}

func NewUserRouter(engine *gin.Engine) *UserRouter {
	return &UserRouter{
		engine: engine,
	}
}

func (u *UserRouter) Handle() {
	userGroup := u.engine.Group("/user")

	userGroup.GET("/login", u.Login)
	userGroup.GET("/register", u.Register)
}

func (u *UserRouter) Login(ctx *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"name": "asheng",
		"age":  21,
	}).Info("hello boy. this is login api")
}

func (u *UserRouter) Register(ctx *gin.Context) {
}
