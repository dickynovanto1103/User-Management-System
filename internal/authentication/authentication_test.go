package authentication

import (
	"testing"

	"github.com/dickynovanto1103/User-Management-System/internal/stringutil"
)

func TestVerifyPassword(t *testing.T) {
	casesFalse := []struct {
		in, notwant string
	}{
		{"pass1", stringutil.CreateRandomString(100)},
		{"pass2", stringutil.CreateRandomString(100)},
	}
	for _, val := range casesFalse {
		got := VerifyPassword(val.notwant, val.in)
		if got {
			t.Errorf("Expected to be not verified")
		}
	}

	casesTrue := []struct {
		in, stored string
	}{
		{"pass0", "9fbb16c4ee6143f5fd0f57b642cfb3fe66277070803f43c41057d28cc474c7663ea674580ef6f4660309a2552b8875f2cbacb561736f7dbe79b87da10f7cd97b80a0dd3c299d718704819036cf018039f32e52e42afe488bc85ae2cdd2d28b1a"},
		{"pass1", "0f4350bb17bbe1592f573f7e326891d1ad2a4a3dc3a7803c8c87e3517440683e73a27d3ec9bf2937c5c348604bed9163d7c171d673f03dd645ea5ab89c6e782588ad9b3cb379c759023ec1890f587330e0dc32eb1e9d8925ff001e8951e6a646"},
	}

	for _, val := range casesTrue {
		got := VerifyPassword(val.stored, val.in)
		if !got {
			t.Errorf("Expected to be verified")
		}
	}
}
