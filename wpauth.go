package wpauth_golang

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

const Itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func cryptPrivate(password, setting string) string {
	output := "*0"
	if setting[:2] == output {
		output = "*1"
	}

	id := setting[:3]
	if id != "$P$" && id != "$H$" {
		return output
	}

	countLog2 := strings.Index(Itoa64, string(setting[3]))
	if countLog2 < 7 || countLog2 > 30 {
		return output
	}

	count := 1 << countLog2
	salt := setting[4:12]
	if len(salt) != 8 {
		return output
	}

	hash := md5.Sum([]byte(salt + password))
	for i := 0; i < count; i++ {
		hash = md5.Sum(append(hash[:], password...))
	}

	output = setting[:12] + encode64(hash[:16], 16)
	return output
}

func encode64(input []byte, count int) string {
	var value int
	i := 0
	output := ""
	for i < count {
		value = int(input[i])
		i++
		output += string(Itoa64[value&0x3f])
		if i < count {
			value |= int(input[i]) << 8
		}
		output += string(Itoa64[(value>>6)&0x3f])
		if i >= count {
			break
		}
		i++
		if i < count {
			value |= int(input[i]) << 16
		}
		output += string(Itoa64[(value>>12)&0x3f])
		if i >= count {
			break
		}
		i++
		output += string(Itoa64[(value>>18)&0x3f])
	}
	return output
}

func WpCheckPassword(password, storedHash string) bool {
	if len(storedHash) <= 32 {
		// Calculate the MD5 hash of the password
		hasher := md5.New()
		hasher.Write([]byte(password))
		passwordHash := hex.EncodeToString(hasher.Sum(nil))

		if passwordHash == storedHash {
			return true
		}

	}
	hashedPassword := cryptPrivate(password, storedHash)
	return hashedPassword == storedHash
}
