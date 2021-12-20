package redis

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GenUUID() string {
	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}
