package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Level struct {
	weapon int
	people int
	place  int
	room   int
	win    int
}

var name string

func generator(level *Level) {
	fmt.Println("Приветствуем на игре...")
	for level.weapon == 0 {
		level.weapon = random_digit(6)
		time.Sleep(1 * time.Second)
		fmt.Println("Генерируется оружие...")
	}
	for level.people == 0 {
		level.people = random_digit(6)
		time.Sleep(1 * time.Second)
		fmt.Println("Генерируются персонажи...")
	}
	for level.place == 0 {
		level.place = random_digit(6)
		time.Sleep(1 * time.Second)
		fmt.Println("Генерируются комнаты...")
	}
}

func init() {
	fmt.Println("Привет, меня зовут Goша, а тебя как?")
	fmt.Scan(&name)
}

func random_digit(n int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	var digit int = rand.Intn(n)
	return digit
}

func main() {
	var level Level
	generator(&level)
	defer goodbye(&level)
	Hi(name)
	loop_room(&level)
	loop_people(&level)
	loop_weapon(&level)
	check_answer(&level)
}

func Hi(name string) {
	var phrase string = "Я очень рад тебя видеть, " + name + "!\nМне нужна твоя помощь! Артур куда-то пропал, не могу его найти.\n Он нас позвал в гости, накрыл на стол и как будто бы растворился в воздухе!\n Поищи его, может тебе удасться его найти!"
	fmt.Println(phrase)
}

func goodbye(level *Level) {
	if r := recover(); r != nil {
		fmt.Println("Игра закончена:", r)
	}
	places := map[int]string{1: "Кухня", 2: "Ванная комната", 3: "Туалет", 4: "Спальня", 5: "Гостинная", 6: "Балкон"}
	peoples := map[int]string{1: "Goша", 2: "Артур", 3: "Мария", 4: "Алекс", 5: "Кира", 6: name}
	weapons := map[int]string{1: "Пистолет", 2: "Нож", 3: "Топор", 4: "Удавка", 5: "Молоток", 6: "Шило"}
	fmt.Println("1. Место преступления:")
	fmt.Println(places[level.place])
	fmt.Println("2. Убийца:")
	fmt.Println(peoples[level.people])
	fmt.Println("3. Орудие:")
	fmt.Println(weapons[level.weapon])
	fmt.Println("Спасибо за игру. Заходи ещё!", name)
	fmt.Println("-----------------------------------------")
	if level.win == 1 {
		fmt.Println("Поздравляю, вы победили!", name)
	} else {
		fmt.Println("К сожалению, вы проиграли!", name)
	}
	fmt.Println("-----------------------------------------")
}

func (l Level) info_room() {
	switch l.room {
	case 1:
		fmt.Println("Керамическая плита ещё теплая. Пахнет мясом и свежим хлебом.")
	case 2:
		fmt.Println("Зеркало запотело. Как будто кто-то совсем недавно принимал горячий душ.")
	case 3:
		fmt.Println("Начищен до блеска. Пахнет стиральным порошком.")
	case 4:
		fmt.Println("Жалюзи прикрыты. Кровать заправлена. Кажется, что время здесь остановилось.")
	case 5:
		fmt.Println("Два кожаных дивана и большой телевизор. Идёт что-то от Вес Андерсона.")
	case 6:
		fmt.Println("Обшит деревом. Ярко красный закат хорошо контрастирует со старым комплектом зимней резины.")
	default:
		fmt.Println("Неверное значение")
	}
}

func info_people(people int) {
	switch people {
	case 1:
		fmt.Println("Коренастый мужчина. На вид 30 лет. С ухоженной бородкой. Говорит с канадским акцентом.")
	case 2:
		fmt.Println("Чванливый англичанин. Любит коллекционировать марки, монеты и пистолеты. Курит трубку. Любит гольф.")
	case 3:
		fmt.Println("Эффектная блондинка на длинных каблуках в клетчатой юбке. Говорит с русским акцентом.\n На руках очаровательный рыжий шпиц, который смотрит влюблёнными глазами на свою хозяйку.")
	case 4:
		fmt.Println("Широкоплечий молодой нарцисс, который не упустит шанса покрасоваться мускулами перед зеркалом. Прилизанные волосы.\n Нарочито поставленный голос, как будто выступает на сцене.")
	case 5:
		fmt.Println("Брюнетка среднего роста. В одежде только сдержанные цвета. Волосы убраны. Тихий голос наполняет комнату лаконичными литературными выражениями.\n Образ дополняют изящные очки, едва заметные при тусклом свете.")
	default:
		fmt.Println("Об этом персонаже ничего не известно...")
	}
}

func info_weapon(weapon int) {
	switch weapon {
	case 1:
		fmt.Println("Раритетный кольт с серебряной каймой, справа небольшая царапина на стволе")
	case 2:
		fmt.Println("Леденящее острое лезвие из стали, рукоятка с обивкой из натуральной кожи")
	case 3:
		fmt.Println("Топорик из закалённой стали с деревянной рукоятью, с небольшими потертостями у основания")
	case 4:
		fmt.Println("Прочная веревка, небольшого диаметра")
	case 5:
		fmt.Println("Рабочий инструмент с небольшими следами ржавчины.")
	case 6:
		fmt.Println("Заострённая спица с рукоятью из клёна")
	default:
		fmt.Println("Неверное значение")
	}
}

func (m *Level) move(numRoom int) {
	m.room = numRoom
	m.info_room()
}

func loop_room(level *Level) {
	var room int
	places := map[int]string{1: "Кухня", 2: "Ванная комната", 3: "Туалет", 4: "Спальня", 5: "Гостинная", 6: "Балкон"}
	// peoples := map[int]string{1: "Goша", 2: "Артур", 3: "Мария", 4: "Алекс", 5: "Кира", 6: name}
	// weapons := map[int]string{1: "Пистолет", 2: "Нож", 3: "Топор", 4: "Удавка", 5: "Молоток", 6: "Шило"}
	fmt.Println("В какую комнату перейдем?")
	fmt.Println(places)
	fmt.Println("Введите 0, если хотите завершить игру")
	fmt.Scan(&room)
	for room != 0 {
		fmt.Println(places[room])
		level.move(room)
		if room == level.place {
			fmt.Println("На полу лежит тело. Артур, что с тобой? Пульса нет.")
			return
		}
		fmt.Scan(&room)
	}
	if room == 0 {
		panic("The End")
	}
}

func loop_people(level *Level) {
	var people int
	peoples := map[int]string{1: "Goша", 2: "Артур", 3: "Мария", 4: "Алекс", 5: "Кира", 6: name}
	fmt.Println("Все сюда, Артуру плохо! Гости прибежали на крик. Что-то здесь не так. Посмотрю я на всех внимательнее. О ком узнать подробнее? ")
	fmt.Println(peoples)
	fmt.Println("Введите 0, если хотите завершить игру")
	fmt.Println("Введите 7, если нашли убийцу")
	fmt.Scan(&people)
	for people != 0 && people != 7 {
		fmt.Println(peoples[people])
		info_people(people)
		if people == level.people {
			fmt.Println("Отводит взгляд в сторону...")
		}
		fmt.Scan(&people)
	}
	if people == 0 {
		panic("The End")
	}
}

func loop_weapon(level *Level) {
	var weapon int
	weapons := map[int]string{1: "Пистолет", 2: "Нож", 3: "Топор", 4: "Удавка", 5: "Молоток", 6: "Шило"}
	fmt.Println("Чем же могли убить Артура? ")
	fmt.Println(weapons)
	fmt.Println("Введите 0, если хотите завершить игру")
	fmt.Println("Введите 7, если знаете ответ, в какой комнате, кто и каким оружием убил Артура")
	fmt.Scan(&weapon)
	for weapon != 0 && weapon != 7 {
		fmt.Println(weapons[weapon])
		info_weapon(weapon)
		if weapon == level.weapon {
			fmt.Println("Пятна крови...")
		}
		fmt.Scan(&weapon)
		if weapon != 7 {
			fmt.Println("Введите 7, если знаете ответ, в какой комнате, кто и каким оружием убил Артура")
		}
	}
	if weapon == 0 {
		panic("The End")
	}
}

func check_answer(level *Level) {
	places := map[int]string{1: "Кухня", 2: "Ванная комната", 3: "Туалет", 4: "Спальня", 5: "Гостинная", 6: "Балкон"}
	peoples := map[int]string{1: "Goша", 2: "Артур", 3: "Мария", 4: "Алекс", 5: "Кира", 6: name}
	weapons := map[int]string{1: "Пистолет", 2: "Нож", 3: "Топор", 4: "Удавка", 5: "Молоток", 6: "Шило"}
	var pl, pe, we int
	fmt.Println("1. Возможные места преступления, введите цифру:")
	fmt.Println(places)
	fmt.Scan(&pl)
	fmt.Println("2. Убийца, введите цифру:")
	fmt.Println(peoples)
	fmt.Scan(&pe)
	fmt.Println("3. Орудие, введите цифру:")
	fmt.Println(weapons)
	fmt.Scan(&we)
	if pl == level.place && pe == level.people && we == level.weapon {
		level.win = 1
	}
}
