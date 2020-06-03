package lenconv

// CMToM converts a Centimeter to Meter.
func CMToM(c Centimeter) Meter { return Meter(c / 1000.0) }

// MToCM converts a Meter to CCentimeter.
func MToCM(m Meter) Centimeter { return Centimeter(m * 1000) }
