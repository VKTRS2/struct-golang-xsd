// Code generated by download. DO NOT EDIT.

package iso20022_camt_010_001_08

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 IBAN,omitempty"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Othr,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Cd,omitempty"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Prtry,omitempty"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Prtry,omitempty"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 AmtWthtCcy,omitempty"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 AmtWthCcy,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 PstlAdr,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Cd,omitempty"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 MmbId"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Dt,omitempty"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 DtTm,omitempty"`
}

type Document struct {
	RtrLmt ReturnLimitV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 RtrLmt"`
}

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Cd,omitempty"`
	Prtry Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Prtry,omitempty"`
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Err"`
	Desc Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Desc,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalAccountIdentification1Code string

// May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// May be no more than 4 items long
type ExternalEnquiryRequestType1Code string

// May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// May be no more than 3 items long
type ExternalMarketInfrastructure1Code string

// May be no more than 4 items long
type ExternalPaymentControlRequestType1Code string

// May be no more than 4 items long
type ExternalSystemErrorHandling1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Cd,omitempty"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Prtry,omitempty"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type Limit7 struct {
	Amt             Amount2Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Amt"`
	CdtDbtInd       CreditDebitCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 CdtDbtInd,omitempty"`
	Sts             LimitStatus1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Sts,omitempty"`
	StartDtTm       DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 StartDtTm,omitempty"`
	UsdAmt          Amount2Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 UsdAmt,omitempty"`
	UsdAmtCdtDbtInd CreditDebitCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 UsdAmtCdtDbtInd,omitempty"`
	UsdPctg         float64                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 UsdPctg,omitempty"`
	RmngAmt         Amount2Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 RmngAmt,omitempty"`
}

type LimitIdentification5 struct {
	SysId          SystemIdentification2Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 SysId,omitempty"`
	BilLmtCtrPtyId BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 BilLmtCtrPtyId,omitempty"`
	Tp             LimitType1Choice                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Tp"`
	AcctOwnr       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 AcctOwnr,omitempty"`
	AcctId         AccountIdentification4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 AcctId,omitempty"`
}

type LimitOrError4Choice struct {
	Lmt    Limit7           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Lmt,omitempty"`
	BizErr []ErrorHandling5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 BizErr,omitempty"`
}

type LimitReport7 struct {
	LmtId    LimitIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 LmtId"`
	LmtOrErr LimitOrError4Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 LmtOrErr"`
}

type LimitReportOrError4Choice struct {
	BizRpt  Limits7          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 BizRpt,omitempty"`
	OprlErr []ErrorHandling5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 OprlErr,omitempty"`
}

// May be one of ENAB, DISA, DELD, REQD
type LimitStatus1Code string

type LimitType1Choice struct {
	Cd    LimitType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Cd,omitempty"`
	Prtry Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Prtry,omitempty"`
}

// May be one of MULT, BILI, MAND, DISC, NELI, INBI, GLBL, DIDB, SPLC, SPLF, TDLC, TDLF, UCDT, ACOL, EXGT
type LimitType3Code string

type Limits7 struct {
	CurLmt  []LimitReport7 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 CurLmt,omitempty"`
	DfltLmt []LimitReport7 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 DfltLmt,omitempty"`
}

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Cd,omitempty"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Prtry,omitempty"`
}

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 34 items long
type Max34Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be no more than 70 items long
type Max70Text string

type MessageHeader7 struct {
	MsgId       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 MsgId"`
	CreDtTm     ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 CreDtTm,omitempty"`
	ReqTp       RequestType4Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 ReqTp,omitempty"`
	OrgnlBizQry OriginalBusinessQuery1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 OrgnlBizQry,omitempty"`
	QryNm       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 QryNm,omitempty"`
}

type OriginalBusinessQuery1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 MsgId"`
	MsgNmId Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 MsgNmId,omitempty"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 CreDtTm,omitempty"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 AdrLine,omitempty"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 PmtCtrl,omitempty"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Enqry,omitempty"`
	Prtry   GenericIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Prtry,omitempty"`
}

type ReturnLimitV08 struct {
	MsgHdr      MessageHeader7            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 MsgHdr"`
	RptOrErr    LimitReportOrError4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 RptOrErr"`
	SplmtryData []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 SplmtryData,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemIdentification2Choice struct {
	MktInfrstrctrId MarketInfrastructureIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 MktInfrstrctrId,omitempty"`
	Ctry            CountryCode                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 Ctry,omitempty"`
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
