cmake_minimum_required(VERSION 3.5)
project(Algorithm)

set(CMAKE_CXX_STANDARD 11)

include_directories(common)

file(GLOB COMMONFILES
        common/*.h
        common/*.cpp
        common/*.c
        common/*.cc)

file(GLOB FILES
        ./*.h
        ./*.cc
        ./*.cpp
        ./*.c
        common/*.h
        common/*.cpp
        common/*.c
        common/*.cc)

add_executable(Def ${COMMONFILES} Main.cpp)

add_executable(Algorithm2 ${FILES} Main.cpp)
