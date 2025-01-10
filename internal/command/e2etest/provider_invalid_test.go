package e2etest

import (
	"testing"

	"github.com/opentofu/opentofu/internal/e2e"
)

// ProviderInvalidTest
func TestProviderInvalid(t *testing.T) {
	if !canRunGoBuild {
		// We're running in a separate-build-then-run context, so we can't
		// currently execute this test which depends on being able to build
		// new executable at runtime.
		//
		// (See the comment on canRunGoBuild's declaration for more information.)
		t.Skip("can't run without building a new provider executable")
	}
	t.Parallel()

	tf := e2e.NewBinary(t, tofuBin, "testdata/provider-invalid")

	_, stderr, err := tf.Run("init")
	if err != nil {
		t.Fatalf("unexpected error: %s\n%s", err, stderr)
	}
}
