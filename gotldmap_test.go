package gotldmap

import "testing"

var findTests = []struct {
	tld      string
	mark     uint
	expected bool
}{
	{"de", 201, true},
	{"com", 165, true},
	{"", 0, false},
	{"zuerich", 840, true},
	{"academy", 5, true},
	{"abb", 1, true},
}

func TestFindByTld(t *testing.T) {
	for _, test := range findTests {
		got, _ := FindByTld(test.tld)
		if got != test.mark {
			t.Errorf("Expected - %t for tld: %q", test.expected, test.tld)
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
