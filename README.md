# memobirdforgo——又一个go语言咕咕机的开发类

## 一.本工具类实现的功能

1.打印纯文本内容

2.通过base64打印图片

3.通过文件打印图片

## 二.使用步骤

1.在咕咕机开放普通获取ak(https://open.memobird.cn/)

2.安装本工具类所使用的依赖

```bash
github.com/doumadou/mahonia
github.com/imroc/req/v3
```

3.获取你的你的设备api

如果你不知道可以使用api获取（格式：http://open.memobird.cn/home/setuserbind?ak=你的ak&timestamp=时间戳&memobirdID=设备编号&useridentifying=一个随机数 ）

4.初始化接口

```go
memobird := newmemobirdapi("你的ak", "你的设备编号", "你的设备api")
```

5.开始打印

```go
//打印纯文本
Result:=memobird.printpaperbytext("你要打印的文字")
//打印base64编码的图片
Result:=memobird.getSignalBase64Pic("你要打印图片的base64")
//打印以文件存储的图片
Result:=memobird.printpaperbyimgfile("你要打印图片的路径")
```

6.查询打印结果

(这个功能官方api炸了,正在等待官方修复)