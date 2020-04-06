package composite

import "testing"

func Test_Composite(t *testing.T) {
	ceo := NewComposite("CEO")
	cho := NewComposite("CHO")
	hr1 := NewComposite("Hr Team 1")
	hr2 := NewComposite("Hr Team 2")
	cho.Add(hr1)
	cho.Add(hr2)

	cto := NewComposite("CTO")
	tech1 := NewComposite("Tech Team 1")
	tech2 := NewComposite("Tech Team 2")
	tech3 := NewComposite("Tech Team 3")
	tech4 := NewComposite("Tech Team 4")

	cto.Add(tech1)
	cto.Add(tech2)
	cto.Add(tech3)
	cto.Add(tech4)

	cfo := NewComposite("CFO")

	ceo.Add(cfo)
	ceo.Add(cho)
	ceo.Add(cto)
	ceo.Display(0)

	cto.Remove(tech3)
	cto.Display(1)
}
