#!/bin/bash
#####   name：编译并打包zdir    #####

OS=$1
ARCH=$2

#Zdir版本号
VERSION=3.3.0

#编译linux
compile_linux() {
    rm -rf *.tar.gz zdir.exe
    #编译程序
    go env -w CGO_ENABLED=0
    go env -w GOOS=linux
    
    #根据参数判断架构
    case $ARCH in
        arm64)
            arch="arm64"
        ;;
        arm)
            ARCH="arm"
        ;;
        *)
            ARCH="amd64"
        ;;
    esac

    
    go env -w GOARCH=${ARCH}
    #删除原有的编译文件
    rm -rf main zdir

    go build -ldflags -w main.go
    
    #压缩程序
    upx -9 main
    #重命名程序
    mv main zdir
    exclude="--exclude=model --exclude=data/db --exclude=.gitignore --exclude=docker --exclude=.git --exclude=*.gz --exclude=cli --exclude=config --exclude=controller --exclude=data/public/* --exclude=logs/* --exclude=router --exclude=compile.sh --exclude=config.ini --exclude=go.mod --exclude=go.sum --exclude=main.go --exclude=run.* --exclude=zdir.exe ."
    tar -zcvf zdir_${VERSION}_linux_${ARCH}.tar.gz ${exclude}
    echo "Compiled successfully.(Linux)"
    reset_golang_env
}

# 编译Windows
compile_windows() {
    rm -rf *.tar.gz
    #编译程序
    go env -w CGO_ENABLED=0
    go env -w GOOS=windows
    go env -w GOARCH=amd64

    #删除原有的编译文件
    rm -rf main.exe zdir.exe

    go build -ldflags -w main.go
    #压缩程序
    upx -9 main.exe
    #重命名程序
    mv main.exe zdir.exe
    #打包程序
    exclude="--exclude=model --exclude=data/db --exclude=.gitignore --exclude=docker --exclude=.git --exclude=sh --exclude=*.gz --exclude=cli --exclude=config --exclude=controller --exclude=data/public/* --exclude=logs/* --exclude=router --exclude=compile.sh --exclude=config.ini --exclude=go.mod --exclude=go.sum --exclude=main.go --exclude=zdir ."
    tar -zcvf zdir_${VERSION}_windows_amd64.tar.gz ${exclude}
    echo "Compiled successfully.(Windows)"
    reset_golang_env
}

# 重置为默认变量
reset_golang_env() {
    go env -w CGO_ENABLED=0
    go env -w GOOS=linux
    go env -w GOARCH=amd64
}

case $1 in
    linux)
        compile_linux
    ;;
    windows)
        compile_windows
    ;;
    *)
        echo "Parameter error!"
    ;;
esac
