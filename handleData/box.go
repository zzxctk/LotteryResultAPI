package handleData

import (
	json "github.com/json-iterator/go"
	"github.com/kirinlabs/HttpRequest"
	"github.com/zllangct/zgo/logger"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	//control    [3]*controlBox
	NewResult  sync.Map
	GameResult map[string]*ResultInfoBox
	//jnd28Result   *ResultInfoBox
	//bj28Result    *ResultInfoBox
	//lastResultStr []string
	//LastMsg       sync.Map
)

type InfoBox struct {
	Opentime  string `json:"opentime"`
	Expect    string `json:"expect"`
	Opencode  string `json:"opencode"`
	NextTime  string
	NextIssue string
}

type ResultInfoBox struct {
	Row         int       `json:"row"`
	Code        string    `json:"code"`
	Data        []InfoBox `json:"data"`
	SleepTime   int
	LastDataStr string
	LastIssue   string
	SetTimerIng bool
	LastMsg     sync.Map
	Lock        sync.RWMutex
}

func PrintLog(gr *ResultInfoBox, err error) bool {
	if err != nil {
		if v, ok := gr.LastMsg.Load("errorInfo"); ok == false {
			gr.LastMsg.Store("errorInfo", err.Error())
			logger.Debug2(3, err.Error())
		} else {
			if err.Error() != v.(string) {
				gr.LastMsg.Store("errorInfo", err.Error())
				logger.Debug2(3, err.Error())
			}
		}
		return false
	}
	return true
}

func GetResult1(name string, gr *ResultInfoBox, res *HttpRequest.Response, delay1, delay2 int, gameName string) {
	b, err := res.Body()
	if PrintLog(gr, err) == false {
		return
	}
	tmp := &ResultInfoBox{}
	err = json.Unmarshal(b, tmp)
	if PrintLog(gr, err) == false {
		gr.SleepTime = delay2
		if strings.Contains(string(b), "3秒内") == false {
			logger.Error(time.Now().Format("2006-01-02 15:04:05"), string(b))
		}
		return
	}
	if len(tmp.Data) > 0 {
		if len(gr.Data) > 0 {
			if tmp.Data[0].Expect > gr.Data[0].Expect {
				gr.Data[0].Opentime = tmp.Data[0].Opentime
				gr.Data[0].Expect = tmp.Data[0].Expect
				gr.Data[0].Opencode = tmp.Data[0].Opencode
				gr.Row = tmp.Row
				gr.Code = tmp.Code
				logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
				gr.SleepTime = delay1
			}
		} else {
			gr.Data = append(gr.Data, tmp.Data[0])
			gr.Row = tmp.Row
			gr.Code = tmp.Code
			logger.Debug(time.Now().Format("2006-01-02 15:04:05"), "获得"+name+"结果：", *gr)
			gr.SleepTime = delay1
		}
		AddNewResult(gr, gameName)
	} else {
		logger.Error("获得的结果数据长度是0")
		gr.SleepTime = delay2
	}
}

func AddNewResult(gr *ResultInfoBox, gameName string) {
	if v, ok := NewResult.Load(gameName); ok {
		dat := v.(*ResultInfoBox)
		dat.Lock.Lock()
		defer dat.Lock.Unlock()
		if len(dat.Data) > 0 {
			if dat.Data[0].Expect < gr.Data[0].Expect {
				dat.Code = gr.Code
				dat.Row = gr.Row
				dat.Data[0].Expect = gr.Data[0].Expect
				dat.Data[0].Opencode = gr.Data[0].Opencode
				dat.Data[0].Opentime = gr.Data[0].Opentime
				dat.Data[0].NextIssue = gr.Data[0].NextIssue
				dat.Data[0].NextTime = gr.Data[0].NextTime
			}
		} else {
			dat.Code = gr.Code
			dat.Row = gr.Row
			dat.Data = append(dat.Data, gr.Data[0])
		}
	}
}

func getBit(s int) int {
	for {
		if s > 100 {
			s -= 100
		} else {
			break
		}
	}
	for {
		if s > 10 {
			s -= 10
		} else {
			break
		}
	}
	if s == 10 {
		s = 0
	}
	return s
}

func To4(arr []string) []string {
	var numStr []string
	if len(arr) < 20 {
		logger.Error("要转换的结果长度不对！", arr)
	}
	var hmint []int
	for _, v := range arr {
		i, _ := strconv.Atoi(v)
		hmint = append(hmint, i)
	}
	s1 := getBit(hmint[1] + hmint[4] + hmint[7] + hmint[10] + hmint[13] + hmint[16])
	s2 := getBit(hmint[2] + hmint[5] + hmint[8] + hmint[11] + hmint[14] + hmint[17])
	s3 := getBit(hmint[3] + hmint[6] + hmint[9] + hmint[12] + hmint[15] + hmint[18])
	s4 := s1 + s2 + s3
	numStr = append(numStr, strconv.FormatInt(int64(s1), 10))
	numStr = append(numStr, strconv.FormatInt(int64(s2), 10))
	numStr = append(numStr, strconv.FormatInt(int64(s3), 10))
	numStr = append(numStr, strconv.FormatInt(int64(s4), 10))
	return numStr
}

func To4V1(arr []string) []string {
	var numStr []string
	if len(arr) < 20 {
		logger.Error("要转换的结果长度不对！", arr)
	}
	var hmint []int
	for _, v := range arr {
		i, _ := strconv.Atoi(v)
		hmint = append(hmint, i)
	}
	s1 := getBit(hmint[2] + hmint[5] + hmint[8] + hmint[11] + hmint[14] + hmint[17])
	s2 := getBit(hmint[3] + hmint[6] + hmint[9] + hmint[12] + hmint[15] + hmint[19])
	s3 := getBit(hmint[4] + hmint[7] + hmint[10] + hmint[13] + hmint[16] + hmint[20])
	s4 := s1 + s2 + s3
	numStr = append(numStr, strconv.FormatInt(int64(s1), 10))
	numStr = append(numStr, strconv.FormatInt(int64(s2), 10))
	numStr = append(numStr, strconv.FormatInt(int64(s3), 10))
	numStr = append(numStr, strconv.FormatInt(int64(s4), 10))
	return numStr
}
