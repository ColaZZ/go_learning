package tempconv

// CTOF 把摄氏度转换为华氏温度
func CTOF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC 把华氏温度转换为摄氏温度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC 把开尔文温度转换为设施温度
func KToC(k Kelvins) Celsius {
	return Celsius(k - 273.15)
}

// CToK 把设施温度转换为开尔文温度
func CToK(c Celsius) Kelvins {
	return Kelvins(c + 273.15)
}
