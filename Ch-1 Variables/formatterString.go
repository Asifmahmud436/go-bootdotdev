package main
import "fmt"

func main(){
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hey %s! Your open rate is %.1f percent\n",name, openRate)
	fmt.Print((msg))
}