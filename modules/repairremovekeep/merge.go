package repairremovekeep

func (ir *ItemRemove) Merge(other ItemRemove) {
	ir.Items = append(ir.Items, other.Items...)
	ir.If = append(ir.If, other.If...)
	ir.Unless = append(ir.Unless, other.Unless...)
}

func (ik *ItemKeep) Merge(other ItemKeep) {
	ik.Items = append(ik.Items, other.Items...)
	ik.If = append(ik.If, other.If...)
	ik.Unless = append(ik.Unless, other.Unless...)
}

func (ak *ArmorKeep) Merge(other ArmorKeep) {
	ak.Items = append(ak.Items, other.Items...)
	ak.If = append(ak.If, other.If...)
	ak.Unless = append(ak.Unless, other.Unless...)
}

func (tr *ToolRepair) Merge(other ToolRepair) {
	tr.Tools = append(tr.Tools, other.Tools...)
	tr.If = append(tr.If, other.If...)
	tr.Unless = append(tr.Unless, other.Unless...)
}
