package util

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomFloat returns a random float between min and max
func RandomFloat(min, max float64) float64 {
	return min * rand.Float64() * (max - min + 1)
}

// RandomInt returns a random int between min and max
func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

// RandomString returns a random string of length n
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// RandomDate returns a random date between 1970-01-01 00:00:00 and 2070-01-01 00:00:00 UTC
func RandomDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	sec := min + rand.Int63n(max-min)
	return time.Unix(sec, 0)
}

// RandomColorHexString generates a random hex color string
func RandomColorHexString() string {
	return "#" + RandomString(6)
}

// RandomName generates a random name
func RandomName() string {
	return RandomString(6)
}

// RandomIdInt generates a random id number
func RandomId() int32 {
	return RandomInt(1, 10)
}

// RandomAmount generates a random amount
func RandomAmount() float64 {
	return RandomFloat(0, 1000)
}
