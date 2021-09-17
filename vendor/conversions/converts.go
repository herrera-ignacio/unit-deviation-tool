package conversions

type Convert func(val float64) float64

func MetersToMiles(meters float64) float64 { return meters / MetersInMiles }

func MilesToMeters(miles float64) float64 { return miles * MetersInMiles }

func LitersToGallons(liters float64) float64 { return liters * GallonsInLiter }

func GallonsToLiters(gallons float64) float64 { return gallons / GallonsInLiter }
