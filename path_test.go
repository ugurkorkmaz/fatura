package fatura_test

import (
	"testing"

	"github.com/ugurkorkmaz/fatura"
)

func TestDispatch(t *testing.T) {
	// create a new instance of Path with debug mode off
	p := fatura.Path{Debug: false}
	expected := "https://earsivportal.efatura.gov.tr/earsiv-services/dispatch"
	if p.Dispatch() != expected {
		t.Errorf("Dispatch() returned unexpected URL. Expected: %s, got: %s", expected, p.Dispatch())
	}

	// change debug mode to on and test again
	p.Debug = true
	expected = "https://earsivportaltest.efatura.gov.tr/earsiv-services/dispatch"
	if p.Dispatch() != expected {
		t.Errorf("Dispatch() returned unexpected URL. Expected: %s, got: %s", expected, p.Dispatch())
	}
}

func TestDownload(t *testing.T) {
	// create a new instance of Path with debug mode off
	p := fatura.Path{Debug: false}
	expected := "https://earsivportal.efatura.gov.tr/earsiv-services/download"
	if p.Download() != expected {
		t.Errorf("Download() returned unexpected URL. Expected: %s, got: %s", expected, p.Download())
	}

	// change debug mode to on and test again
	p.Debug = true
	expected = "https://earsivportaltest.efatura.gov.tr/earsiv-services/download"
	if p.Download() != expected {
		t.Errorf("Download() returned unexpected URL. Expected: %s, got: %s", expected, p.Download())
	}
}

func TestLogin(t *testing.T) {
	// create a new instance of Path with debug mode off
	p := fatura.Path{Debug: false}
	expected := "https://earsivportal.efatura.gov.tr/earsiv-services/assos-login"
	if p.Login() != expected {
		t.Errorf("Login() returned unexpected URL. Expected: %s, got: %s", expected, p.Login())
	}

	// change debug mode to on and test again
	p.Debug = true
	expected = "https://earsivportaltest.efatura.gov.tr/earsiv-services/assos-login"
	if p.Login() != expected {
		t.Errorf("Login() returned unexpected URL. Expected: %s, got: %s", expected, p.Login())
	}
}

func TestReferrer(t *testing.T) {
	// create a new instance of Path with debug mode off
	p := fatura.Path{Debug: false}
	expected := "https://earsivportal.efatura.gov.tr/intragiris.html"
	if p.Referrer() != expected {
		t.Errorf("Referrer() returned unexpected URL. Expected: %s, got: %s", expected, p.Referrer())
	}

	// change debug mode to on and test again
	p.Debug = true
	expected = "https://earsivportaltest.efatura.gov.tr/intragiris.html"
	if p.Referrer() != expected {
		t.Errorf("Referrer() returned unexpected URL. Expected: %s, got: %s", expected, p.Referrer())
	}
}

func TestEsign(t *testing.T) {
	// create a new instance of Path with debug mode off
	p := fatura.Path{Debug: false}
	expected := "https://earsivportal.efatura.gov.tr/earsiv-services/esign"
	if p.Esign() != expected {
		t.Errorf("Esign() returned unexpected URL. Expected: %s, got: %s", expected, p.Esign())
	}

	// change debug mode to on and test again
	p.Debug = true
	expected = "https://earsivportaltest.efatura.gov.tr/earsiv-services/esign"
	if p.Esign() != expected {
		t.Errorf("Esign() returned unexpected URL. Expected: %s, got: %s", expected, p.Esign())
	}
}
