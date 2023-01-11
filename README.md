# Study---例子：
包括文件操作，TCP，HTTP，并发编程
# Tools---工具包：
包括文件操作，时间格式转换，bcrypt加密，sm2加密，crc加密

##### 1.数据类型转换 --StrToBytes--BytesToStr--

零拷贝字符串转数组，数组转字符串
//直接使用括号强转，是直接copy数据到其他类型，如果不修改数据，仅转换类型，即可避开复制

通过unsafe.Pointer（指针转换）和uintptr（指针运算）实现转换

StrToBytes只需要构建[3]uintptr{ptr,len,len}

BytesToStr直接转换类型

##### 2.uint16转数组 --Uint16ToBytes--