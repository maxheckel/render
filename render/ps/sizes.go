package ps

import (
	"render/render/render"
	"strings"
)

func Thumbnail() render.Size {
	return render.Size{
		Width:  200,
		Height: 200,
	}
}

func Standard() render.Size {
	return render.Size{
		Width:  484,
		Height: 484,
	}
}

func Large() render.Size {
	return render.Size{
		Width:  550,
		Height: 550,
	}
}

func Giant() render.Size {
	return render.Size{
		Width:  2000,
		Height: 2000,
	}
}

func StringToSize(size string) render.Size {
	size = strings.ToLower(size)
	switch size {
	case "thumbnail", "thumb":
		return Thumbnail()
	case "standard":
		return Standard()
	case "large":
		return Large()
	case "giant":
		return Giant()
	default:
		return Thumbnail()
	}
}

