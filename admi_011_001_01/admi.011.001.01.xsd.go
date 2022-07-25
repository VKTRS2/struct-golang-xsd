// Code generated by download. DO NOT EDIT.

package iso20022_admi_011_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Document struct {
	SysEvtAck SystemEventAcknowledgementV01 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 SysEvtAck"`
}

type Event1 struct {
	EvtCd    Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 EvtCd"`
	EvtParam []Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 EvtParam,omitempty"`
	EvtDesc  Max350Text           `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 EvtDesc,omitempty"`
	EvtTm    ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 EvtTm,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemEventAcknowledgementV01 struct {
	MsgId       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 MsgId"`
	OrgtrRef    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 OrgtrRef,omitempty"`
	SttlmSsnIdr Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 SttlmSsnIdr,omitempty"`
	AckDtls     Event1                 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 AckDtls,omitempty"`
	SplmtryData []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01 SplmtryData,omitempty"`
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
