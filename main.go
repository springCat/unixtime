package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	flag.Parse()
	Args := flag.Args()
	timeStr := strings.Join(Args," ")

	if(timeStr == "") {
		now := time.Now()
		fmt.Println(now.Unix())
	}else {
		var parse time.Time
		var e error

		parse, e = time.ParseInLocation("20060102150405",timeStr, time.Local)
		if e == nil {
			fmt.Println(parse.Unix())
		}

		parse, e = time.ParseInLocation("2006-01-02 15:04:05",timeStr, time.Local)
		if e == nil {
			fmt.Println(parse.Unix())
		}

		parse, e = time.ParseInLocation("2006/01/02 15:04:05",timeStr, time.Local)
		if e == nil {
			fmt.Println(parse.Unix())
		}

		i, e := strconv.ParseInt(timeStr, 10, 32)
		if e == nil {
			parse = time.Unix(i,0)
			fmt.Println(parse.Format("2006-01-02 15:04:05"))
		}

	}
}
