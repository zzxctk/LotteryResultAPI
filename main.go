package main

import (
	"LotteryResultAPI/handleData"
	"bn/goPostMan"
	"fmt"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"github.com/kirinlabs/HttpRequest"
	"github.com/pkg/errors"
	"github.com/zllangct/zgo/logger"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type controlBox struct {
	SleepTime int
}

func InputEvent(url string, names map[string]int, data map[string]string, req *goPostMan.WebBox) {
	return
}

func ResultEvent(name string, res *HttpRequest.Response, req *goPostMan.WebBox) {
	gr := handleData.GameResult[name]
	if gr == nil {
		return
	}
	gr.Lock.Lock()
	defer gr.Lock.Unlock()

	if name == "加拿大28" || name == "北京28" {
		handleData.GetResult1(name, gr, res, 9000, 7000, "加拿大28")
	}
	if name == "古早" {
		handleData.GetResult1(name, gr, res, 9000, 7000, "北京28")
	}

	if name == "蛋蛋北京28" {
		b, err := res.Body()
		if handleData.PrintLog(gr, err) == false {
			return
		}
		type boxinfo struct {
			Issue string
			Time  string
			C1    string
			C2    string
			C3    string
			C4    int
			Kj    string
		}
		type box struct {
			Ret  int
			Data []boxinfo
			Msg  string
		}
		mdata := &box{}
		if err := json.Unmarshal(b, mdata); err == nil {
			if len(mdata.Data) > 0 {
				qh := ""
				m := make(map[string]boxinfo, len(mdata.Data))
				for _, v := range mdata.Data {
					if v.Issue != "" && v.Issue > qh {
						qh = v.Issue
					}
					m[v.Issue] = v
				}
				if len(gr.Data) > 0 {
					if gr.Data[0].Expect < qh {
						gr.Data[0].Opencode = m[qh].C1 + "," + m[qh].C2 + "," + m[qh].C3
						gr.Data[0].Expect = qh
						gr.Row = 1
						gr.Code = "ddbj28"
						logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
					}
				} else {
					gr.Data = make([]handleData.InfoBox, 1)
					gr.Data[0].Opencode = m[qh].C1 + "," + m[qh].C2 + "," + m[qh].C3
					gr.Data[0].Expect = qh
					gr.Row = 1
					gr.Code = "ddbj28"
					logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
				}
				handleData.AddNewResult(gr, "北京28")
				gr.SleepTime = 10000
			} else {
				handleData.PrintLog(gr, errors.New("蛋蛋北京28给的data是空的..."))
				gr.SleepTime = 10000
				return
			}
		} else {
			handleData.PrintLog(gr, err)
			gr.SleepTime = 10000
			return
		}
	}

	if name == "大白加拿大28" {
		b, err := res.Body()
		if handleData.PrintLog(gr, err) == false {
			return
		}
		type boxinfo struct {
			Issue string
			Time  string
			C1    string
			C2    string
			C3    string
			C4    int
			Kj    string
		}
		type box struct {
			Ret  int
			Data []boxinfo
			Msg  string
		}
		mdata := &box{}
		if err := json.Unmarshal(b, mdata); err == nil {
			if len(mdata.Data) > 0 {
				qh := ""
				m := make(map[string]boxinfo, len(mdata.Data))
				for _, v := range mdata.Data {
					if v.Issue != "" && v.Issue > qh {
						qh = v.Issue
					}
					m[v.Issue] = v
				}
				if len(gr.Data) > 0 {
					if gr.Data[0].Expect < qh {
						gr.Data[0].Opencode = m[qh].C1 + "," + m[qh].C2 + "," + m[qh].C3
						gr.Data[0].Expect = qh
						gr.Row = 1
						gr.Code = "dbjnd28"
						logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
					}
				} else {
					gr.Data = make([]handleData.InfoBox, 1)
					gr.Data[0].Opencode = m[qh].C1 + "," + m[qh].C2 + "," + m[qh].C3
					gr.Data[0].Expect = qh
					gr.Row = 1
					gr.Code = "dbjnd28"
					logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
				}
				handleData.AddNewResult(gr, "加拿大28")
				gr.SleepTime = 10000
			} else {
				handleData.PrintLog(gr, errors.New("大白加拿大28给的data是空的..."))
				gr.SleepTime = 10000
				return
			}
		} else {
			handleData.PrintLog(gr, err)
			gr.SleepTime = 10000
			return
		}
	}
}

func getMyDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

type Person struct {
	Name string
	Age  uint64
	KKK  map[string]interface{}
	in   string
}

type Person2 struct {
	Name string
	KKK  map[string]interface{}
	In   string
}

func main() {
	myPath := ""
	handleData.GameResult = make(map[string]*handleData.ResultInfoBox)
	handleData.NewResult.Store("加拿大28", new(handleData.ResultInfoBox))
	handleData.NewResult.Store("北京28", new(handleData.ResultInfoBox))
	handleData.NewResult.Store("英伦快5", new(handleData.ResultInfoBox))

	var lock sync.RWMutex
	goCount := 0
	IncGo := func() int {
		lock.Lock()
		defer lock.Unlock()
		goCount++
		return goCount
	}
	//time.Sleep(time.Second * 2)
	//go func() {
	//	name := "www.kjwapi.com加拿大28-1"
	//	gr := new(handleData.ResultInfoBox)
	//	gr.SleepTime = 100
	//	handleData.GameResult[name] = gr
	//	go func() {
	//		time.Sleep(time.Second * 60)
	//		if _, ok := gr.LastMsg.Load("errorInfo"); ok {
	//			gr.LastMsg.Delete("errorInfo")
	//		}
	//	}()
	//	w := goPostMan.CreatePostMan(IncGo(), myPath+"./urlJson/www.duourl.com.postman_collection.json",
	//		myPath+"./urlJson/www.duourl.com.postman_test_run.json", InputEvent, handleData.HandleJND28)
	//	for {
	//		time.Sleep(time.Millisecond * time.Duration(handleData.GameResult[name].SleepTime))
	//		w.RunName(name, false)
	//	}
	//}()

	time.Sleep(time.Second * 1)
	go func() {
		name := "www.caipiaoapi.com北京28-01"
		gr := new(handleData.ResultInfoBox)
		gr.SleepTime = 100
		handleData.GameResult[name] = gr
		go func() {
			time.Sleep(time.Second * 60)
			if _, ok := gr.LastMsg.Load("errorInfo"); ok {
				gr.LastMsg.Delete("errorInfo")
			}
		}()
		w := goPostMan.CreatePostMan(IncGo(), myPath+"./urlJson/www.duourl.com.postman_collection.json",
			myPath+"./urlJson/www.duourl.com.postman_test_run.json", InputEvent, handleData.HandleBJ28)
		for {
			time.Sleep(time.Millisecond * time.Duration(handleData.GameResult[name].SleepTime))
			w.RunName(name, false)
		}
	}()

	time.Sleep(time.Second * 1)
	go func() {
		name := "www.caipiaoapi.com加拿大28-01"
		gr := new(handleData.ResultInfoBox)
		gr.SleepTime = 100
		handleData.GameResult[name] = gr
		go func() {
			time.Sleep(time.Second * 60)
			if _, ok := gr.LastMsg.Load("errorInfo"); ok {
				gr.LastMsg.Delete("errorInfo")
			}
		}()
		w := goPostMan.CreatePostMan(IncGo(), myPath+"./urlJson/www.duourl.com.postman_collection.json",
			myPath+"./urlJson/www.duourl.com.postman_test_run.json", InputEvent, handleData.Handle2JND28)
		for {
			time.Sleep(time.Millisecond * time.Duration(handleData.GameResult[name].SleepTime))
			w.RunName(name, false)
		}
	}()


	//time.Sleep(time.Second * 1)
	//go func() {
	//	name := "user.lzwt.net蛋蛋28"
	//	gr := new(handleData.ResultInfoBox)
	//	gr.SleepTime = 100
	//	handleData.GameResult[name] = gr
	//	go func() {
	//		time.Sleep(time.Second * 60)
	//		if _, ok := gr.LastMsg.Load("errorInfo"); ok {
	//			gr.LastMsg.Delete("errorInfo")
	//		}
	//	}()
	//	w := goPostMan.CreatePostMan(IncGo(), myPath+"./urlJson/www.duourl.com.postman_collection.json",
	//		myPath+"./urlJson/www.duourl.com.postman_test_run.json", InputEvent, handleData.HandleBJ28_2)
	//	for {
	//		time.Sleep(time.Millisecond * time.Duration(handleData.GameResult[name].SleepTime))
	//		w.RunName(name, false)
	//	}
	//}()
	//time.Sleep(time.Second * 1)
	//go func() {
	//	name := "user.lzwt.net加拿大28"
	//	gr := new(handleData.ResultInfoBox)
	//	gr.SleepTime = 100
	//	handleData.GameResult[name] = gr
	//	go func() {
	//		time.Sleep(time.Second * 60)
	//		if _, ok := gr.LastMsg.Load("errorInfo"); ok {
	//			gr.LastMsg.Delete("errorInfo")
	//		}
	//	}()
	//	w := goPostMan.CreatePostMan(IncGo(), myPath+"./urlJson/www.duourl.com.postman_collection.json",
	//		myPath+"./urlJson/www.duourl.com.postman_test_run.json", InputEvent, handleData.HandleJND28_2)
	//	for {
	//		time.Sleep(time.Millisecond * time.Duration(handleData.GameResult[name].SleepTime))
	//		w.RunName(name, false)
	//	}
	//}()
	//time.Sleep(time.Second * 1)
	//go func() {
	//	name := "https://www.caijuapi.com北京28"
	//	gr := new(handleData.ResultInfoBox)
	//	gr.SleepTime = 100
	//	handleData.GameResult[name] = gr
	//	go func() {
	//		time.Sleep(time.Second * 60)
	//		if _, ok := gr.LastMsg.Load("errorInfo"); ok {
	//			gr.LastMsg.Delete("errorInfo")
	//		}
	//	}()
	//	w := goPostMan.CreatePostMan(IncGo(), myPath+"./urlJson/www.duourl.com.postman_collection.json",
	//		myPath+"./urlJson/www.duourl.com.postman_test_run.json", InputEvent, handleData.HandleBJ28_3)
	//	for {
	//		time.Sleep(time.Millisecond * time.Duration(handleData.GameResult[name].SleepTime))
	//		w.RunName(name, false)
	//	}
	//}()
	//time.Sleep(time.Second * 1)
	//go func() {
	//	name := "https://www.caijuapi.com加拿大28"
	//	gr := new(handleData.ResultInfoBox)
	//	gr.SleepTime = 100
	//	handleData.GameResult[name] = gr
	//	go func() {
	//		time.Sleep(time.Second * 60)
	//		if _, ok := gr.LastMsg.Load("errorInfo"); ok {
	//			gr.LastMsg.Delete("errorInfo")
	//		}
	//	}()
	//	w := goPostMan.CreatePostMan(IncGo(), myPath+"./urlJson/www.duourl.com.postman_collection.json",
	//		myPath+"./urlJson/www.duourl.com.postman_test_run.json", InputEvent, handleData.HandleJND28_3)
	//	for {
	//		time.Sleep(time.Millisecond * time.Duration(handleData.GameResult[name].SleepTime))
	//		w.RunName(name, false)
	//	}
	//}()

	time.Sleep(time.Second * 1)
	go func() {
		name := "英伦快5www.caipiaozixun888.com"
		gr := new(handleData.ResultInfoBox)
		gr.SleepTime = 100
		handleData.GameResult[name] = gr
		go func() {
			time.Sleep(time.Second * 60)
			if _, ok := gr.LastMsg.Load("errorInfo"); ok {
				gr.LastMsg.Delete("errorInfo")
			}
		}()
		w := goPostMan.CreatePostMan(IncGo(), myPath+"./urlJson/www.duourl.com.postman_collection.json",
			myPath+"./urlJson/www.duourl.com.postman_test_run.json", InputEvent, handleData.YLK5)
		for {
			time.Sleep(time.Millisecond * time.Duration(handleData.GameResult[name].SleepTime))
			w.RunName(name, false)
		}
	}()

	router := gin.New()
	router.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	router.GET("/GetResult/JND28", func(c *gin.Context) {
		if v, ok := handleData.NewResult.Load("加拿大28"); ok {
			c.JSON(200, v.(*handleData.ResultInfoBox))
		}
	})
	router.GET("/GetResult/BJ28", func(c *gin.Context) {
		if v, ok := handleData.NewResult.Load("北京28"); ok {
			c.JSON(200, v.(*handleData.ResultInfoBox))
		}
	})
	router.GET("/GetResult/YLK5", func(c *gin.Context) {
		if v, ok := handleData.NewResult.Load("英伦快5"); ok {
			c.JSON(200, v.(*handleData.ResultInfoBox))
		}
	})
	err := router.Run(":7710")
	//err := router.Run("127.0.0.1:7710")
	if err != nil {
		logger.Error(err.Error())
	}
	//for {
	//	time.Sleep(time.Second)
	//}
}
