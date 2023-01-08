#!/bin/bash
#####   name:Zdir一键安装脚本   #####

#声明Zdir版本
VERSION="3.1.1"

#检查是否已经安装过
is_install() {
    if [ -f "/etc/systemd/system/zdir.service" ]
    then
        echo "Zir已经安装，无需重复安装！"
        exit
    fi
}

#初始化
init() {
    if [ -e "/usr/bin/yum" ]
	then
		yum -y install tar wget
	elif [ -e "/usr/bin/apt-get" ]
    then
		#更新软件，否则可能make命令无法安装
		apt-get update
		apt-get install -y wget
	fi
}

#获取用户参数
get_parameters() {
    read -p "输入Zdir安装路径（请填写绝对地址，默认/data/apps/zdir）:" zdirpath
    read -p "输入Zdir监听端口（默认为6080）:" zdirport

    #设置默认路径和端口
    if [ -z "${zdirpath}" ]
	then
		zdirpath='/data/apps/zdir'
	fi

    if [ -z "${zdirport}" ]
	then
		zdirport='6080'
	fi
}

#下载Zdir
download() {
    mkdir -p /tmp/zdir
    cd /tmp/zdir
    wget http://soft.xiaoz.org/zdir/${VERSION}/zdir_${VERSION}_linux_amd64.tar.gz
    #解压
    tar -xvf zdir_${VERSION}_linux_amd64.tar.gz
}

#自动放行端口
chk_firewall(){
	if [ -e "/etc/firewalld/zones/public.xml" ]
	then
		firewall-cmd --zone=public --add-port=${zdirport}/tcp --permanent
		firewall-cmd --reload
	elif [ -e "/etc/ufw/before.rules" ]
	then
		ufw allow ${zdirport}/tcp
	fi
}

#安装Zddir
install_zdir() {
    #创建文件夹
    mkdir -p ${zdirpath}
    # 复制zdir文件
    cp -ar /tmp/zdir/* ${zdirpath}
    # 初始化zdir
    cd ${zdirpath} && ./zdir init
    # 替换端口
    sed -i "s/6080/${zdirport}/g" ${zdirpath}/data/config.ini
    #删除临时文件
    rm -rf /tmp/zdir
    # 启动服务
    systemctl start zdir
    # 获取访问地址
    myip=$(curl -s ip.rss.ink)
    echo "---------------------------------------------"
    echo "Zdir安装成功，请访问："
    echo "http://${myip}:${zdirport}"
    echo "或访问："
    echo "http://IP:${zdirport}"
    echo "---------------------------------------------"
}

# 卸载Zdir
uninstall_zdir() {
    # 停止Zdir
    systemctl stop zdir
    # 获取Zdir安装目录
    zirpath=`grep "ExecStart=" /etc/systemd/system/zdir.service|sed 's/ExecStart=//'|sed 's/zdir start//'`
    # 删除服务
    rm -rf /etc/systemd/system/zdir.service
    # 重载服务
    systemctl daemon-reload
    echo "Zdir卸载完成，请手动执行:rm -rf ${zirpath}删除Zdir文件夹！"
}


# 获取参数
case $1 in
    install) 
        #安装Zdir
        is_install && init && get_parameters && download && chk_firewall && install_zdir
    ;;
    uninstall)
        # 卸载Zdir
        uninstall_zdir
    ;;
    test)
        get_parameters
    ;;
    *)
        echo "参数错误!"
        exit
    ;;
esac
