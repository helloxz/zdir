# zdir
Zdir是一款使用PHP开发的目录列表程序，无需数据库，体积小巧，功能完善。

![](https://imgurl.org/upload/1806/349f3b54028d58d6.png)



### 环境要求

* PHP >= 5.6(建议PHP 7.2)
* 文件索引依赖于`curl`组件
* 如果需要获得更高级的文件管理功能，依赖于Fileinfo, iconv, zip, tar and mbstring组件（非必须）

### 使用方法

* 下载源码放到站点根目录
* 将`config.simple.php`修改为`config.php`
* 修改`config.php`设置自己的网站标题/关键词/描述/密码等信息
* 如果需要排除某个目录，再次修改`config.php` ，里面有说明
* 更多说明请查看帮助文档：[https://dwz.ovh/zdir](https://dwz.ovh/zdir)

### Docker方式部署

```bash
docker run --name="zdir"  \
    -d -p 1080:80 --restart=always \
    -v /data/wwwroot/default:/data/wwwroot/default \
    helloz/zdir \
    /usr/sbin/run.sh
```

更多说明可参考：[https://www.xiaoz.me/archives/14809](https://www.xiaoz.me/archives/14809)

### 主要功能
- [x] 目录浏览
- [x] MarkDown文件预览
- [x] CSS/JavaScript一键复制
- [x] 文件hash
- [x] 图片预览
- [x] 视频播放（支持.mp4 .ts .m3u8等部分格式）
- [x] 音频播放（支持`.mp3` `.wav` `.flac` `.ape` ）
- [x] 文本查看器（支持的格式有.txt .sh .py .go .c .cpp）
- [x] Office在线预览（支持.doc .docx .xls .xlsx .ppt .pptx）,**注意：如果是内网或IP访问方式不支持预览**
- [x] 文件索引
- [x] 二维码生成
- [x] 文件管理（上传/删除/编辑等）
- [x] 文件搜索
- [x] 密码验证
- [x] 支持中文显示

### Demo
* 演示地址：[http://soft.xiaoz.org/](http://soft.xiaoz.org/)
* 备用演示：[https://wget.ovh/](https://wget.ovh/)

### 获取捐赠版
扫描下方二维码，捐赠大于`30元`以上可获得捐赠版，捐赠版可享受首次技术支持及去除广告，捐赠后请联系我的QQ:337003006获取。

![](https://imgurl.org/upload/1712/cb349aa4a1b95997.png)

### 联系我
* Blog: [https://www.xiaoz.me/](https://www.xiaoz.me/)
* QQ: 337003006
* 社区支持： [https://ttt.sh/](https://ttt.sh/category/15/)
* 帮助文档：[https://dwz.ovh/zdir](https://dwz.ovh/zdir)



### 感谢

Zdir的诞生离不开以下开源项目，在此表示感谢。

* [tinyfilemanager](https://github.com/prasathmani/tinyfilemanager)
* [parsedown](https://github.com/erusev/parsedown)
* [LayUI](https://github.com/sentsin/layui)