package filter

type Person struct {
	Name          string
	Sex           string
	MaritalStatus string
}
type Persons []*Person
type Criteria interface {
	meetCriteria(Persons) Persons
}

type CriteriaFemale struct {
}

func (female *CriteriaFemale) meetCriteria(p Persons) (r Persons) {
	for _, v := range p {
		if v.Sex == "female" {
			r = append(r, v)
		}
	}
	return
}

type CriteriaMale struct {
}

func (male *CriteriaMale) meetCriteria(p Persons) (r Persons) {
	for _, v := range p {
		if v.Sex == "male" {
			r = append(r, v)
		}
	}
	return
}

type CriteriaSingle struct {
}

func (s *CriteriaSingle) meetCriteria(p Persons) (r Persons) {
	for _, v := range p {
		if v.MaritalStatus == "signle" {
			r = append(r, v)
		}
	}
	return
}

type AndCriteria struct {
	One     Criteria
	Another Criteria
}

func (and *AndCriteria) meetCriteria(persons Persons) (n Persons) {
	n = and.One.meetCriteria(persons)
	n = and.Another.meetCriteria(n)
	return n
}

func NewAndCriteria(one Criteria, other Criteria) *AndCriteria {
	return &AndCriteria{
		one,
		other,
	}
}

type OrCriteria struct {
	One     Criteria
	Another Criteria
}

func (or *OrCriteria) meetCriteria(persons Persons) (n Persons) {
	one := or.One.meetCriteria(persons)
	another := or.Another.meetCriteria(persons)
	n = append(n, one...)
	for _, v1 := range another {
		contains := false
		for _, v2 := range one {
			if v1 == v2 { //v1,v2struct的指针类型, 可以用==, 判断是否指向同一个对象
				contains = true
				break
			}
		}
		if contains == false {
			n = append(n, v1)
		}
	}
	return
}

func NewOrCriteria(one Criteria, other Criteria) *OrCriteria {
	return &OrCriteria{
		one,
		other,
	}
}
