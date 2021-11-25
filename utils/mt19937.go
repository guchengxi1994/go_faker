package utils

import (
	"time"
)

const (
	W = 32
	N = 624
	M = 397
	R = 31
	A = 0x9908B0DF
	F = 1812433253
	U = 11
	D = 0xFFFFFFFF
	S = 7
	B = 0x9D2C5680
	T = 15
	C = 0xEFC60000
	L = 18
)

var (
	MASK_LOWER = (1 << R) - 1
	MASK_UPPER = (1 << R)
	index      int
	isInit     bool
	MT         [N]int
)

func srand(seed int) {
	index = 0
	isInit = true
	MT[0] = seed

	for i := 1; i < N; i++ {
		t := 1812433253*(MT[i-1]^(MT[i-1]>>30)) + i
		MT[i] = t & 0xffffffff
	}
}

func generate() {
	for i := 0; i < N; i++ {
		y := (MT[i] & MASK_UPPER) + (MT[(i+1)%N] & MASK_LOWER)
		MT[i] = MT[(i+M)%N] ^ (y >> 1)
		if y&1 != 0 {
			MT[i] ^= A
		}
	}
}

func Rand() int {
	if !isInit {
		seed := int(time.Now().Unix())
		srand(seed)
	}

	if index == 0 {
		generate()
	}

	y := MT[index]
	y = y ^ (y >> U)
	y = y ^ ((y << S) & B)
	y = y ^ ((y << T) & C)
	y = y ^ (y >> L)
	index = (index + 1) % N
	return y
}

func Randn(max int) int {
	if max <= 1 {
		return 0
	}
	return Rand() % max
}
