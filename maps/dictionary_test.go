package dictionary

import "testing"

func assertString(t testing.TB, got, want string){
    t.Helper()

    if got != want {
        t.Errorf("got %q but wanted %q", got, want)
    }
}

func TestSearch(t *testing.T){
    dictionary := map[string]string{"test": "this is a test"}

    got := search(dictionary, "test") 
    want := "this is a test"

    assertString(t, got, want)
}
