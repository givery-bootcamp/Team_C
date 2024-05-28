FROM mysql:8.0.35

ENV TZ "Asia/Tokyo"

# MySQLのrpmパッケージの公開鍵を明示的に追加
# https://updraft.hatenadiary.com/entry/2022/01/18/150311
RUN rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2023

RUN microdnf update -y \
	&& microdnf install -y glibc-locale-source \
	&& localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8

ENV LANG ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8
