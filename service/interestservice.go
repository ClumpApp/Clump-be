package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

const (
	threshold = 3
)

func (obj *Service) CreateInterest(interestDTO model.InterestDTO) {
	interest := model.Interest{
		Title:   interestDTO.Title,
		Picture: interestDTO.Picture,
	}
	obj.db.Create(&model.Interest{}, &interest)
}

func (obj *Service) GetInterests() []model.InterestDTO {
	var interestDTOs []model.InterestDTO
	obj.db.Query(&model.Interest{}, &model.Interest{SubInterestCount: 0}, &interestDTOs)
	for index := range interestDTOs {
		interestDTOs[index].Picture = utility.GetStorage().GetURL() + interestDTOs[index].Picture
	}
	return interestDTOs
}

func (obj *Service) AddInterests(interestsDTO []model.InterestDTO, userid float64) {
	userID := uint(userid)
	ieUserInterests := []model.IEUserInterest{}
	for _, interestDTO := range interestsDTO {
		var interest model.Interest
		uuid := utility.ConvertUUID(interestDTO.UUID)
		obj.db.Query(&model.Interest{}, &model.Interest{UUID: uuid}, &interest)
		ieUserInterests = append(ieUserInterests, model.IEUserInterest{
			UserID:     userID,
			InterestID: interest.ID,
		})
	}
	obj.db.Create(&model.IEUserInterest{}, &ieUserInterests)
}

func (obj *Service) FindMatchingGroup(userid float64) bool {
	userID := uint(userid)

	var ieUserInterests []model.IEUserInterest
	obj.db.Query(&model.IEUserInterest{}, &model.IEUserInterest{UserID: userID}, &ieUserInterests)

	interestIDs := []uint{}
	for _, ieUserInterest := range ieUserInterests {
		interestIDs = append(interestIDs, ieUserInterest.InterestID)
	}

	var interests []model.Interest
	obj.db.QueryWithPreload(&model.Interest{}, &interestIDs, &interests)

	groups := make(map[uint]int)
	for _, interest := range interests {
		for _, ieGroupInterest := range interest.GroupInterests {
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

	if bestMatchRank < threshold { // Create a group of the user alone
		var newGroup model.Group
		obj.db.Create(&model.Group{}, &newGroup)
		ieGroupInterests := []model.IEGroupInterest{}
		for _, ieUserInterest := range ieUserInterests {
			ieGroupInterest := model.IEGroupInterest{GroupID: newGroup.ID, InterestID: ieUserInterest.InterestID}
			ieGroupInterests = append(ieGroupInterests, ieGroupInterest)
		}
		obj.db.Create(&model.IEGroupInterest{}, &ieGroupInterests)
		obj.db.Create(&model.IEUserGroup{}, &model.IEUserGroup{GroupID: newGroup.ID, UserID: userID})
		obj.db.Update(&model.User{}, userID, &model.User{GroupID: &newGroup.ID})
		return false
	}

	// Remove additional interests from a group if there is only 1 user
	var userCount int64
	obj.db.Count(&model.User{}, &model.User{GroupID: &bestMatch}, &userCount)
	if userCount == 1 {
		var ieGroupInterests []model.IEGroupInterest
		obj.db.Query(&model.IEGroupInterest{}, &model.IEGroupInterest{GroupID: bestMatch}, &ieGroupInterests)

		ieGroupInterestsToBeDeleted := []uint{}
		for _, ieGroupInterest := range ieGroupInterests {
			isCommon := false
			for _, ieUserInterest := range ieUserInterests {
				if ieUserInterest.InterestID == ieGroupInterest.InterestID {
					isCommon = true
					break
				}
			}
			if !isCommon {
				ieGroupInterestsToBeDeleted = append(ieGroupInterestsToBeDeleted, ieGroupInterest.ID)
			}
		}

		obj.db.Delete(&model.IEGroupInterest{}, &ieGroupInterestsToBeDeleted)
	}

	obj.db.Create(&model.IEUserGroup{}, &model.IEUserGroup{GroupID: bestMatch, UserID: userID})
	obj.db.Update(&model.User{}, userID, &model.User{GroupID: &bestMatch})
	return true
}

/*
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
*/
