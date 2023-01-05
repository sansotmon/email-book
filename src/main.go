package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Mail struct {
	To   string `json:"To"`
	From string `json:"From"`
	Date string `json:"Date"`
}

type Connection struct {
	Index string  `json:"Index"`
	Mails []*Mail `json:"records"`
}

func main() {
	var mails []*Mail

	//path := "/home/santiago/Downloads/enron_mail_20110402"
	path := "/home/santiago/Downloads/fedex"
	err := filepath.WalkDir(path, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			file, err := os.Open(s)
			if err != nil {
				log.Fatalf("failed to open")
			}

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)

			var text []string
			for scanner.Scan() {
				text = append(text, scanner.Text())
			}

			file.Close()

			for _, each_ln := range text {
				mails = append(mails, &Mail{To: each_ln, From: each_ln, Date: each_ln})
			}
		}
		return nil
	})
	if err != nil {
		return
	}

	p, _ := json.Marshal(Connection{Index: "soto", Mails: mails})

	url := "http://localhost:4080/api/_bulkv2"
	fmt.Println("URL:>", url)

	credentials := "admin:Complexpass#123"
	credentials_base64 := base64.StdEncoding.EncodeToString([]byte(credentials))
	all_word := "Basic " + credentials_base64
	fmt.Println(all_word)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(p))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", all_word)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
