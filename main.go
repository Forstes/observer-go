package main

func main() {
	moscowRO := newRecruitmentOffice("Военкомат г. Москва", "пр-т Мира, 15, Москва")
	petersburgRO := newRecruitmentOffice("Военкомат г. Санкт-Петербург", "Английский пр., 8/10, Санкт-Петербург")

	recruit1 := newCitizen("050902778817", "Александров Денис")
	recruit2 := newCitizen("190301181115", "Васильев Станислав")
	recruit3 := newCitizen("29099861719", "Дегтярев Михаил")
	recruit4 := newCitizen("141299181912", "Пупкин Василий")

	moscowRO.register(recruit1)
	moscowRO.register(recruit2)
	petersburgRO.register(recruit3)
	petersburgRO.register(recruit4)

	moscowRO.announceMobilization()
	petersburgRO.announceMobilization()

	moscowRO.deregister(recruit2)
}
