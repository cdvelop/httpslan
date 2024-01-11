package httpslan_test

import (
	"testing"

	"github.com/cdvelop/httpslan"
)

func Test_GenerateCert(t *testing.T) {

	cert := httpslan.AddHttpsLan(nil)

	cert.BuildSSlCertificate()

}
