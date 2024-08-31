package algorithms

import (
	"testing"
)

func TestPadding(t *testing.T) {
	var str string

	str = padStr("1001101", "0", 8, true)
	if str != "01001101" {
		t.Errorf("padStr(\"1001101\", \"0\", 8, true) = %s; want \"01001101\"", str)
	}

	str = padStr("1001101", "=", 10, false)
	if str != "1001101===" {
		t.Errorf("padStr(\"1001101\", \"0\", 8, true) = %s; want \"1001101===\"", str)
	}
}

func TestBase64Encoding(t *testing.T) {
	helperTestBase64(t, "Many hands make light work.", "TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu")
	helperTestBase64(t, "light work.", "bGlnaHQgd29yay4=")
	helperTestBase64(t, "light work", "bGlnaHQgd29yaw==")
	helperTestBase64(t, "M", "TQ==")
}

func helperTestBase64(t *testing.T, input string, output string) {
	var str string = CustomBase64(input)
	if str != output {
		t.Errorf("customBase64(\"%s\") = %s; want \"%s\"", input, str, output)
	}
}
