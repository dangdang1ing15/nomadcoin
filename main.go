package main

import (
	"github.com/dangdang1ing15/nomadcoin/rest"
)

func main() {
	rest.Start(4000)
}

// #1 스트링을 바꿔준다는 거에서 번거롭, 자동화가 안되어있어 블록마다 복붙해야하는 문제 있음, 해결예정
// #2 main 함수에서 분리하는 작업 진행, block 배열을 만들고, 이전 해시를 리턴하는 함수, 블록을 만드는 함수, 블록의 정보를 출력하는 함수, 그리고 메인에선 블록을 추가하는 명령어들로 구현되어있음. 그러나 더하는 함수쪽에 총 3개지 기능 (블록 만들기, 해싱, 해싱값 slice에 추가)있어 분리예정

// #3 서버사이드 랜더링 시작, Explorer / ResponseWriter에 간단한 text를 써서 user에게 보내는 방법 해봄
// #4 html 랜더링 시작/ 템플릿 활용
// #5 분활, 리액트에서 컴포넌트처럼(?)
// #6 REST API 만들기
// #7 url을 완전한 url로 바꾸고 싶음, struct가 encode될 때 url을 추가하고 싶음
// #8 멀티플렉서를 직접 만들어서 핸들링, GorillaMux 설치
