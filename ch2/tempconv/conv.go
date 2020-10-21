package tempconv

func CToF(c Celcius) Farenheit { return Farenheit(c*9/5 + 32) }
func FToC(f Farenheit) Celcius { return Celcius((f-32) * 5 / 9) }

