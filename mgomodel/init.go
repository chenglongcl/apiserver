package mgomodel

import (
	"github.com/chenglongcl/log"
	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
)

type GlobalSessions struct {
	self   *mgo.Session
	docker *mgo.Session
}

var globalSessions *GlobalSessions

func Init() {
	globalSessions = &GlobalSessions{
		self:   initDB("mongo_db"),
		docker: initDB("mongo_docker_db"),
	}
}

func initDB(name string) *mgo.Session {
	if name == "" {
		return nil
	}
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{viper.GetString(name + ".addr")},
		Source:   viper.GetString(name + ".source"),
		Username: viper.GetString(name + ".username"),
		Password: viper.GetString(name + ".password"),
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Errorf(err, "create %s mgo session error", name)
	}
	return session
}

func Close() {
	globalSessions.self.Close()
	globalSessions.docker.Close()
}

func GetSession(hostName string) *DefaultSession {
	var globalS *mgo.Session
	switch hostName {
	case "self":
		globalS = globalSessions.self
	case "docker":
		globalS = globalSessions.docker
	}
	if globalS == nil {
		panic("mongodb session is nil, go to connect mongodb first")
	}
	return &DefaultSession{
		newSession: globalS.Copy(),
	}
}

type DefaultSession struct {
	newSession *mgo.Session
	db         *mgo.Database
	collection *mgo.Collection
}

func (a *DefaultSession) DB(db string) *DefaultSession {
	if a.newSession == nil {
		panic("mongodb copy session is nil, go to connect mongodb first")
	}
	a.db = a.newSession.DB(db)
	return a
}

func (a *DefaultSession) Collection(collection string) *DefaultSession {
	if a.db == nil {
		panic("use GetSession(hostname).DB(db).Collection(collection)")
	}
	a.collection = a.db.C(collection)
	return a
}

func (a *DefaultSession) IsExist(query interface{}) bool {
	defer a.newSession.Close()
	count, _ := a.collection.Find(query).Count()
	return count > 0
}

func (a *DefaultSession) Count(query interface{}) (int, error) {
	defer a.newSession.Close()
	return a.collection.Find(query).Count()
}

func (a *DefaultSession) Insert(docs ...interface{}) error {
	defer a.newSession.Close()
	return a.collection.Insert(docs...)
}

func (a *DefaultSession) FindOne(query, selector, result interface{}) error {
	defer a.newSession.Close()
	return a.collection.Find(query).Select(selector).One(result)
}

func (a *DefaultSession) FindAll(collection string, query, selector, result interface{}) error {
	defer a.newSession.Close()
	return a.collection.Find(query).Select(selector).All(result)
}

func (a *DefaultSession) FindList(Offset, limit int, query, selector, result interface{}, sort ...string) error {
	defer a.newSession.Close()
	return a.collection.Find(query).Select(selector).Skip(Offset).Limit(limit).Sort(sort...).All(result)
}

func (a *DefaultSession) FindIter(query interface{}) *mgo.Iter {
	defer a.newSession.Close()
	return a.collection.Find(query).Iter()
}

func (a *DefaultSession) Update(selector, update interface{}) error {
	defer a.newSession.Close()
	return a.collection.Update(selector, update)
}

func (a *DefaultSession) Upsert(selector, update interface{}) error {
	defer a.newSession.Close()
	_, err := a.collection.Upsert(selector, update)
	return err
}

func (a *DefaultSession) UpdateAll(selector, update interface{}) error {
	defer a.newSession.Close()
	_, err := a.collection.UpdateAll(selector, update)
	return err
}

func (a *DefaultSession) Remove(selector interface{}) error {
	defer a.newSession.Close()
	return a.collection.Remove(selector)
}

func (a *DefaultSession) RemoveAll(selector interface{}) error {
	defer a.newSession.Close()
	_, err := a.collection.RemoveAll(selector)
	return err
}

//insert one or multi documents
func (a *DefaultSession) BulkInsert(docs ...interface{}) (*mgo.BulkResult, error) {
	defer a.newSession.Close()
	bulk := a.collection.Bulk()
	bulk.Insert(docs...)
	return bulk.Run()
}

func (a *DefaultSession) BulkRemove(selector ...interface{}) (*mgo.BulkResult, error) {
	defer a.newSession.Close()
	bulk := a.collection.Bulk()
	bulk.Remove(selector...)
	return bulk.Run()
}

func (a *DefaultSession) BulkRemoveAll(selector ...interface{}) (*mgo.BulkResult, error) {
	defer a.newSession.Close()
	bulk := a.collection.Bulk()
	bulk.RemoveAll(selector...)
	return bulk.Run()
}

func (a *DefaultSession) BulkUpdate(pairs ...interface{}) (*mgo.BulkResult, error) {
	defer a.newSession.Close()
	bulk := a.collection.Bulk()
	bulk.Update(pairs...)
	return bulk.Run()
}

func (a *DefaultSession) BulkUpdateAll(pairs ...interface{}) (*mgo.BulkResult, error) {
	defer a.newSession.Close()
	bulk := a.collection.Bulk()
	bulk.UpdateAll(pairs...)
	return bulk.Run()
}

func (a *DefaultSession) BulkUpsert(pairs ...interface{}) (*mgo.BulkResult, error) {
	defer a.newSession.Close()
	bulk := a.collection.Bulk()
	bulk.Upsert(pairs...)
	return bulk.Run()
}
