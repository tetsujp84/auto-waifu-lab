package mywaifulab

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Base64Data struct {
	girl string `json:"girl"`
}

func GetImage() {
	URL := "https://api.waifulabs.com/generate_big"

	rand.Seed(time.Now().UnixNano())

	// パラメータ生成
	values := `{"currentGirl" : [`
	param := ""
	val := strconv.Itoa(rand.Intn(999999))
	for i := 0; i < 16; i++ {
		param += val + ","
	}
	param += "0"
	values += param
	values += "]}"
	fmt.Println(values)

	req, err := http.NewRequest("POST", URL, bytes.NewBufferString(values))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	// API実行
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	// Res解析
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	if err == nil {
		// Base64デコード
		base64Data := new(Base64Data)
		if err := json.Unmarshal(b, &base64Data); err != nil {
			fmt.Println(err)
			return
		}
		data, err := base64.StdEncoding.DecodeString(base64Data.girl)
		if err != nil {
			fmt.Println(err)
		}
		file, err := os.Create("/tmp/today" + ".png")
		if err != nil {
			fmt.Println(err)
			defer file.Close()
			return
		}
		file.Write(data)
		defer file.Close()
	}
	defer resp.Body.Close()
}

// https://dev.classmethod.jp/articles/struct-json/
// https://syossan.hateblo.jp/entry/2017/06/25/005837
func (u *Base64Data) UnmarshalJSON(data []byte) (err error) {
	var value map[string]string
	json.Unmarshal(data, &value)
	u.girl = string(value["girl"])
	return nil
}
