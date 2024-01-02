package	main

import(
	"fmt"
	"net/http"
	"sync"
)

func makeRequests(url string,wg *sync.WaitGroup,ch chan int){
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		// fmt.Println("Some Error Occurred : ",err)
		ch <- -1
		return
	}
	ch <- response.StatusCode
}


func main(){
	var url string
	var numberOfRequests int
	fmt.Println("Add url = ")
	fmt.Scanln(&url)
	
	fmt.Println("Add number of Requests = ")
	fmt.Scanln(&numberOfRequests)

	var wg sync.WaitGroup 
	ch := make(chan int, numberOfRequests)
	
	for i :=0; i< numberOfRequests; i++{
		wg.Add(1)
		go makeRequests(url, &wg, ch)
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	var successCnt,failuresCnt int

	for StatusCode := range ch{
		if StatusCode == 200{
			successCnt+=1
		}else {
			failuresCnt +=1
		}
	}
	fmt.Println("Results \n")
	fmt.Println("Total Requests = ",successCnt+failuresCnt)
	fmt.Println("Successful Requests = ",successCnt)
	fmt.Println("Failed Requests = ",failuresCnt)
	
}