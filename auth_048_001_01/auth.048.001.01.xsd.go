// Code generated by download. DO NOT EDIT.

package iso20022_auth_048_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CountryCodeAndName3 struct {
	Cd CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 Cd"`
	Nm Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 Nm"`
}

type CurrencyCodeAndName1 struct {
	Cd ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 Cd"`
	Nm Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 Nm"`
}

type Document struct {
	FinInstrmRptgCcyCdRpt FinancialInstrumentReportingCurrencyCodeReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 FinInstrmRptgCcyCdRpt"`
}

type FinancialInstrumentReportingCurrencyCodeReportV01 struct {
	CcyData     []SecuritiesCurrencyIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 CcyData"`
	SplmtryData []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 SplmtryData,omitempty"`
}

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

// May be no more than 350 items long
type Max350Text string

// May be no more than 70 items long
type Max70Text string

// May be one of NOCH, MODI, DELE, ADDD
type Modification1Code string

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 ToDt"`
}

type Period4Choice struct {
	Dt       ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 Dt,omitempty"`
	FrDt     ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 FrDt,omitempty"`
	ToDt     ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 ToDt,omitempty"`
	FrDtToDt Period2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 FrDtToDt,omitempty"`
}

type SecuritiesCurrencyIdentification2 struct {
	Ccy       CurrencyCodeAndName1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 Ccy"`
	FrctnlDgt float64              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 FrctnlDgt,omitempty"`
	CtryDtls  CountryCodeAndName3  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 CtryDtls"`
	PreEuro   bool                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 PreEuro"`
	Mod       Modification1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 Mod,omitempty"`
	VldtyPrd  Period4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 VldtyPrd"`
	LastUpdtd ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 LastUpdtd,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.048.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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
