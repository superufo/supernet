package logic

import (
	"testing"
)

func Test(t *testing.T) {
	he := HashExplosion{
		Seed:          8572516295, // 10000000000-99999999999   1672516295
		ExplosionSets: make(map[string][]string),
	}

	//hash := he.generalHash()
	//t.Logf("hash %s", hash)

	//hash := "92cc3c0e7d5ac05ca79ef52989a1559e017acf5b907dc55c34c1dec7"
	//hash := "06d2d6ffd3865f2631b7d997e6fbc6a27ab8f9ced55da779a3ac774e1054adc7"  “2cc3c0e7d5ac05ca79ef52989a1559e017acf5b907dc55c34c1dec7”
	res, err := he.hashResult("07ff47e891dbab61cf34f4b26e5ae1b36e935015017769f2a37ba47f84810c0c")

	if err != nil {
		t.Fatalf("测试错误%s", err.Error())
	}

	if res > 0.99 && res < 100000 {
		t.Logf("测试结果%f", res)
	}

	//he.HashSets("631b7d997e6fbc6a27ab8f", 1000)
	//
	//for i, v := range he.ExplosionSets["631b7d997e6fbc6a27ab8f_0.99_1.99"] {
	//	res, _ := he.hashResult(v)
	//	t.Log(i, v, res)
	//}

	explosion := he.GeneralExplosion()
	t.Log(explosion)
}
