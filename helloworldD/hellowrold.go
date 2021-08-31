package helloworldD

import "fmt"
import helloworldd "github.com/luzejia/go-mod-test-d/helloworld"

func HelloWorld() {
    fmt.Println("I am b, I call on helloworld d")
    helloworldd.HelloWorld()
}
