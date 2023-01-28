package main

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"
)

func basename(s string) string {
	// "a/b/c.go" => "c"
	// "c.d.go" => "c.d"
	// "abc" => "abc"

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basenameV2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	// 12345 => 12,345
	// 12 => 12
	n := len(s)
	if n < 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToStrings(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')

	return buf.String()
}

func IsUp(v net.Flags) bool {
	fmt.Printf("TRY: %b&%b = %b \n", v, net.FlagUp, v&net.FlagUp)
	return v&net.FlagUp == net.FlagUp
}
func TurnDown(v *net.Flags) {
	*v &^= net.FlagUp
}
func SetBroadcast(v *net.Flags) {
	*v |= net.FlagBroadcast
}
func IsCast(v net.Flags) bool {
	return v&(net.FlagBroadcast|net.FlagMulticast) != 0
}

func main() {
	//fmt.Println(basename("a/b/c.go"))
	//fmt.Println(basename("c.d.go"))
	//fmt.Println(basename("abc"))
	//fmt.Println(comma("12345"))
	//
	//fmt.Println(intsToStrings([]int{1, 2, 3}))
	//fmt.Println(strconv.FormatInt(int64(123), 2))
	//strconv.Atoi("123")
	//strconv.ParseInt("123", 10, 64)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)

	var v net.Flags = net.FlagMulticast | net.FlagUp
	fmt.Printf("> %b \n", net.FlagUp)
	fmt.Printf("%b %t \n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))

}
