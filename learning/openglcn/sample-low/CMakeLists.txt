#添加cmake的版本要求
cmake_minimum_required(VERSION 3.0)

#设置工程名
project(sample-low)

# set(CMAKE_C_FLAGS "-m32")
# set(CMAKE_CXX_FLAGS "-m32")

#添加CMakeLists.txt相对位置目录include到头文件设置中
include_directories(../freeglut/include/GL ../glew-2.1.0/include/GL)
# include_directories(../freeglut/include/GL)
# include_directories(../glew-2.1.0/include/GL)

#链接目录
link_directories(../freeglut/lib/x64 ../glew-2.1.0/lib/)
# link_directories(../freeglut/lib/x64)
# link_directories(../glew-2.1.0/lib/)

#递归查找所有含main函数的cpp文件，保存到mains列表中
file(GLOB_RECURSE mains RELATIVE "${CMAKE_CURRENT_SOURCE_DIR}" "${CMAKE_CURRENT_SOURCE_DIR}/sample-*.cpp")

#循环遍历每个含main函数的文件，分别编译生成exe
foreach(mainfile IN LISTS mains)
    #获取无文件后缀的文件名放在mainname中
    get_filename_component(mainname ${mainfile} NAME_WE)
    add_executable(${mainname} ${mainfile})
    #链接库文件
    target_link_libraries(${mainname} opengl32 freeglut glew32.dll)
endforeach()
