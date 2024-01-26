package card

import (
	"cqu-backend/src/bo"
	"fmt"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"strings"
)

func ParseRecord(json string) bo.ConsumptionRecord {
	data := gjson.Get(json, "rows")
	records := bo.ConsumptionRecord{}
	for _, record := range data.Array() {
		records = append(records, bo.OBJData{
			Time:        record.Get("tranDt").String(),
			Address:     strings.TrimSpace(record.Get("mchAcctName").String()),
			Transaction: strings.TrimSpace(record.Get("tranName").String()),
			Cost:        fmt.Sprintf("%.2f", cast.ToFloat32(record.Get("tranAmt").Int())/100.0),
			Balance:     fmt.Sprintf("%.2f", cast.ToFloat32(record.Get("acctAmt").Int())/100.0),
		})
	}
	return records
}

type campusFee struct {
	feeId string
	feeBo func(string) *bo.ElectricCharge
}

var (
	oldCampusFee = campusFee{
		feeId: oldFeeItemId,
		feeBo: func(fee string) *bo.ElectricCharge {
			return &bo.ElectricCharge{
				Balance: gjson.Get(fee, "现金余额").String(),
				Bonus:   gjson.Get(fee, "补贴余额").String(),
			}
		},
	}
	newCampusFee = campusFee{
		feeId: newFeeItemId,
		feeBo: func(fee string) *bo.ElectricCharge {
			return &bo.ElectricCharge{
				Balance: gjson.Get(fee, "剩余金额").String(),
				Bonus:   gjson.Get(fee, "电剩余补助").String(),
			}
		},
	}
)
