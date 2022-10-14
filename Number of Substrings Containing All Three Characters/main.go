func numberOfSubstrings(s string) int {
    var ans, a, b, c, i, j int = 0, 0, 0, 0, 0, 0

	for i < len(s) {

		if s[i] == 'a' {
			a++
		}
		if s[i] == 'b' {
			b++
		}
		if s[i] == 'c' {
			c++
		}
		if a > 0 && b > 0 && c > 0 {
			ans += len(s) - i
		}

		t := j

		for j < i && a > 0 && b > 0 && c > 0 {
			if s[j] == 'a' {
				a--
			}
			if s[j] == 'b' {
				b--
			}
			if s[j] == 'c' {
				c--
			}
			j++
		}

		if t != j {
			ans += (j - t - 1) * (len(s) - i)
		}

		i++
	}

    return ans
}
