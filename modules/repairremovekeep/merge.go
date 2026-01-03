package repairremovekeep

func (ir *ItemRemove) Merge(other ItemRemove) {
	ir.Items = append(ir.Items, other.Items...)
}

func (ik *ItemKeep) Merge(other ItemKeep) {
	ik.Items = append(ik.Items, other.Items...)
}

func (ak *ArmorKeep) Merge(other ArmorKeep) {
	ak.Items = append(ak.Items, other.Items...)
}

func (tr *ToolRepair) Merge(other ToolRepair) {
	tr.Tools = append(tr.Tools, other.Tools...)
}
