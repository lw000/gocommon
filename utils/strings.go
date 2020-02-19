package tyutils

func Reverse(src string) string {
	s := []rune(src)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func Split(buf []byte, limit int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/limit+1)
	for len(buf) >= limit {
		chunk, buf = buf[:limit], buf[limit:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}

func Include(a string, s ...string) bool {
	for _, v := range s {
		if v == a {
			return true
		}
	}
	return false
}

func Merge(s1 []string, s2 ...string) []string {
	return []string{}
}
