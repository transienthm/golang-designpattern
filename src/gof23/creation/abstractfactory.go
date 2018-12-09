package creation

type PlantFactory interface {
	MakePlant() Plant
	MakePicker() Picker
}

type Plant interface {
}
type Picker interface {
}

type PearFactory struct {
}

type Pear struct {
	desc string
}

type PearPicker struct {
	desc string
}


func (this PearFactory) MakePlant() Plant {
	return Pear{"我是一个梨"}
}

func (this PearFactory) MakePicker() Picker {
	return PearPicker{"采摘一个梨"}
}