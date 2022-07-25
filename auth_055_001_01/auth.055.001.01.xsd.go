// Code generated by download. DO NOT EDIT.

package iso20022_auth_055_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type Amount3 struct {
	OrgnlAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 OrgnlAmt,omitempty"`
	RptgAmt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 RptgAmt"`
}

type AmountAndDirection102 struct {
	Amt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Amt"`
	Sgn bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Sgn"`
}

type CCPMemberRequirementsReportV01 struct {
	IntraDayRqrmntAmt []IntraDayRequirement1    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 IntraDayRqrmntAmt"`
	IntraDayMrgnCall  []IntraDayMarginCall1     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 IntraDayMrgnCall,omitempty"`
	EndOfDayRqrmnt    []EndOfDayRequirement2    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 EndOfDayRqrmnt"`
	DfltFndRqrmnt     []DefaultFundRequirement1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 DfltFndRqrmnt"`
	SplmtryData       []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 SplmtryData,omitempty"`
}

type DefaultFundRequirement1 struct {
	ClrMmbId GenericIdentification165 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 ClrMmbId"`
	SvcId    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 SvcId,omitempty"`
	Amt      ActiveCurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Amt"`
}

type Document struct {
	CCPMmbRqrmntsRpt CCPMemberRequirementsReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 CCPMmbRqrmntsRpt"`
}

type EndOfDayRequirement2 struct {
	InitlMrgnRqrmnts InitialMarginRequirement1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 InitlMrgnRqrmnts"`
	VartnMrgnRqrmnts AmountAndDirection102     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 VartnMrgnRqrmnts"`
	MrgnAcctId       GenericIdentification165  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 MrgnAcctId"`
}

type GenericIdentification165 struct {
	Id      Max256Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Id"`
	Desc    Max140Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Desc,omitempty"`
	Issr    Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Issr,omitempty"`
	SchmeNm SchemeIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 SchmeNm,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type InitialMarginExposure1 struct {
	Amt     Amount3           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Amt"`
	Tp      MarginType2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Tp"`
	CoreInd bool              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 CoreInd"`
}

type InitialMarginRequirement1 struct {
	InitlMrgnXpsr []InitialMarginExposure1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 InitlMrgnXpsr"`
	Cdt           ActiveCurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Cdt"`
}

type IntraDayMarginCall1 struct {
	MrgnAcctId   GenericIdentification165 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 MrgnAcctId"`
	IntraDayCall ActiveCurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 IntraDayCall"`
	TmStmp       ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 TmStmp"`
}

type IntraDayRequirement1 struct {
	IntraDayMrgnCall   ActiveCurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 IntraDayMrgnCall"`
	PeakInitlMrgnLblty ActiveCurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 PeakInitlMrgnLblty"`
	PeakVartnMrgnLblty ActiveCurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 PeakVartnMrgnLblty"`
	AggtPeakLblty      ActiveCurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 AggtPeakLblty"`
	MrgnAcctId         GenericIdentification165 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 MrgnAcctId"`
}

type MarginType2Choice struct {
	Cd    MarginType2Code         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Cd,omitempty"`
	Prtry GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Prtry,omitempty"`
}

// May be one of ADFM, COMA, CEMA, SEMA, SCMA, UFMA, MARM, SORM, WWRM, BARM, LIRM, CRAM, CVMA, SPMA, JTDR, DRAO, OTHR
type MarginType2Code string

// May be no more than 140 items long
type Max140Text string

// May be no more than 256 items long
type Max256Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be one of MARG, COLL, POSI, CLIM
type SchemeIdentificationType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.055.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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
