package helloworldC

import "fmt"
import helloworldc "github.com/luzejia/go-mod-test-c/helloworld"

func HelloWorld() {
    fmt.Println("I am b, I call on helloworld c")
    helloworldc.HelloWorld()
}
