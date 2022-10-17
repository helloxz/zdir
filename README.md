# Zdir

使用Golang + Vue3开发的轻量级目录列表程序，支持Linux、Docker、Windows部署，适合个人或初创公司文件分享使用，亦可作为轻量级网盘使用。

![](https://img.rss.ink/imgs/2022/10/17/10d74765a20fdc7a.png)

## 功能特点

- [x] 目录列表
- [x] MarkDown预览
- [x] 支持搜索当前目录与全局搜索（备注：全局搜索仅Linux支持）
- [x] 视频预览（支持H.264编码的`.mp4`格式及`.m3u8`）
- [x] 音频预览
- [x] 图片预览
- [x] 代码与文本预览，支持部分代码高亮
- [x] CSS/JavaScript一键复制
- [x] Office在线预览（支持.doc .docx .xls .xlsx .ppt .pptx）,**注意：如果是内网或IP访问或非标准多端口方式不支持预览**
- [x] 二维码生成
- [x] 支持中文显示
- [x] 支持Linux、Docker、Windows等多种部署方式
- [ ] 后台管理（站点信息设置等）
- [ ] 文件管理（上传、删除、编辑、复制、移动）
- [ ] 私有文件
- [ ] 私有文件分享
- [ ] 音乐播放列表

## 快速开始

**Linux一键安装：**

如果您想快速安装Zdir，可以使用Zdir官方提供的一键安装脚本，只需要执行下面的命令：

```bash
# CentOS系统
yum -y install curl
curl "http://soft.xiaoz.org/zdir/sh/zir.sh" | bash -s install
# Debian or Ubuntu系统
apt-get install curl
curl "http://soft.xiaoz.org/zdir/sh/zir.sh" | bash -s install
```

* 默认安装路径为：`/data/apps/zdir`
* 需要公开的文件列表路径位于`/data/apps/zdir/data/public`

安装完毕后访问`http://IP:6080`即可

___

一键安装脚本适合对Linux系统不太熟悉或者想快速体验Zdir的朋友，您可以参考帮助文档：[https://doc.xiaoz.me/books/zdir-3](https://doc.xiaoz.me/books/zdir-3) 获取更多安装方式。

## 文档 & Demo

* 帮助文档：[https://doc.xiaoz.me/books/zdir-3](https://doc.xiaoz.me/books/zdir-3)
* Demo：[http://soft.xiaoz.org/](http://soft.xiaoz.org/)

## 问题反馈

* 论坛：[https://xiawen.cc/t/zdir](https://xiawen.cc/t/zdir)
* QQ:446199062
* QQ群：932795364
* TG:xiaozme



