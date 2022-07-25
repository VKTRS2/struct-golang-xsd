// Code generated by download. DO NOT EDIT.

package iso20022_fxtr_016_001_04

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

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AgreedRate3 struct {
	XchgRate float64            `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 XchgRate"`
	UnitCcy  ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 UnitCcy,omitempty"`
	QtdCcy   ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 QtdCcy,omitempty"`
}

type AgreementConditions1 struct {
	AgrmtCd Max6AlphaText     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AgrmtCd"`
	Dt      ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Dt,omitempty"`
	Vrsn    Exact4NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Vrsn,omitempty"`
}

// May be one of POST, PREA, UNAL
type AllocationIndicator1Code string

type AmountsAndValueDate1 struct {
	TradgSdBuyAmt  ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TradgSdBuyAmt"`
	TradgSdSellAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TradgSdSellAmt"`
	SttlmDt        ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 SttlmDt"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type ClearingBrokerIdentification1 struct {
	SdInd     SideIndicator1Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 SdInd"`
	ClrBrkrId Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ClrBrkrId"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Cd,omitempty"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Prtry,omitempty"`
}

// May be one of FULL, ONEW, PART, UNCO
type CollateralisationIndicator1Code string

type ContactInformation1 struct {
	Nm       Max350Text  `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Nm,omitempty"`
	FaxNb    PhoneNumber `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 FaxNb,omitempty"`
	TelNb    PhoneNumber `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TelNb,omitempty"`
	EmailAdr Max256Text  `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 EmailAdr,omitempty"`
}

// May be one of L, A, C, I, F, O, R, U
type CorporateSectorIdentifier1Code string

type CounterpartySideTransactionReporting1 struct {
	RptgJursdctn     Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 RptgJursdctn,omitempty"`
	RptgPty          PartyIdentification73Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 RptgPty,omitempty"`
	CtrPtySdUnqTxIdr []UniqueTransactionIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CtrPtySdUnqTxIdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Dt,omitempty"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 DtTm,omitempty"`
}

type Document struct {
	FXTradInstrCxl ForeignExchangeTradeInstructionCancellationV04 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 FXTradInstrCxl"`
}

// Must match the pattern [a-zA-Z0-9]{2}
type Exact2AlphaNumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// May be no more than 4 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type ForeignExchangeTradeInstructionCancellationV04 struct {
	TradInf             TradeAgreement15                 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TradInf"`
	TradgSdId           TradePartyIdentification6        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TradgSdId"`
	CtrPtySdId          TradePartyIdentification6        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CtrPtySdId"`
	AgrdRate            AgreedRate3                      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AgrdRate,omitempty"`
	NDFConds            NonDeliverableForwardConditions1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 NDFConds,omitempty"`
	TradgSdSttlmInstrs  SettlementParties29              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TradgSdSttlmInstrs,omitempty"`
	CtrPtySdSttlmInstrs SettlementParties29              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CtrPtySdSttlmInstrs,omitempty"`
	OptnlGnlInf         GeneralInformation5              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 OptnlGnlInf,omitempty"`
	TradAmts            AmountsAndValueDate1             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TradAmts"`
	RgltryRptg          RegulatoryReporting6             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 RgltryRptg,omitempty"`
	SplmtryData         []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 SplmtryData,omitempty"`
}

type FundIdentification4 struct {
	FndId         PartyIdentification60       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 FndId"`
	AcctIdWthCtdn Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AcctIdWthCtdn,omitempty"`
	CtdnId        PartyIdentification73Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CtdnId,omitempty"`
}

type GeneralInformation5 struct {
	BlckInd            bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 BlckInd,omitempty"`
	RltdTradRef        Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 RltdTradRef,omitempty"`
	DealgMtd           Trading1MethodCode          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 DealgMtd,omitempty"`
	BrkrId             PartyIdentification73Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 BrkrId,omitempty"`
	CtrPtyRef          Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CtrPtyRef,omitempty"`
	BrkrsComssn        ActiveCurrencyAndAmount     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 BrkrsComssn,omitempty"`
	SndrToRcvrInf      Max210Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 SndrToRcvrInf,omitempty"`
	DealgBrnchTradgSd  PartyIdentification73Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 DealgBrnchTradgSd,omitempty"`
	DealgBrnchCtrPtySd PartyIdentification73Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 DealgBrnchCtrPtySd,omitempty"`
	CtctInf            ContactInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CtctInf,omitempty"`
	AgrmtDtls          AgreementConditions1        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AgrmtDtls,omitempty"`
	DefsYr             ISOYear                     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 DefsYr,omitempty"`
	BrkrsRef           Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 BrkrsRef,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

type ISOYear time.Time

func (t *ISOYear) UnmarshalText(text []byte) error {
	return (*xsdGYear)(t).UnmarshalText(text)
}
func (t ISOYear) MarshalText() ([]byte, error) {
	return xsdGYear(t).MarshalText()
}

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Cd,omitempty"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Prtry,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type MatchingSystemReference1Choice struct {
	MtchgSysUnqRef Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 MtchgSysUnqRef,omitempty"`
	RltdRef        Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 RltdRef,omitempty"`
}

// May be no more than 105 items long
type Max105Text string

// May be no more than 10 items long
type Max10Text string

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 210 items long
type Max210Text string

// May be no more than 256 items long
type Max256Text string

// May be no more than 34 items long
type Max34Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be no more than 4 items long
type Max4Text string

// May be no more than 52 items long
type Max52Text string

// Must match the pattern [a-zA-Z]{1,6}
type Max6AlphaText string

// May be no more than 70 items long
type Max70Text string

type NDFOpeningFixing1Choice struct {
	OpngConds   OpeningConditions1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 OpngConds,omitempty"`
	OpngConfRef Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 OpngConfRef,omitempty"`
}

type NameAndAddress8 struct {
	Nm         Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Nm"`
	Adr        PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Adr,omitempty"`
	AltrntvIdr []Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AltrntvIdr,omitempty"`
}

type NonDeliverableForwardConditions1 struct {
	OpngInd      bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 OpngInd"`
	OpngFxgConds NDFOpeningFixing1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 OpngFxgConds"`
}

type OpeningConditions1 struct {
	SttlmCcy     ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 SttlmCcy"`
	ValtnDt      ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ValtnDt"`
	SttlmRateSrc []SettlementRateSource1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 SttlmRateSrc"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Tp"`
}

type PartyIdentification44 struct {
	AnyBIC     AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AnyBIC"`
	AltrntvIdr []Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AltrntvIdr,omitempty"`
}

type PartyIdentification59 struct {
	PtyNm      Max34Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 PtyNm,omitempty"`
	AnyBIC     PartyIdentification44               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AnyBIC,omitempty"`
	AcctNb     Max34Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AcctNb,omitempty"`
	Adr        Max105Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Adr,omitempty"`
	ClrSysId   ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ClrSysId,omitempty"`
	LglNttyIdr LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 LglNttyIdr,omitempty"`
}

type PartyIdentification60 struct {
	FndId      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 FndId"`
	NmAndAdr   NameAndAddress8 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 NmAndAdr,omitempty"`
	LglNttyIdr LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 LglNttyIdr,omitempty"`
}

type PartyIdentification73Choice struct {
	NmAndAdr NameAndAddress8       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 NmAndAdr,omitempty"`
	AnyBIC   PartyIdentification44 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AnyBIC,omitempty"`
	PtyId    PartyIdentification59 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 PtyId,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Ctry"`
}

// Must match the pattern [a-zA-Z]{3}[0-9]{1,2}
type RateSourceText string

type RegulatoryReporting6 struct {
	TradgSdTxRptg          []TradingSideTransactionReporting1      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TradgSdTxRptg,omitempty"`
	CtrPtySdTxRptg         []CounterpartySideTransactionReporting1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CtrPtySdTxRptg,omitempty"`
	CntrlCtrPtyClrHs       PartyIdentification73Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CntrlCtrPtyClrHs,omitempty"`
	ClrBrkr                PartyIdentification73Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ClrBrkr,omitempty"`
	ClrXcptnPty            PartyIdentification73Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ClrXcptnPty,omitempty"`
	ClrBrkrId              ClearingBrokerIdentification1           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ClrBrkrId,omitempty"`
	ClrThrshldInd          bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ClrThrshldInd,omitempty"`
	ClrdPdctId             Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ClrdPdctId,omitempty"`
	UndrlygPdctIdr         UnderlyingProductIdentifier1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 UndrlygPdctIdr,omitempty"`
	AllcnInd               AllocationIndicator1Code                `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AllcnInd,omitempty"`
	CollstnInd             CollateralisationIndicator1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CollstnInd,omitempty"`
	ExctnVn                Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ExctnVn,omitempty"`
	ExctnTmstmp            DateAndDateTimeChoice                   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ExctnTmstmp,omitempty"`
	NonStdFlg              bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 NonStdFlg,omitempty"`
	LkSwpId                string                                  `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 LkSwpId,omitempty"`
	FinNtrOfTheCtrPtyInd   bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 FinNtrOfTheCtrPtyInd,omitempty"`
	CollPrtflInd           bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CollPrtflInd,omitempty"`
	CollPrtflCd            Max10Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CollPrtflCd,omitempty"`
	PrtflCmprssnInd        bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 PrtflCmprssnInd,omitempty"`
	CorpSctrInd            CorporateSectorIdentifier1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CorpSctrInd,omitempty"`
	TradWthNonEEACtrPtyInd bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TradWthNonEEACtrPtyInd,omitempty"`
	NtrgrpTradInd          bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 NtrgrpTradInd,omitempty"`
	ComrclOrTrsrFincgInd   bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ComrclOrTrsrFincgInd,omitempty"`
	FinInstrmId            SecurityIdentification19                `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 FinInstrmId,omitempty"`
	ConfDtAndTmstmp        ISODateTime                             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ConfDtAndTmstmp,omitempty"`
	ClrTmstmp              ISOTime                                 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ClrTmstmp,omitempty"`
	AddtlRptgInf           Max210Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AddtlRptgInf,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Desc,omitempty"`
}

type SettlementParties29 struct {
	DlvryAgt    PartyIdentification73Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 DlvryAgt,omitempty"`
	Intrmy      PartyIdentification73Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Intrmy,omitempty"`
	RcvgAgt     PartyIdentification73Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 RcvgAgt"`
	BnfcryInstn PartyIdentification73Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 BnfcryInstn,omitempty"`
}

type SettlementRateSource1 struct {
	RateSrc RateSourceText         `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 RateSrc"`
	Tm      Exact4NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Tm,omitempty"`
	CtryCd  CountryCode            `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CtryCd,omitempty"`
	LctnCd  Exact2AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 LctnCd,omitempty"`
}

// May be one of CCPL, CLNT
type SideIndicator1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeAgreement15 struct {
	TradDt        ISODate                        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TradDt"`
	OrgtrRef      Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 OrgtrRef"`
	MtchgSysRef   MatchingSystemReference1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 MtchgSysRef"`
	CmonRef       Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 CmonRef,omitempty"`
	AmdOrCclRsn   Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 AmdOrCclRsn,omitempty"`
	OprTp         Max4Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 OprTp,omitempty"`
	OprScp        Max4Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 OprScp,omitempty"`
	PdctTp        Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 PdctTp,omitempty"`
	SttlmSsnIdr   Exact4AlphaNumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 SttlmSsnIdr,omitempty"`
	PmtVrssPmtInd bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 PmtVrssPmtInd,omitempty"`
}

type TradePartyIdentification6 struct {
	SubmitgPty PartyIdentification73Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 SubmitgPty"`
	TradPty    PartyIdentification73Choice `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TradPty,omitempty"`
	FndId      []FundIdentification4       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 FndId,omitempty"`
}

// May be one of ELEC, PHON, BROK
type Trading1MethodCode string

type TradingSideTransactionReporting1 struct {
	RptgJursdctn    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 RptgJursdctn,omitempty"`
	RptgPty         PartyIdentification73Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 RptgPty,omitempty"`
	TradgSdUnqTxIdr []UniqueTransactionIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 TradgSdUnqTxIdr,omitempty"`
}

// May be one of FORW, NDFO, SPOT, SWAP
type UnderlyingProductIdentifier1Code string

type UniqueTransactionIdentifier2 struct {
	UnqTxIdr    Max52Text   `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 UnqTxIdr"`
	PrrUnqTxIdr []Max52Text `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 PrrUnqTxIdr,omitempty"`
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

type xsdGYear time.Time

func (t *xsdGYear) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006")
}
func (t xsdGYear) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006")
}
func (t xsdGYear) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYear) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
