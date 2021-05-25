package servertools

import (
	"fmt"
	"net/http"
)

func GenerateHtmlHeader(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/html, charset=UTF-8")

	fmt.Fprintln(w, "<!DOCTYPE html>")
	fmt.Fprintln(w, "<html lang=\"en\">")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")
}

func GenerateHtmlFooter(w http.ResponseWriter) {
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")

}

// func GenerateLoginFormHtml(w http.ResponseWriter, showLoginFailed bool) {
// 	if showLoginFailed {
// 		fmt.Fprintln(w, "<span class=\"error\">Login failed</span><br /><br />")
// 	}
// 	fmt.Fprintln(w, "<form method=\"post\" id=\"loginForm\" role=\"form\" data-qa=\"login-form\">")
// 	fmt.Fprintln(w, "<h2>Login</h2>")
// 	fmt.Fprintln(w, "Username:<br />")
// 	fmt.Fprintln(w, "<input type=\"text\" name=\"username\" data-qa=\"login-username\" /><br /><br />")
// 	fmt.Fprintln(w, "Password:<br />")
// 	fmt.Fprintln(w, "<input type=\"password\" name=\"password\" data-qa=\"login-password\" /><br /><br />")
// 	fmt.Fprintln(w, "<button id=\"btn-login\" type=\"submit\" data-qa=\"login-button\">Log In</button>")
// 	fmt.Fprintln(w, "</form>")
// }

// func GenerateAllowDenyAppHtml(w http.ResponseWriter, clientId string, scope string, state string, redirectUri string, devApp *dto.DeveloperModel) {
// 	fmt.Fprintln(w, "<form method=\"post\" id=\"validateForm\" role=\"form\" data-qa=\"validate-form\">")
// 	fmt.Fprintf(w, "<h2>%s</h2>\n", devApp.AppName)
// 	fmt.Fprintf(w, "%s <br />", devApp.ProjectUrl)
// 	fmt.Fprintf(w, "This app would like to have the following scope: %s<br /> \n", scope)
// 	// TODO: enumerate through the scopes
// 	fmt.Fprintln(w, "<button id=\"btn-deny\" type=\"submit\" data-qa=\"oauth-deny-button\">Deny</button>")
// 	fmt.Fprintln(w, "<button id=\"btn-allow\" type=\"submit\" data-qa=\"oauth-allow-button\">Allow</button>")
// 	fmt.Fprintln(w, "</form>")
// }
