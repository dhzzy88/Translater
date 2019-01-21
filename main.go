package main

import (
	"bufio"
	bytes2 "bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"baliance.com/gooxml/document"
)

var Filename = "/home/leo/Downloads/wenxian.docx"
var From = "en"
var To = "zh"

//百度api返回信息json结构
//申请的信息
var appID = 2015063000000001
var password = "12345678"
var channel = make(chan bool)

//百度翻译api接口
var Url = "http://api.fanyi.baidu.com/api/trans/vip/translate"

type configfile struct {
	From  string `json:"from"`
	To    string `json:"to"`
	AppID string `json:"appid"`
	Key   string `json:"key"`
}

//配置文件的json结构
type inline struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type back struct {
	From        string   `json:"from"`
	To          string   `json:"to"`
	Tran_result []inline `json:"trans_result"`
}

type TranslateModel struct {
	Q     string
	From  string
	To    string
	Appid int
	Salt  int
	Sign  string
}

func NewTranslateModeler(q, from, to string) TranslateModel {
	tran := TranslateModel{
		Q:    q,
		From: from,
		To:   to,
	}
	tran.Appid = appID
	tran.Salt = time.Now().Second()
	content := strconv.Itoa(appID) + q + strconv.Itoa(tran.Salt) + password
	sign := SumString(content) //计算sign值
	tran.Sign = sign
	return tran
}

func (tran TranslateModel) ToValues() url.Values {
	values := url.Values{
		"q":     {tran.Q},
		"from":  {tran.From},
		"to":    {tran.To},
		"appid": {strconv.Itoa(tran.Appid)},
		"salt":  {strconv.Itoa(tran.Salt)},
		"sign":  {tran.Sign},
	}
	return values
}

//计算文本的md5值
func SumString(content string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(content))
	bys := md5Ctx.Sum(nil)
	//bys := md5.Sum([]byte(content))//这个md5.Sum返回的是数组,不是切片哦
	value := hex.EncodeToString(bys)
	//fmt.Println(value)
	return value
}

func Coredeal(str string) string {
	tran := NewTranslateModeler(str, From, To)
	values := tran.ToValues()
	resp, err := http.PostForm(Url, values)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func JsonfromByteconfig(bytes []byte) (configfile, error) {
	//用来匹配配置信息的函数

	var b = configfile{}
	var c = new(bytes2.Buffer)
	json.Compact(c, bytes)

	err := json.Unmarshal(c.Bytes(), &b)
	if err != nil {
		return b, err
	}
	return b, nil
}

func JsonfromByte(bytes []byte) (back, error) {
	var b = back{}
	b.Tran_result = []inline{}
	var c = new(bytes2.Buffer)
	json.Compact(c, bytes)
	err := json.Unmarshal(c.Bytes(), &b)
	if err != nil {
		return b, err
	}

	return b, nil
}

func getfilepath(filename string) string {
	filedir := filepath.Dir(filename)
	//get the workdir
	//fmt.Println("给定的路径",filename)
	//fmt.Println("得到的路径", filedir)
	return filedir

}
func getfilereal(fileame string) string {
	filerealname := filepath.Base(fileame)
	return filerealname
}

func Bauduapi() int {
	//打开文件并且将文件格式处理
	fmt.Println("开始处理...")
	file, err := document.Open(Filename)
	file.Styles.Clear()
	if err != nil {
		fmt.Println("error:", err)
	}
	filecopy := document.New()
	filetran := document.New()
	if err != nil {

		fmt.Println(err)
		return -1
	} else {
		for _, i := range file.Paragraphs() {

			for _, k := range i.Runs() {
				//fmt.Println(k.Text())
				jsonpre := Coredeal(k.Text())
				a, _ := JsonfromByte([]byte(jsonpre))
				if len(a.Tran_result) <= 0 {
					continue
				}

				file1 := filecopy.AddParagraph().InsertRunAfter(k)
				file1.Properties().SetFontFamily("宋体")
				file1.AddText(a.Tran_result[0].Dst)
				files := filecopy.AddParagraph().InsertRunAfter(k)
				files.Properties().SetFontFamily("宋体")
				files.AddText(string(k.Text()))

				k.ClearContent()
				k.AddText(a.Tran_result[0].Dst)
				//filetran.AddParagraph().InsertRunAfter(k).AddText(a.Tran_result[0].Dst)
			}
		}
		file.SaveToFile(path.Join(getfilepath(Filename) + "\\replace" + getfilereal(Filename)))
		filetran.SaveToFile(path.Join(getfilepath(Filename) + "\\translated" + getfilereal(Filename)))
		//filetran.SaveToFile("tran" + Filename)
		filecopy.SaveToFile(path.Join(getfilepath(Filename) + "\\copyed" + getfilereal(Filename)))
		//filecopy.SaveToFile("copy" + Filename)
		//fmt.Println(a.Tran_result[0].Dst)
	}
	return 0
}

func config() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("配置文件打开错误")
		return
	}
	var s = string(file)
	s = strings.Replace(s, "  ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	congfig, err := JsonfromByteconfig([]byte(s))
	fmt.Println(s, congfig)
	if err != nil {
		fmt.Println("错误", err)
	}
	if congfig.AppID != "" {
		AppID, _ := strconv.Atoi(congfig.AppID)
		appID = AppID
	}
	if congfig.Key != "" {
		password = congfig.Key
	}
	if congfig.From != "en" {
		From = congfig.From
	}
	if congfig.To != "zh" {
		To = congfig.To
	}

}
func B() int {
	fmt.Println("b")
	time.Sleep(time.Second)
	return 0
}

func wait() {

	go func() {
		for {
			select {
			case <-channel:
				return

			default:
				{
					fmt.Printf("%s", "请等待")
					time.Sleep(time.Nanosecond * 500000000)
					fmt.Printf("\r\r\r")
					fmt.Printf("%s", "请等待....")
					time.Sleep(time.Nanosecond * 500000000)
					fmt.Printf("\r\r\r")
					fmt.Printf("%s", "正在翻译.........")
					time.Sleep(time.Nanosecond * 500000000)
					fmt.Printf("\r\r\r\r\r")
				}
			}
		}

	}()

	if Bauduapi() == 0 {
		channel <- true
		return
	}
}

func Pathstring(pathstring string) string {
	return strings.Trim(strconv.Quote(pathstring), "\"")

}

func logforfile() {
	file, err := os.Create(".\\log.log")
	if err != nil {
		fmt.Println("日志文件创建错误", "err")
	}

	log.SetOutput(file)
	defer file.Close()
}

func IExist(pathname string) bool {
	files, err := os.Open(pathname)
	defer files.Close()
	return !os.IsNotExist(err)
}

func main() {

	config() //读取配置文件
	go logforfile()
	for {
	begun:
		fmt.Println("１．输入路径时请输入绝对路径，建议从文件属性中复制路径位置")
		fmt.Println("２．可以到这个网站转化pdf     https://www.pdf2go.com/zh/pdf-to-word")
		fmt.Println("３．事先对需翻译的文件排版以及去除多余格式会获得好的结果")
		fmt.Println("４．本程序支持中英互译，默认英译中，若有需要请在文件config.txt中配置")
		fmt.Println("５．请输入word文件路径或者将文件拖入本窗口（输入前请事先排版去除分栏等不必要的格式）")
		fmt.Println("６．路径：")
		var a = bufio.NewReader(os.Stdin)
		l, _, err := a.ReadLine()
		if err != nil {
			fmt.Println("输入错误", err)
		}
		inputpath := Pathstring(string(l))
		//fmt.Println(inputpath)
		if !filepath.IsAbs(inputpath) {
			fmt.Println("请输入绝对路径")
			goto begun
		}
		Filename = inputpath

		fmt.Println("你输入的路径为：", inputpath)
		fmt.Println("文件输入是否有效", IExist(inputpath))
		fmt.Println("路径", getfilepath(Filename))

	//	fmt.Println("信息", appID, "      ", password)
		//fmt.Println(path.Join(getfilepath(Filename) + "\\copy" + getfilereal(Filename)))
		//fmt.Println("翻译文件已经生成translated+" + path.Join(getfilepath(Filename), getfilereal(Filename)))
		wait()

		//等待

		fmt.Println("翻译文件已经生成:" + path.Join(getfilepath(Filename)+"\\translated"+getfilereal(Filename)))
		fmt.Println("对照文件已经生成:" + path.Join(getfilepath(Filename)+"\\copyed"+getfilereal(Filename)))

	}
}
