package common

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"time"

	"golang.org/x/crypto/blake2b"
)

func base64Encode(payload []byte) string {
	return base64.StdEncoding.EncodeToString(payload)
}

func base64Decode(payload string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(payload)
}

func signRequest(privateKey string, payload []byte, currentTime int, ttl int) (string, error) {

	//compute blake 512 hash over the payload
	hash := blake2b.Sum512(payload)
	digest := base64Encode(hash[:])

	//create a signature and then sign with ed25519 private key
	signatureBody := fmt.Sprintf("(created): %d\n(expires): %d\ndigest: BLAKE-512=%s", currentTime, (currentTime + ttl), digest)
	decodedKey, err := base64Decode(privateKey)
	if err != nil {
		fmt.Println("Error decoding signing private key", err)
		return "", err
	}
	signature := ed25519.Sign(decodedKey, []byte(signatureBody))
	return base64Encode(signature), nil
}

func GetAuthHeader(requestBody []byte, signingPrivKey string) (string, error) {
	var authHeader string

	currentTime := int(time.Now().Unix())

	//ttl we are using is 30 seconds
	ttl := 30

	signature, err := signRequest(signingPrivKey, requestBody, currentTime, ttl)
	if err != nil {
		fmt.Println("Could not compute signature", err)
		return authHeader, err
	}
	fmt.Println(string(requestBody))
	subscriberID := "app.mfapis.in"
	if subscriberID == "" {
		fmt.Println("Could not compute signature", err)
		return authHeader, errors.New("subscriber_id can not be empty")
	}

	uniqueKeyID := "3875f37b-df73-4592-9cb6-0887adf9b962"
	authHeader = fmt.Sprintf(`Signature keyId="%s|%s|ed25519",algorithm="ed25519",created="%d",expires="%d",headers="(created) (expires) digest",signature="%s"`,
		subscriberID, uniqueKeyID, currentTime, currentTime+ttl, signature)

	return authHeader, nil
}

func parseAuthHeader(authHeader string) (string, string, string, string, error) {
	signatureRegex := regexp.MustCompile(`keyId=\"(.+?)\".+?created=\"(.+?)\".+?expires=\"(.+?)\".+?signature=\"(.+?)\"`)
	groups := signatureRegex.FindAllStringSubmatch(authHeader, -1)
	if len(groups) > 0 && len(groups[0]) > 4 {
		return groups[0][1], groups[0][2], groups[0][3], groups[0][4], nil
	}
	fmt.Println("Error parsing auth header. Please make sure that the auh headers passed as command line argument is valid")
	return "", "", "", "", errors.New("error parsing auth header")
}

// Only for Tetsing purpose
func VerifyRequest(authHeader string, requestBody []byte) bool {

	publicKey := "xoJcsMp2CVDyaOgf7wef0cJ9l0gCpSk+bJ0LqbzlOP0="

	_, created, expires, signature, err := parseAuthHeader(authHeader)
	if err != nil {
		return false
	}

	//compute blake 512 hash over the payload
	hash := blake2b.Sum512(requestBody)
	digest := base64Encode(hash[:])

	//create a signature and then sign with ed25519 private key
	computedMessage := fmt.Sprintf("(created): %s\n(expires): %s\ndigest: BLAKE-512=%s", created, expires, digest)
	publicKeyBytes, err := base64Decode(publicKey)
	if err != nil {
		fmt.Println("Error decoding public key", err)
		return false
	}
	receivedSignature, err := base64Decode(signature)
	if err != nil {
		fmt.Println("Unable to base64 decode received signature", err)
		return false
	}
	return ed25519.Verify(publicKeyBytes, []byte(computedMessage), receivedSignature)
}
