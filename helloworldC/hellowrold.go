package helloworldC

import "fmt"
import "github.com/luzejia/go-mod-test-c/helloworld"

func HelloWorld() {
    fmt.Println("I am b, I call on helloworld c")
    helloworld.HelloWorld()
}
