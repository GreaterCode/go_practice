/**
@author:admin
@date:2023/2/17
@note
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var logFile string = "./test.log"
func main() {
	SetupLogger()
	HttpGet("https://www.baidu.com")
	HttpGet("https://cn.bing.com")

}

func HttpGet(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s: %v", url, err.Error())
	}else {
		log.Printf("Fetching url %s,  status: %d", url, res.StatusCode)
	}
	defer res.Body.Close()
}

func SetupLogger() {

	logFileLocation, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0744)
	if err != nil {
		fmt.Errorf("Open log file %s: %v\n", logFile, err)
	}
	log.SetOutput(logFileLocation)

	/**
	* log.SetFlags()仅支持Ldate, Ltime, and so on. log.Llongfile
	*/
	log.SetFlags(log.Lshortfile)
 }
