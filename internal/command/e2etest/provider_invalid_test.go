package e2etest

import (
	"strings"
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

	// err is expected to be non-nil, so we can ignore in this test
	stdout, stderr, _ := tf.Run("init")

	if !strings.Contains(stderr, "Duplicate required providers configuration") {
		t.Fatalf("wrong output:\nstdout:%s\nstderr%s", stdout, stderr)
	}
}
