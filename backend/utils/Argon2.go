package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GenerateHashArgon2(password string)(string , error){
	var (
		memory      uint32 = 64 * 1024 // 64 MB
		time        uint32 = 3
		threads     uint8  = 2
		keyLen      uint32 = 32
		saltLength         = 16
	)
	salt := make([]byte , saltLength)
	if _,err := rand.Read(salt);err != nil{
		return "",err
	}
	hash := argon2.IDKey([]byte(password) , salt , time , memory , threads , keyLen)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",argon2.Version , memory , time , threads , b64Salt , b64Hash)

	return encodedHash , nil
}

func ValidArgon2(passwordInput string , encodedHash string)(bool , error){
	params , salt , storedHash , err := decodeHash(encodedHash)
	if err != nil{
		return false , err
	}
	newHash := argon2.IDKey(
		[]byte(passwordInput),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		params.KeyLength,
	)

	if subtle.ConstantTimeCompare(storedHash, newHash) == 1 {
        return true, nil
    }

	return false , fmt.Errorf("senhas não conferem")
}

type Argon2Params struct {
    Memory      uint32
    Iterations  uint32
    Parallelism uint8
    SaltLength  uint32
    KeyLength   uint32
}

func decodeHash(encodedHash string)(*Argon2Params , []byte , []byte , error){
	// separa o hash em partes
	parts := strings.Split(encodedHash , "$")
	// verifica se a quantidade de partes é valida
	if len(parts) != 6{
		return nil , nil , nil , fmt.Errorf("formato de hash inválido") 
	}
	// verifica se é um argon2
	if parts[1] != "argon2id"{
		return nil , nil , nil , fmt.Errorf("algorítimo não suportado") 
	}

	var version int
	_ , err := fmt.Sscanf(parts[2] , "v=%d" , &version)
	if err != nil {
        return nil, nil, nil, err
    }
    if version != argon2.Version {
        return nil, nil, nil, fmt.Errorf("versão incompatível")
    }

	params := &Argon2Params{}
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", 
        &params.Memory, &params.Iterations, &params.Parallelism)
    if err != nil {
        return nil, nil, nil, err
    }

	salt , err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil{
		return nil, nil, nil, err
	}

	hash , err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
        return nil, nil, nil, err
    }

    params.KeyLength = uint32(len(hash))

	return params , salt , hash , nil
}