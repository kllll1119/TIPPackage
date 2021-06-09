package TIPPackage

import (
	"testing"

	"github.com/kllll1119/TIPPackage"
)

type Data struct {
	Ip     string
	Expect string
}

func TestGetAddress(t *testing.T) {
	s := []Data{
		Data{"192.168.0.1", "局域网"},
		Data{"118.113.243.111", "四川省"},
		Data{"17.15.144.12", "法国"},
		Data{"7.55.44.11", "美国"},
	}

	for i := 0; i < len(s); i++ {
		if ans := TIPPackage.GetAddress(s[i].Ip); ans != s[i].Expect {
			t.Errorf("[%d]ip:%s,expect:%s,real:%s", i, s[i].Ip, s[i].Expect, ans)
		}
	}
}
