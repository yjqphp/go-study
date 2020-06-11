package main

import "fmt"

/**
	for是Go中的 “while”
	此时你是可以去掉分号，因为C的while在Go中叫做for
 */
func main(){
	sum:=1
	for sum<1000 {
		sum += sum
	}
	fmt.Println(sum)
}
