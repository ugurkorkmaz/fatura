package fatura

// Earsiv portal paths for production and test environments
type Path struct {
	// debug mode on/off
	Debug bool
}

var _portal = map[bool]string{
	true:  "https://earsivportaltest.efatura.gov.tr",
	false: "https://earsivportal.efatura.gov.tr",
}

// Returns the dispatch path
func (ap *Path) Dispatch() string {
	return _portal[ap.Debug] + "/earsiv-services/dispatch"
}

// Returns the download path
func (ap *Path) Download() string {
	return _portal[ap.Debug] + "/earsiv-services/download"
}

// Returns the login path
func (ap *Path) Login() string {
	return _portal[ap.Debug] + "/earsiv-services/assos-login"
}

// Returns the referrer path
func (ap *Path) Referrer() string {
	return _portal[ap.Debug] + "/intragiris.html"
}

// Returns the esign path
func (ap *Path) Esign() string {
	return _portal[ap.Debug] + "/earsiv-services/esign"
}
