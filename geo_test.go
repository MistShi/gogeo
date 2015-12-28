package gogeo

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	{
		geo, _ := New("./ip_sample.txt", 0)
		fmt.Println(geo.Find_ip("1.0.167.5"))
		fmt.Println(geo.Find_ip("1.0.165.5"))
		fmt.Println(geo.Find_ip("1.0.8.5"))
		geo.Clean()
	}

	fmt.Println("--------alias begin--------")
	{
		geo, _ := New("./ip_sample.txt", GEO_F_ALIAS)
		fmt.Println(geo.Find_ip("1.0.167.5"))
		fmt.Println(geo.Find_ip("1.0.165.5"))
		fmt.Println(geo.Find_ip("1.0.8.5"))
		geo.Clean()
	}
}
