package main

import "testing"

func TestHello(t *testing.T){
    t.Run("Say hello to person when name is passed as an argument with no language", func(t *testing.T){
        got := Hello("", "Anav")
        want := "Hello Anav"

        assertCorrectMesage(t, got, want)
    })

    t.Run("Say Hello World when empty strings are passed as arguments", func(t *testing.T){
        got := Hello("", "")
        want := "Hello World"

        assertCorrectMesage(t, got, want)
    })

    t.Run("In Spanish", func(t *testing.T){
        got := Hello("Spanish", "Anav")
        want:= "Hola Anav"

        assertCorrectMesage(t, got, want)
    })

    t.Run("In French", func(t *testing.T){
        got := Hello("French", "Anav")
        want:= "Bonjour Anav"

        assertCorrectMesage(t, got, want)
    })
}

func assertCorrectMesage(t testing.TB, got, want  string){
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
