package scanner

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Run(url string, pocs []string) {
	for _, poc := range pocs {
		if !strings.HasPrefix(url, "/") {
			url = url + "/"
		}
		target := url + poc

		req, err := http.NewRequest(http.MethodGet, target, nil)
		if err != nil {
			fmt.Println("error: ", err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error: ", err)
		}
		strBody := string(body)

		if strings.Contains(strBody, "GAOJIMicrosoft") || resp.StatusCode == 500 {
			fmt.Println(target)
			fmt.Println(strBody)
			fmt.Print("\n\n\n\n")
		}
	}
}
