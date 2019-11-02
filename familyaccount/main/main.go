package main
import(
	"go_code/familyaccount/utils"
	"fmt"
)

func main() {
	fmt.Println("这是面向对象的方式完成的~~")
	utils.NewFamilyAccount().MainMenu()
}