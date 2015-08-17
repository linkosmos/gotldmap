package gotldmap

import "testing"

var findTests = []struct {
	tld      string
	mark     int
	expected bool
}{
	{"de", 225, true},
	{"com", 185, true},
	{"", 0, false},
	{"zuerich", 1039, true},
	{"academy", 5, true},
	{"abb", 1, true},
	{"zw", 1040, true},
	{"wiki", 907, true},
	{"seek", 1045, true},
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
