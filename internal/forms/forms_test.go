package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}

}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "A")
	postedData.Add("b", "B")
	postedData.Add("c", "C")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	IsValid := form.Has("a")
	if IsValid {
		t.Error("forms shows has field a when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "A")

	//r, _ = http.NewRequest("POST", "/whatever", nil)

	//r.PostForm = postedData
	form = New(postedData)
	IsValid = form.Has("a")
	if !IsValid {
		t.Error("shows form does not have a field a when it should")
	}

}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	//first test
	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	//second test
	postedData := url.Values{}
	postedData.Add("first_name", "John")
	form = New(postedData)

	IsValid := form.MinLength("first_name", 3)
	if !IsValid {
		t.Errorf("John should have passed MinLength. %s", form.Errors)
	}

	isError = form.Errors.Get("first_name")
	if isError != "" {
		t.Error("should not have an error but got one")
	}

	//third test
	postedData = url.Values{}
	postedData.Add("first_name", "aa")
	form = New(postedData)

	IsValid = form.MinLength("first_name", 3)
	if IsValid {
		t.Error("aa should have failed MinLength")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	//first test
	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	//second test
	email := "abc"
	postedData = url.Values{}
	postedData.Add("email", email)
	form = New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Errorf("%s is not a valid email but is shown as valid", email)
	}

	//third test
	email = "abc@test.com"
	postedData = url.Values{}
	postedData.Add("email", email)
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Errorf("%s is a valid email but is shown as invalid", email)
	}
}
