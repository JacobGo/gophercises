# Gophercises: Quiz Game

### Questions
* Why do I need to handle `fmt.Fprint` errors? See [StackOverflow](https://stackoverflow.com/a/43976633)
* How can I use types more efficiently to hold my scores?
* Would it be better to select from input and timer channels in the control flow instead of arbitrarily exiting the program?

### Neat Features
* Added a Scoreboard Server to track multiuser attempts
* Identified clients through the [machineid go module](https://github.com/denisbrodbeck/machineid).

### Ideas
* Customizable time limit
* Randomly generate questions on the Scoreboard Server
* Show the full Scoreboard in an ASCII table on the client
* Randomly generate names for machine ids
    * maybe based on [BitCoin's seed phrase implementation] (https://github.com/tyler-smith/go-bip39/blob/master/bip39.go)
    