package main

import "testing"

func TestHello(t *testing.T){
    t.Run("Say hello to person when name is passed as an argument", func(t *testing.T){
        got := Hello("Anav")
        want := "Hello Anav"

        assertCorrectMesage(t, got, want)
    })
    t.Run("Say Hello World when empty string is passed as an argument", func(t *testing.T){
        got := Hello("")
        want := "Hello World"

        assertCorrectMesage(t, got, want)
    })
}

func assertCorrectMesage(t testing.TB, got, want  string){
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
