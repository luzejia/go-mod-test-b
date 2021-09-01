package main

// import "github.com/luzejia/go-mod-test-b/helloworld"

import helloworldC  "github.com/luzejia/go-mod-test-c/helloworld"
import helloworldD  "github.com/luzejia/go-mod-test-d/helloworld"

func main() {
	// helloworld.HelloWorldC()
	// helloworld.HelloWorldD()
        helloworldC.HelloWorld()
        helloworldD.HelloWorld()
}
