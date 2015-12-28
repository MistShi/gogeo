package gogeo_test

import (
	"fmt"

	"github.com/yubo/gogeo"
)

func Example() {
	geo, _ := gogeo.New("./ip_sample.txt", gogeo.GEO_F_ALIAS)
	fmt.Println(geo.Find_ip("1.0.8.5"))
	geo.Clean()
	//Output:2 21 <nil>
}
