package pointersanderrors

import "testing"

func assertBalance(t testing.TB, got, want Bitcoin) {
	t.Helper()

	if (got != want) {
		t.Errorf("got %s want %s", got, want)
	}

}

func asserError(t testing.TB, err error, want string) {
	t.Helper()

	if err == nil {
		t.Fatal("wanted error but didn't get one")
	}

	if err.Error() != want {
		t.Errorf("got %s want %s", err, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got error but didn't want one")
	}
}