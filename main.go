package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Fio      string `json:"name"`
	Age      int    `json:"age"`
	Job      string `json:"job"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
	//card
	//car num
	//mail

}

var nouns = []string{
	"Apple", "Banana", "Apricot", "Peach", "Melon", "Watermelon",
	"Strawberry", "Cranberry", "Avocado", "Raspberry", "Cherry", "Apple"}

var adj = []string{
	"Beautiful", "Joyful", "Comfortable", "Big", "Hot", "Great", "Powerful", "Ideal",
	"Strong", "Tall", "Small", "Smart", "Dear"}

var Job = []string{
	"Учитель", "Медсестра", "Инженер", "Менеджер", "Повар",
	"Доктор", "Художник", "Учёный", "Бухгалтер", "Архитектор",
	"Писатель", "Музыкант", "Фотограф", "Продавец", "Электрик",
	"Сантехник", "Механик", "Веб-разработчик", "Аналитик данных", "Руководитель проекта",
	"HR-специалист", "Маркетолог", "Исследователь", "Фармацевт",
	"Ветеринар", "Социальный работник", "Графический дизайнер", "UX/UI-дизайнер",
	"Финансовый аналитик", "Консультант", "Риелтор", "Бариста",
	"Су-шеф", "Копирайтер", "Организатор мероприятий", "Специалист по связям с общественностью",
	"Программист", "Дизайнер", "Менеджер по продажам", "Специалист по рекламе",
	"Журналист", "Блогер", "Флорист", "Клинический психолог", "Специалист по IT-безопасности",
	"Специалист по охране труда", "Аудитор", "Антикризисный менеджер", "Сомелье", "Массажист",
	"Автомеханик", "Сценарист", "Тренер", "Косметолог", "Экономист",
	"Арт-директор", "Политолог", "Туристический агент", "Секретарь", "Специалист по автосервису",
	"Кастинг-директор", "Археолог", "Психотерапевт",
}

var letters = []string{
	"A", "B", "C", "D", "E", "F", "G",
	"H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T", "U",
	"V", "W", "X", "Y", "Z",
	"a", "b", "c", "d", "e", "f", "g",
	"h", "i", "j", "k", "l", "m", "n",
	"o", "p", "q", "r", "s", "t", "u",
	"v", "w", "x", "y", "z",
}

var numbers = []string{
	"0", "1", "2", "3", "4",
	"5", "6", "7", "8", "9",
}

var femaleNames = []string{
	"Анна", "Екатерина", "Мария", "София", "Татьяна",
	"Дарья", "Анастасия", "Вероника", "Наталья", "Ксения",
	"Юлия", "Полина", "Арина", "Людмила", "Валерия",
	"Ирина", "Алиса", "Ольга", "Кристина", "Светлана",
	"Галина", "Эмилия", "Милена", "Тамара", "Зоя",
	"Елизавета", "Лилия", "Диана", "Виктория", "Серафима",
	"Нина", "Ксении", "Лариса", "Грета", "Сабина",
}

var maleNames = []string{
	"Максим", "Антон", "Игорь", "Дмитрий", "Сергей",
	"Илья", "Владимир", "Евгений", "Александр", "Роман",
	"Анатолий", "Никита", "Даниил", "Иван", "Михаил",
	"Станислав", "Василий", "Кирилл", "Павел", "Юрий",
	"Степан", "Григорий", "Алексей", "Леонид", "Арсений",
	"Денис", "Константин", "Филипп", "Егор", "Семен",
}

var femaleSurnames = []string{
	"Смирнова", "Иванова", "Кузнецова", "Попова", "Соколова",
	"Михайлова", "Новикова", "Федорова", "Морозова", "Волкова",
	"Лебедева", "Алексеева", "Петрова", "Сорокина", "Борисова",
	"Киселева", "Зайцева", "Орехова", "Филиппова", "Семенова",
	"Григорьева", "Тихонова", "Масленникова", "Абрамова", "Наумова",
	"Ульянова", "Богданова", "Терешева", "Ковалева", "Булгакова",
}

var maleSurnames = []string{
	"Смирнов", "Иванов", "Кузнецов", "Попов", "Соколов",
	"Михайлов", "Новиков", "Федоров", "Морозов", "Волков",
	"Лебедев", "Алексеев", "Петров", "Сорокин", "Борисов",
	"Киселев", "Зайцев", "Орехов", "Филиппов", "Семенов",
	"Григорьев", "Тихонов", "Масленников", "Абрамов", "Наумов",
	"Ульянов", "Богданов", "Терешев", "Ковалев", "Булгаков",
}

var femalePatronymics = []string{
	"Ивановна", "Алексеевна", "Дмитриевна", "Сергеевна", "Петровна",
	"Максимовна", "Денисовна", "Станиславовна", "Ильинична", "Анатольевна",
	"Владимировна", "Егоровна", "Александровна", "Васильевна", "Григорьевна",
	"Кирилловна", "Ярославовна", "Антоновна", "Филипповна", "Леонидовна",
}

var malePatronymics = []string{
	"Иванович", "Алексеевич", "Дмитриевич", "Сергеевич", "Петрович",
	"Максимович", "Денисович", "Станиславович", "Ильинич", "Анатольевич",
	"Владимирович", "Егорович", "Александрович", "Васильевич", "Григорьевич",
	"Кириллович", "Ярославович", "Антонович", "Филиппович", "Леонидович",
}

func generateRandomAge(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func genereePhone() string {
	var sl []string
	sl = append(sl, "7927")

	for n := 0; n < 6; n++ {
		sl = append(sl, strconv.Itoa(rand.Intn(10)))
	}

	phone := strings.Join(sl, "")
	return phone
}

func generateRandomPassword(length int) string {
	password := ""
	for i := 0; i < length; i++ {
		if rand.Intn(2) == 0 {
			index := rand.Intn(len(letters))
			password += letters[index]
		} else {
			index := rand.Intn(len(numbers))
			password += numbers[index]
		}
	}
	return password
}

func generateFullName(isFemale bool) string {
	if isFemale {
		firstName := femaleNames[rand.Intn(len(femaleNames))]
		surname := femaleSurnames[rand.Intn(len(femaleSurnames))]
		patronymic := femalePatronymics[rand.Intn(len(femalePatronymics))]
		return surname + " " + firstName + " " + patronymic
	} else {
		firstName := maleNames[rand.Intn(len(maleNames))]
		surname := maleSurnames[rand.Intn(len(maleSurnames))]
		patronymic := malePatronymics[rand.Intn(len(malePatronymics))]
		return surname + " " + firstName + " " + patronymic
	}
}

func data(c *gin.Context) {
	isFemale := rand.Intn(2) == 0
	fullName := generateFullName(isFemale)

	newTodo := Todo{
		Username: adj[rand.Intn(len(adj))] + "_" + nouns[rand.Intn(len(nouns))],
		Password: generateRandomPassword(12),
		Job:      Job[rand.Intn(len(Job))],
		Age:      generateRandomAge(18, 45),
		Fio:      fullName,
		Phone:    genereePhone(),
	}

	c.JSON(http.StatusOK, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/data", data)
	router.Run(":8080")
}
