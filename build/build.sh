#!/usr/bin/env bash

current=`pwd`

rm -rf ../production
mkdir -p ../production

# copy default Corefile to production root
cd ${current}
\cp ./Corefile ../production/

# build production
git clone https://github.com/coredns/coredns.git
cd coredns
\cp ././plugin.cfg ./

go generate
make