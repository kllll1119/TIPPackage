package TIPPackage

import (
	"bytes"
	"strconv"
	"strings"
)

func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}
func IpIntToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}

// GetAddress 通过IP获取地址
func GetAddress(strIP string) string {
	n := StringIpToInt(strIP)
	nLen := uint32(len(AddressList))
	return check(0, nLen, uint32(n))
}

// 二分法
func check(nLeft uint32, nRight uint32, nValue uint32) string {
	nMid := (nLeft + nRight) / 2
	if AddressList[nMid].Left > nValue {
		return check(nLeft, nMid, nValue)

	}

	if AddressList[nMid].Right < nValue {
		return check(nMid, nRight, nValue)

	}
	return AddressList[nMid].S
}
