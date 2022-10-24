module github.com/j2gg0s/go-mod-example/main-new-module

go 1.17

require (
	github.com/j2gg0s/go-mod-example/new-module v1.0.2
	github.com/j2gg0s/go-mod-example/old-module v1.0.2
)

require (
	github.com/j2gg0s/go-mod-example/used-new-module v1.0.1 // indirect
	github.com/j2gg0s/go-mod-example/used-old-module v1.0.1 // indirect
)
