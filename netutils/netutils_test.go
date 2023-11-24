package netutils_test

import (
	"fmt"
	"testing"

	"github.com/daqiancode/myutils/netutils"
	"github.com/stretchr/testify/assert"
)

func TestGetIP(t *testing.T) {
	localIPs := netutils.GetLocalIPs()
	assert.True(t, len(localIPs) > 0)
	fmt.Println(localIPs)
}
