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
	"Самой милой красотке самого доброго утра 😎",
	"Хәерле иртэ!",
	"Доброе утро, Алсушка ^^ 💋💋💋",
	"Кто это у нас еще спит? Или не спит...\nДоброе утро!!!!!",
	"Опа, Алсундекс, это ты? Доброе утро! 💋🏄",
}

var morningImages = [...]string{
	"https://i.mycdn.me/i?r=AyH4iRPQ2q0otWIFepML2LxRjFo4tAvRWweiiLE_CLPYA",
	"https://i.pinimg.com/originals/1d/6f/c1/1d6fc18e48aab88f978a30da6592b3d4.jpg",
	"https://i.pinimg.com/564x/8e/09/31/8e0931c73721302a67b0f521ae936314.jpg",
	"https://photowords2.ru/pics_max/photowords_ru_8957.jpg",
	"https://klike.net/uploads/posts/202111/1637049855_3.jpg",
}
