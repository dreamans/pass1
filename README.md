## 密码混淆工具

    原始密码(任意长度字符) => 单向加密 => 混淆后密码(16字符)

## 算法

    md5Pwd = md5(md5(原始密码 + 盐))

    base64Pwd = base64(md5Pwd)

    混淆后的密码 = base64Pwd前10字符 + '0' + 'Z' + 'a' + md5Pwd前3字符

## 语言实现

* Go 
* JavaScript (+HTML)
