package address

import "testing"

func TestGetAddressByIP(t *testing.T) {
	addr := GetAddressByIP("175.11.202.69")
	t.Log(addr)
}
