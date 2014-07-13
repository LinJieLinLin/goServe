package uploadDemo
import (
	"fmt"
	"net/http"
	"github.com/Centny/Cny4go/log"
	"config"
)

func Upload(w http.ResponseWriter, r *http.Request) {
//	fmt.Println(r)
	fmt.Println(r.Form)
	fmt.Println(r.FormValue("data"))
	re:=ReData{}
	re.Code = 0
	if "POST" == r.Method {
		//设置内存缓存大小，貌似这个是最大了
		r.ParseMultipartForm(32 << 20)
		//读取文件
		file, handler, err := r.FormFile("File")
		if err != nil {
			log.E("读取不到文件",err)
			re.Code = 1
			config.Send(re,w)
			return
		}
		//关闭文件
		defer file.Close()
		//检查文件大小并保存，返回没有“www”的文件路径和错误信息
		saveUrl, err := CheckFile(file, handler, r)
		log.I("返回路径：",saveUrl)
		if err != nil {
			log.E("上传文件失败",err)
			re.Code = 2
			config.Send(re,w)
			return
		}
		re.Msg = "上传成功！"
		re.Data = saveUrl
		config.Send(re,w)
	}
}


