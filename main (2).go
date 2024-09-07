package main
import ("fmt"
	"math/rand"
	"time")

func sieve(n int) {

	var i int
	var prime []int
	
	for (i = 0; i <= n; i++){
		prime[i] = true
	}
	for (k := 2; k*k <= n; k++){
		if prime[i] == true {
			for (p:= k*k; p <= n; p += k){
				prime[i] = false	
			}
		}
	}
	for (i = 2; i <= n; i++){
		if (prime[i] == true){
			fmt.Print(i)
		}
	}
	
	
}


func main() {
	var num int
	fmt.Println("Enter number: ")
	fmt.Scanf("%d",&num)
	sieve(num)

}
