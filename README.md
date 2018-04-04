# gox
Tic-tac-toe game in go

## Documentation
Use `go doc` to generate documentation. Allmost all exported fields are in-line
documented.

## Structure
This repository is split into four packages:
+   main:
    Contains code for the binary with flag parsing game-mode selection and the
    actual gamemodes.
+   utils:
    Some functions that makes handling user interaction easy.
+   gox:
    Contains all type definitons for the board and game
+   ai:
    This is a tiny artifical inteligence. It rates fields on a simple pre-
    defiend value but most importantly it checks for wins or ties. It then uses
    the negamax- algortihm with varable search depth to determine the next move.
    It is suprisingly stupid.

## Tests
Currently the only test ist broken because the function is now obsolete and
removed. This uses standard `go test` tools.

## go get
I you should be able to get the code by entering
    
    go get -u -v github.com/D1CED/gox
    
in your terminal. Otherwise just clone the repo.
