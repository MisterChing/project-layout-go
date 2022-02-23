package netx

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestIsLan(t *testing.T) {
    ipStr := "192.168.1.2"
    isLan := IsLan(ipStr)
    assert.True(t, isLan)
    ipStr = "210.168.1.2"
    isLan = IsLan(ipStr)
    assert.False(t, isLan)
}

func TestGetLocalIP(t *testing.T) {
    localIP := GetLocalIP()
    isLan := IsLan(localIP)
    assert.True(t, isLan)
}
