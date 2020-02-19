package tyutils

import (
	"bytes"
	crand "crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"io"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

func UUID() string {
	u1 := uuid.NewV4()
	return u1.String()
}

func GenerateSID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(crand.Reader, b); err == nil {
		return base64.URLEncoding.EncodeToString(b)
	}
	return ""
}

func RandomString(n int) string {
	buf := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// var result []byte
	result := make([]byte, 0, n)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, buf[r.Intn(len(buf))])
		// result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

// TODO: 1 <= c and c <= 19
func RandomIntger(c int) uint64 {
	if c <= 0 {
		return 0
	}

	if c > 19 {
		return 0
	}

	buf := []byte("0123456789")

	result := make([]byte, c)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ln := len(buf)
	for i := 0; i < c; i++ {
		result[i] = buf[r.Intn(ln)]
	}

	v, er := strconv.ParseUint(string(result), 10, 64)
	if er != nil {
		return 0
	}

	return v
}

func CompareMapStringString(m1, m2 map[string]string) bool {
	// if len(m1) != len(m2) {
	// 	return false
	// }
	//
	// for k, v := range m1 {
	// 	if v1, ok := m2[k]; !ok || v1 != v {
	// 		return false
	// 	}
	// }

	return reflect.DeepEqual(m1, m2)
}

func Substr(str string, start int, end int) (string, error) {
	bys := []byte(str)
	length := len(str)
	if start < 0 || start > length {
		return "", errors.New("start is wrong")
	}

	if end < 0 || end > length {
		return "", errors.New("end is wrong")
	}

	return string(bys[start:end]), nil
}

func BigHashCode(s string) string {
	v := crc64.Checksum([]byte(s), crc64.MakeTable(5849493))
	return fmt.Sprintf("%d", v)
}

func HashCode(s string) uint32 {
	v := crc32.ChecksumIEEE([]byte(s))
	return v
}

// Strings hashes a list of strings to a unique hashcode.
func HashCodes(strs ...string) string {
	var buf bytes.Buffer
	for _, s := range strs {
		buf.WriteString(fmt.Sprintf("%s-", s))
	}
	return fmt.Sprintf("%d", HashCode(buf.String()))
}
