/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-9
* Time: 上午10:58
* */
package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	SessionMap sync.Map
)

type SessionNode struct {
	Name string
	Time int64
	OutTime int64
}

func main() {
	// 用户登陆就设置cookie
	http.HandleFunc("/login",login)
	// 用户访问主业验证cookie
	http.HandleFunc("/home",home)

	err := http.ListenAndServe(":9085", nil)
	if err!= nil {
		panic(err.Error())
	}
}

func login(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	name := r.PostForm.Get("name")
	password := r.PostForm.Get("password")
	// 用户验证
	if name == "dollarkiller" && password == "password" {
		//用户登陆成功设置cookie
		id := generateSessionId()
		cookie := &http.Cookie{Name: "session",Value:id}
		http.SetCookie(w,cookie)
		// 存储用户session到map中
		outtime := time.Now().Unix() + 6*60*60
		i := &SessionNode{
			Name:name,
			Time:time.Now().Unix(),
			OutTime:outtime,
		}
		SessionMap.Store(id,i)

		fmt.Println(i)
		fmt.Println(id)

		w.WriteHeader(200)
		w.Header().Set("Content-Type","application/json")
		resp := map[string]interface{} {
			"msg":"login OK",
		}
		bytes, _ := json.Marshal(resp)
		w.Write(bytes)
		return
	}
	w.WriteHeader(400)
	w.Header().Set("Content-Type","application/json")
	resp := map[string]interface{} {
		"msg":"error",
	}
	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}


// 用户携带sessionId前来
func home(w http.ResponseWriter,r *http.Request) {
	// 检测sessionId
	cookie, e := r.Cookie("session")
	if e != nil {
		fmt.Println("错误:",e.Error())
		http.Redirect(w,r,"/login",301)
		return
	}
	session := cookie.Value
	// 根据session去map中查询是否存在
	value, ok := SessionMap.Load(session)
	if ok!=true {
		fmt.Println("用户不存在")
		http.Redirect(w,r,"/login",301)
		return
	}
	nowTime := time.Now().Unix()
	i := value.(*SessionNode)
	ctime := i.Time
	outTime :=	i.OutTime
	fmt.Println("ss: ",nowTime)
	fmt.Println("ssb: ",outTime)
	if nowTime > ctime && nowTime < outTime {
		i2 := map[string]interface{} {
			"msg":"ok",
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type","application/json")
		bytes, _ := json.Marshal(i2)
		w.Write(bytes)
		return
	}
	http.Redirect(w,r,"/login",301)
}


// 生成session id
func generateSessionId() string {
	intn := rand.Intn(100000) // 随机数
	unix := time.Now().UnixNano() // 时间
	encode := Md5Encode(strconv.FormatInt(unix, 10) + strconv.Itoa(intn))
	return encode
}

func Md5Encode(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}