#!/usr/bin/env bash

URL=https://github.com/infoforcefeed/OlegDB
mkdir libs
echo "Download from ${URL}"
cd libs
git clone $URL
cd OlegDB
make liboleg
export CGO_LDFLAGS="-L/${PWD}/libs/OlegDB"
export LD_LIBRARY_PATH="$LD_LIBRARY_PATH:${PWD}/libs/OlegDB"
echo $LD_LIBRARY_PATH