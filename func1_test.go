package controFunc

import (
	"testing"
)

func TestRevomeSpecialChars(t *testing.T) {

	const str = "want: %q, got: %q"

	t.Run("verifying outcome (return)", func(t *testing.T) {

		got := RevomeSpecialChars("C@l%i*en(tAut#he*nt!ica$tio$%nTo@ken")

		want := "ClientAuthenticationToken"

		if got != want {

			t.Errorf(str, want, got)
		}
	})

}
