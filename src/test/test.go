package test
import (
	"github.com/Centny/Cny4go/log"
	"fmt"
	"net/http"
)
const  (
	a=100
)
func Test(w http.ResponseWriter, r *http.Request){
	log.I("INOF:","日志启动")
	fmt.Fprintln(w, `<p style="text-align: left;">启动成功</p><br/>​`)
}
func init(){
	fmt.Println("init是初始化函数，只要引入这个包就会自动运行")
}

