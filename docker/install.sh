#!/bin/sh

VERSION=3.1.1

# 初始化环境
init() {
    #更新软件
    apk update
    #安装timezone
    apk add -U tzdata
    #安装必要软件
    apk add curl wget
    #查看时区列表
    ls /usr/share/zoneinfo
    #拷贝需要的时区文件到localtime
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
    #查看当前时间
    date
    #为了精简镜像，可以将tzdata删除了
    apk del tzdata

    #拷贝运行文件
    chmod +x /root/run.sh
    cp /root/run.sh /usr/sbin/
}

#下载Zdir
download(){
    cd /root && mkdir zdir && cd zdir
    name=zdir_${VERSION}_linux_amd64.tar.gz
    wget http://soft.xiaoz.org/zdir/${VERSION}/${name}

    #解压
    tar -xvf ${name}
    #拷贝文件
    cp -ar /root/zdir/* /data/apps/zdir
    #添加执行权限
    chmod +x /data/apps/zdir/zdir
    #删除压缩文件
    rm -rf /root/${name}
    rm -rf /data/apps/zdir/*.tar.gz

    # 下载ARM64架构
    cd /tmp && mkdir zdir && cd zdir
    name=zdir_${VERSION}_linux_arm64.tar.gz
    wget http://soft.xiaoz.org/zdir/${VERSION}/${name}
    #解压
    tar -xvf ${name}
    #拷贝文件
    cp -ar /tmp/zdir/zdir /data/apps/zdir/zdir_arm64
    #删除临时目录
    rm -rf /tmp/*

    # 下载ARM架构
    cd /tmp && mkdir zdir && cd zdir
    name=zdir_${VERSION}_linux_arm.tar.gz
    wget http://soft.xiaoz.org/zdir/${VERSION}/${name}
    #解压
    tar -xvf ${name}
    #拷贝文件
    cp -ar /tmp/zdir/zdir /data/apps/zdir/zdir_arm
    #删除临时目录
    rm -rf /tmp/*
}


init && download