# 🇿🇦 South African ID Validator for Go

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/jacovdloo/rsaid)

> Inspired by https://www.npmjs.com/package/south-african-id-validator

Validate ID numbers for the Republic of South Africa - taking eligibility age (16 years) into account.

The following information can be derived if the ID Number is valid:

- Gender
- Citizenship
- Date of birth

## Details

A South African ID number is a 13-digit number with the following format: `YYMMDDSNNNCAZ`.

The first six digits `YYMMDD` indicates a person's birth date. For example, 24 June 1995 becomes **950624**.

`{950624} SNNNCAZ`

> Although rare, it can happen that someone’s birth date does not correspond with their ID number.


The next digit `S` indicates a person's gender/sex. Females have a number of `0 to 4`, while males are `5 to 9`.

`{950624} {5} NNNCAZ`

The next three digits `NNN` indicates the position of birth by registry based on the date of birth and gender.

If for example, the number is 120, it means that person was the 120th person to be registered as having been born on that particular day, for that gender.

`{950624} {5} {120} CAZ`

The next digit `C` indicates citizenship. `0` if the person is a South African citizen, or `1` if the person is a permanent resident.

`{950624} {5} {120} {0} AZ`

The next digit `A` was used until the late 1980s to indicate a person’s race. This has been eliminated and old ID numbers were reissued to remove this.

`{950624} {5} {120} {0} {0} Z`

The last digit `Z` is a checksum digit, used to check that the number sequence is accurate using the [Luhn](https://en.wikipedia.org/wiki/Luhn_algorithm) algorithm.

`{950624} {5} {120} {0} {0} {8}`

So, ID number `9506245120008` will reflect as the `120th` `male` South African `citizen` born/registered on the `24th of June, 1995`

## Install

```go
go get github.com/jacovdloo/rsaid
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/jacovdloo/rsaid"
)

func main() {

	id := rsaid.IdentityNumber {
		Number: "9506245120008",
	}

	err := id.IsValid();
	cit, _ := id.IsCitizen()
	gen, _ := id.Gender()
	dob, _ := id.DateOfBirth()

	fmt.Println("Valid:", err == nil)
	fmt.Println("Citizen:", cit)
	fmt.Println("Male:", gen == rsaid.GenderMale)
	fmt.Println("DOB:", dob)
}
```
