package main

import "fmt"

 const (
     englishGreeting = "Hello "
     spanishGreeting = "Hola "
     frenchGreeting = "Bonjour "

     spanish = "Spanish"
     french = "French"
 )

func Hello(languge, name string) string {
    if name == "" {
        name = "World"
    }

    return greetPrefix(languge) + name
}

func greetPrefix(language string)(prefix string){
    switch language {
    case spanish: 
        prefix = spanishGreeting
    case french:
        prefix = frenchGreeting
    default:
        prefix = englishGreeting
    }

    return

}

func main(){
    fmt.Println(Hello("", "Anav"))
}
