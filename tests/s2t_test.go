package opencc

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/longbridgeapp/opencc"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func readFile(filename string) string {
	pwd, _ := os.Getwd()
	data, err := os.ReadFile(path.Join(pwd, "fixtures", filename))
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(data))
}

func Test_s2t(t *testing.T) {
	raw := readFile("html-raw.txt")
	expected := readFile("html-s2t.txt")

	cc, _ := opencc.New("s2t")
	out, _ := cc.Convert(raw)

	fmt.Println(out)

	if strings.TrimSpace(expected) != strings.TrimSpace(out) {
		dmp := diffmatchpatch.New()

		diffs := dmp.DiffMain(expected, out, true)

		t.Errorf(dmp.DiffPrettyText(diffs))
	}

}

func TestFinance_s2hk_finance(t *testing.T) {
	raw := readFile("html-raw.txt")
	expected := readFile("html-s2hk-finance.txt")

	cc, _ := opencc.New("s2hk-finance")
	out, _ := cc.Convert(raw)

	fmt.Println(out)

	if strings.TrimSpace(expected) != strings.TrimSpace(out) {
		dmp := diffmatchpatch.New()

		diffs := dmp.DiffMain(expected, out, true)

		t.Errorf(dmp.DiffPrettyText(diffs))
	}

}

func assertCases(t *testing.T, cases map[string]string) {
	t.Helper()
	cc, _ := opencc.New("s2hk")

	for raw, expected := range cases {
		out, _ := cc.Convert(raw)
		if strings.TrimSpace(expected) != strings.TrimSpace(out) {
			t.Errorf("expected %s, got %s", expected, out)
		}
	}
}

// Special hotfix in this project
func TestSelfSpecialHotfix(t *testing.T) {
	cases := map[string]string{
		"来自于汇丰，以及汇丰银行，汇入的款项": "來自於滙豐，以及滙豐銀行，匯入的款項",
		"汇业银行集团": "滙業銀行集團",
	}

	assertCases(t, cases)
}
