package handleData

import (
	"LotteryResultAPI/tool"
	"bn/goPostMan"
	"errors"
	json "github.com/json-iterator/go"
	"github.com/kirinlabs/HttpRequest"
	"github.com/zllangct/zgo/logger"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func setTimer(gr *ResultInfoBox, n int) {
	if !gr.SetTimerIng {
		gr.SleepTime = n
		gr.SetTimerIng = true
	}
}

func HandleJND28(name string, res *HttpRequest.Response, req *goPostMan.WebBox) {
	gr := GameResult[name]
	gr.SetTimerIng = false
	if name == "www.kjwapi.com加拿大28-1" || name == "www.kjwapi.com加拿大28-2" {
		b, err := res.Body()
		if PrintLog(gr, err) == false {
			return
		}
		type boxinfo struct {
			Gid   string
			Time  string
			Award string
		}
		type box struct {
			Data    []boxinfo
			Message string
		}
		type box0 struct {
			ErrorCode int
			Message   string
			Result    box
		}
		mdata := &box0{}
		if err := json.Unmarshal(b, mdata); err == nil {
			if mdata.ErrorCode != 0 {
				PrintLog(gr, errors.New(mdata.Message))
				setTimer(gr, 8000)
				return
			}
			if len(mdata.Result.Data) > 0 {
				qh := ""
				m := make(map[string]boxinfo, len(mdata.Result.Data))
				for _, v := range mdata.Result.Data {
					if v.Gid != "" && v.Gid > qh {
						qh = v.Gid
					}
					m[v.Gid] = v
				}
				if len(gr.Data) > 0 {
					if gr.Data[0].Expect < qh {
						c := strings.Split(m[qh].Award, ",")
						gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
						gr.Data[0].Expect = qh
						gr.Row = 1
						gr.Code = "jnd28"
						logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
						setTimer(gr, 180000)
					}
				} else {
					gr.Data = make([]InfoBox, 1)
					c := strings.Split(m[qh].Award, ",")
					gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
					gr.Data[0].Expect = qh
					gr.Row = 1
					gr.Code = "jnd28"
					logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
					setTimer(gr, 180000)
				}
				AddNewResult(gr, "加拿大28")
				setTimer(gr, 8000)
			} else {
				PrintLog(gr, errors.New(name+"data是空的..."))
				setTimer(gr, 8000)
				return
			}
		} else {
			PrintLog(gr, err)
			setTimer(gr, 8000)
			return
		}
	}

	if name == "api.ip5i.com加拿大28" {
		b, err := res.Body()
		if PrintLog(gr, err) == false {
			return
		}
		type boxinfo struct {
			First    int
			Second   int
			Third    int
			Opentime string
			Expect   string
		}
		type box struct {
			Open []boxinfo
			Code string
		}
		type box0 struct {
			Errcode int
			Data    box
		}
		mdata := &box0{}
		if err := json.Unmarshal(b, mdata); err == nil {
			if mdata.Errcode != 0 {
				setTimer(gr, 8000)
				return
			}
			if len(mdata.Data.Open) > 0 {
				qh := ""
				m := make(map[string]boxinfo, len(mdata.Data.Open))
				for _, v := range mdata.Data.Open {
					if v.Expect != "" && v.Expect > qh {
						qh = v.Expect
					}
					m[v.Expect] = v
				}
				if len(gr.Data) > 0 {
					if gr.Data[0].Expect < qh {
						c := []string{
							strconv.FormatInt(int64(m[qh].First), 10),
							strconv.FormatInt(int64(m[qh].Second), 10),
							strconv.FormatInt(int64(m[qh].Third), 10),
						}
						gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
						gr.Data[0].Expect = qh
						gr.Row = 1
						gr.Code = "jnd28"
						logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
						setTimer(gr, 180000)
					}
				} else {
					gr.Data = make([]InfoBox, 1)
					c := []string{
						strconv.FormatInt(int64(m[qh].First), 10),
						strconv.FormatInt(int64(m[qh].Second), 10),
						strconv.FormatInt(int64(m[qh].Third), 10),
					}
					gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
					gr.Data[0].Expect = qh
					gr.Row = 1
					gr.Code = "jnd28"
					logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
					setTimer(gr, 180000)
				}
				AddNewResult(gr, "加拿大28")
				setTimer(gr, 8000)
			} else {
				PrintLog(gr, errors.New(name+"data是空的..."))
				setTimer(gr, 8000)
				return
			}
		} else {
			PrintLog(gr, err)
			setTimer(gr, 8000)
			return
		}
	}
}

func HandleBJ28(name string, res *HttpRequest.Response, req *goPostMan.WebBox) {
	gr := GameResult[name]
	gr.SetTimerIng = false
	if name == "www.caipiaoapi.com北京28-01" || name == "www.caipiaoapi.com北京28-02" {
		b, err := res.Body()
		if PrintLog(gr, err) == false {
			return
		}
		type boxinfo struct {
			Gid   string
			Time  string
			Award string
		}
		type box struct {
			Data    []boxinfo
			Message string
		}
		type box0 struct {
			ErrorCode int
			Message   string
			Result    box
		}
		mdata := &box0{}
		if err := json.Unmarshal(b, mdata); err == nil {
			if mdata.ErrorCode != 0 {
				PrintLog(gr, errors.New(mdata.Message))
				setTimer(gr, 8000)
				return
			}
			if len(mdata.Result.Data) > 0 {
				qh := ""
				m := make(map[string]boxinfo, len(mdata.Result.Data))
				for _, v := range mdata.Result.Data {
					if v.Gid != "" && v.Gid > qh {
						qh = v.Gid
					}
					m[v.Gid] = v
				}
				if len(gr.Data) > 0 {
					if gr.Data[0].Expect < qh {
						c := strings.Split(m[qh].Award, ",")
						gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
						gr.Data[0].Expect = qh
						gr.Row = 1
						gr.Code = "bj28"
						logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
						setTimer(gr, 240000)
					}
				} else {
					gr.Data = make([]InfoBox, 1)
					c := strings.Split(m[qh].Award, ",")
					gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
					gr.Data[0].Expect = qh
					gr.Row = 1
					gr.Code = "bj28"
					logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
					setTimer(gr, 240000)
				}
				AddNewResult(gr, "北京28")
				setTimer(gr, 8000)
			} else {
				PrintLog(gr, errors.New(name+"data是空的..."))
				setTimer(gr, 8000)
				return
			}
		} else {
			PrintLog(gr, err)
			setTimer(gr, 8000)
			return
		}
	}
}

func HandleJND28_2(name string, res *HttpRequest.Response, req *goPostMan.WebBox) {
	gr := GameResult[name]
	gr.SetTimerIng = false
	if name == "user.lzwt.net加拿大28" {
		b, err := res.Body()
		if PrintLog(gr, err) == false {
			return
		}
		type box struct {
			Opentime string
			Opencode string
			Expect   string
		}
		type box0 struct {
			Code string
			Data []box
		}
		mdata := &box0{}
		if err := json.Unmarshal(b, mdata); err == nil {
			if mdata.Code != "jnd28" {
				setTimer(gr, 8000)
				return
			}
			if len(mdata.Data) > 0 {
				qh := mdata.Data[0].Expect
				if len(gr.Data) > 0 {
					if gr.Data[0].Expect < qh {
						c := strings.Split(mdata.Data[0].Opencode, ",")
						gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
						gr.Data[0].Expect = qh
						gr.Row = 1
						gr.Code = "jnd28"
						logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
						setTimer(gr, 180000)
					}
				} else {
					gr.Data = make([]InfoBox, 1)
					c := strings.Split(mdata.Data[0].Opencode, ",")
					gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
					gr.Data[0].Expect = qh
					gr.Row = 1
					gr.Code = "jnd28"
					logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
					setTimer(gr, 180000)
				}
				AddNewResult(gr, "加拿大28")
				setTimer(gr, 8000)
			} else {
				PrintLog(gr, errors.New(name+"data是空的..."))
				setTimer(gr, 8000)
				return
			}
		} else {
			PrintLog(gr, err)
			setTimer(gr, 8000)
			return
		}
	}
}

func HandleBJ28_2(name string, res *HttpRequest.Response, req *goPostMan.WebBox) {
	gr := GameResult[name]
	gr.SetTimerIng = false
	if name == "user.lzwt.net蛋蛋28" {
		b, err := res.Body()
		if PrintLog(gr, err) == false {
			return
		}
		type box struct {
			Opentime string
			Opencode string
			Expect   string
		}
		type box0 struct {
			Code string
			Data []box
		}
		mdata := &box0{}
		if err := json.Unmarshal(b, mdata); err == nil {
			if mdata.Code != "jnd28" {
				setTimer(gr, 8000)
				return
			}
			if len(mdata.Data) > 0 {
				qh := mdata.Data[0].Expect
				if len(gr.Data) > 0 {
					if gr.Data[0].Expect < qh {
						c := strings.Split(mdata.Data[0].Opencode, ",")
						gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
						gr.Data[0].Expect = qh
						gr.Row = 1
						gr.Code = "bj28"
						logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
						setTimer(gr, 260000)
					}
				} else {
					gr.Data = make([]InfoBox, 1)
					c := strings.Split(mdata.Data[0].Opencode, ",")
					gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
					gr.Data[0].Expect = qh
					gr.Row = 1
					gr.Code = "bj28"
					logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
					setTimer(gr, 260000)
				}
				AddNewResult(gr, "北京28")
				setTimer(gr, 8000)
			} else {
				PrintLog(gr, errors.New(name+"data是空的..."))
				setTimer(gr, 8000)
				return
			}
		} else {
			PrintLog(gr, err)
			setTimer(gr, 8000)
			return
		}
	}
}

func Handle2JND28(name string, res *HttpRequest.Response, req *goPostMan.WebBox) {
	gr := GameResult[name]
	gr.SetTimerIng = false
	if name == "http://main.caipiaoapi.com" || name == "www.caipiaoapi.com加拿大28-01" || name == "www.caipiaoapi.com加拿大28-02" {
		b, err := res.Body()
		if PrintLog(gr, err) == false {
			return
		}
		type boxinfo struct {
			Gid           string
			Time          string
			Award         string
			NextOpenIssue string
			NextOpenTime  string
		}
		type box struct {
			Data    []boxinfo
			Message string
		}
		type box0 struct {
			ErrorCode int
			Message   string
			Result    box
		}
		mdata := &box0{}
		if err := json.Unmarshal(b, mdata); err == nil {
			if mdata.ErrorCode != 0 {
				PrintLog(gr, errors.New(mdata.Message))
				setTimer(gr, 8000)
				return
			}
			if len(mdata.Result.Data) > 0 {
				qh := ""
				m := make(map[string]boxinfo, len(mdata.Result.Data))
				for _, v := range mdata.Result.Data {
					if v.Gid != "" && v.Gid > qh {
						qh = v.Gid
					}
					m[v.Gid] = v
				}
				if len(gr.Data) > 0 {
					if gr.Data[0].Expect < qh {
						c := strings.Split(m[qh].Award, ",")
						gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
						gr.Data[0].Expect = qh
						gr.Data[0].Opentime = m[qh].Time
						gr.Data[0].NextTime = m[qh].NextOpenTime
						gr.Data[0].NextIssue = m[qh].NextOpenIssue
						gr.Row = 1
						gr.Code = "jnd28"
						logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
						setTimer(gr, 170000)
					}
				} else {
					gr.Data = make([]InfoBox, 1)
					c := strings.Split(m[qh].Award, ",")
					gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
					gr.Data[0].Expect = qh
					gr.Data[0].Opentime = m[qh].Time
					gr.Data[0].NextTime = m[qh].NextOpenTime
					gr.Data[0].NextIssue = m[qh].NextOpenIssue
					gr.Row = 1
					gr.Code = "jnd28"
					logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
					setTimer(gr, 170000)
				}
				AddNewResult(gr, "加拿大28")
				setTimer(gr, 8000)
			} else {
				PrintLog(gr, errors.New(name+"data是空的..."))
				setTimer(gr, 8000)
				return
			}
		} else {
			PrintLog(gr, err)
			setTimer(gr, 8000)
			return
		}
	}
}

func HandleJND28_3(name string, res *HttpRequest.Response, req *goPostMan.WebBox) {
	gr := GameResult[name]
	gr.SetTimerIng = false
	if name == "https://www.caijuapi.com加拿大28" {
		b, err := res.Body()
		if PrintLog(gr, err) == false {
			return
		}
		type box struct {
			Time    string
			Kjcodes string
			Issue   string
		}
		type box0 struct {
			Code  string
			Datas []box
		}
		mdata := &box0{}
		if err := json.Unmarshal(b, mdata); err == nil {
			if mdata.Code != "jnd28" {
				setTimer(gr, 8000)
				return
			}
			if len(mdata.Datas) > 0 {
				qh := mdata.Datas[0].Issue
				if len(gr.Data) > 0 {
					if gr.Data[0].Expect < qh {
						c := strings.Split(mdata.Datas[0].Kjcodes, ",")
						gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
						gr.Data[0].Expect = qh
						gr.Row = 1
						gr.Code = "jnd28"
						logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
						setTimer(gr, 180000)
					}
				} else {
					gr.Data = make([]InfoBox, 1)
					c := strings.Split(mdata.Datas[0].Kjcodes, ",")
					gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
					gr.Data[0].Expect = qh
					gr.Row = 1
					gr.Code = "jnd28"
					logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
					setTimer(gr, 180000)
				}
				AddNewResult(gr, "加拿大28")
				setTimer(gr, 8000)
			} else {
				PrintLog(gr, errors.New(name+"data是空的..."))
				setTimer(gr, 8000)
				return
			}
		} else {
			PrintLog(gr, err)
			setTimer(gr, 8000)
			return
		}
	}
}

func HandleBJ28_3(name string, res *HttpRequest.Response, req *goPostMan.WebBox) {
	gr := GameResult[name]
	gr.SetTimerIng = false
	if name == "https://www.caijuapi.com北京28" {
		b, err := res.Body()
		if PrintLog(gr, err) == false {
			return
		}
		type box struct {
			Time    string
			Kjcodes string
			Issue   string
		}
		type box0 struct {
			Code  string
			Datas []box
		}
		mdata := &box0{}
		if err := json.Unmarshal(b, mdata); err == nil {
			if mdata.Code != "xy28" {
				setTimer(gr, 8000)
				return
			}
			if len(mdata.Datas) > 0 {
				qh := mdata.Datas[0].Issue
				if len(gr.Data) > 0 {
					if gr.Data[0].Expect < qh {
						c := strings.Split(mdata.Datas[0].Kjcodes, ",")
						gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
						gr.Data[0].Expect = qh
						gr.Row = 1
						gr.Code = "bj28"
						logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
						setTimer(gr, 240000)
					}
				} else {
					gr.Data = make([]InfoBox, 1)
					c := strings.Split(mdata.Datas[0].Kjcodes, ",")
					gr.Data[0].Opencode = c[0] + "," + c[1] + "," + c[2]
					gr.Data[0].Expect = qh
					gr.Row = 1
					gr.Code = "bj28"
					logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
					setTimer(gr, 240000)
				}
				AddNewResult(gr, "北京28")
				setTimer(gr, 8000)
			} else {
				PrintLog(gr, errors.New(name+"data是空的..."))
				setTimer(gr, 8000)
				return
			}
		} else {
			PrintLog(gr, err)
			setTimer(gr, 8000)
			return
		}
	}
}

func YLK5(name string, res *HttpRequest.Response, req *goPostMan.WebBox) {
	gr := GameResult[name]
	gr.SetTimerIng = false
	if name == "英伦快5www.caipiaozixun888.com" {
		b, err := res.Body()
		if PrintLog(gr, err) == false {
			return
		}
		mdata := make(map[string]interface{})
		res := make(map[string]interface{})
		res["Result"] = false
		if err := json.Unmarshal(b, &mdata); err == nil {
			if _, ok := mdata["sta"]; !ok {
				logger.Error("英伦快5 sta没有")
				return
			}
			if tool.InterfaceToInt(mdata["sta"]) != 1 {
				res["Msg"] = mdata["msg"]
			}
			if reflect.TypeOf(mdata["data"]).String() == "[]interface {}" {
				var data map[string]interface{}
				var next map[string]interface{}
				if arr, ok := mdata["data"].([]interface{}); ok {
					if v, ok := arr[0].(map[string]interface{}); ok {
						data = v
					} else {
						logger.Error("数据类型不符合：", arr)
						res["Msg"] = "数据类型不符合"
					}
				}
				if arr, ok := mdata["next"].([]interface{}); ok {
					if v, ok := arr[0].(map[string]interface{}); ok {
						next = v
					} else {
						logger.Error("数据类型不符合：", arr)
						res["Msg"] = "数据类型不符合"
					}
				}
				getData := func(key string, dat interface{}) bool {
					if msg, v := tool.GetMapDataCall(data, key); msg != "" {
						res["Msg"] = msg
						return false
					} else {
						*dat.(*string) = v.(string)
						return true
					}
				}
				getNext := func(key string, dat interface{}) bool {
					if msg, v := tool.GetMapDataCall(next, key); msg != "" {
						res["Msg"] = msg
						return false
					} else {
						*dat.(*string) = v.(string)
						return true
					}
				}
				qh := ""
				if !getData("qs", &qh) {
					return
				}
				if len(gr.Data) > 0 {
					if gr.Data[0].Expect < qh {
						if !getData("kjhm", &gr.Data[0].Opencode) {
							return
						}
						if !getData("kjtime", &gr.Data[0].Opentime) {
							return
						}
						if !getNext("xqs", &gr.Data[0].NextIssue) {
							return
						}
						if !getNext("xqkjjztime", &gr.Data[0].NextTime) {
							return
						}
						gr.Data[0].Expect = qh
						gr.Row = 1
						gr.Code = "ylk5"
						logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
						setTimer(gr, 240000)
					}
				} else {
					gr.Data = make([]InfoBox, 1)
					if !getData("kjhm", &gr.Data[0].Opencode) {
						return
					}
					if !getData("kjtime", &gr.Data[0].Opentime) {
						return
					}
					if !getNext("xqs", &gr.Data[0].NextIssue) {
						return
					}
					if !getNext("xqkjjztime", &gr.Data[0].NextTime) {
						return
					}
					gr.Data[0].Expect = qh
					gr.Row = 1
					gr.Code = "ylk5"
					logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
					setTimer(gr, 240000)
				}
				AddNewResult(gr, "英伦快5")
				setTimer(gr, 10000)
			} else {
				PrintLog(gr, errors.New(name+"data是空的..."))
				setTimer(gr, 10000)
				return
			}
		} else {
			PrintLog(gr, err)
			setTimer(gr, 10000)
			return
		}
	}
}
