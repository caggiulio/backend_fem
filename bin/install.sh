#!/bin/bash


#
#   	Script build backend	
#        Copyright (C) 2015+  Gabriele Baldoni
#

export GOPATH=$GOPATH:$HOME/backend_fem

export APPPATH=$HOME/backend_fem

go build access
go build server


go install access
go install server


go build -o $APPPATH/bin/fembackend $APPPATH/src/main/main.go


echo "Compiled!"

echo "installing"

#cp gdgbackend /etc/init.d/gdgbackend

cp $APPPATH/bin/fembackend /usr/bin
#cp $APPPATH/bin/gdgserverstarter /usr/local/bin

#chkconfig gdgbackend on


#inizializza configurazione default
#if [ ! -d /etc/gdgbackend ]; then
#    mkdir /etc/gdgbackend
#  if [ ! -f /etc/gdgbackend/backend.conf ]; then
#
#    touch /etc/gdgbackend/backend.conf
#
#	printf "address 127.0.0.1\nport 8080\n" > /etc/gdgbackend/backend.conf
#
#fi
#fi


## copia configurazione default

if [ ! -d /etc/gdgbackend ]; then
       mkdir /etc/gdgbackend
  if [ ! -f /etc/gdgbackend/backend.conf ]; then

    cp $APPPATH/var/backend.conf /etc/gdgbackend/backend.conf
fi

if [ ! -f /etc/gdgbackend/keys ]; then

    cp $APPPATH/var/keys /etc/gdgbackend/keys
fi

fi



cp $APPPATH/bin/gdgbackend /usr/bin
cp $APPPATH/bin/gdgserverstarter /usr/bin

chkconfig gdgbackend on
