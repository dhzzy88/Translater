# Translastion Of Literature


## 此程序是为中国大学生阅读文献的辅助工具

![本程序](https://raw.githubusercontent.com/dhzzy88/Translation_Of_Literaure/master/main.png)
![Running](https://raw.githubusercontent.com/dhzzy88/Translation_Of_Literaure/master/runtime.png)
  
## 只能支持word(.docx)文件  
  由于使用的百度翻译的通用api翻译的准确度不是十分好，仅作为辅助工具使用，由于解析PDF比较复杂，所以本程序暂时没有提供翻译PDF文件的功能，需要将所需要的文档转换成为word文档（.docx）。 比较常用的是迅捷<http://app.xunjiepdf.com/pdf2word>在线转换可能对文件大小有限制(2M)。或者smallpdf<https://smallpdf.com/cn/pdf-to-word>对文件没有限制，功能免费，但是比较慢。
  
  
  
## 配置文件
   
   使用之前需要申请百度翻译通用api的appid和key，只要有百度账号可以直接登录申请（一般都有百度云盘估计都注册了）。具体申请流程百度经验<https://jingyan.baidu.com/article/3f16e00305bb552591c10304.html>，具体点击了解。
   配置文件为`config.json`
   所给的内容为配置的格式，尤其冒号两边不能有空格，如下：
   >
   {
   "from":"en",
   "to":"zh",
   "appid":"2015063000000001",
   "key": "12345678"
   }
   >
   
   最后一行没有逗号，注意信息要加英文输入法的双引号（冒号括号一样的）
   from和to为翻译的语言，一般不用改默认英语翻译成中文
   
   
   
## 使用直接将文件拖入框内即可
   如果文件名中有中文注意检查路径在拖入后是否包在在双引号内，如果有将双引号删除
![如图](https://raw.githubusercontent.com/dhzzy88/Translation_Of_Literaure/master/yinhao.png)
   
   
##  关于生成的文件
    
    1. 生成的translate+.....为只有翻译结果的
    2. 生成的copy+...为中英对照的
    3. 生成的replace 为正式的翻译文件，与源文献格式相同，需要排版
    建议选择最后replace+...文件使用

![生成文件](https://raw.githubusercontent.com/dhzzy88/Translation_Of_Literaure/master/result.png)

    
## Example

![翻译前](https://raw.githubusercontent.com/dhzzy88/Translation_Of_Literaure/master/fanyiqian.png)
![翻译前](https://raw.githubusercontent.com/dhzzy88/Translation_Of_Literaure/master/fanyiqian2.png)
   
   
![翻译后](https://raw.githubusercontent.com/dhzzy88/Translation_Of_Literaure/master/fanyigou1.png)
![翻译后](https://raw.githubusercontent.com/dhzzy88/Translation_Of_Literaure/master/fanyihou2.png)
