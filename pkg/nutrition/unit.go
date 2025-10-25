package nutrition

type MeasurementUnit uint16

const (
	InvalidUnit MeasurementUnit = iota
	Gramm
	Milliliter
	Unit
)

func NewMeasurementUnit(iunit int32) MeasurementUnit {
	switch iunit {
	case 1:
		return Gramm
	case 2:
		return Milliliter
	case 3:
		return Unit
	default:
		return InvalidUnit
	}
}

func (unit MeasurementUnit) String() string {
	switch unit {
	case Gramm:
		return "Gramm"
	case Milliliter:
		return "Milliliter"
	case Unit:
		return "Unit"
	default:
		return "InvalidUnit"
	}
}
