package main

import (
	"regexp"
	"strconv"
)

func (s *validate) Required() *validate {
	if s.r == true {
		if s.i == "" || len(s.i) <= 0 {
			s.r = false
			s.e = "Is Required"
		}
	}
	return s
}

func (s *validate) minLength(length int) *validate {
	if s.r == true {
		if len(s.i) < length {
			s.r = false
			s.e = "Minimum length " + strconv.Itoa(length) + " characters allowed"
		}
	}
	return s
}

func (s *validate) maxLength(length int) *validate {
	if s.r == true {
		if len(s.i) > length {
			s.r = false
			s.e = "Maximum length " + strconv.Itoa(length) + " characters allowed"
		}
	}
	return s
}

func (s *validate) isEmail() *validate {
	if s.r == true {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		var match = re.MatchString(s.i)
		if !match {
			s.r = false
			s.e = "invalid email"
		}
	}
	return s
}

func (s *validate) oneLowerCase() *validate {
	if s.r == true {
		re := regexp.MustCompile("[a-z]+")
		var match = re.MatchString(s.i)
		if !match {
			s.r = false
			s.e = "does not contain atleast lowercase letter"
		}
	}
	return s
}

func (s *validate) allLowerCase() *validate {
	if s.r == true {
		re := regexp.MustCompile("^[a-z]+$")
		var match = re.MatchString(s.i)
		if !match {
			s.r = false
			s.e = "does not contain lowercase letter"
		}
	}
	return s
}

func (s *validate) oneUpperCase() *validate {
	if s.r == true {
		re := regexp.MustCompile("[A-Z]+")
		var match = re.MatchString(s.i)
		if !match {
			s.r = false
			s.e = "does not contain atleast one uppercase letter"
		}
	}
	return s
}

func (s *validate) allUpperCase() *validate {
	if s.r == true {
		re := regexp.MustCompile("^[A-Z]+$")
		var match = re.MatchString(s.i)
		if !match {
			s.r = false
			s.e = "all letters are not uppercase"
		}
	}
	return s
}

func (s *validate) oneNumber() *validate {
	if s.r == true {
		re := regexp.MustCompile("[0-9]+")
		var match = re.MatchString(s.i)
		if !match {
			s.r = false
			s.e = "does not contain atleast one numeric character"
		}
	}
	return s
}

func (s *validate) isSpecialCharacter() *validate {
	if s.r == true {
		re := regexp.MustCompile("\\`|\\~|\\!|\\@|\\#|\\$|\\%|\\^|\\&|\\*|\\(|\\)|\\+|\\=|\\[|\\{|\\]|\\}|\\||\\|\\'|\\<|\\,|\\.|\\>|\\?|\\/|\"|\\;|\\:|\\s")
		var match = re.MatchString(s.i)
		if !match {
			s.r = false
			s.e = "does not contain atleast one special character"
		}
	}
	return s
}

func (s *validate) check() bool {
	return s.r
}
