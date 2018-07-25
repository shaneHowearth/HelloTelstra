package main

import (
	"net/http/httptest"
	"testing"
)

func TestReturnHello(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	HelloWorldHandler(w, req)

	// check "hello world" was sent back
	expected := "hello world"
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}

func TestWrongMethod(t *testing.T) {
	req := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	HelloWorldHandler(w, req)

	// check "hello world" was sent back
	expected := ""
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}

func TestWrongPath(t *testing.T) {
	req := httptest.NewRequest("GET", "/jkhgaf", nil)
	w := httptest.NewRecorder()
	HelloWorldHandler(w, req)

	// check "hello world" was sent back
	expected := ""
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}
