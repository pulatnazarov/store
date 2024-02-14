package helper

import (
	"fmt"
	"math/rand"
)

func PhoneGenerate() string {
	phoneNumbers := make(map[string]bool)

	for {
		randNum := fmt.Sprintf("%09d", rand.Intn(1000000000))
		numberWithRandom := "+998" + randNum

		if !phoneNumbers[numberWithRandom] {
			phoneNumbers[numberWithRandom] = true // bu nomer uje ishlatildi
			return numberWithRandom
		}
	}
}
