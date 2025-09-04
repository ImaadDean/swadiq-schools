package students

import (
	"fmt"
	"time"
)

func GenerateStudentID() string {
	return fmt.Sprintf("STU%d", time.Now().Unix())
}
