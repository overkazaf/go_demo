package main
import (
	"strconv"
	"log"
	"fmt"
	"net/http"
	"io"
	"os"
	"strings"
	"math/rand"
	"time"
)

var (
	id = 1
)

func genId() int {
	id++
	return id
}

func getNextFileName(prefix string, name string, suffix string) string {
	
	filename := strings.Join([]string {
		prefix, 
		name,
		suffix,
	}, "")
	fmt.Println("Generating filename::", name, filename)
	return filename
}


func checkAndLog(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func downloadImage(url string, name string) {
	resp, e := http.Get(url)

	defer resp.Body.Close()
	checkAndLog(e)

	fmt.Println("name", name)
	var filename = getNextFileName("images/", name, ".png")
	fd, e := os.Create(filename)
	checkAndLog(e)

	b, e := io.Copy(fd, resp.Body)
	checkAndLog(e)

	fmt.Println(url, " has been successfully downloaded, size::", b)
	defer fd.Close()
}

func multiDownload(c chan int, url string, name string) {
	fmt.Println("Start download task...")

	downloadImage(url, name)
	c <- 1

	// fmt.Println("Finish download task...", v)
}

func main() {
	t := time.Now()
	rand.Seed(t.UTC().UnixNano())
	urls := []string {
		"https://img0.bdstatic.com/static/searchresult/img/logo-2X_b99594a.png", 
		"https://img1.bdstatic.com/static/common/widget/shitu/images/camera_new_off_a552294.png", 
		"https://img2.bdstatic.com/static/common/widget/shitu/images/camera_new_on_4e3e250.png", 
		"https://img2.bdstatic.com/static/common/widget/shitu/images/mark_b68ff2e.png",
	}

	for i := 0; i < len(urls); i++ {
		// 有缓冲的channel，先喂再吐
		// 无缓冲的channel， 先吐再喂
		// 无缓冲的区别在于，生产者产出后，如果无消费者立即消费，则生产者线程会阻塞
		c := make(chan int)
		var in int = rand.Intn(32) + 1
		name := strconv.Itoa(in)
		fmt.Println("raw name", in, name)
		
		go multiDownload(c, urls[i], name)
		<-c
		time.Sleep(1 * time.Second)
	}

	fmt.Println("1 * time.Second", 1 * time.Second)

	// for i := 0; i < 10; i++ {
	// 	id := genId()
	// 	fmt.Println(id)
	// }

	fmt.Println("Mutithread image download task has been complete successfully...")
}