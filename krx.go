// Copyright 2021, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krx

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// GetSiseKr returns 한글 시세 정보
func GetSiseKr(code string) (*Stockprice, error) {
	return getSise(siseURLKrFmt, code)
}

// GetSiseEng returns 영문 시세 정보
func GetSiseEng(code string) (*Stockprice, error) {
	return getSise(siseURLEngFmt, code)
}

// GetDisKr returns 한글 공시 정보
func GetDisKr(code string) (*DisclosureMain, error) {
	return getDis(disInfoURLKrFmt, code)
}

// GetDisEng returns 영문 공시 정보
func GetDisEng(urlFmt, code string) (*DisclosureMain, error) {
	return getDis(disInfoURLEngFmt, code)
}

// ----

func getSise(urlFmt, code string) (*Stockprice, error) {
	resp, err := http.Get(fmt.Sprintf(urlFmt, code))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := xml.NewDecoder(resp.Body)
	var sise Stockprice
	dec.Decode(&sise)

	return &sise, nil
}

func getDis(urlFmt, code string) (*DisclosureMain, error) {
	resp, err := http.Get(fmt.Sprintf(urlFmt, code))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := xml.NewDecoder(resp.Body)
	var dis DisclosureMain
	dec.Decode(&dis)

	return &dis, nil
}
