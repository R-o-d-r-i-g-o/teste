package controFunc

import "testing"

func TestLetOnlyNumbers(t *testing.T) {

	const str = "want: %q, got: %q"

	t.Run("verifying outcome (return)", func(t *testing.T) {

		got := LetOnlyNumbers("1g2h3#4m5n6v7c8U9$10")

		want := "12345678910"

		if got != want {

			t.Errorf(str, want, got)
		}
	})

}
