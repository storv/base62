package main

import (
	"flag"
	"fmt"
	"os"
)
import base62 "zidanduanxin.com/common/base62/core"

func main() {
	var  message string
	flag.StringVar(&message,"message","","待加密或者解密字符串")
	var alphabet string
	flag.StringVar(&alphabet,"alphabet","","可选字符集,字符长度62【必须】")
	var optionType string
	flag.StringVar(&optionType,"optionType","","可选类别 encode 或者 decode")
	flag.Parse()
	base62 := base62.GetBase62Instance("")
	switch {
	case optionType == "encode":
		fmt.Printf("Encode - Before:%s After:%s \n",message,base62.Encode(message))
	case optionType == "decode":
		fmt.Printf("Decode - Before:%s After:%s \n",message,base62.Decode(message))
	default:
		fmt.Printf("无效参数 optionType:%s \n",optionType)
		os.Exit(2)
	}

}