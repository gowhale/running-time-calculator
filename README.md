# running-time-calculator

## Introduction
If I want to run a 5KM run in 24 minutes, what average pace will I need to aim for? What time should I be achieveing when I reach the second KM? 

I'm aiming to create Go script that will take in 2 inputs: 
1. DISTANCE 
1. TIME

Then will output the following: 
1. an average pace to aim for
1. KM by KM times to achieve

Just to practice using Go and maybe help me get a running PB 👀

## Running the main code for yourself

To be able to run the main sequence of code:

1. clone this repo
1. cd into the this repo
1. run the command: "brew install go"
1. run the command: "go run running_time_calculator.go"
1. The program should now ask you for inputs and give you metrics to follow to achieve your running dreams!

## To run the tests for yourself

To be able to run the tests in this repo you will need to do the following:

1. clone this repo
1. cd into the this repo
1. run the command: "brew install go"
1. run the command: "go get github.com/stretchr/testify/assert" (this is equivilant to "pip install xyz")
    1. Make sure you .gitconfig file is up to date for this
1. finally to run the tests run: "go test -v"
1. Tests will now run.  
