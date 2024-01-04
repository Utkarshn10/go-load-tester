package	main

import(
	"fmt"
	"net/http"
	"sync"
)


func makeRequests(url string,requestMethod string,wg *sync.WaitGroup,ch chan int){
	defer wg.Done()

	switch{
	case requestMethod == '1':
		response, err := http.Get(url)
	case requestMethod == '2':
		response, err := http.Post(url)
	case requestMethod == '3':
		response, err := http.Put(url)
	case requestMethod == '4':
		response, err := http.Patch(url)

	}
	if err != nil {
		// fmt.Println("Some Error Occurred : ",err)
		ch <- -1
		return
	}
	ch <- response.StatusCode
}


func main(){
	var url,requestMethod string
	var numberOfRequests int
	fmt.Print("Enter url = ")
	fmt.Scanln(&url)
	
	fmt.Print("Enter number of Requests = ")
	fmt.Scanln(&numberOfRequests)

	fmt.Println("Enter Request method")
	fmt.Println("1. GET ")
	fmt.Println("2. POST ")
	fmt.Println("3. PUT ")
	fmt.Println("4. PATCH ")
	fmt.Scanln(&requestMethod)

	var wg sync.WaitGroup 
	ch := make(chan int, numberOfRequests)
	
	for i :=0; i< numberOfRequests; i++{
		wg.Add(1)
		go makeRequests(url,requestMethod, &wg, ch)
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
	fmt.Println("\n Results \n")
	fmt.Println("Total Requests = ",successCnt+failuresCnt)
	fmt.Println("Successful Requests = ",successCnt)
	fmt.Println("Failed Requests = ",failuresCnt)
	
}
