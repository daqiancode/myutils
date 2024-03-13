package nets_test

import (
	"fmt"
	"testing"

	"github.com/daqiancode/myutils/nets"
	"github.com/stretchr/testify/assert"
)

func TestGetIP(t *testing.T) {
	localIPs := nets.GetLocalIPs()
	assert.True(t, len(localIPs) > 0)
	fmt.Println(localIPs)
}
