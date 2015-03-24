package samples

import (
    "crypto/rand"
)
// RandString function is used for generating randomly generated strings initialized with a given prefix
// with the length of given number.
// This function is called when creating a new Alert.
func RandString(n int) string {
    const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes = make([]byte, n)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = alphanum[b % byte(len(alphanum))]
    }
    return string(bytes)
}

func RandStringWithPrefix(prefix string, n int) string {
    return prefix + "-" + RandString(n)
}
