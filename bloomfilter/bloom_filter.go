package bloomfilter

type BloomFilter struct {
	filter_len int
	bitMask    int
}

const (
	hash1Const = 17
	hash2Const = 223
)

func (bf BloomFilter) Hash1(s string) int {
	hashValue := 0
	for _, char := range s {
		code := int(char)
		hashValue = (hashValue*hash1Const + code) % bf.filter_len
	}
	return hashValue
}

func (bf BloomFilter) Hash2(s string) int {
	hashValue := 0
	for _, char := range s {
		code := int(char)
		hashValue = (hashValue*hash2Const + code) % bf.filter_len
	}
	return hashValue
}

func (bf *BloomFilter) Add(s string) {
	index1 := bf.Hash1(s)
	index2 := bf.Hash2(s)

	bf.bitMask |= 1 << index1
	bf.bitMask |= 1 << index2
}

func (bf *BloomFilter) IsValue(s string) bool {
	index1 := bf.Hash1(s)
	index2 := bf.Hash2(s)

	return bf.bitMask&(1<<index1) > 0 && bf.bitMask&(1<<index2) > 0
}
