package route

import (
	"encoding/json"
	"fmt"
	"html/template"
	"message/global"
	"message/model"
	"message/util"
	"net/http"
	"strconv"
	"time"
)

type H map[string]interface{}

const (
	CodeSuccess = 1000
	CodeError   = 1001
)

// MessageList 显示首页
func MessageList(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() == "/favicon.ico" {
		return
	}

	var messages []map[string]interface{}
	global.Db.Model(&model.Message{}).Order("id DESC").Limit(5).Find(&messages)
	for _, v := range messages {
		thisTime := int64(v["created_at"].(uint32))
		v["created_at"] = util.Date(util.Layout, thisTime)
	}

	templateFile := "./template/index.tmpl"
	tmpl := template.Must(template.ParseFiles(templateFile))
	tmpl.Execute(w, map[string]interface{}{
		"list": messages,
	})
}

// inputJson 输出json信息
func inputJson(w http.ResponseWriter, obj H) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(obj); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// MessageCreate 创建留言 返回json
func MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "只接受 POST 请求,当前请求 %s\n", r.Method)
		return
	}
	name := r.FormValue("name")
	if name == "" {
		inputJson(w, H{
			"code": CodeError,
			"msg":  "主题不能为空",
		})
		return
	}
	content := r.FormValue("content")
	if content == "" {
		inputJson(w, H{
			"code": CodeError,
			"msg":  "内容不能为空",
		})
		return
	}
	// 创建记录
	message := model.Message{Name: name, Content: &content, CreatedAt: uint32(time.Now().Unix())}
	result := global.Db.Create(&message)
	if result.Error != nil {
		inputJson(w, H{
			"code": CodeError,
			"msg":  fmt.Sprintf("插入数据异常,err:%s", result.Error),
		})
		return
	}
	inputJson(w, H{
		"code": CodeSuccess,
		"msg":  "ok",
		"data": H{
			"id": message.Id,
		},
	})
}

// MessageDel 删除对应留言
func MessageDel(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "只接受 POST 请求,当前请求 %s\n", r.Method)
		return
	}
	idS := r.FormValue("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		inputJson(w, H{
			"code": CodeError,
			"msg":  "id参数异常",
		})
		return
	}
	global.Db.Delete(&model.Message{}, id)
	inputJson(w, H{
		"code": CodeSuccess,
		"msg":  "ok",
		"data": H{
			"id": id,
		},
	})

}
