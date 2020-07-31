// これはmain パッケージの中の main 関数と決まっているので、これから書くコードは main の package だよ、と書いてあげて main 関数をこのような記法で書いてあげてます。
package main 

// モジュールのリスト
import (
	"fmt"
	"math/rand"
	"time"
)

// mainパッケージの中のmain関数が最初に実行される
func main() {
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(10) + 1
	count := 0
	for {
		// それから、他のパッケージで読み込んだ命令の最初の一文字は必ず大文字になるので、それも覚えておくと良いでしょう。
		fmt.Print("Your guess?")
		var guess int
		count++
		fmt.Scanf("%d", &guess)
		if guess == answer {
			fmt.Printf("Bingo! It took %v guesses!¥n", count)
			return
		} else if guess < answer {
			fmt.Println("The Answer is higher!")
		} else {
			fmt.Println("The Answer is lower!")
		}
	}
}