# go-baseconv

[![GoDoc](https://godoc.org/github.com/bryanjeal/go-baseconv?status.svg)](https://godoc.org/github.com/bryanjeal/go-baseconv)
[![Go Report Card](https://goreportcard.com/badge/github.com/bryanjeal/go-baseconv)](https://goreportcard.com/report/github.com/bryanjeal/go-baseconv)

Convert uint64 into Base62 or BaseNoAmbiguous

## Base62
Base62 is a-z A-Z 0-9

## BaseNoAmbiguous (currently Base47)
The goal of BaseNoAmbiguous is to remove the Ambiguous Characters from Base62.

Currently BaseNoAmbiguous is equivalent to Base62 without "B8G6I1l0OQDS5Z2"
