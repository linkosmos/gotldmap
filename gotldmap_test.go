package gotldmap

import "testing"

var findTests = []struct {
	tld      string
	mark     int
	expected bool
}{
	{"de", 1225, true},
	{"com", 1185, true},
	{"", 0, false},
	{"zuerich", 2039, true},
	{"academy", 1005, true},
	{"abb", 1001, true},
	{"zw", 2040, true},
	{"wiki", 1907, true},
	{"seek", 2045, true},
	{"buy", 2124, true},
	{"author", 2121, true},
	{"zero", 2144, true},
	{"pars", 2145, true},
	{"xn--pbt977c", 2165, true},
	{"analytics", 2166, true},
	{"safety", 2176, true},
	{"pid", 2174, true},
	{"clinique", 2182, true},
	{"promo", 2184, true},
	{"baidu", 2185, true},
	{"volkswagen", 2188, true},
	{"tube", 2189, true},
	{"weather", 2190, true},
	{"softbank", 2198, true},
	{"lifeinsurance", 2200, true},
	{"edeka", 2201, true},
	{"tmall", 2205, true},
	{"pamperedchef", 2202, true},
	{"adac", 2206, true},
	{"weatherchannel", 2211, true},
	{"deloitte", 2212, true},
	{"unicom", 2213, true},
	{"quest", 2215, true},
	{"tvs", 2222, true},
	{"bcg", 2228, true},
	{"gmbh", 2231, true},
	{"stream", 2254, true},
	{"shaw", 2256, true},
	{"barefoot", 2258, true},
}

func TestFindByTld(t *testing.T) {
	for _, test := range findTests {
		got, _ := FindByTld(test.tld)
		if got != test.mark {
			t.Errorf("Expected - %t %d for tld: %q", test.expected, test.mark, test.tld)
		}
	}
}

func TestFindByMark(t *testing.T) {
	for _, test := range findTests {
		got, _ := FindByMark(test.mark)
		if got != test.tld {
			t.Errorf("Expected - %t for tld: %q", test.expected, test.tld)
		}
	}
}

func TestTldExist(t *testing.T) {
	for _, test := range findTests {
		got := TldExist(test.tld)
		if got != test.expected {
			t.Errorf("Expected - %t for tld: %q", test.expected, test.tld)
		}
	}
}

func TestNoDuplicates(t *testing.T) {
	m := make(map[int]int)

	for _, position := range Map {
		m[position]++
	}

	for tld, duplicated := range m {
		if duplicated > 1 {
			t.Errorf("tld mark %d is duplicated times %d", tld, duplicated)
		}
	}
}
