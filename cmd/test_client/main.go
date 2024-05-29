package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/pircuser61/media_main_task/config"
)

func main() {
	strUrl := "http://" + config.GetAddr()
	cl := http.Client{}
	str := `{
  "amount": 400,
  "banknotes": [5000, 2000, 1000, 500, 200, 100, 50]
}}`
	fmt.Println("send request to:", strUrl)
	resp, err := cl.Post(strUrl, "text/json", bytes.NewBufferString(str))
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Println("Respnose code:", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error on read resp body:", err.Error())
	}
	strResp := string(body)
	fmt.Println("Response:", strResp)
}
