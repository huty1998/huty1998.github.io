---
layout: post
title: "Unicode(UTF-8) & ASCII"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - CS
---


### **计算机怎么知道三个字节表示一个字符，而不是分别表示三个字符呢？**  

**ASCII码用一个字节表示字符**，ASCII 码一共规定了128个字符的编码，大写的字母`A`
是65（二进制`01000001`）。这128个符号只占用了一个字节的后面7位，最前面的一位统一规定为`0`。

**UTF-8 是 Unicode 的一种变长编码方式，使用1~4个字节表示一个字符**，UTF-16用两个字节或四个字节表示一个字符； UTF-32用四个字节表示一个字符  
如果一个字节的第一位是0，则这个字节单独就是一个字符；如果第一位是1，则连续有多少个1，就表示当前字符占用多少个字节。Go中的字符串是UTF-8编码的。


  
计算机在存储器中排列字节有两种方式：大端法和小端法
大端法就是将第一个字节放在低地址处（前面），如0x1234
