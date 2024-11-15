package entity

import (
	"fmt"
	"strconv"
	"sync_btim/utils"

	"gorm.io/gorm"
)

type IRepository interface {
	SyncClientID(email string, offset int, limit int) (bool, error)
}

type repository struct {
	DB1 *gorm.DB
	DB2 *gorm.DB
}

func NewTablesNameRepository(db1 *gorm.DB, db2 *gorm.DB) *repository {
	return &repository{db1, db2}
}

func (r *repository) SyncClientID(email string, offset int, limit int) (bool, error) {
	tx := r.DB1.Begin()
	if tx.Error != nil {
		return false, tx.Error
	}

	// Defer rollback in case of error - will be ignored if committed successfully
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if email == "" {
		// ambil data dari blink v3,
		var userBlink []utils.UserBlink
		err := r.DB1.Raw(fmt.Sprintf("SELECT kyc.KycID, kyc.ClientID, kyc.Email, kyc.IdentityNo FROM client_tbl_kyc kyc LIMIT %d OFFSET %d", limit, offset)).Scan(&userBlink).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		fmt.Println(userBlink)

		// looping data dan mencari kesamaan dengan data BTIM client
		for _, v := range userBlink {
			var clientIDBTIM utils.ClientIDBTIM
			err = r.DB2.Raw(fmt.Sprintf("SELECT c.ClientID FROM client c WHERE c.Email = '%s' AND c.IDNo = '%s'", v.Email, v.IdentityNo)).Scan(&clientIDBTIM).Error
			if err != nil {
				tx.Rollback()
				panic(err)
			}
			fmt.Println(clientIDBTIM)

			// upadte ClientIDBTIM
			BTIMID := strconv.Itoa(clientIDBTIM.ClientID)
			if clientIDBTIM.ClientID != 0 {
				err := r.DB1.Raw(fmt.Sprintf("UPDATE client_tbl_kyc SET ClientIDBTIM = %s WHERE KycID = %d", BTIMID, v.KycID)).Scan(&clientIDBTIM).Error
				if err != nil {
					tx.Rollback()
					panic(err)
				}
				fmt.Println("udpate kycID", v.KycID)
			} else {
				fmt.Println(v.Email + "belum terdaftar di BTIM")
			}

		}
	}

	if email != "" {
		// ambil data dari blink v3,
		var userBlink utils.UserBlink
		err := r.DB1.Raw(fmt.Sprintf("SELECT kyc.KycID, kyc.ClientID, kyc.Email, kyc.IdentityNo FROM client_tbl_kyc kyc WHERE kyc.Email = %s", email)).Scan(&userBlink).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		fmt.Println(userBlink)

		var clientIDBTIM utils.ClientIDBTIM
		err = r.DB2.Raw(fmt.Sprintf("SELECT c.ClientID FROM client c WHERE c.Email = '%s' AND c.IDNo = '%s'", userBlink.Email, userBlink.IdentityNo)).Scan(&clientIDBTIM).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		fmt.Println(clientIDBTIM)

		// upadte ClientIDBTIM
		BTIMID := strconv.Itoa(clientIDBTIM.ClientID)
		if clientIDBTIM.ClientID != 0 {
			err := r.DB1.Raw(fmt.Sprintf("UPDATE client_tbl_kyc SET ClientIDBTIM = %s WHERE KycID = %d", BTIMID, userBlink.KycID)).Scan(&clientIDBTIM).Error
			if err != nil {
				tx.Rollback()
				panic(err)
			}
			fmt.Println("udpate kycID", userBlink.KycID)
		} else {
			fmt.Println(userBlink.Email + "belum terdaftar di BTIM")
		}
	}

	return true, nil

}
