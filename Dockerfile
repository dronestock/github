FROM storezhang/alpine:3.17.2


LABEL author="storezhang<华寅>" \
email="storezhang@gmail.com" \
qq="160290688" \
wechat="storezhang" \
description="Drone持续集成Github插件，增加以下功能：1、发布"


# 复制文件
COPY github /bin


RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/github \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 执行命令
ENTRYPOINT /bin/github
