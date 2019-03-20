#!/bin/bash

echo "Installing packages:"

echo "- accounting"
go get -u github.com/leekchan/accounting

echo "- easycsv"
go get -u github.com/yunabe/easycsv

echo "Done"
