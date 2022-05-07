package main

import (
	"project-layout-go/pkg/utils/debugutil"

	"github.com/go-resty/resty/v2"
)

func main() {

	client := resty.New()
	debugutil.DebugPrint(client, 0)
	req := client.R()
	debugutil.DebugPrint(req, 0)
	resp, _ := req.EnableTrace().Get("http://www.bai.com")

	debugutil.DebugPrint(resp, 0)
	debugutil.DebugPrint(resp.Request.TraceInfo(), 0)
}
