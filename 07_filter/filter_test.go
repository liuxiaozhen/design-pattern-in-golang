package filter

import (
	"fmt"
	"testing"
)

var persons Persons

func init() {
	persons = append(persons, &Person{"Tom", "male", "signle"})
	persons = append(persons, &Person{"Lily", "female", "married"})
	persons = append(persons, &Person{"Laura", "female", "married"})
	persons = append(persons, &Person{"Juice", "male", "signle"})
	persons = append(persons, &Person{"Lucy", "male", "married"})
	persons = append(persons, &Person{"Diana", "female", "signle"})
}

func Test_Filter(t *testing.T) {
	maleCriteria := &CriteriaMale{}
	femaleCriteria := &CriteriaFemale{}
	signleCriteria := &CriteriaSingle{}
	fmt.Println("=== male ===")
	printPersons(maleCriteria.meetCriteria(persons))

	fmt.Println("=== female and signle===")
	signleFemale := NewAndCriteria(femaleCriteria, signleCriteria)
	printPersons(signleFemale.meetCriteria(persons))

	fmt.Println("=== female or signle===")
	signleOrFemale := NewOrCriteria(femaleCriteria, signleCriteria)
	printPersons(signleOrFemale.meetCriteria(persons))
}

func printPersons(persons Persons) {
	for _, v := range persons {
		fmt.Println(v)
	}
}
