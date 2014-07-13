/**
 * Created by Lin on 14-6-8.
 */
//ajax请求
var LinHttp = function (arg_method, arg_url, arg_data, arg_callBack) {
    var sendHttp = {
        Start: function () {
            $.ajax({
                type: arg_method,
                url: arg_url,
//                timeout:arg_timeOut,
                data: arg_data,
                dataType: "json",
                success: sendHttp.CallBackData,
                error: sendHttp.Error
            });
        },
        CallBackData: function (arg_data) {
            arg_callBack(arg_data);
        },
        Error: function (arg_errMsg) {
            arg_callBack({code: 1});
        }
    };
    sendHttp.Start();
}
//获取浏览器信息（返回客户端的名字）
function getBrowser()
{
    if(navigator.userAgent.indexOf("MSIE")>0) {
        return "MSIE";
    }
    if(isFirefox=navigator.userAgent.indexOf("Firefox")>0){
        return "Firefox";
    }
    if(isSafari=navigator.userAgent.indexOf("Safari")>0) {
        return "Safari";
    }
    if(isCamino=navigator.userAgent.indexOf("Camino")>0){
        return "Camino";
    }
    if(isMozilla=navigator.userAgent.indexOf("Gecko/")>0){
        return "Gecko";
    }

}