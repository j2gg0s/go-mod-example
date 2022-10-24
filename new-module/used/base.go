package used

import (
	"fmt"

	module "github.com/j2gg0s/go-mod-example/used-new-module/used"
)

func Say() {
	fmt.Println("new-module/used")
	module.Say()
}
