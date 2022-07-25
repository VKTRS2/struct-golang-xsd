// Code generated by download. DO NOT EDIT.

package iso20022_reda_023_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of INSE, UPDT, DELT
type DataModification1Code string

type Document struct {
	SctiesAcctModReq SecuritiesAccountModificationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 SctiesAcctModReq"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 SchmeNm,omitempty"`
}

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type MarketSpecificAttribute1 struct {
	Nm  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 Nm"`
	Val Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 Val"`
}

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be no more than 70 items long
type Max70Text string

type MessageHeader1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 MsgId"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 CreDtTm,omitempty"`
}

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 Nm,omitempty"`
}

type SecuritiesAccountModification2 struct {
	ScpIndctn DataModification1Code                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 ScpIndctn"`
	ReqdMod   SecuritiesAccountModification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 ReqdMod"`
}

type SecuritiesAccountModification2Choice struct {
	SysSctiesAcct SystemSecuritiesAccount5 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 SysSctiesAcct,omitempty"`
	SysRstrctn    SystemRestriction1       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 SysRstrctn,omitempty"`
	MktSpcfcAttr  MarketSpecificAttribute1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 MktSpcfcAttr,omitempty"`
}

type SecuritiesAccountModificationRequestV01 struct {
	MsgHdr      MessageHeader1                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 MsgHdr,omitempty"`
	AcctId      SecuritiesAccount19              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 AcctId"`
	Mod         []SecuritiesAccountModification2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 Mod"`
	SplmtryData []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 SplmtryData,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemRestriction1 struct {
	VldFr ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 VldFr"`
	VldTo ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 VldTo,omitempty"`
	Tp    Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 Tp"`
}

type SystemSecuritiesAccount5 struct {
	ClsgDt       ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 ClsgDt,omitempty"`
	HldInd       bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 HldInd,omitempty"`
	NegPos       bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 NegPos,omitempty"`
	EndInvstrFlg Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 EndInvstrFlg,omitempty"`
	PricgSchme   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:reda.023.001.01 PricgSchme,omitempty"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
