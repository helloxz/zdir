#!/bin/sh

#####   name:Zdir3升级脚本      #####

PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/bin:/sbin
export PATH

#获取Zdir路径
zdir_path=`grep "ExecStart=" /etc/systemd/system/zdir.service|sed 's/ExecStart=//'|sed 's/zdir start//'`
#获取Zdir版本
VERSION=$1

# 检查函数
check(){
    # 如果zdir目录为空，则直接停止
    if [ -z $zdir_path ]
    then
        echo "未检测到Zdir安装，升级停止！"
        exit
    fi

    # 如果版本号为空
    if [ -z $VERSION ]
    then
        echo "请输入要升级的版本号！"
        exit
    fi
}

# 升级函数
upgrade(){
    # 下载对应版本
    cd $zdir_path
    #删除原有的压缩包
    rm -rf *.tar.gz
    name=zdir_${VERSION}_linux_amd64.tar.gz
    wget http://soft.xiaoz.org/zdir/${VERSION}/${name}

    #停止Zdir
    systemctl stop zdir

    #获取 zdir pid
    pid=`ps -ef|grep zdir|grep -v grep|awk '{print $2}'`
    if [ ! -z $pid ]
    then
        echo "正在停止Zdir进程:"$pid
        kill -9 $pid
        sleep 10
    fi

    # 删除原来的静态文件
    rm -rf $zdir_path"data/dist/assets/*"

    #解压Zdir
    tar -xvf $name

    # 添加执行权限
    chmod +x ./zdir

    # 重新启动
    systemctl start zdir

    echo "--------------------------------------------------"
    echo "Zdir $VERSION升级完毕！"
    echo "如需帮助请访问：https://dwz.ovh/24330"
    echo "--------------------------------------------------"
}

check && upgrade