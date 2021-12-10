package service

import (
	"github.com/clumpapp/clump-be/model"
	//"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) AddInterests(interests []string, userid uint) {
	var intr model.Interest
	var intrs []model.Interest
	for _, interest := range interests {
		obj.db.Query(&model.Interest{}, &model.Interest{Title: interest}, &intr)
		intrs = append(intrs, intr)
	}
	obj.db.Update(&model.User{}, userid, &model.User{UserInterests: intrs})
}

func CalculateInterest() {
	//subInterest addition: 50 vs. penalty: -5
	//midInterest addition: 20 vs. penalty: -10
	//topInterest addition: 10 vs. penalty: -20
	//this is calculated when we already have groups

}
