#!/bin/sh

####    name:zdir运行脚本   #####

# 检查目录文件
check_dir() {
    if [ ! -d "/data/apps/zdir/data/dist" ]
    then
        cp -ar /root/zdir/data/dist /data/apps/zdir/data/
    fi

    if [ ! -d "/data/apps/zdir/data/public" ]
    then
        cp -ar /root/zdir/data/public /data/apps/zdir/data/
    fi

    if [ ! -f "/data/apps/zdir/data/config.ini" ]
    then
        cp /root/zdir/config.simple.ini /data/apps/zdir/data/config.ini
    fi
}

# 运行zdir
run() {
    cd /data/apps/zdir/
    ./zdir start
}

check_dir && run