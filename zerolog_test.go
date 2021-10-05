package zerolog

import "testing"

func TestInit(t *testing.T) {
	z := ZerologExtension{}
	t.Logf("%+v\n", z)

	z.InitConfig()
}
