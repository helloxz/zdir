FROM alpine:3.16
RUN mkdir -p /data/apps/zdir
WORKDIR /data/apps/zdir
#复制启动文件
COPY * /root/
# 安装
RUN sh install.sh
# 暴露文件夹和端口
VOLUME /data/apps/zdir/data
EXPOSE 6080
CMD ["/usr/sbin/run.sh"]