package main
import ("fmt"
	"math/rand"
	"time")


func main() {
	
	var i int
	var values []int
	rand.Seed(time.Now().UTC().UnixNano())
	
	for z := 0; z < 20; z++ {
		num := rand.Intn(7)
		values = append(values, num)
	}
	

	inRun := false
	for := 0; i < 20; i++{
    	if inRun == true {
			if values[i] != values[i-1] {
				fmt.Print(")")
				inRun = false
			}
		}	
		if inRun == false {
			if values[i] == values[i+1]{
				fmt.Print("(")
				inRun = true
			}
		}
		fmt.Print(values[i])
		if inRun {
			fmt.Print(values[len(values-1)]
			fmt.Print(")")
		}
		else{
			fmt.Print(values[len(values-1))
		}
 	}
}
