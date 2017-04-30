package models

import "gopkg.in/mgo.v2/bson"

type RuleSet struct  {
	Id bson.ObjectId
	Relation string
	Options  []string
}
