// Code generated by download. DO NOT EDIT.

package iso20022_casp_013_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorRejection2 struct {
	RjctRsn  RejectReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 RjctRsn"`
	AddtlInf Max500Text        `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 AddtlInf,omitempty"`
	MsgInErr Max100KBinary     `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 MsgInErr,omitempty"`
}

type Document struct {
	SaleToPOIMsgRjctn SaleToPOIMessageRejectionV02 `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 SaleToPOIMsgRjctn"`
}

type GenericIdentification177 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 Id"`
	Tp       PartyType33Code    `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 Tp,omitempty"`
	Issr     PartyType33Code    `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 ShrtNm,omitempty"`
	RmotAccs NetworkParameters7 `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 RmotAccs,omitempty"`
	Glctn    Geolocation1       `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 Glctn,omitempty"`
}

type Geolocation1 struct {
	GeogcCordints GeolocationGeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 GeogcCordints,omitempty"`
	UTMCordints   GeolocationUTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 UTMCordints,omitempty"`
}

type GeolocationGeographicCoordinates1 struct {
	Lat  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 Lat"`
	Long Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 Long"`
}

type GeolocationUTMCoordinates1 struct {
	UTMZone    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 UTMZone"`
	UTMEstwrd  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 UTMEstwrd"`
	UTMNrthwrd Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 UTMNrthwrd"`
}

type Header41 struct {
	MsgFctn    RetailerMessage1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 MsgFctn"`
	PrtcolVrsn Max6Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 PrtcolVrsn"`
	XchgId     Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 XchgId"`
	CreDtTm    ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 CreDtTm"`
	InitgPty   GenericIdentification177 `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 InitgPty"`
	RcptPty    GenericIdentification177 `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 RcptPty,omitempty"`
	Tracblt    []Traceability8          `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 35 items long
type Max35Text string

// May be no more than 500 items long
type Max500Text string

// May be no more than 6 items long
type Max6Text string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type NetworkParameters7 struct {
	Adr        []NetworkParameters9 `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 SctyPrfl,omitempty"`
}

type NetworkParameters9 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 NtwkTp"`
	AdrVal Max500Text       `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 AdrVal"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS, MTMG, TAXH, TMGT
type PartyType33Code string

// May be one of UNPR, IMSG, PARS, SECU, INTP, RCPP, DPMG, VERS, MSGT
type RejectReason1Code string

// May be one of SSAB, SAAQ, SAAP, SDDR, SDDP, SSEN, SSMQ, SSMR, SSRJ, SARQ, SARP, SFRP, SFRQ, SFSQ, SFSP, SASQ, SASP
type RetailerMessage1Code string

type SaleToPOIMessageRejectionV02 struct {
	Hdr  Header41           `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 Hdr"`
	Rjct AcceptorRejection2 `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 Rjct"`
}

type Traceability8 struct {
	RlayId      GenericIdentification177 `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 RlayId"`
	PrtcolNm    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 PrtcolNm,omitempty"`
	PrtcolVrsn  Max6Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 PrtcolVrsn,omitempty"`
	TracDtTmIn  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 TracDtTmIn"`
	TracDtTmOut ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:casp.013.001.02 TracDtTmOut"`
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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