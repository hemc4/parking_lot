# Parking Lot
This repository includes the source code for the assigment. Required automated testing suite is conencted to build file and required commnds for the automated testing connects to project build systema that utilize the `make`

There are 2 package in the project, `main` and `dao` Dao have been used to seperate the storage logic.

## Installation and Compilation 
`make all  ` command will build and test. It should output a exec file `parking_lot` in the  project directory.

There is one external library used for testing `github.com/stretchr/testify` In order to run it properly this package needs to be downloaded. To download the package golang's internal tool `go get` is used and connected by make file.

## Testing

Both packages have unit tests cases, so `make test` command run the test cases for both the packages. Unit test utilize the internal package `testing` and exteranal package for more complex test cases.

## Run 
Binary can be run in the parent directory with of without file aragumant  by command `./parking_lot` or `./parking_lot {filepath}`

