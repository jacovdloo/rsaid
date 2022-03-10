package rsaid_test

import (
	"testing"

	"github.com/jacovdloo/rsaid"
)

var validMale = "9506245120008"
var validFemale = "9506244120009"
var invalidDOB = "9502305120008"
var invalidIDN = "9506245120009"
var nonCitizen = "9506245120107"
var shortNumber = "950624"

func Test_IsValid(t *testing.T) {

	if err := rsaid.IsValid(invalidIDN); err == nil {
		t.Errorf("Does not determine valid id correctly")
	}

	if err := rsaid.IsValid(validMale); err != nil {
		t.Errorf("Does not determine valid id correctly")
	}

	if err := rsaid.IsValid(shortNumber); err == nil {
		t.Errorf("Does not determine valid id correctly")
	}
}

func Test_Gender(t *testing.T) {

	man, man_err := rsaid.Gender(validMale)
	woman, wom_err := rsaid.Gender(validFemale)
	if man != rsaid.Male || man_err != nil {
		t.Errorf("Does not determine gender correctly")
	}

	if woman != rsaid.Female || wom_err != nil {
		t.Errorf("Does not determine gender correctly")
	}
}

func Test_IsCitizen(t *testing.T) {

	cit, cit_err := rsaid.IsCitizen(validMale)
	pem, pem_err := rsaid.IsCitizen(nonCitizen)

	if cit != true || cit_err != nil {
		t.Errorf("Does not determine citizenship correctly")
	}

	if pem != false || pem_err != nil {
		t.Errorf("Does not determine citizenship correctly")
	}
}

func Test_BirthDate(t *testing.T) {

	dob, err := rsaid.DateOfBirth(validMale)

	if dob.Year() != 1995 || err != nil {
		t.Errorf("Does not determine date of birth correctly")
	}
	if dob.Month() != 6 || err != nil {
		t.Errorf("Does not determine date of birth correctly")
	}
	if dob.Day() != 24 || err != nil {
		t.Errorf("Does not determine date of birth correctly")
	}

	_, dob_err := rsaid.DateOfBirth(invalidDOB)
	if dob_err == nil {
		t.Errorf("Does not determine date of birth correctly")
	}
}

func Test_Parse(t *testing.T) {

	person, err := rsaid.Parse(validMale)

	if err != nil {
		t.Errorf("Does not parse id number correctly: %s", err.Error())
	}
	if person.Gender != rsaid.Male {
		t.Errorf("Does not parse gender correctly")
	}
	if person.Citizen != true {
		t.Errorf("Does not parse citizenship correctly")
	}
	if person.DOB.Year() != 1995 || person.DOB.Month() != 6 || person.DOB.Day() != 24 {
		t.Errorf("Does not parse birth date correctly")
	}
}
