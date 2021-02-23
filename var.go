// Copyright 2021, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package krx

var (
	// DungRakMap represents 등락 번호의 의미
	DungRakMap = map[string]string{"1": "상한", "2": "상승", "3": "보합", "4": "하한", "5": "보합"}
	// titleMap   = map[string]string{"JongName": "종목명", "CurJuka": "현재가", "Debi": "전일대비", "DungRak": "등락", "PrevJuka": "전일종가", "Volume": "거래량", "Money": "거래대금", "StartJuka": "시가", "HighJuka": "고가", "LowJuka": "저가", "High52": "52주최고", "Low52": "52주최저", "UpJuka": "상한가", "DownJuka": "하한가", "Per": "PER", "Amount": "상장주식수", "FaceJuka": "액면가"}
	// dailyMap   = map[string]string{"day_Date": "일자", "day_EndPrice": "종가", "day_getDebi": "전일대비", "day_Dungrak": "전일대비(등락)", "day_Start": "시가", "day_High": "고가", "day_Low": "저가", "day_Volume": "거래량", "day_getAmount": "거래대금"}
	// askMap     = map[string]string{"member_memdoMem": "매도상위 증권사", "member_memdoVol": "매도 거래량", "member_memsoMem": "매수상위 증권사", "member_mesuoVol": "매수 거래량"}

	// 실시간시세
	siseURLKrFmt  = "http://asp1.krx.co.kr/servlet/krx.asp.XMLSise?code=%s"
	siseURLEngFmt = "http://asp1.krx.co.kr/servlet/krx.asp.XMLSiseEng?code=%s"

	// 공시정보
	disInfoURLKrFmt  = "http://asp1.krx.co.kr/servlet/krx.asp.DisList4MainServlet?code=%s&gubun=K"
	disInfoURLEngFmt = "http://asp1.krx.co.kr/servlet/krx.asp.DisList4MainServlet?code=%s&gubun=E"
)
