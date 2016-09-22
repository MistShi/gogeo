package gogeo

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {

	fmt.Println("--------find begin--------")
	{
		geo, _ := New("./ip_sample.txt", GEO_N_ALIAS)
		fmt.Println(geo.Find_ip("1.0.167.5"))
		fmt.Println(geo.Find_ip("1.0.165.5"))
		fmt.Println(geo.Find_ip("1.0.8.5"))
		geo.Clean()
	}

	fmt.Println("--------find alias begin--------")
	{
		geo, _ := New("./ip_sample.txt", GEO_F_ALIAS)
		fmt.Println(geo.Find_ip("1.0.167.5"))
		fmt.Println(geo.Find_ip("1.0.165.5"))
		fmt.Println(geo.Find_ip("1.0.8.5"))
		geo.Clean()
	}

	fmt.Println("--------geo begin--------")
	{
		geo, _ := New("./ip_sample.txt", GEO_F_ALIAS)
		fmt.Println(geo.FindGeo_ip("1.0.167.5"))
		fmt.Println(geo.FindGeo_ip("1.0.165.5"))
		fmt.Println(geo.FindGeo_ip("1.0.8.5"))
		geo.Clean()
	}

	fmt.Println("--------geo alias begin--------")
	{
		geo, _ := New("./ip_sample.txt", GEO_F_ALIAS)
		fmt.Println(geo.FindGeo_ip("1.0.167.5"))
		fmt.Println(geo.FindGeo_ip("1.0.165.5"))
		fmt.Println(geo.FindGeo_ip("1.0.8.5"))
		geo.Clean()
	}
}
