package scheduler

import "math/rand"

func RandomGoodMorningImage() string {
	return morningImages[rand.Int()%len(morningImages)]
}

func RandomGoodMorningText() string {
	return morningTexts[rand.Int()%len(morningTexts)]
}

// TODO: get texts on http request

var morningTexts = [...]string{
	"Знаю, что хорошо потусила вчера 😏\nНу а теперь с добрым утром! 😘",
}

var morningImages = [...]string{
	"https://khlopot.net/wp-content/uploads/2019/09/generalnaya-uborka-doma-3.jpg",
}
