package tron_objects_api

import (
	"fmt"
	_ "github.com/kattana-io/tron-objects-api/pkg/api"
	_ "github.com/kattana-io/tron-objects-api/pkg/justomoney"
	_ "github.com/kattana-io/tron-objects-api/pkg/sunswap"
	_ "github.com/kattana-io/tron-objects-api/pkg/trc20"
	_ "github.com/kattana-io/tron-objects-api/pkg/url"
)

func main() {
	fmt.Println("Success")
}
