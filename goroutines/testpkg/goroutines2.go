package testpkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func ShowMsgWithTime(msg interface{}, a int) {
	for i := 0; i < a; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("%v\n", msg)
	}
}

func responseSize(url string) {
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func ForEachPrintUntil(msg string, a int) {
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Millisecond * 100)
		if i == a {
			runtime.Goexit() //退出当前goroutine
		}
		fmt.Printf("%v\n", strings.Join([]string{msg, strconv.Itoa(i)}, ""))
	}
}
