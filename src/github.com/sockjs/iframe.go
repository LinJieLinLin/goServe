package sockjs

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"text/template"
)

var tmpl, _ = template.New("asdas").Parse(iframe_body)

func (ctx *context) iframeHandler(rw http.ResponseWriter, req *http.Request) {
	etag_req := req.Header.Get("If-None-Match")
	hash := md5.New()
	hash.Write([]byte(iframe_body))
	etag := fmt.Sprintf("%x", hash.Sum(nil))
	if etag == etag_req {
		rw.WriteHeader(http.StatusNotModified)
		return
	}
	setContentType(rw.Header(), "text/html; charset=UTF-8")
	disableCache(rw.Header())
	setExpires(rw.Header())
	rw.Header().Add("ETag", etag)
	tmpl.Execute(rw, ctx.SockjsUrl)
}

var iframe_body = `<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <script>
    document.domain = document.domain;
    _sockjs_onload = function(){SockJS.bootstrap_iframe();};
  </script>
  <script src="{{.}}"></script>
</head>
<body>
  <h2>Don't panic!</h2>
  <p>This is a SockJS hidden iframe. It's used for cross domain magic.</p>
</body>
</html>`
