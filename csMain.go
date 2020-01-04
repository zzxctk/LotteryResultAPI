package main

import (
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"github.com/kirinlabs/HttpRequest"
	"github.com/zllangct/zgo/logger"
	"sync"
	"time"
)

type InfoBox struct {
	Data map[string]interface{}
	lock sync.RWMutex
}

func (this *InfoBox) Get() map[string]interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.Data
}

func (this *InfoBox) Set(m map[string]interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.Data = m
}

func main() {
	info := new(InfoBox)
	go func() {
		for {
			time.Sleep(time.Second * 5)
			c := HttpRequest.NewRequest()
			c.SetTimeout(time.Second * 7)
			if res, err := c.Get("http://www.caipiaozixun888.com/Json/getylkw.aspx?m=hqylkw"); err != nil {
				logger.Error(err.Error())
			} else {
				if b, err := res.Body(); err != nil {
					logger.Error(err.Error())
				} else {
					m := make(map[string]interface{})
					if err := json.Unmarshal(b, &m); err != nil {
						logger.Error(err.Error())
					} else {
						info.Set(m)
					}
				}
			}
		}
	}()
	router := gin.New()
	router.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	router.GET("/getYLK5", func(c *gin.Context) {
		c.JSON(200, info.Get())
		c.Status(200)
	})
	router.Run(":1234")

	ch := make(chan bool)
	for {
		ok := <-ch
		if ok {
			return
		}
	}
}
