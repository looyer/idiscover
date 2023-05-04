#!/bin/bash -e
# author	: looyer
# date		: 2022/12/04 23:49
# note      : cmake mingw

rm -rf ./build 
mkdir -p build
cd build
cmake ../ -G "MinGW Makefiles"

mingw32-make 

echo ""
read -s -n1 -p "press any key exit..."
