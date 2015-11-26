package gogeo

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	{
		ips, _ := Open_ips("./ip_sample.txt", 0)
		fmt.Println(Find_ip(ips, "1.0.167.5"))
		fmt.Println(Find_ip(ips, "1.0.165.5"))
		fmt.Println(Find_ip(ips, "1.0.8.5"))
		Clean_ips(ips)
	}

	fmt.Println("--------alias begin--------")
	{
		ips, _ := Open_ips("./ip_sample.txt", GEO_F_ALIAS)
		fmt.Println(Find_ip(ips, "1.0.167.5"))
		fmt.Println(Find_ip(ips, "1.0.165.5"))
		fmt.Println(Find_ip(ips, "1.0.8.5"))
		Clean_ips(ips)
	}
}
