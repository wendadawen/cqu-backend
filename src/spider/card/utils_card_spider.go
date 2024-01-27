package card

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/object"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"log"
	"net/url"
	"os/exec"
	"regexp"
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

func (this *cardTemplate) eleChargeAt(campus campusFee) (*bo.ElectricCharge, error) {
	room := this.account.Room
	location, err := this.tryGetLocation()
	if err != nil {
		log.Printf("[CardSpider eleChargeAt Error] %+v\n", err)
		return nil, err
	}
	parse, _ := url.Parse(location)
	token := parse.Query().Get("token")
	if token == "" {
		log.Printf("[CardSpider eleChargeAt Error] %+v\n", object.CardTokenError)
		return nil, object.CardTokenError
	}
	post, err := resty.New().
		SetHeader("synjones-auth", "bearer "+token).
		SetFormData(map[string]string{
			"feeitemid": cast.ToString(campus.feeId), "type": "IEC", "level": "2", "room": room,
		}).R().Post(feeEleUrl)

	if err != nil {
		log.Printf("[CardSpider eleChargeAt Error] %+v\n", err)
		return nil, err
	}
	postFeeJson := post.String()
	data := gjson.Get(postFeeJson, "map.showData").String()
	if data == "" {
		log.Printf("[CardSpider eleChargeAt Error] %+v\n", object.CardRoomError)
		return nil, object.CardRoomError
	} else {
		return campus.feeBo(data), nil
	}
}

func (this *cardTemplate) tryGetLocation() (string, error) {
	ticket := this.hallticket
	if ticket == "" {
		return "", object.CardTicketError
	}
	refer := fmt.Sprintf(feeReferUrl, ticket) /**/
	curlResult, err := exec.Command("curl", "-I", refer).Output()
	if err != nil {
		log.Printf("[CardSpider tryGetLocation Error] %+v\n", err)
		return "", err
	}
	location := regexp.MustCompile("Location: (.*)").FindStringSubmatch(string(curlResult))
	if len(location) == 0 {
		return "", object.CardLocationError
	}
	return strings.TrimSpace(location[1]), nil
}
