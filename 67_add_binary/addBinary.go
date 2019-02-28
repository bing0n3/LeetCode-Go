package addbinary

func addBinary(a string, b string) string {
	res := []byte{}
	m := len(a)
	n := len(b)
	upgrade := false

	i := 1
	for ; i <= m && i <= n; i++ {
		tmp := 0
		if upgrade == true {
			tmp++
		}
		if a[m-i] == '1' && b[n-i] == '1' {
			upgrade = true
			res = append([]byte{byte(tmp) + '0'}, res[:]...)
			continue
		}
		if a[m-i]-'0'+b[n-i]-'0' == 1 {
			tmp++
		}

		if tmp >= 2 {
			upgrade = true
			tmp = tmp - 2
		} else {
			upgrade = false
		}

		res = append([]byte{byte(tmp) + '0'}, res[:]...)
	}

	for i <= m {
		tmp := 0
		if upgrade && a[m-i]-'0' == 1 {
			upgrade = true
			tmp = 0
		} else if upgrade {
			upgrade = false
			tmp = 1
		} else {
			upgrade = false
			tmp = int(a[m-i] - '0')
		}
		res = append([]byte{byte(tmp) + '0'}, res[:]...)
		i++
	}

	for i <= n {
		tmp := 0
		if upgrade && b[n-i]-'0' == 1 {
			upgrade = true
			tmp = 0
		} else if upgrade {
			upgrade = false
			tmp = 1
		} else {
			upgrade = false
			tmp = int(b[n-i] - '0')
		}
		res = append([]byte{byte(tmp) + '0'}, res[:]...)
		i++
	}

	if upgrade == true {
		res = append([]byte{'1'}, res[:]...)
	}

	return string(res)
}

// optimize
func addBinary_op(a string, b string) string {
	res := []byte{}
	carry := 0
	m := len(a) - 1
	n := len(b) - 1

	for m >= 0 || n >= 0 {
		sum := carry
		if m >= 0 {
			sum += int(a[m] - '0')
			m--
		}

		if n >= 0 {
			sum += int(b[n] - '0')
			n--
		}

		res = append([]byte{byte(sum%2) + '0'}, res...)
		carry = sum / 2
	}

	if carry > 0 {
		res = append([]byte{'1'}, res...)
	}

	return string(res)
}
