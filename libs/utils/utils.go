package utils

import (
  "crypto/sha256"
  "fmt"
  )


//it should be a function from the user struct
func GetHash(str string) string {
    return fmt.Sprintf("%x",sha256.Sum256([]byte(str)))
}
