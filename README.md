# Guess Age

## Description
This project is a mini project for me to learn test-driven development (TDD) using golang.
In this project, user will be asked to guess a `person` age.
The age of the `person` is random 0 to 100 inclusive at initalization.

There are three conditions:
* when user guesses too high, system will return message "too high"
* when user guesses too low, system will return message "too low"
* when user guesses right, system will return message "correct"

## Environment Setup
If go isn't installed yet, you can install it by going to [this link](https://golang.org/dl/)
Choose go version `go1.11.4` based on your system.

To check the version, type this command on your terminal:
```
go version
```

## How to test
To test the project, go to this project root directory. In the directory, type the following command in your terminal:
```
go test --cover ./...
```

## How to build
To build the project, go to this project root directory. In the directory, type the following command in your terminal:
```
go build -o guess_person main.go
```

## How to run
To run the project, go to this project root directory. If it's not built yet, build the project. After that, in the directory, type the following command in your terminal:
```
./guess_person
```
