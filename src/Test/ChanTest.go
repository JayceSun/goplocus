package main

import (

	"strings"
)

func main() {

	//ch := make(chan int)
	//
	//go func() {
	//	for i := 0;i < 10 ; i++ {
	//		ch<-i
	//	}
	//	close(ch)
	//}()
	//
	//<-ch
	//<-ch
	//<-ch
	//<-ch
	//<-ch
	//<-ch
	//
	//for value := range ch {
	//
	//	fmt.Println(value)
	//}

	//a := map[string]interface{}{"loop": [2]int{1, 950}, "Rule": "生成请求"}
	//
	//fmt.Println(a)
	//
	//loop := a["loop"].([2]int)
	//
	//fmt.Println(loop)


	//str := "{\"retcode\":20000000,\"msg\":\"\",\"data\":{\"crossdomainlist\":{\"weibo.com\":\"https:\\/\\/passport.weibo.com\\/sso\\/crossdomain?entry=mweibo&action=login&proj=1&ticket=ST-NjM1MTY5MTY0OQ%3D%3D-1503976883-tc-3B4896A9ACF4E39CF09B31C14184CB03-1\",\"sina.com.cn\":\"https:\\/\\/login.sina.com.cn\\/sso\\/crossdomain?entry=mweibo&action=login&proj=1&ticket=ST-NjM1MTY5MTY0OQ%3D%3D-1503976883-tc-CA894C8CC7F26E95DA84988D829A1C50-1\",\"weibo.cn\":\"https:\\/\\/passport.sina.cn\\/sso\\/crossdomain?entry=mweibo&action=login&ticket=ST-NjM1MTY5MTY0OQ%3D%3D-1503976883-tc-5CE442AEAA1C7B160C69B1DC078EF1CC-1"},"loginresulturl":"","uid":"6351691649"}}
	//"

	str := "昵称:财经网认证:财经网官方微博性别:男地区:北京 朝阳区生日:2009-01-01认证信息：财经网官方微博简介:商务合作： 廖先生 QQ ：3196598336 邮箱：kailiao@caijing.com.cn标签:财经网 融资 基金 更多>>"

	i := strings.LastIndex(str,"标签")
	z := strings.LastIndex(str,"更多")

	//昵称
	str2 := str[i+7:z]

	//详细信息
	str3 := str[0:i]




	if str2=="" {
		println("kkk")
	}


	print(str2)
	println()
	print(str3)








}