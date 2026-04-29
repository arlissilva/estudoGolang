package pacoteHttp

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func HttpTimeout() {

	client := &http.Client{
		Timeout: time.Microsecond, //Define um tempo limite para a solicitação HTTP.
	}
	resp, err := client.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
