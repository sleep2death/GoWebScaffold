package server

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	UsersCollection    = "users"
	SessionsCollection = "sessions"

	DBNotFound = "not found"

	DBIndexKey = "username"
)

var (
	dbname    string
	dbsession *mgo.Session
)

func ConnectDB(name string, usr string, pwd string, port int) error {
	dial := fmt.Sprintf("%s:%s@localhost:%d/%s", usr, pwd, port, name)
	dbname = name

	session, err := mgo.Dial(dial)
	if err != nil {
		return err
	}

	//try to ping to the DB
	err = session.Ping()
	if err != nil {
		return err
	}

	//session.SetMode(mgo.Monotonic, true)

	index := mgo.Index{
		Key:    []string{DBIndexKey},
		Unique: true,
	}

	err = session.DB(dbname).C(UsersCollection).EnsureIndex(index)
	if err != nil {
		return err
	}

	dbsession = session

	return nil
}

func FindUserExistedByMail(mail string) (n int, err error) {
	s := dbsession.Copy()
	defer s.Close()

	n, err = s.DB(dbname).C(UsersCollection).Find(bson.M{DBIndexKey: mail}).Count()
	return
}

func SaveUser(user *User) error {
	s := dbsession.Copy()
	defer s.Close()

	_, err := s.DB(dbname).C(UsersCollection).Upsert(bson.M{DBIndexKey: user.Email}, bson.M{"$set": user})
	return err
}
