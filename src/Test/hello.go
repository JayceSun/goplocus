package main

import "fmt"
import (
	"math"
	"strings"
	"net/http"
	"log"

)


func sayhelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()  //解析参数，默认是不会解析的
	//fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println(r.Form["name"])
	fmt.Fprintf(w, "Hello astaxie!")
	fmt.Fprintf(w,  strings.Join(r.Form["name"], ""))

}


	var a byte = 65

const cd  = 1



func main()  {

	http.HandleFunc("/login", sayhelloName) //设置访问的路由
	  err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


    fmt.Println(string(a))

   // b := 5

    //println(a,b)

    var i int
    var zhizheng * int
    i = 100
    zhizheng = &i
    println(*zhizheng)

    var str string = "!!!"
    var ii  = len(str)
    println(ii)
    var pi = math.Pi
    println(pi)
	b := int(pi)
	println(b)

	for j := 0;j < ii ;j++  {
		println(str[j])
	}

	var sss string
	//fmt.Scan(&sss)
	println(sss)


	println(strings.Contains("hello","lo"))
	println(strings.Count("abcdeafag","a"))

	//t := time.Now()
	//println(t.Format("2010 年 6月 10日"))
	//println(t.Weekday().String())

	//var hang,lie int
	//fmt.Scan(&hang ,&lie)
	//println(hang,lie)

	for k := 1 ; k <= 9 ; k++ {
		for j := 1;j<=k ; j++ {
			fmt.Printf("%dx%d=%d  ",k,j,k*j)

		}
		fmt.Printf("\n")
	}


}