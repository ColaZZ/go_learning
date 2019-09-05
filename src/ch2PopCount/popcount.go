package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	ret := byte(0)
	for i := uint8(0); i < 8; i++ {
		ret += pc[byte(x>>(i*8))]
	}
	return int(ret)
}

func PopCount2(x uint64) int {
	ret := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			ret++
		}
		x = x >> 1
	}
	return ret
}

func PopCount3(x uint64) int {
	ret := 0
	for x != 0 {
		x = x & (x - 1)
		ret++
	}
	return ret
}
