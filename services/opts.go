package opts

import (
	"github.com/ibtehalahmed/Go-Challenge/models"
)

type RulSet struct {
	ruleSet    models.RuleSet
}

//interface that implements AddDep() so i can call addDep like this ruleSet.addDep()
type IRulSet interface {
	AddDep(option1 string, option2 string) (ruleSet models.RuleSet)
}
// initialize new Rule
func NewRuleSet() *RulSet {
	var ruleSet models.RuleSet
	ruleSet.Id = "1"
	r := RulSet{}
	r.ruleSet = ruleSet
	return &r
}
//adding dependency for a given r of type *RulSet
func (r *RulSet) AddDep(option1 string, option2 string) (ruleSet models.RuleSet) {
	r.ruleSet.Options = make([]string, 2)

	r.ruleSet.Relation = "Dependency"
	r.ruleSet.Options[0] = option1
	r.ruleSet.Options[1] = option2

	return  r.ruleSet
}
// check for relationship type
func (r *RulSet) IsCoherent() bool{
	if r.ruleSet.Relation == "Dependency" && r.ruleSet.Options[0] == r.ruleSet.Options[1]{
		return true
	}else {
		return  false
	}
}

