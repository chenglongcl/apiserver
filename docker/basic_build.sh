#! /bin/bash

build(){

    sudo mkdir -p /dockerMaps/data \
        /dockerMaps/conf \
        /dockerMaps/logs

    sudo mkdir -p /dockerMaps/data/mysql \
         /dockerMaps/conf/mysql \
         /dockerMaps/logs/mysql

    sudo mkdir -p /dockerMaps/data/redis \

    sudo mkdir -p /webroot/go/apiserver \
        /dockerMaps/tool

    sudo docker network create custom
}

build
echo "done!"