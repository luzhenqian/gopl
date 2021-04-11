# 名称

## 命名规则

以字母或下划线开头，后面可以有任意数量的字符、数字和下划线。
区分大小写：name 和 Name 是不同的。

## 关键字

共 25 个
break default func interface select case defer go map struct chan else goto package switch const fallthrough if range type continue for import return var

内置的预声明常量、类型和函数：

## 常量

true false iota nil

## 类型

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex128 complex64
bool byte rune string error

## 函数

make len cap new append copy close delete complex real imag panic recover

## 作用域

函数中声明的实体只在函数的局部范围有效。

函数外声明的实体在整个包内可见。

首字母大写的实体可以跨包。

## 命名风格

Go 语言推荐短命名，比如用 i 表示 index。

由单词组成的名称，使用驼峰式。

像 HTML 这类首字母缩写的字母使用相同的大小写，比如 htmlEscape、HTMLEscape 或者 escapeHTML，不会是 escapeHtml。
