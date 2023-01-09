#!/bin/sh
####    name:zdir运行脚本   #####

# zdir配置文件默认目录
zdir_config_file="/data/apps/zdir/data/config/config.ini"

# 检查目录文件
check_dir() {
    # 3.2.0不再需要dist目录
    # if [ ! -d "/data/apps/zdir/data/dist" ]
    # then
    #     cp -ar /root/zdir/data/dist /data/apps/zdir/data/
    # fi

    if [ ! -d "/data/apps/zdir/data/public" ]
    then
        cp -ar /root/zdir/data/public /data/apps/zdir/data/
    fi

    if [ ! -f ${zdir_config_file} ]
    then
        cp /root/zdir/config.simple.ini ${zdir_config_file}
    fi
}

# 运行zdir
run() {
    cd /data/apps/zdir/
    # 判断架构
    get_arch=$(arch)
    if [[ "${get_arch}" == "x86_64" ]]
    then
        ./zdir start
    elif [[ "${get_arch}" == "aarch64" ]]
    then
        ./zdir_arm64 start
    else
        ./zdir_arm start
    fi
}

check_dir && run