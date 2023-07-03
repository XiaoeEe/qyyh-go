package robotService

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"qyyh-go/database/table"
	"qyyh-go/model"
	"qyyh-go/utils"
	"regexp"
	"strings"
)

func getEntryList(message model.CQMessage) {
	var strs []string
	for _, entry := range table.GetEntryList() {
		strs = append(strs, entry.Name)
	}
	if len(strs) == 0 {
		SendGroupMsg(message, "词条库为空")
	} else {
		SendGroupMsg(message, "词条列表\n"+strings.Join(strs, ", "))
	}
}

func getEntry(message model.CQMessage, m string) {
	entry := table.GetEntry(m)
	if entry.Name != "" {
		SendGroupMsg(message, entry.Value)
	} else {
		entrys := table.GetEntryListByLikeName(m)
		if len(entrys) == 0 {
			SendGroupMsg(message, "未查询到相关词条")
		} else {
			var strs []string
			for _, e := range entrys {
				strs = append(strs, e.Name)
			}
			SendGroupMsg(message, fmt.Sprintf("查询到以下类似词条:\n%s", strings.Join(strs, "\n")))
		}
	}
}

func setEntry(message model.CQMessage, m []string) {
	name := m[1]
	value := ""
	for i := 2; i < len(m); i++ {
		value += m[i] + " "
	}
	value = fomatMsg(value)
	if err := table.SaveEntry(table.Entry{
		Name:  name,
		Value: value,
	}); err != nil {
		SendGroupMsg(message, err.Error())
	} else {
		SendGroupMsg(message, "设置词条成功")
	}
}

func delEntry(message model.CQMessage, m string) {
	if err := table.DeleteEntry(m); err != nil {
		SendGroupMsg(message, "删除失败："+err.Error())
	} else {
		SendGroupMsg(message, "删除完成")
	}
}

func fomatMsg(msg string) string {
	s := subCQCode(msg)
	if len(s) == 0 {
		return msg
	}
	var list []string
	for _, item := range s {
		if item["type"] != "CQ:image" {
			continue
		}
		filepath, err := saveImg(item["url"])
		if err != nil {
			fmt.Println("缓存图片错误: " + err.Error())
			return ""
		}
		list = append(list, fmt.Sprintf("[CQ:image,file=https://file.qyyh.net/robot/entry/%s,cache=0]", filepath))
	}
	return replaceCQCode(msg, list)
}

func subCQCode(cq string) []map[string]string {
	reg := regexp.MustCompile(`\[([^]\[\r\n]*)]`)
	var list []map[string]string

	for _, s := range reg.FindAllString(cq, -1) {
		ss := strings.Split(strings.ReplaceAll(strings.ReplaceAll(s, "[", ""), "]", ""), ",")
		tmpMap := map[string]string{}
		for _, item := range ss {
			c := strings.Split(item, "=")
			if len(c) == 1 {
				tmpMap["type"] = c[0]
			} else {
				tmpMap[c[0]] = c[1]
			}
		}
		list = append(list, tmpMap)
	}
	return list
}

func replaceCQCode(cq string, list []string) string {
	reg := regexp.MustCompile(`\[([^]\[\r\n]*)]`)
	s := reg.ReplaceAllString(cq, "[待替换文本]")
	for _, l := range list {
		s = strings.Replace(s, "[待替换文本]", l, 1)
	}
	return strings.TrimSpace(s)
}

func saveImg(url string) (filename string, err error) {
	utils.MakeDir("./file/entry")
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	fmt.Println()
	render, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	ss := strings.Split(res.Header.Get("Content-Type"), "/")
	if ss[0] == "image" {
		filename = fmt.Sprintf("%s.%s", uuid.NewString(), ss[1])
		err = os.WriteFile(fmt.Sprintf("./file/entry/%s", filename), render, 0777)
		if err != nil {
			return
		}
	}
	return
}
