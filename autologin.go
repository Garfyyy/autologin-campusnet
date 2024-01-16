package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

/*
根据需要修对应内容
- `DDDDD`: 账号
- `upass`: 密码
- `R3`: 选择网络类型
  - `0`：校园网
  - `1`：电信
  - `2`：联通
  - `3`：移动
- `CHECK_INTERVAL`: 检查网络状态间隔时间
*/

const (
	DDDDD         = "xxx"
	UPASS         = "xxx"
	R3            = 0
	checkInterval = 10 * time.Minute
	maxRetry      = 17
	pingTarget    = "www.baidu.com" // ipv4.ip.sb
)

var loginURL = fmt.Sprintf(
	"https://yc.gxnu.edu.cn/drcom/login?callback=dr1003&DDDDD=%s&upass=%s&0MKKey=123456&R1=0&R2=&R3=%d&R6=0&para=00&v6ip=&terminal_type=1&lang=en",
	DDDDD,
	UPASS,
	R3,
)

func ping() bool {
	log.Printf("ping %s...", pingTarget)
	cmd := exec.Command("ping", "-4", "-n", "2", pingTarget)
	return cmd.Run() == nil
}

func login() {
	resp, err := http.Get(loginURL)
	if err != nil {
		log.Printf("requests error: %v", err)
		return
	}
	resp.Body.Close()
}

func main() {
	for {
		cnt := 0
		flag := true
		for !ping() {
			cnt++
			log.Printf("Ping failure, relogin %d...", cnt)
			login()
			if cnt >= maxRetry {
				flag = false
				break
			}
		}
		if flag {
			log.Println("\033[1;32mLogin Success\033[0m")
		} else {
			log.Printf("Ping \033[1;31mfailure\033[0m after %d attempts, will retry in %.2fmin...", maxRetry, checkInterval.Minutes())
		}
		time.Sleep(checkInterval)
	}
}
