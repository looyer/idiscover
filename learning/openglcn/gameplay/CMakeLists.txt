#添加cmake的版本要求
cmake_minimum_required(VERSION 3.0)

#设置工程名
project(gameplay)

#添加CMakeLists.txt相对位置目录src 到变量demo_src中
aux_source_directory(./src/ demo_src)

#设置变量demo_sss
set(demo_sss ${demo_src})

#添加CMakeLists.txt相对位置目录include到头文件设置中
include_directories(./include/)

add_library(gameplay ${demo_sss})
