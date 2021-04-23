package popcount

const (
    BYTESOF_INT64 = 8
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

// PopCount returms the population count (number of set bits) of x
func PopCount(x unit64) int {
    count := 0

    for i := 0; i < BYTESOF_INT64; i++ {
        count += pc[byte(x>>i*8)]
    }
    return count
}
