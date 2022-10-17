#!/bin/bash
#####   name:Zdir一键安装脚本   #####

#声明Zdir版本
VERSION="3.0.0"

#初始化
init() {
    if [ -e "/usr/bin/yum" ]
	then
		yum -y install tar wget
	elif [ -e "/usr/bin/apt-get" ]
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
		downpath='/data/apps/zdir'
	fi

    if [ -z "${zdirport}" ]
	then
		zdirport='6080'
	fi

}

#下载Zdir
download() {
    mkdir -p /tmp/zdir
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
		sudo ufw allow ${zdirport}/tcp
	fi
}

#安装Zddir
install_zdir() {
    # 复制zdir文件
    cp -ar /tmp/zdir/* ${zdirpath}
    # 初始化zdir
    cd ${zdirpath} && ./zdir init
    # 替换端口
    sed -i "s/:6080/${zdirport}/g" ${zdirpath}/data/config.ini
    # 启动服务
    systemctl start zdir
    # 获取访问地址
    myip=$(curl ip.rss.ink)
    echo "------------------------------------"
    echo "Zdir安装成功，请访问："
    echo "${myip}:${zdirport} 或 http://IP:${zdirport}"
}

# 卸载Zdir
uninstall_zdir() {
    # 停止Zdir
    systemctl stop zdir
    # 删除服务
    rm -rf /etc/systemd/system/zdir.service
    # 重载服务
    systemctl daemon-reload
    echo "Zdir卸载完成，请手动删除Zdir文件夹！"
}

# 获取参数
case $1 in
    install) 
        #安装Zdir
        init && get_parameters && download && chk_firewall && install_zdir
    ;;
    uninstall)
        # 卸载Zdir
        uninstall_zdir
    ;;
    *)
        echo "参数错误!"
        exit
    ;;
esac
