package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from Home!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// #1 스트링을 바꿔준다는 거에서 번거롭, 자동화가 안되어있어 블록마다 복붙해야하는 문제 있음, 해결예정
// #2 main 함수에서 분리하는 작업 진행, block 배열을 만들고, 이전 해시를 리턴하는 함수, 블록을 만드는 함수, 블록의 정보를 출력하는 함수, 그리고 메인에선 블록을 추가하는 명령어들로 구현되어있음. 그러나 더하는 함수쪽에 총 3개지 기능 (블록 만들기, 해싱, 해싱값 slice에 추가)있어 분리예정
