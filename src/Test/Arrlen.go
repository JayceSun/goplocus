package main

import (


	"time"
	"strconv"
	"math/big"
	"crypto/rand"
	//"database/sql"
	//"fmt"
	"fmt"
	"github.com/henrylee2cn/pholcus/common/mahonia"
	"encoding/json"
	"strings"
)

func main()  {
	//var a [3][5] int

	//var a [32] byte

	//var a [10] int

	//a := [] int {1,2,3,4,5,6,7}
	//b :=[] int {8,9,10}
	//copy(a,b)
	//
	//fmt.Println(a)

	//a:="31126304,30885904"
	//a =strings.Split(a,",")[0]
	//newa :=a[:16]
	//fmt.Println(len(a))

	//for i:=0;i<len(a) ;i++  {
	//	fmt.Println(a[i])
	//}
	//for _, v := range a {
	//	fmt.Println(v)
	//}
	//arrlength :=len(a)
	//fmt.Println(arrlength)

	//a := "/f/index/forumpark?cn=%E6%B8%AF%E5%8F%B0%E4%B8%9C%E5%8D%97%E4%BA%9A%E6%98%8E%E6%98%9F&ci=0&pcn=%E5%A8%B1%E4%B9%90%E6%98%8E%E6%98%9F&pci=0&ct=1&st=new&pn=3"
	//b  :=strings.Split(a,"pn=")
	//for i:=0 ; i<30 ; i++ {
	//	fmt.Println(b[0]+"pn="+strconv.Itoa(i))
	//}

	//newstime :="09月14日12:05:46 abc"//新闻时间
	//splits1 :=strings.Split(newstime,"月")
	//splits2 :=strings.Split(splits1[1],"日")
	//newsTime := strconv.Itoa(2017)+"-"+splits1[0]+"-"+splits2[0]+" "+splits2[1]
	//fmt.Println(string(newstime[0:18]))

	//month := time.Now().Unix()
	//a := month.Format("2006-01-02")
	//aa := strings.Split(a,"-")
	////day := time.Now().Day()
	//fmt.Println(month)

	//t := time.Now().Unix()
	//str :=strconv.FormatInt(t, 10)
	//fmt.Println(str)

	//h := md5.New()
	//h.Write([]byte("1507862667")) // 需要加密的字符串为 sharejs.com
	//aa := hex.EncodeToString(h.Sum(nil))
	//fmt.Println(aa)
	//newaa :=aa[:5]
	//fmt.Printf("%s\n", newaa) // 输出加密结果

	//for i:=0;i<10;i++{
	//	time.Sleep( 10 * time.Second)
	//	fmt.Println(i)
	//}

	//a := "10月19日 11:00:00"
	//reg := regexp.MustCompile("月|日")
	//b:= reg.ReplaceAllString(a,"-")
	//fmt.Println(b)

	//sText := "中文"
 	//textQuoted := strconv.QuoteToASCII(sText)
 	//textUnquoted := textQuoted[1 : len(textQuoted)-1]
	//fmt.Println(textUnquoted)

	//textUnquoted := "\u675c\u660e\u793c\u548c\u67e5\u5764\u6700\u540e\u5173\u5934\u7684\u751f\u6b7b\u4e0d\u79bb\uff0c\u8ba9\u53cd\u6d3e\u5145\u6ee1\u4e86\u4eba\u60c5\u5473\",\"abstract\":\"\u660e\u793c\u8ba9\u67e5\u5764\u5148\u8dd1\uff0c\u4e34\u8d70\u524d\u4ed6\u56de\u5934\u90a3\u4e00\u773c\uff0c\u5176\u5b9e\u662f\u5728\u4e0e\u67e5\u5764\u8bc0\u522b\u3002\u53ef\u662f\u67e5\u5764\u6ca1\u6709\u8d70\uff0c\u4ed6\u4e00\u76f4\u4ee5\u4e3a\u4ed6\u6548\u5fe0\u7684\u662f\u738b\u7237\uff0c\u5176\u5b9e\u4ed6\u662f\u5fe0\u5fc3\u4e8e\u660e\u793c\uff0c\u4ed6\u4fe9\u4ece\u5c0f\u4e00\u8d77\u76f8\u4e92\u6276\u6301\u957f\u5927\uff0c\u4e00\u8d77\u5403\u8fc7\u82e6\uff0c\u4e00\u8d77\u53d7\u8fc7\u7f6a\u3002\u4e5f\u8bb8\u5728\u660e\u793c\u4e3a\u4ed6\u625b\u4e0b\u7b2c\u4e00\u987f\u677f\u5b50\u7684\u65f6\u5019\uff0c\u4ed6\u5c31\u4e0d\u518d\u53ea\u662f\u4ed6\u7684\u8001\u677f\uff0c\u66f4\u662f\u4ed6\u8981\u6c38\u8fdc\u5b88\u62a4\u7684\u4eba\u3002\u4ed6\u5bf9\u6240\u6709\u4eba\u90fd\u72e0\uff0c\u552f\u72ec\u5bf9\u660e\u793c\u4e00\u7247\u75f4\u5fc3\u3002\u6240\u4ee5\u4ed6\u8df3\u51fa\u6765\uff0c\u4ee5\u4e00\u654c\u56db\uff0c\u4e5f\u8981\u6551\u4e0b\u660e\u793c\u3002\u660e\u793c\u5462\uff1f\u4ed6\u6015\u662f\u76f4\u5230\u8fd9\u4e00\u523b\uff0c\u624d\u660e\u767d\u539f\u6765\u67e5\u5764\u624d\u662f\u4ed6\u6700\u5728\u4e4e\u7684\u4eba\u5427\u3002\u4ed6\u8ff7\u604b\u548f\u6885\uff0c\u4f46\u5728\u94f6\u5b50\u4e0e\u548f\u6885\u4e4b\u95f4\uff0c\u4ed6\u9009\u62e9\u4e86\u94f6...\",\"content\":\"<p>\u660e\u793c\u8ba9\u67e5\u5764\u5148\u8dd1\uff0c\u4e34\u8d70\u524d\u4ed6\u56de\u5934\u90a3\u4e00\u773c\uff0c\u5176\u5b9e\u662f\u5728\u4e0e\u67e5\u5764\u8bc0\u522b\u3002\u53ef\u662f\u67e5\u5764\u6ca1\u6709\u8d70\uff0c\u4ed6\u4e00\u76f4\u4ee5\u4e3a\u4ed6\u6548\u5fe0\u7684\u662f\u738b\u7237\uff0c\u5176\u5b9e\u4ed6\u662f\u5fe0\u5fc3\u4e8e\u660e\u793c\uff0c\u4ed6\u4fe9\u4ece\u5c0f\u4e00\u8d77\u76f8\u4e92\u6276\u6301\u957f\u5927\uff0c\u4e00\u8d77\u5403\u8fc7\u82e6\uff0c\u4e00\u8d77\u53d7\u8fc7\u7f6a\u3002\u4e5f\u8bb8\u5728\u660e\u793c\u4e3a\u4ed6\u625b\u4e0b\u7b2c\u4e00\u987f\u677f\u5b50\u7684\u65f6\u5019\uff0c\u4ed6\u5c31\u4e0d\u518d\u53ea\u662f\u4ed6\u7684\u8001\u677f\uff0c\u66f4\u662f\u4ed6\u8981\u6c38\u8fdc\u5b88\u62a4\u7684\u4eba\u3002\u4ed6\u5bf9\u6240\u6709\u4eba\u90fd\u72e0\uff0c\u552f\u72ec\u5bf9\u660e\u793c\u4e00\u7247\u75f4\u5fc3\u3002\u6240\u4ee5\u4ed6\u8df3\u51fa\u6765\uff0c\u4ee5\u4e00\u654c\u56db\uff0c\u4e5f\u8981\u6551\u4e0b\u660e\u793c\u3002"
	//sUnicodev := strings.Split(textUnquoted, "\\u")
	//var context string
	//for _, v := range sUnicodev {
	//if len(v) < 1 {
	//          continue
	//    }
	//temp, err := strconv.ParseInt(v, 16, 32)
	// if err != nil {
	// panic(err)
	// }
	//	context += fmt.Sprintf("%c", temp)
	//}
	//fmt.Println(context)
	//
	//a :=(214/30)+1
	//print(a)

	//var urlStr  = "中国平安"
	//src:="中国平安"
	//
	//enc := mahonia.NewEncoder("GBK")
	//
	//output := enc.ConvertString(src)
	//
	//a:= url.QueryEscape (output)

	//b :=url.QueryEscape(a)
	//
	//c:= url.EscapedPath()
	//fmt.Println(a)
	//
	//
	////c ,err4 := url.Parse(a)
	//d ,_:=url.ParseQuery(urlStr)
	//
	//fmt.Println(d)
	//fmt.Println(err3)
	//fmt.Println(err4)
	//fmt.Println(l3.Query().Encode()+"--")

	//
	//t := time.Now().AddDate(0,0,-(1))
	//t.Clock()
	//const timeLayout = "2006/01/02"
	//today := t.Format(timeLayout)
	//println(today)
	//if strings.EqualFold(today, "2018/01/27") {
	//	println("1111111")
	//}

	//t, _ := time.Parse("2006-01-02 15:04:05", "2017-11-10 13:5:00")
	//fmt.Println(today)

	//t :=time.Now()
	//hour :=t.Hour()
	//min :=t.Minute()-30
	//sec := t.Second()
	//const timeLayout = "2006-01-02 15"
	//t1 :=t.Format(timeLayout)
	//fmt.Println( t1+":"+strconv.Itoa(min)+":"+strconv.Itoa(sec))

	//timeNow := time.Now()
	//t :=timeNow.Unix()
	//
	//a :=30*60
	////时间戳int64转int
	//string:=strconv.FormatInt(t,10)
	//b ,_:=strconv.Atoi(string)
	//
	//string =strconv.Itoa(b-a)
	//c, _ := strconv.ParseInt(string, 10, 64)
	//resultdate :=time.Unix(c,0).Format("2006-01-02 15:04:05")
	//fmt.Println(resultdate)
	//
	//var a = "科比                詹姆斯    教育部"
	//
	//reg,_ := regexp.Compile("\\s+")
	//b:= reg.ReplaceAllString(a," ")
	//println(b)

	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i:=0; i<10; i++ {
	//	fmt.Println(r.Intn(10))
	//}

	//i := 1
	//for i:=0;i<=50 ;i++  {
	//	ranNum()
	//}

	//for e := b.Front(); e != nil; e = e.Next() {
	//	fmt.Printf("元素%d:%d\n", i, e.Value)
	//	i++
	//	if i==10 {
	//		//fmt.Println(e.Value)
	//		//break
	//	}
	//}

	//for i:=0; i<=50;i++  {
	//	fmt.Println(RandInt64(0,10))
	//	bb :=RandInt64(0,10)
	//	string := strconv.FormatInt(bb,10)
	//	b,_:=strconv.Atoi(string)
	//
	//	fmt.Println(b)
	//}

	//mysqlClient()
	//bianma()
	//
	//cur := time.Now()
	//timestamp := cur.UnixNano()
	//println(timestamp/1000000)

	//jsonstring()


	var a = "01-01"
	b:=string([]rune(a)[:5])
	//c:=strings.Index(a, "瓶")
	println(b)

}

func min(minnum int) string {
	timeNow := time.Now()
	tu :=timeNow.Unix()

	am :=minnum*60
	//时间戳int64转int
	sf:=strconv.FormatInt(tu,10)
	ba ,_:=strconv.Atoi(sf)

	si:=strconv.Itoa(ba-am)
	sp, _ := strconv.ParseInt(si, 10, 64)
	resultdate :=time.Unix(sp,0).Format("2006-01-02 15:04:05")

	return resultdate
}
//func ranNum() list.List{
//
//	a := list.List{}
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	//for i:=0; i<50; i++ {
//		a.PushBack(r.Intn(10))
//		b :=r.Intn(10)
//		println(b)
//	//}
//	return a
//}

func RandInt64(min,max int64) int64{
	maxBigInt:=big.NewInt(max)
	i,_:=rand.Int(rand.Reader,maxBigInt)
	if i.Int64()<min{
		RandInt64(min,max)
	}
	return i.Int64()
}
//
//func mysqlClient()  {
//	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/pholcus?charset=utf8")
//	checkErr(err)
//
//	rows, err := db.Query("SELECT * FROM 新浪微博用户____用户详情")
//	checkErr(err)
//
//	for rows.Next() {
//		var id int
//		//var userName string
//		//var userAge int
//		//var userSex int
//		rows.Columns()
//		err = rows.Scan(&id)
//		checkErr(err)
//		fmt.Println(id)
//	}
//}

//func checkErr(err error) {    if err != nil {        panic(err)    } }

func bianma(){

	//"你好，世界！"的GBK编码

	utfStr := "中国平安"

	var enc mahonia.Encoder

	enc = mahonia.NewEncoder("gbk")

	if ret, ok := enc.ConvertStringOK(utfStr); ok {

		fmt.Println(ret)

	}

}

func jsonstring()  {
	a  := `{"from":"en","to":"zh","trans_result":{"src":"today","dst":"\u4eca\u5929"},"result":["src","today","dst","\u4eca\u5929"]}`
	b, _ := json.Marshal(a)
	println(b)


}

/**
保留两位小数转百分数，10%
 */
func float2per(f float64) string {

	str := fmt.Sprintf("%0.2f",f)

	sp := strings.Split(str,".")

	var int1 = 0

	if(sp[0]=="0"){
		int1 ,_=strconv.Atoi(sp[0])
		int1 = int1*100
	}
	int2 ,_:= strconv.Atoi(sp[1])

	per := strconv.Itoa(int1+int2)+"%"

	return  per
}


