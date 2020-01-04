package tool

import (
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"reflect"
	"strconv"
	"strings"
)

var (
	MAXTX_COUNT int
)

func scale(in io.Reader, out io.Writer, width, height, quality int) error {
	origin, fm, err := image.Decode(in)
	if err != nil {
		return err
	}
	if width == 0 || height == 0 {
		width = origin.Bounds().Max.X
		height = origin.Bounds().Max.Y
	}
	if quality == 0 {
		quality = 100
	}
	canvas := resize.Thumbnail(uint(width), uint(height), origin, resize.Lanczos3)

	switch fm {
	case "jpg":
		return jpeg.Encode(out, canvas, &jpeg.Options{quality})
	case "jpeg":
		return jpeg.Encode(out, canvas, &jpeg.Options{quality})
	case "png":
		return png.Encode(out, canvas)
	case "gif":
		return gif.Encode(out, canvas, &gif.Options{})
	default:
		return errors.New("ERROR FORMAT")
	}
}

func ArrMapFloatToStr(m []interface{}) []interface{} {
	var arr []interface{}
	for _, v0 := range m {
		r := make(map[string]interface{})
		for k, v := range v0.(map[string]interface{}) {
			lx := reflect.TypeOf(v).String()
			if lx == "map[string]interface {}" {
				r[k] = MapFloatToStr(v.(map[string]interface{}))
				continue
			}
			if lx == "float32" {
				r[k] = strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32)
				continue
			}
			if lx == "float64" {
				r[k] = strconv.FormatFloat(v.(float64), 'f', -1, 64)
				continue
			}
			r[k] = v
		}
		arr = append(arr, r)
	}
	return arr
}

func MapFloatToStr(m map[string]interface{}) map[string]interface{} {
	r := make(map[string]interface{})
	for k, v := range m {
		lx := reflect.TypeOf(v).String()
		if lx == "map[string]interface {}" {
			r[k] = MapFloatToStr(v.(map[string]interface{}))
			continue
		}
		if lx == "float32" {
			r[k] = strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32)
			continue
		}
		if lx == "float64" {
			r[k] = strconv.FormatFloat(v.(float64), 'f', -1, 64)
			continue
		}
		r[k] = v
	}
	return r
}

func GetMapDataType(m map[string]interface{}, name string, typeStr string) (interface{}, bool, error) {
	if v, ok := m[name]; !ok {
		return nil, false, errors.New("错误的字段" + name)
	} else {
		if typeStr == "" {
			return v, true, nil
		} else {
			if v == nil {
				return nil, false, errors.New("字段nil")
			}
			if strings.ToUpper(reflect.TypeOf(v).String()) != strings.ToUpper(typeStr) {
				return nil, false, errors.New("错误的类型" + name + " " + reflect.TypeOf(v).String())
			} else {
				return v, true, nil
			}
		}
	}
}

func InterfaceToInt(data interface{}) int {
	lx := reflect.TypeOf(data).String()
	if lx == "uint8" {
		return int(data.(uint8))
	}
	if lx == "uint16" {
		return int(data.(uint16))
	}
	if lx == "uint32" {
		return int(data.(uint32))
	}
	if lx == "int8" {
		return int(data.(int8))
	}
	if lx == "int32" {
		return int(data.(int32))
	}
	if lx == "int64" {
		return int(data.(int64))
	}
	if lx == "int" {
		return data.(int)
	}
	if lx == "float64" {
		return int(data.(float64))
	}
	if lx == "float32" {
		return int(data.(float32))
	}
	panic("不是数字类型！是：" + lx)
}

func InterfaceToFloat(data interface{}) float64 {
	lx := reflect.TypeOf(data).String()
	if lx == "uint8" {
		return float64(data.(uint8))
	}
	if lx == "uint16" {
		return float64(data.(uint16))
	}
	if lx == "uint32" {
		return float64(data.(uint32))
	}
	if lx == "int8" {
		return float64(data.(int8))
	}
	if lx == "int32" {
		return float64(data.(int32))
	}
	if lx == "int64" {
		return float64(data.(int64))
	}
	if lx == "int" {
		return float64(data.(int))
	}
	if lx == "float64" {
		return data.(float64)
	}
	if lx == "float32" {
		return float64(data.(float32))
	}
	panic("不是数字类型！是：" + lx)
}

func GetMapDataCall(m map[string]interface{}, key string) (errMsg string, value interface{}) {
	if v, ok := m[key]; !ok {
		return "字段不存在", nil
	} else {
		return "", v
	}
}
