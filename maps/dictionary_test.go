package dictionary

import "testing"

func assertString(t testing.TB, got, want string){
    t.Helper()

    if got != want {
        t.Errorf("got %q but wanted %q", got, want)
    }
}

func assertError(t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error %q want %q", got, want)
    }
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }

    assertString(t, got, definition)
}

func TestSearch(t *testing.T){
    dictionary := Dictionary{"test": "this is a test"}

    t.Run("Known word", func(t *testing.T) {
        got, _ := dictionary.Search("test") 
        want := "this is a test"

        assertString(t, got, want)
    })

    t.Run("Unknown word", func(t *testing.T) {
        _, got := dictionary.Search("unknown") 

        if got == nil {
            t.Fatal("Expected to get error")
        }

        assertError(t, got, ErrNotFound)
    })
}

func TestAdd (t *testing.T) {
    dictionary := Dictionary{}

    word := "go"
    definition := "golang"

    dictionary.Add(word, definition)

    assertDefinition(t, dictionary, word, definition)
}
