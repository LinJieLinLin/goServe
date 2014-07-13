/**
 * Created by Lin on 14-6-8.
 */
var browser = getBrowser();
var upScope = "";
function uploadCtrl($scope) {
    upScope = $scope;
    upScope.upImgUrl = "";
}
var picUploader = {};
picUploader.postFrom = function (arg_file,arg_id) {
    try {
        $("#"+arg_id).ajaxSubmit({
            dataType: 'json', //数据格式为json
            resetForm: true,//布尔标志，表示如果表单提交成功是否进行重置。
            clearForm: true,//布尔标志，表示如果表单提交成功是否清除表单数据。
            beforeSubmit: function (xmlHttp) { //上传前检测
//                document.getElementById(imgId).src="../../../../preImg/ad_pic/loading.gif"
//                setImgClick(false)
            },
            success: function (data) {
                alert(angular.toJson(data));
                upScope.upImgUrl = data.data;
                upScope.$apply();
            },
            error: function (data, e, status) { //上传失败
                alert("上传失败");
            }
        });
    } catch (error) {
        alert("上传出错" + error)
    }
}
picUploader.uploadCheck = function (arg_data, arg_types, arg_size) {
    if (!arg_data || !arg_data.fileType || !arg_types || !arg_size) {
        alert("参数有误！");
        return false
    }
    if (arg_types.lastIndexOf(arg_data.fileType.toLowerCase()) == -1) {
        alert("不支持该文件格式！");
        return false
    }
    if (arg_data.fileSize > arg_size * 1024 * 1024) {
        alert("文件大小不能超过"+arg_size+"Mb!")
        return false
    }
    return true
}
function upload(arg_fileId,arg_formId) {
    if(!arg_fileId||!arg_formId){
        alert("传入数据有误！")
        return false
    }
    var picMsg = document.getElementById(arg_fileId);
    var fileType = picMsg.value.substring(picMsg.value.lastIndexOf(".") + 1, picMsg.value.length);
    var fileSize = 0;
    if (browser != "MSIE") {
        fileSize = picMsg.files[0].size;
    }
    var types = "jpg,jpeg,png,bmp";
    //size单位M
    var size = 1;
    var res = picUploader.uploadCheck({fileType: fileType, fileSize: fileSize}, types, size)
    if (!res) {
        return false
    }
    picUploader.postFrom(picMsg,arg_formId)
}