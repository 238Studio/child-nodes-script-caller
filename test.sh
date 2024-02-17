#!/bin/bash

# 接受两个字符串作为参数
str1=$1
str2=$2

# 检查是否提供了两个参数
if [ -z "$str1" ] || [ -z "$str2" ]; then
  echo "请提供两个字符串作为参数"
  exit 1
fi

# 拼接字符串
result="$str1 $str2"

# 输出结果到终端
echo "拼接结果为: $result"
echo "拼接结果为: $result"

# 将结果写入文本文件test.txt
echo "$result" > test.txt
