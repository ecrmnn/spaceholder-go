package providers

import (
	"strings"
)

var (
	m = map[string]func(string) string{
		"loremPixel": loremPixel,
		"unsplash": unsplash,
		"dummyImage": dummyImage,
		"fakeImg": fakeImg,
		"placeImg": placeImg,
		"placeholdIt": placeholdIt,
	}
)

func GetImageFromProvider(provider, size string) string {
	if (provider == "random") {
		provider = randomProvider()
	}

	return m[provider](size)
}

func randomProvider() string {
	randomProvider := ""
	for key := range m {
		randomProvider = key
	}

	return randomProvider
}

func loremPixel(size string) string {
	pieces := strings.Split(size, "x")

	return "http://lorempixel.com/" + pieces[0] + "/" + pieces[1] + "/"
}

func unsplash(size string) string {
	pieces := strings.Split(size, "x")

	return "https://unsplash.it/" + pieces[0] + "/" + pieces[1];
}

func dummyImage(size string) string {
	return "https://dummyimage.com/" + size + "/000/fff"
}

func fakeImg(size string) string {
	return "https://fakeimg.pl/" + size + "/384f66/ecf0f1/?text=Spaceholder&font=lobster"
}

func placeImg(size string) string {
	pieces := strings.Split(size, "x")

	return "http://placeimg.com/" + pieces[0] + "/" + pieces[1] + "/any"
}

func placeholdIt(size string) string {
	pieces := strings.Split(size, "x")

	return "http://placeholdit.imgix.net/~text?txtsize=40&txt=Spaceholder&w=" + pieces[0] + "&h=" + pieces[1]
}