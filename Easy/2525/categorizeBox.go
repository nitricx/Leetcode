package categorizeBox

func categorizeBox(length int, width int, height int, mass int) string {
	massLimit := 100
	bulkyLimit := 10000
	volumeLimit := 1000000000
	bulkyFlag := false
	heavyFlag := false
	if length >= bulkyLimit || width >= bulkyLimit || height >= bulkyLimit || (length*width*height) >= volumeLimit {
		bulkyFlag = true
	}
	if mass >= massLimit {
		heavyFlag = true
	}
	if bulkyFlag && heavyFlag {
		return "Both"
	} else if bulkyFlag {
		return "Bulky"
	} else if heavyFlag {
		return "Heavy"
	}
	return "Neither"
}