#! /bin/sh

mkdir /etc/polo
wget https://olivetree.oss-cn-hangzhou.aliyuncs.com/polo/config.toml -O /etc/polo/config.toml
wget https://olivetree.oss-cn-hangzhou.aliyuncs.com/polo/polo -O /usr/local/bin/polo
chmod +x /usr/local/bin/polo
