package problem1352

import "testing"

func TestExecute(t *testing.T) {

  pon := Constructor()

  pon.Add(3)
  pon.Add(0)
  pon.Add(2)
  pon.Add(5)
  pon.Add(4)

  if pon.GetProduct(2) != 20 {
    t.Errorf("[FAILED] expecting 20, get %d", pon.GetProduct(2))
  }

  if pon.GetProduct(3) != 40 {
    t.Errorf("[FAILED] expecting 40, get %d", pon.GetProduct(3))
  }

  if pon.GetProduct(4) != 0 {
    t.Errorf("[FAILED] expecting 0, get %d", pon.GetProduct(4))
  }

  pon.Add(8)

  if pon.GetProduct(2) != 32 {
    t.Errorf("[FAILED] expecting 32, get %d", pon.GetProduct(2))
  }

}
