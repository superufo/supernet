package logic

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/supernet/common/utils"
	"math/rand"
	"strconv"
	"time"
)

type HashExplosion struct {
	Seed float32

	ExplosionSets map[string][]string
}

func (he *HashExplosion) seed() {

}

func (he *HashExplosion) hashResult(hash string) (float64, error) {
	tp, err := strconv.ParseInt(string([]byte(hash)[:8]), 16, 64)
	fmt.Println(hash)

	//if tp%100 == 0 {
	//	return 1, err
	//}

	// Div(decimal.NewFromInt(10))
	tc := decimal.New(tp, 0)
	res, _ := decimal.NewFromFloat32(he.Seed).Div(tc.Add(decimal.NewFromInt(1)).Mul(decimal.NewFromFloat32(0.99))).Truncate(2).Float64()
	res, _ = decimal.NewFromFloat(res).Sub(decimal.NewFromInt(1)).Float64()

	if decimal.NewFromFloat(res).LessThan(decimal.NewFromFloat(0.99)) == true {
		res = 0.99
	}

	fmt.Println("转换的值：", res)
	return res, err
}

func (he *HashExplosion) generalHash() string {
	str := utils.GenRandString(32)
	hash := utils.GenerateSHA256(str)

	return hash
}

func (he *HashExplosion) HashSets(gameId string, num int) {
	for count := 0; count < num; count++ {
		str := utils.GenRandString(32)
		hash := utils.GenerateSHA256(str)

		res, _ := he.hashResult(hash)

		key := fmt.Sprintf("%s_0.99_1.99", gameId)
		if he.Between(res, 0.99, 1.99) {
			key = fmt.Sprintf("%s_0.99_1.99", gameId)
		} else if he.Between(res, 1.99, 2.99) {
			key = fmt.Sprintf("%s_1.99_2.99", gameId)
		} else if he.Between(res, 2.99, 3.99) {
			key = fmt.Sprintf("%s_2.99_3.99", gameId)
		} else if he.Between(res, 3.99, 4.99) {
			key = fmt.Sprintf("%s_3.99_4.99", gameId)
		} else if he.Between(res, 4.99, 5.99) {
			key = fmt.Sprintf("%s_4.99_5.99", gameId)
		} else if he.Between(res, 5.99, 9.99) {
			key = fmt.Sprintf("%s_5.99_9.99", gameId)
		} else if he.Between(res, 9.99, 1000) {
			key = fmt.Sprintf("%s_9.99_1000", gameId)
		}

		he.ExplosionSets[key] = append(he.ExplosionSets[key], hash)
	}
}

func (he *HashExplosion) Between(res float64, min float64, max float64) bool {
	return decimal.NewFromFloat(res).LessThan(decimal.NewFromFloat(max)) && decimal.NewFromFloat(min).LessThan(decimal.NewFromFloat(res))
}

func (he *HashExplosion) GeneralExplosion() (explosion float64) {
	// 飞行的微秒数
	var s int64

	// 飞行描述的概率分布  0.99+0.5*0.03*8.2*8.2 = 1.9986  飞行高度  0.99 - 1.9986
	gl := rand.Int63n(1000)
	if gl < 300 {
		s = rand.Int63n(8200) // 对应飞行高度  0.99 - 1.9986
	}

	// 飞行描述的概率分布  0.99+0.5*0.03*11.6*11.6 = 3.0084  飞行高度 1.9986 - 3.0084
	if gl > 300 && gl < 500 {
		s = 8200 + rand.Int63n(11600-8200) // 对应飞行高度  1.9986 - 3.0084
	}

	// 飞行描述的概率分布  0.99+0.5*0.03*14.16*14.16 = 3.997584  飞行高度  3.0084 - 3.997584
	if gl > 500 && gl < 700 {
		s = 11600 + rand.Int63n(14160-11600) // 对应飞行高度  3.0084 - 3.997584
	}

	// 飞行描述的概率分布  0.99+0.5*0.03*16.35*16.35 = 4.9998 飞行高度 3.997584 -4.9998
	if gl > 700 && gl < 800 {
		s = 14160 + rand.Int63n(16350-14160) // 对应飞行高度 3.997584 - 4.9998
	}

	// 飞行描述的概率分布  0.99+0.5*0.03*18.25*18.25 = 5.000 飞行高度 4.9998 - 5.9859
	if gl > 800 && gl < 900 {
		s = 16350 + rand.Int63n(18250-16350) // 对应飞行高度 4.9998 -5.9859
	}

	// 飞行描述的概率分布  0.99+0.5*0.03*24.49*24.49 = 9.98 飞行高度 5.9859 - 9.98
	if gl > 900 && gl < 990 {
		s = 18250 + rand.Int63n(24490-18250) // 对应飞行高度 5.9859 -9.98
	}

	// 飞行描述的概率分布  0.99+0.5*0.03*258*258 = 999.45  飞行高度 9.98 - 999.45
	if gl > 990 && gl < 1000 {
		s = 18250 + rand.Int63n(24490-18250) // 对应飞行高度 9.98 - 999.45
	}

	//tm := time.NewTimer(300 * time.Millisecond)
	//stop := make(chan bool)
	//
	//go func(t *time.Timer) {
	//	defer t.Stop()
	//	for {
	//		select {
	//		case <-t.C:
	//			t.Reset(300 * time.Millisecond)
	//		case <-stop:
	//			break
	//		}
	//	}
	//}(tm)
	//stop <- true

	// 初始化日志文件

	time.Sleep(time.Duration(s) * time.Millisecond)
	// 1/2*0.03
	explosion, _ = decimal.NewFromFloat(0.99).Add(decimal.NewFromFloat(0.015).Mul(decimal.NewFromInt(s).Mul(decimal.NewFromInt(s)))).Float64()
	explosion, _ = decimal.NewFromFloat(explosion).Div(decimal.NewFromInt(1000)).Div(decimal.NewFromInt(1000)).Truncate(2).Float64()

	if explosion < 0.99 {
		explosion = 0.99
	}

	return explosion
}
