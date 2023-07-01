package phonenumber

import (
	"fmt"
	"testing"
)

//func TestCheckPhoneNumber(t *testing.T) {
//	phone, err := CheckAndGetPhoneNumber("+639611429606")
//	fmt.Println(phone, err)
//}
//
//func TestMakePhoneNumber(t *testing.T) {
//	pNumber, err := MakePhoneNumberHelper(""+
//		"588021430", "CN")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println(pNumber.GetNormalizeDigits())
//}

func TestMakePhoneNumber2(t *testing.T) {
	// 63 969 025 1456
	// 63 995 659 1464
	pNumber, err := MakePhoneNumberHelper("+63 969 659 1464", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pNumber)
	fmt.Println(pNumber.GetNormalizeDigits(), ", ", pNumber.GetRegionCode(), ", ", pNumber.GetCountryCode())
}
