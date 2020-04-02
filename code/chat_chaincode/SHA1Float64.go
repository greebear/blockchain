package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Sha1WithRandFloat64(input string)(hashResult string, randNumber string, err error){
	// set random seed
	rand.Seed(time.Now().UnixNano())
	randNumStr := strconv.FormatFloat(rand.Float64(), 'g', -1, 64)
	input = input + randNumStr
	hashRes, err := Sha1(input)
	if err != nil {
		return "", "", err
	}
	return fmt.Sprintf("%x", hashRes), randNumStr, nil
}

func Sha1(input string) ([]byte, error) {
	var err error
	// sha1
	hash := sha1.New()
	_, err = hash.Write([]byte(input))
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}
