package oauth

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type MyOauth struct {
	Config             *oauth2.Config
	GenerateStateToken func() string
}

func (o MyOauth) MyLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := o.GenerateStateToken()

	// Store state in session handler
	fmt.Println(fmt.Errorf("Store the state where?"))

	url := o.Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func GenerateStateToken() string {
	return "testtest"
}
