package main

import (
	"fmt"
	"time"
)

type RecruitmentOffice struct {
	observerList       []Observer
	name               string
	address            string
	isСonscriptionTime bool
}

func newRecruitmentOffice(name string, address string) *RecruitmentOffice {
	return &RecruitmentOffice{
		name:    name,
		address: address,
	}
}

func (r *RecruitmentOffice) announceMobilization() {
	r.isСonscriptionTime = true
	fmt.Printf("Военкомат %s начинает призыв\n", r.name)
	r.notifyAll()
}

func (r *RecruitmentOffice) register(o Observer) {
	r.observerList = append(r.observerList, o)
}

func (i *RecruitmentOffice) deregister(o Observer) {
	i.observerList = removeFromslice(i.observerList, o)
	c, ok := o.(*Citizen)
	if ok {
		fmt.Printf("Гражданин %s снят с воинского учета", c.name)
	} else {
		panic("Cast failed")
	}
}

func (r *RecruitmentOffice) notifyAll() {
	for _, observer := range r.observerList {
		observer.update(r.address + " " + time.Now().Format(time.Stamp))
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
