package utils

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func ToString(in any) string {
	return fmt.Sprintf("%v", in)
}

func FileIsExisted(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func MakeDir(dir string) {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			fmt.Printf("创建文件夹%s失败%s\n", dir, err.Error())
		}
	}
}

func ListDir(dir string) (files []string) {
	d, _ := os.ReadDir(dir)
	for _, item := range d {
		if !item.IsDir() {
			files = append(files, dir+"/"+item.Name())
		}
	}
	return
}

func ReadTxt(path string) string {
	file, _ := os.Open(path)
	defer file.Close()
	reader := bufio.NewReader(file)
	s := ""
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		s += str
	}
	return s
}

func CopyFile(src, dest string) {
	f1, _ := os.Open(src)
	f2, _ := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer f2.Close()
	defer f1.Close()
	_, _ = io.Copy(f2, f1)
}

func RandInt(count, size int) []int {
	var nums []int
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		if size != -1 {
			nums = append(nums, rand.Intn(size))
		} else {
			nums = append(nums, rand.Int())
		}
	}
	return nums
}

func MicsSlice[T any](origin []T, count int) []T {
	var indexs []int
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := len(origin)
	for len(indexs) < count {
		i := rand1.Intn(size - 1)
		if Isin(indexs, i) {
			continue
		} else {
			indexs = append(indexs, i)
		}
	}
	result := make([]T, 0, count)
	for _, i := range indexs {
		result = append(result, origin[i])
	}
	return result
}

func Isin[T string | int | int8 | int16 | int32 | int64 | float32 | float64](list []T, item T) bool {
	for _, t := range list {
		if t == item {
			return true
		}
	}
	return false
}

func POST(url string, parameter any, result any) error {
	var (
		body io.Reader
		err  error
	)
	dataBytes, err := json.Marshal(parameter)
	if err != nil {
		return err
	}
	body = bytes.NewReader(dataBytes)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := &http.Client{}
	client.Timeout = 2 * time.Minute
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if result == nil {
		return errors.New("result is nil")
	}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return err
	}
	return err
}

func Get(url string, parameter any, result any) error {
	data, err := json.Marshal(&parameter)
	if err != nil {
		return err
	}
	p := make(map[string]any)
	if err = json.Unmarshal(data, &p); err != nil {
		return err
	}
	if len(p) > 0 {
		url += "?"
		for k, v := range p {
			url = fmt.Sprintf("%s%s=%v&", url, k, v)
		}
		url = url[:len(url)-1]
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	client.Timeout = 5 * time.Minute
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if result == nil {
		return errors.New("result is nil")
	}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(responseBytes, result)
	return err
}

func GetChp() string {
	chp := struct {
		Data struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"data"`
	}{}
	err := Get("https://api.shadiao.pro/chp", nil, &chp)
	if err != nil {
		return "获取失败" + err.Error()
	}
	return strings.TrimSpace(strings.Split(strings.ToLower(chp.Data.Text), "by")[0])
}

func GetDu() string {
	chp := struct {
		Data struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"data"`
	}{}
	err := Get("https://api.shadiao.pro/du", nil, &chp)
	if err != nil {
		return "获取失败" + err.Error()
	}
	return strings.TrimSpace(strings.Split(strings.ToLower(chp.Data.Text), "by")[0])
}

func GetPyq() string {
	chp := struct {
		Data struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"data"`
	}{}
	err := Get("https://api.shadiao.pro/pyq", nil, &chp)
	if err != nil {
		return "获取失败" + err.Error()
	}
	return strings.TrimSpace(strings.Split(strings.ToLower(chp.Data.Text), "by")[0])
}

func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
