package helloworld

import "fmt"
import helloworldc "github.com/luzejia/go-mod-test-c/helloworld"
import helloworldd "github.com/luzejia/go-mod-test-d/helloworld"

func HelloWorldC() {
    fmt.Println("I am b, I call on helloworld c")
    helloworldc.HelloWorld()
}

func HelloWorldD() {
    fmt.Println("I am b, I call on helloworld d")
    helloworldd.HelloWorld()
}
