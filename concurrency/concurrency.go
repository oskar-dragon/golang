package concurrency

type WebsiteChecker func(string)bool
type Result struct {
	string
	bool
}


func CheckWebsites(wc WebsiteChecker, urls []string)map[string]bool {
	result := make(map[string]bool)
	resultsChannel := make(chan Result)

	for _,url := range urls {
		go func(u string) {
			resultsChannel <- Result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <- resultsChannel
		result[r.string] = r.bool
	}

	return result
}