# 微信消息转发机器人

## 介绍

把微信消息全部转发到已添加的好友，适用于多个微信号管理的场景，不用来回切换微信号，只需在一个微信号上操作即可。

注意：目前在手机上切换微信号时，会导致微信网页版退出登录，需要重新扫码登录。

## 使用

### 下载

```shell
git clone https://github.com/t3ls/wechat-forward-bot.git
```

### 修改你的目标账号

修改`docker-compose.yml`文件中的`forward_target_username`字段，填入你的目标账号昵称。

### 然后运行

```bash
cd wechat-forward-bot && docker compose up -d
```

