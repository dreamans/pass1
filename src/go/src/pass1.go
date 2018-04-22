package main

import (
    "fmt"
    "strings"
    "crypto/md5"
    "encoding/hex"
    "encoding/base64"
)

func main() {
    var rawPass string
    fmt.Printf("\nPlease input raw password:\n")
    if _, err := fmt.Scanf("%s", &rawPass); err != nil {
        fmt.Printf("%s\n", err)
        return
    }

    var salt string
    fmt.Printf("\nPlease input salt:\n")
    if _, err := fmt.Scanf("%s", &salt); err != nil {
        fmt.Printf("%s\n", err)
        return
    }

    s := strings.Join([]string{rawPass, salt}, "")
    md5Pwd := makeMd5(s)
    base64Pwd := makeBase64(md5Pwd)

    rawPwd := string([]byte(base64Pwd)[:10]) + "0" + "Z" + "a" + string([]byte(md5Pwd)[:3])

    fmt.Printf("\nYour password is:\n\n\t%s\n\n", rawPwd)
}

func makeMd5(s string) string {
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(s))
    return hex.EncodeToString(md5Ctx.Sum(nil))
}

func makeBase64(s string) string {
    return base64.StdEncoding.EncodeToString([]byte(s))
}

