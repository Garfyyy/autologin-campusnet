# autologin-campusnet

## 介绍

自动检测网络状态，自动登录 $GXNU$ 校园网。

## 下载

1. 使用 `git clone`

    ```shell
    git clone https://github.com/Garfyyy/autologin-campusnet.git
    ```

2. 直接下载 `autologin.go` 文件

## 使用

根据需要修改 `autologin.go` 中对应内容

- `DDDDD`: 账号
- `upass`: 密码
- `R3`: 选择网络类型
  - `0`：校园网
  - `1`：电信
  - `2`：联通
  - `3`：移动
- `CHECK_INTERVAL`: 检查网络状态间隔时间

修改后运行代码:

```shell
go run autologin.go
```

或者编译后使用:

```shell
go build autologin.go -o autologin
./autologin
```

## *PS*

个人用于放假回家远控实验室电脑，避免电脑断网。

(如果你也有同样的需求，可以把编译后的文件放到开机启动项中)
