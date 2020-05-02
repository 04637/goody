package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func BenchmarkList(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			list()
		}
	})
}

func BenchmarkInsert(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insert()
		}
	})
}

func TestList(t *testing.T) {
	list()
}

func TestInsert(t *testing.T) {
	insert()
}

func list() {
	resp, err := http.Get("http://deva.wiki:18080/goody-v1/userInfo/select?name=hello1")

	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(body))
}

func insert() {
	resp, err := http.Get("http://deva.wiki:8099/goody-v1/userInfo/insert")

	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(body))
}

func insert2() {
	resp, err := http.Post("http://localhost:8080/user/create",
		"application/json",
		strings.NewReader("{\"name\": \"sam\", \"age\": \"13\", \"bio\": \"im a 13 man\"}"))

	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(body))
}
