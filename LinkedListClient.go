package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func StartHttpClientCLI() {
	/* Code handler for HTTP */
	fmt.Println("\nThis is a HTTP Client CLI!")
	for true {
		var URLString, MagicNum, MagicNumDesc, method string
		var Method, BodyNeeded int
		var body io.Reader

		fmt.Println("\nTo send a HTTP Request, please enter following details:")
		fmt.Print("Enter URL: ")
		fmt.Scanf("%s", &URLString)
		fmt.Print("Choose Method(1-GET, 2-PUT, 3-DELETE, 4-POST): ")
		for true {
			fmt.Scanf("%d", &Method)
			if Method <= 4 && Method >= 1 {
				switch Method {
				case 1:
					method = http.MethodGet
				case 2:
					method = http.MethodPut
				case 3:
					method = http.MethodDelete
				case 4:
					method = http.MethodPost
				default:

				}
				break
			}
			fmt.Println("Incorrect Input. Please Choose Method again")
			fmt.Print("Choose Method(1-GET, 2-PUT, 3-DELETE, 4-POST): ")
		}

		if method == http.MethodPut || method == http.MethodDelete || method == http.MethodPost {
			fmt.Print("Body to be set? (1-YES, 2-NO): ")
			fmt.Scanf("%d", &BodyNeeded)
			if BodyNeeded == 1 {
				fmt.Print("Enter data to be set in the body.\n")
				fmt.Print("Enter Magic Number: ")
				fmt.Scanf("%s", &MagicNum)
				fmt.Print("Enter Magic Number Description: ")
				fmt.Scanf("%s", &MagicNumDesc)
				magicdata := MagicNumber{MagicNum, MagicNumDesc}
				jsondata, _ := json.Marshal(magicdata)
				body = bytes.NewBuffer(jsondata)
			}
		}
		SendHttpRequest(method, URLString, body)
	}
	return
}

func SendHttpRequest(method string, url string, body io.Reader) error {
	switch method {
	case http.MethodGet:
		resp, err := http.Get(url)
		if err != nil {
			return err
		} else {
			fmt.Println("HTTP GET Request Sent")
		}
		fmt.Printf("\nHTTP Response received Status code: %d", resp.StatusCode)
		if resp.StatusCode == 200 {
			body, err1 := ioutil.ReadAll(resp.Body)
			if err1 != nil {
				return err1
			}
			defer resp.Body.Close()
			fmt.Printf("\nHTTP Response Body:\n%s", body)
		}
	case http.MethodPut:
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPut, url, body)
		if err != nil {
			return err
		}
		if body != nil {
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
		}
		resp, err1 := client.Do(req)
		if err1 != nil {
			return err1
		} else {
			fmt.Println("HTTP PUT Request Sent")
		}
		fmt.Printf("\nHTTP Response received Status code: %d\n", resp.StatusCode)
	case http.MethodDelete:
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodDelete, url, body)
		if err != nil {
			return err
		}
		if body != nil {
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
		}
		resp, err1 := client.Do(req)
		if err1 != nil {
			return err1
		} else {
			fmt.Println("HTTP DELETE Request Sent")
		}
		fmt.Printf("\nHTTP Response received Status code: %d\n", resp.StatusCode)
	case http.MethodPost:
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPost, url, body)
		if err != nil {
			return err
		}
		if body != nil {
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
		}
		resp, err1 := client.Do(req)
		if err1 != nil {
			return err1
		} else {
			fmt.Println("HTTP POST Request Sent")
		}
		fmt.Printf("\nHTTP Response received Status code: %d\n", resp.StatusCode)
	default:
		fmt.Println("Incorrect Method")
	}
	return nil
}
