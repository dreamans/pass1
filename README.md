## 密码混淆工具

    原始密码 + 盐 = 混淆后密码

## 算法

    md5Pwd = md5(原始密码 + 盐)

    base64Pwd = base64(md5Pwd)

    混淆后的密码 = base64Pwd前10字符 + '0' + 'Z' + 'a' + md5Pwd前3字符
