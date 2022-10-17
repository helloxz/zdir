#!/bin/bash
#####   name：编译并打包zdir    #####

#编译linux
compile_linux() {
    rm -rf *.tar.gz *.exe
    #编译程序
    go env -w CGO_ENABLED=0
    go env -w GOOS=linux
    go env -w GOARCH=amd64
    #删除原有的编译文件
    rm -rf main zdir

    go build -ldflags -w main.go
    
    #压缩程序
    upx -9 main
    #重命名程序
    mv main zdir
    tar -zcvf zdir_3.0.0_linux_amd64.tar.gz --exclude=*.gz --exclude=cli --exclude=config --exclude=controller --exclude=data/public/* --exclude=logs/* --exclude=router --exclude=compile.sh --exclude=config.ini --exclude=go.mod --exclude=go.sum --exclude=main.go --exclude=run.* --exclude=zdir.exe .
    echo "Compiled successfully.(Linux)"
    reset_golang_env
}

# 编译Windows
compile_windows() {
    rm -rf *.tar.gz *.exe
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
    tar -zcvf zdir_3.0.0_windows_amd64.tar.gz --exclude=sh --exclude=*.gz --exclude=cli --exclude=config --exclude=controller --exclude=data/public/* --exclude=logs/* --exclude=router --exclude=compile.sh --exclude=config.ini --exclude=go.mod --exclude=go.sum --exclude=main.go --exclude=zdir .
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
