package hdea

type FitnessNode struct {
	left *FitnessNode
	right *FitnessNode
	parent *FitnessNode
	selfId int32
	parentId int32

	individuum *Individuum

	noVisit float64
	dimension int32
	threshold float64
}

type FitnessTree struct {
	root *FitnessNode
	currentId int32

	noDimension int32
	noLeaf int32
	interval []float64

	maxArchiveSize int32
	noAxisBoundary int32
	AxisBoundary_Set []int32

	curInterval *[][]float64;

	noRevisit int32

	noVisit int32
	maxNoVisit int32
}

func NewFitnessNode(parent *FitnessNode) (*FitnessNode) {
	node := new(FitnessNode)
	node.left = nil
	node.right = nil
	node.parent = parent

	return node
}
