package gogeo_test

import (
	"fmt"

	"github.com/yubo/gogeo"
)

func Example() {
	ips, _ := gogeo.Open_ips("./ip_sample.txt", gogeo.GEO_F_ALIAS)
	fmt.Println(gogeo.Find_ip(ips, "1.0.8.5"))
	gogeo.Clean_ips(ips)
	//Output:2 21 <nil>
}
