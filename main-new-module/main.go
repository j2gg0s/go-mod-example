package main

import (
	newModule "github.com/j2gg0s/go-mod-example/new-module/used"
	oldModule "github.com/j2gg0s/go-mod-example/old-module/used"
)

func main() {
	newModule.Say()
	oldModule.Say()
}
