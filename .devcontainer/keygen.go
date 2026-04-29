// wolfram mathematica 14.3 keygen in Go
// ported from https://github.com/thedeepdeepsky/mathematica_keygen/blob/dbafdf8db96679effeede2acdead2fd9c32eaada/keygen_testv0.1_pub.py

package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	hashCode1   = 0b1000001011100001 // 33505 / 0x82E1
	hashCode2   = 0b1000001100100101 // 33573 / 0x8325
	magicNumber = 59222
)

// MathPass holds the state for a single keygen operation.
type MathPass struct {
	version       [3]int
	mathID        string
	activationKey string
	password      string
}

// NewMathPass creates a MathPass for the given Math ID and version.
// If customKey is non-empty and matches the required format it is used;
// otherwise a random activation key is generated using rng.
func NewMathPass(mathID, version, customKey string, rng *rand.Rand) *MathPass {
	mp := &MathPass{}
	mp.version = parseVersion(version)
	mp.setMathID(mathID)
	keyFmt := mp.activationKeyFormat()
	if customKey != "" && checkFormat(keyFmt, customKey) {
		mp.activationKey = customKey
	} else {
		mp.activationKey = randomActivationKey(keyFmt, rng)
	}
	return mp
}

func (mp *MathPass) versionAtLeast(major, minor, patch int) bool {
	v := mp.version
	return v[0] > major ||
		(v[0] == major && v[1] > minor) ||
		(v[0] == major && v[1] == minor && v[2] >= patch)
}

func (mp *MathPass) activationKeyFormat() string {
	if mp.versionAtLeast(14, 1, 0) {
		return "xxxx-xxxx-aaaaaa"
	}
	return "xxxx-xxxx-xxxxxx"
}

func (mp *MathPass) setMathID(mathID string) bool {
	if checkFormat("xxxx-xxxxx-xxxxx", mathID) {
		mp.mathID = mathID
		return true
	}
	return false
}

// GeneratePassword generates the password for the given math number and expiry
// date.  Empty strings fall back to the defaults ("800001" and 999 days from
// now respectively).
func (mp *MathPass) GeneratePassword(mathNum, expireDate string) bool {
	if mathNum == "" {
		mathNum = "800001"
	}
	if expireDate == "" {
		expireDate = dateAfter(999)
	}
	if mp.versionAtLeast(14, 1, 0) {
		return mp.generatePasswordV14_1_0(mathNum, expireDate)
	}
	return false
}

func (mp *MathPass) generatePasswordV14_1_0(mathNum, expireDate string) bool {
	strVal := mp.mathID + "@" + expireDate + "$" + mathNum + "&" + mp.activationKey
	chars := reverseString(strVal)
	hc := magicNumber
	n0 := encodingCharacters(hashCode1, hc, chars)
	n1 := (n0 + 0x72FA) % 65536
	hc = encodingHash(n1)
	n2 := encodingCharacters(hashCode2, hc, chars)
	mp.password = constructPassword(n1, n2) + "::" + mathNum + ":" + expireDate
	return true
}

// hasher processes one byte through the CRC-like hash step.
func hasher(hasherCode, hashVal, byteVal int) int {
	for i := 0; i < 8; i++ {
		bit := byteVal & 1
		if hashVal%2 == bit {
			hashVal >>= 1
		} else {
			hashVal >>= 1
			hashVal ^= hasherCode
		}
		byteVal >>= 1
	}
	return hashVal
}

// splitHex maps a 16-bit value to a 5-digit decimal representation and returns
// the digits from least-significant to most-significant (index 0 = ones place).
func splitHex(hexVal int) [5]int {
	n := int(math.Floor(float64(hexVal) * 99999.0 / 0xFFFF))
	var d [5]int
	for i := 0; i < 5; i++ {
		d[i] = n % 10
		n /= 10
	}
	return d
}

// encodingHash computes the secondary hash value derived from n1.
func encodingHash(n1 int) int {
	n := int(math.Floor(float64(n1) * 99999.0 / 0xFFFF))
	n01 := n % 100
	n -= n01
	n2 := n % 1000
	n -= n2
	n += n01*10 + n2/100 // integer division is exact since n2 is always a multiple of 100
	temp := int(math.Ceil(float64(n) * 65535.0 / 99999))
	return hasher(hashCode2, hasher(hashCode2, 0, temp&0xFF), temp>>8)
}

// encodingCharacters searches for a 16-bit value whose two bytes, appended to
// the hashed character stream, produce 0xA5B6.
func encodingCharacters(hasherCode, hashVal int, chars []int) int {
	for _, c := range chars {
		hashVal = hasher(hasherCode, hashVal, c)
	}
	var c1, c2 int
	for c1 = 0; c1 < 256; c1++ {
		for c2 = 0; c2 < 256; c2++ {
			if hasher(hasherCode, hasher(hasherCode, hashVal, c1), c2) == 0xA5B6 {
				return c1 | (c2 << 8)
			}
		}
	}
	return c1 | (c2 << 8)
}

// constructPassword assembles the printable password from n1 and n2.
//
// Python uses splitHex(n)[::-1] so its index 0 is the most-significant digit.
// Go's splitHex stores the least-significant digit at index 0, so we reverse
// the mapping: Python index i  →  Go index (4 - i).
//
// Password pattern (Python indices, both arrays reversed):
//   n2[3] n1[3] n1[1] n1[0] - n2[4] n1[2] n2[0] - n2[2] n1[4] n2[1]
//
// Translated to Go (splitHex, LSB-first) indices:
//   n2[1] n1[1] n1[3] n1[4] - n2[0] n1[2] n2[4] - n2[2] n1[0] n2[3]
func constructPassword(n1, n2 int) string {
	a := splitHex(n1) // a[0]=ones, a[4]=ten-thousands
	b := splitHex(n2)
	return fmt.Sprintf("%d%d%d%d-%d%d%d-%d%d%d",
		b[1], a[1], a[3], a[4],
		b[0], a[2], b[4],
		b[2], a[0], b[3])
}

// reverseString returns the byte values of s in reverse order.
func reverseString(s string) []int {
	out := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		out[i] = int(s[len(s)-1-i])
	}
	return out
}

// checkFormat validates s against a format string where:
//
//	'x' = ASCII digit, 'a' = uppercase letter, 'b' = digit or uppercase letter,
//	any other character must match literally.
func checkFormat(format, s string) bool {
	if len(format) != len(s) {
		return false
	}
	for i := 0; i < len(format); i++ {
		switch format[i] {
		case 'x':
			if s[i] < '0' || s[i] > '9' {
				return false
			}
		case 'a':
			if s[i] < 'A' || s[i] > 'Z' {
				return false
			}
		case 'b':
			if !((s[i] >= '0' && s[i] <= '9') || (s[i] >= 'A' && s[i] <= 'Z')) {
				return false
			}
		default:
			if format[i] != s[i] {
				return false
			}
		}
	}
	return true
}

// randomActivationKey generates a random key matching the given format string.
func randomActivationKey(format string, rng *rand.Rand) string {
	var sb strings.Builder
	for i := 0; i < len(format); i++ {
		switch format[i] {
		case 'x':
			sb.WriteByte(byte('0' + rng.Intn(10)))
		case 'a':
			sb.WriteByte(byte('A' + rng.Intn(26)))
		default:
			sb.WriteByte(format[i])
		}
	}
	return sb.String()
}

// parseVersion converts "14.1.0" into [14, 1, 0].
func parseVersion(version string) [3]int {
	parts := strings.Split(version, ".")
	var v [3]int
	for i, p := range parts {
		if i >= 3 {
			break
		}
		v[i], _ = strconv.Atoi(p)
	}
	return v
}

// dateAfter returns the date that is n days from today in "YYYYMMDD" format.
func dateAfter(days int) string {
	return time.Now().AddDate(0, 0, days).Format("20060102")
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [mathid [activation_key]]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "  mathid         Math ID in the format xxxx-xxxxx-xxxxx")
		fmt.Fprintln(os.Stderr, "  activation_key Activation key in the format xxxx-xxxx-aaaaaa (optional)")
		fmt.Fprintln(os.Stderr, "\nIf arguments are omitted, the program runs in interactive mode.")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()

	var mathID, customKey, expireDate string

	switch len(args) {
	case 0:
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Math ID (xxxx-xxxxx-xxxxx): ")
		mathID, _ = reader.ReadString('\n')
		mathID = strings.TrimSpace(mathID)

		fmt.Print("Activation Key (leave blank to generate one, format xxxx-xxxx-aaaaaa): ")
		customKey, _ = reader.ReadString('\n')
		customKey = strings.TrimSpace(customKey)

		fmt.Print("Expiry Date (YYYYMMDD, default 999 days from now): ")
		expireDate, _ = reader.ReadString('\n')
		expireDate = strings.TrimSpace(expireDate)
	case 1:
		mathID = args[0]
	default:
		mathID = args[0]
		customKey = args[1]
	}

	mp := NewMathPass(mathID, "14.1.0", customKey, rng)
	mp.GeneratePassword("800001", expireDate)

	fmt.Printf("Activation Key: %s\n", mp.activationKey)
	fmt.Printf("Password: %s\n", mp.password)
}