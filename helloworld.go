package main

import helloworldC  "github.com/luzejia/go-mod-test-c/helloworld"
import helloworldD  "github.com/luzejia/go-mod-test-d/helloworld"

func HelloWorld() {
	// helloworld.HelloWorldC()
	// helloworld.HelloWorldD()
        helloworldC.HelloWorld()
        helloworldD.HelloWorld()
}




