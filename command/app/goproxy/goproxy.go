package goproxy

import (
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
)

/*
https://goproxy.cn/
GOPROXY 设置为 http://localhost:8080 来试用它。另外，我们也建议你把 GO111MODULE 设置为 on
*/
func Run() {
	http.ListenAndServe("localhost:8080", &goproxy.Goproxy{
		GoBinEnv: append(
			os.Environ(),
			"GOPROXY=https://goproxy.cn,direct", // 使用 Goproxy.cn 作为上游代理
			"GOPRIVATE=gitlab.com/*,gitlab.umcasual.cn/*,github.com/Ho-J/*", // 解决私有模块的拉取问题（比如可以配置成公司内部的代码源）
		),
		ProxiedSUMDBs: []string{
			"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org", // 代理默认的校验和数据库
		},
	})
}
