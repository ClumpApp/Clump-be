package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

const (
	threshold = 3
)

func (obj *Service) CreateInterest(interestDTO model.InterestDTO) {
	var interest model.Interest
	utility.Convert(interestDTO, interest)
	obj.db.Create(&model.Interest{}, &interest)
}

func (obj *Service) GetInterests() []model.InterestDTO {
	var interests []model.Interest
	obj.db.Query(&model.Interest{}, &model.Interest{InterestID: nil}, &interests)
	var interestDTOs []model.InterestDTO
	for _, interest := range interests {
		interestDTOs = append(interestDTOs, model.InterestDTO{
			UUID:    utility.ConvertString(interest.UUID),
			Title:   interest.Title,
			Picture: interest.Picture,
		})
	}
	return interestDTOs
}

func (obj *Service) AddInterests(interestsDTO []model.InterestDTO, userid float64) {
	var interest model.Interest
	for _, interestDTO := range interestsDTO {
		obj.db.Query(&model.Interest{}, &model.Interest{
			UUID: utility.ConvertUUID(interestDTO.UUID),
		}, &interest)
		obj.db.Create(&model.IEUserInterest{}, &model.IEUserInterest{
			UserID:     uint(userid),
			InterestID: interest.ID,
		})
	}
}

func (obj *Service) FindGroup(userid float64) bool {
	var user model.User
	obj.db.QueryWithPreload(&model.User{}, uint(userid), &user)
	ieUserInterests := user.UserInterests

	groups := make(map[uint]int)
	for _, ieUserInterest := range ieUserInterests {
		var ieGroupInterests []model.IEGroupInterest
		obj.db.Query(&model.IEGroupInterest{}, &model.IEGroupInterest{InterestID: ieUserInterest.ID}, &ieGroupInterests)
		for _, ieGroupInterest := range ieGroupInterests {
			groups[ieGroupInterest.GroupID] += 1
		}
	}

	var bestMatch uint
	bestMatchRank := 0
	for id, rank := range groups {
		if rank > bestMatchRank {
			bestMatch = id
			bestMatchRank = rank
		}
	}

	if bestMatchRank < threshold {
		return false
	}

	obj.db.Update(&model.User{}, uint(userid), &model.User{GroupID: bestMatch})
	return true
}

func (obj *Service) FindInterests(userid uint) []model.Interest {
	var intr model.Interest
	var intrs []model.Interest
	var userintrs []model.IEUserInterest
	obj.db.Query(&model.User{}, userid, &userintrs)
	for _, userinterest := range userintrs {
		obj.db.Query(&model.Interest{}, userinterest.ID, &intr)
		intrs = append(intrs, intr)
	}
	return intrs
}

func (obj *Service) FindInterestTitles(userid uint) []string {
	var intr model.Interest
	var intrtitles []string
	var userintrs []model.IEUserInterest
	obj.db.Query(&model.User{}, userid, &userintrs)
	for _, userinterest := range userintrs {
		obj.db.Query(&model.Interest{}, userinterest.ID, &intr)
		intrtitles = append(intrtitles, intr.Title)
	}
	return intrtitles
}

// Calculate using Loss Metric
func (obj *Service) CalculateLevel(groupintrs []string, userintrs []string) {

}

func CalculateInterest() {
	//subInterest addition: 50 vs. penalty: -5
	//midInterest addition: 20 vs. penalty: -10
	//topInterest addition: 10 vs. penalty: -20
	//this is calculated when we already have groups

}
