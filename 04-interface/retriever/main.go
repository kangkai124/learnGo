package main

import (
	"fmt"
	"learnGo/04-interface/retriever/mock"
	"learnGo/04-interface/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com",
		map[string]string{
			"name":   "kk",
			"course": "golang",
		})
}

const url = "http://www.baidu.com"

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked baidu",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"hello mock contents"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Chrome/63",
		Timeout:   time.Second,
	}
	inspect(r)

	// type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.Timeout)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("try a session")
	fmt.Println(session(&retriever))

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type Switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
