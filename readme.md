
# JP Morgan Technical Test


## Introduction

This project generates the following reports, aggregating instructions to buy or sell sent by various clients to JP Morgan to execute in the international market:

- Amount in USD settled incoming everyday;
- Amount in USD settled outgoing everyday;
- Ranking of entities based on incoming and outgoing amount.


## Prerequisites

This project was built using [Golang 1.12 for Ubuntu](https://github.com/golang/go/wiki/Ubuntu)


## Installation

This project requires the following packages to be installed:

- [accounting](https://github.com/leekchan/accounting): to format money and currency in reports
- [easycsv](https://github.com/yunabe/easycsv): to parse the CSV file of instructions

You can install them at once by using the following script:

```sh
./go_get.sh
```


## Usage

Use the following commands from within the `instructions` folder to compile then run the binary for this project:

```sh
# Compile the binary
go build

# Run the compiled binary
./instructions
```

The binary will be using the sample CSV file located within the same folder. Optionally, you can specify your own CSV file as follows:

```sh
./instructions -file=sample.csv
```


## Unit tests

Use the following command to run the unit tests:

```sh
go test -v
```
