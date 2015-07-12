package db

import(
	r "github.com/dancannon/gorethink"
	"../config"
	"log"
)

var Session *r.Session

func Connect() {
	session, err := r.Connect(r.ConnectOpts{
	    Address: config.Get().RethinkDB.Address,
	    Database: config.Get().RethinkDB.Database,
	    MaxIdle: 10,
	    MaxOpen: 10,
	})
	Session = session
	if err != nil {
	    log.Fatalln(err.Error())
	}
}


