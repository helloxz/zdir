# 关于Zdir

Zdir是一款使用使用Golang + Vue3开发的多功能私有存储程序，集成了文件索引、在线预览和分享功能，支持WebDAV和离线下载，非常适合安装在NAS设备或大容量VPS上，是个人、工作室和小团队分享文件的理想选择。

![333bf4d33c30d70c.png](https://img.rss.ink/imgs/2024/06/16/333bf4d33c30d70c.png)

## 起源

### 2018年

许多年前，公司运维部门利用nginx搭建了一个文件索引平台，用于分享常用软件供员工下载。然而，nginx的索引功能极为有限，界面简陋、功能单一。xiaoz作为一个喜欢折腾VPS的爱好者，我萌生了自建文件分享平台的想法。经过尝试Fancy Index、PHP Directory Lister、h5ai等解决方案，我依然不满意。最终在2018年，我用PHP正式开发了Zdir `1.x`。

> 2021年，`2.x.x`版本正式结束了其生命周期，至此我们不再维护PHP版本的Zdir！

### 2022年

随着时间的推移，`2.x.x`版本的问题日益凸显，且由于PHP的局限性难以解决。于是，在2022年，我采用新技术Golang和Vue3进行了全面重构，发布了全新的`3.x.x`版本。

### 2024年

今天，我们自豪地宣布Zdir `4.x`版本的到来。`4.x`在`3.x.x`版本的基础上进行了升级优化，前端界面重新布局和设计，部分风格参考了Alist，整体依旧以简洁实用为主。

> 从4.x版本开始，我们将不再开源。之前的源码仍然可以在Github上找到：[https://github.com/helloxz/zdir](https://github.com/helloxz/zdir "https://github.com/helloxz/zdir")

## 4.x版本功能

- [x] 离线下载
- [x] 音乐列表模式
- [x] 文件索引
- [x] 文件说明
- [x] 文件预览（支持图片、文档、音乐、视频等预览）
- [x] 私有文件
- [x] 私有文件分享
- [x] 文件管理（上传、下载、删除、重命名）
- [x] 图片预览
- [x] **API支持**
- [x] **WebDAV服务端支持**
- [x] **全局文件搜索**
- [x] 前后台一体化
- [x] 文本编辑
- [x] 文件移动、复制
- [x] 多语言
- [ ] Ofiice预览


## 团队成员

目前我们的团队有2位成员，其中一位成员是[mTab作者@tushan](https://shop.xiuping.net/buy/1.html "mTab作者") ，以及[OneNav作者@xiaoz](https://shop.xiuping.net/onenav/index "OneNav作者")，其中tushan负责前端开发，xiaoz负责后端开发。

## 联系我们

* 技术支持：请添加微信`xiaozme`

## 其他

* 官网：[https://www.zdir.pro/](https://www.zdir.pro/ "https://www.zdir.pro/")
* 演示：[https://soft.xiaoz.org/#/](https://soft.xiaoz.org/#/ "https://soft.xiaoz.org/#/")
* 购买授权：[https://shop.xiuping.net/zdir/index](https://shop.xiuping.net/zdir/index "https://shop.xiuping.net/zdir/index")
