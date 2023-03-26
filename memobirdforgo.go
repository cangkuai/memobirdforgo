package memobirdforgo

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"time"

	"github.com/doumadou/mahonia"
	"github.com/imroc/req/v3"
)

// 记录初始化数据
type memobirdbody struct {
	ak     string
	did    string
	userID string
}

// 获取打印结果的结构体
type printpaperResultstruct struct {
	Showapi_res_code  int    `json:"showapi_res_code"`
	Showapi_res_error string `json:"showapi_res_error"`
	Result            int    `json:"result"`
	SmartGuid         string `json:"smartGuid"`
	Printcontentid    string `json:"printcontentid"`
}

// 获取点位图结果的结构体
type SignalBase64PicResultstruct struct {
	Showapi_res_code  int    `json:"showapi_res_code"`
	Showapi_res_error string `json:"showapi_res_error"`
	Result            string `json:"result"`
}

// 获取打印状态结果的结构体
type getprintstateResultstruct struct {
	Showapi_res_code  int    `json:"showapi_res_code"`
	Showapi_res_error string `json:"showapi_res_error"`
	Printflag         string `json:"printflag"`
	Printcontentid    string `json:"printcontentid"`
}

var urls = "http://open.memobird.cn/"

// 初始化接口
func newmemobirdapi(ak string, did string, userids string) *memobirdbody {
	return &memobirdbody{ak: ak, did: did, userID: userids}
}

// 转换为gbk的base64编码（官方要求）
func Text(text string) string {
	enc := mahonia.NewEncoder("gbk")
	if ret, ok := enc.ConvertStringOK(text); ok {
		return "T:" + base64.StdEncoding.EncodeToString([]byte(ret))
	} else {
		panic(ok)
	}
}

// 获取时间戳
func gettimestamp() string {
	t := time.Now()
	timestamp := t.Format("2006-01-02") + " " + t.Format("15:04:05")
	return timestamp
}

// 获取点位图
func (c memobirdbody) getSignalBase64Pic(textss string) string {
	var result SignalBase64PicResultstruct
	client := req.C()
	client.R().SetFormData(map[string]string{"ak": c.ak, "imgBase64String": textss}).SetSuccessResult(&result).Post("http://open.memobird.cn/home/getSignalBase64Pic")
	return result.Result
}

// 纯文字打印函数
func (c memobirdbody) printpaperbytext(textss string) printpaperResultstruct {
	var result printpaperResultstruct
	client := req.C()
	client.R().SetFormData(map[string]string{"ak": c.ak, "timestamp": gettimestamp(), "printcontent": Text(textss), "memobirdID": c.did, "userID": c.userID}).SetSuccessResult(&result).Post("http://open.memobird.cn/home/printpaper")
	return result
}

// 通过图片base64打印
func (c memobirdbody) printpaperbyimgbase64(textss string) printpaperResultstruct {
	var result printpaperResultstruct
	client := req.C()
	client.R().SetFormData(map[string]string{"ak": c.ak, "timestamp": gettimestamp(), "printcontent": "P:" + c.getSignalBase64Pic(textss), "memobirdID": c.did, "userID": c.userID}).SetSuccessResult(&result).Post("http://open.memobird.cn/home/printpaper")
	return result
}

// 通过图片文件打印
func (c memobirdbody) printpaperbyimgfile(path string) printpaperResultstruct {
	file, _ := os.Open(path)
	defer file.Close()
	imgByte, _ := ioutil.ReadAll(file)
	base64string := base64.StdEncoding.EncodeToString(imgByte)
	return c.printpaperbyimgbase64(base64string)
}

// 检查打印状态
func (c printpaperResultstruct) getprintstate(ak string) getprintstateResultstruct {
	var result getprintstateResultstruct
	client := req.C()
	client.R().SetFormData(map[string]string{"ak": ak, "timestamp": gettimestamp(), "printcontentid": c.Printcontentid}).SetSuccessResult(&result).Post("http://open.memobird.cn/home/getprintstatus")
	return result
}
