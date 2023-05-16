package config

import (
	"net/http"

	controllers "github.com/Totoies/Totoies/src/Application/Controllers"
)

/*
Perpose - Have the details of all the roots
*/

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/": controllers.Home.Load,
}
