package strategy

type Strategy interface {
	doOperation(int, int) int
}

type StrategyContext struct {
	num1     int
	num2     int
	strategy Strategy
}

func (s *StrategyContext) executeStrategy() int {
	return s.strategy.doOperation(s.num1, s.num2)
}
func NewStrategyContext(num1 int, num2 int, strategy Strategy) *StrategyContext {
	return &StrategyContext{
		num1,
		num2,
		strategy,
	}
}

type AddStrategy struct {
}

func (s *AddStrategy) doOperation(n1, n2 int) int {
	return n1 + n2
}

type SubStrategy struct {
}

func (s *SubStrategy) doOperation(n1, n2 int) int {
	return n1 - n2
}

type MulStrategy struct {
}

func (s *MulStrategy) doOperation(n1, n2 int) int {
	return n1 * n2
}
