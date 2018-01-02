package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yohcop/openid-go"
)

type nilDiscoveryCache struct{}

func (n *nilDiscoveryCache) Put(id string, info openid.DiscoveredInfo) {}
func (n *nilDiscoveryCache) Get(id string) openid.DiscoveredInfo       { return nil }

type nilNonceStore struct{}

func (n *nilNonceStore) Accept(endpoint, nonce string) error { return nil }

func handler(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("../website/dist/"))
	if strings.Contains(r.URL.String(), ".") == false {
		r.URL.Path = "/"
	}
	fs.ServeHTTP(w, r)
}

func steamLoginHandler(w http.ResponseWriter, r *http.Request) {
	if url, err := openid.RedirectURL(
		"http://steamcommunity.com/openid",
		"http://"+webAddress+"/api/steamcallback",
		"http://"+webAddress); err == nil {
		http.Redirect(w, r, url, http.StatusSeeOther)
	} else {
		log.Printf("Failed to setup Redirect URL: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func steamCallbackHandler(w http.ResponseWriter, r *http.Request) {
	fullurl := "http://" + webAddress + r.URL.String()
	log.Println("CallBack:", fullurl)
	id, err := openid.Verify(fullurl, &nilDiscoveryCache{}, &nilNonceStore{})
	if err != nil {
		log.Println("error Verifying:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// we have confirmed the user has logged in
	// lets create a token for future requests and store
	// this token to match future requests.
	Claims := jwt.MapClaims{
		"steamid": id,
		// documentation about the use of exp for validation @
		// https://github.com/dgrijalva/jwt-go/blob/master/map_claims.go#L70
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims).SignedString(jwtsecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to create jwt, err: %s\n", err)
		return
	}
	s := `<head>
		<script>
			var persist = localStorage.getItem('persistlogin');
			console.log(persist);
			var token='` + token + `';
			console.log(token);
			if(persist==null) {
				sessionStorage.auth=token;
				localStorage.removeItem('auth');
			} else {
				localStorage.auth=token;
				sessionStorage.removeItem('auth');
			}
			window.location='http://` + webAddress + `/';
		</script>
	</head>`
	w.Write([]byte(s))

}
