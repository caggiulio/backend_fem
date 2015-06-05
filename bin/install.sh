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

chkconfig fembackend on
