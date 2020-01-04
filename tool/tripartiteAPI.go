package tool

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
)

type tmpArr struct {
	Key   string
	Value interface{}
}

type IntSlice []tmpArr

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i].Key < s[j].Key }

func md5V3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

//法币支付那边的sign获得方式
func GetSignUrl(m map[string]interface{}) string {
	var arr IntSlice
	for k, v := range m {
		arr = append(arr, tmpArr{k, v})
	}
	sort.Sort(arr)
	var str string
	for _, v := range arr {
		str += v.Key + "=" + v.Value.(string) + "&"
	}
	sign := md5V3(str)
	str += "sign=" + sign
	return str
}
