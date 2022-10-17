#!/bin/bash

#####   name:Linux 服务注册 #####

create_service(){
    #获取当前目录
    zdir_path=`echo $(pwd) | sed 's/sh//g'`
    #二进制文件
    zdir_bin=${zdir_path}/zdir
    #创建服务文件
    touch /etc/systemd/system/zdir.service
    # 写入服务文件
    cat >>/etc/systemd/system/zdir.service << EOF
[Unit]
Description=zdir
After=network.target
[Service]
WorkingDirectory=${zdir_path}
ExecStart=${zdir_bin} start
[Install]
WantedBy=multi-user.target
EOF
    #重载服务
    systemctl daemon-reload
}

create_service