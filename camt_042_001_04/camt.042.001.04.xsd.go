// Code generated by download. DO NOT EDIT.

package iso20022_camt_042_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification1 struct {
	Prtry SimpleIdentificationInformation `xml:"urn:swift:xsd:camt.042.001.04 Prtry"`
}

type ActiveCurrencyAnd13DecimalAmount struct {
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

type AdditionalParameters1 struct {
	Ctry    CountryCode                  `xml:"urn:swift:xsd:camt.042.001.04 Ctry,omitempty"`
	Ccy     ActiveOrHistoricCurrencyCode `xml:"urn:swift:xsd:camt.042.001.04 Ccy,omitempty"`
	GeoArea Max35Text                    `xml:"urn:swift:xsd:camt.042.001.04 GeoArea,omitempty"`
}

type AdditionalReference3 struct {
	Ref     Max35Text                  `xml:"urn:swift:xsd:camt.042.001.04 Ref"`
	RefIssr PartyIdentification2Choice `xml:"urn:swift:xsd:camt.042.001.04 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:swift:xsd:camt.042.001.04 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateSecurityIdentification1 struct {
	Id         Max35Text   `xml:"urn:swift:xsd:camt.042.001.04 Id"`
	DmstIdSrc  CountryCode `xml:"urn:swift:xsd:camt.042.001.04 DmstIdSrc,omitempty"`
	PrtryIdSrc Max35Text   `xml:"urn:swift:xsd:camt.042.001.04 PrtryIdSrc,omitempty"`
}

type AmountOrRate3Choice struct {
	Amt  ActiveCurrencyAnd13DecimalAmount `xml:"urn:swift:xsd:camt.042.001.04 Amt,omitempty"`
	Rate float64                          `xml:"urn:swift:xsd:camt.042.001.04 Rate,omitempty"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be no more than 35 items long
type BloombergIdentifier string

type BreakdownByCountry2 struct {
	Ctry       CountryCode        `xml:"urn:swift:xsd:camt.042.001.04 Ctry"`
	CshInFcst  []CashInForecast5  `xml:"urn:swift:xsd:camt.042.001.04 CshInFcst,omitempty"`
	CshOutFcst []CashOutForecast5 `xml:"urn:swift:xsd:camt.042.001.04 CshOutFcst,omitempty"`
	NetCshFcst []NetCashForecast4 `xml:"urn:swift:xsd:camt.042.001.04 NetCshFcst,omitempty"`
}

type BreakdownByCurrency2 struct {
	Ccy        ActiveOrHistoricCurrencyCode `xml:"urn:swift:xsd:camt.042.001.04 Ccy"`
	CshOutFcst []CashOutForecast5           `xml:"urn:swift:xsd:camt.042.001.04 CshOutFcst,omitempty"`
	CshInFcst  []CashInForecast5            `xml:"urn:swift:xsd:camt.042.001.04 CshInFcst,omitempty"`
	NetCshFcst []NetCashForecast4           `xml:"urn:swift:xsd:camt.042.001.04 NetCshFcst,omitempty"`
}

type BreakdownByParty3 struct {
	Pty         InvestmentAccount42   `xml:"urn:swift:xsd:camt.042.001.04 Pty"`
	AddtlParams AdditionalParameters1 `xml:"urn:swift:xsd:camt.042.001.04 AddtlParams,omitempty"`
	CshInFcst   []CashInForecast5     `xml:"urn:swift:xsd:camt.042.001.04 CshInFcst,omitempty"`
	CshOutFcst  []CashOutForecast5    `xml:"urn:swift:xsd:camt.042.001.04 CshOutFcst,omitempty"`
	NetCshFcst  []NetCashForecast4    `xml:"urn:swift:xsd:camt.042.001.04 NetCshFcst,omitempty"`
}

type BreakdownByUserDefinedParameter3 struct {
	Pty        InvestmentAccount42          `xml:"urn:swift:xsd:camt.042.001.04 Pty,omitempty"`
	Ctry       CountryCode                  `xml:"urn:swift:xsd:camt.042.001.04 Ctry,omitempty"`
	Ccy        ActiveOrHistoricCurrencyCode `xml:"urn:swift:xsd:camt.042.001.04 Ccy,omitempty"`
	UsrDfnd    DataFormat2Choice            `xml:"urn:swift:xsd:camt.042.001.04 UsrDfnd,omitempty"`
	CshInFcst  []CashInForecast5            `xml:"urn:swift:xsd:camt.042.001.04 CshInFcst,omitempty"`
	CshOutFcst []CashOutForecast5           `xml:"urn:swift:xsd:camt.042.001.04 CshOutFcst,omitempty"`
	NetCshFcst []NetCashForecast4           `xml:"urn:swift:xsd:camt.042.001.04 NetCshFcst,omitempty"`
}

type CashInForecast5 struct {
	CshSttlmDt       ISODate                           `xml:"urn:swift:xsd:camt.042.001.04 CshSttlmDt"`
	SubTtlAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:camt.042.001.04 SubTtlAmt,omitempty"`
	SubTtlUnitsNb    FinancialInstrumentQuantity1      `xml:"urn:swift:xsd:camt.042.001.04 SubTtlUnitsNb,omitempty"`
	XcptnlCshFlowInd bool                              `xml:"urn:swift:xsd:camt.042.001.04 XcptnlCshFlowInd,omitempty"`
	CshInBrkdwnDtls  []FundCashInBreakdown3            `xml:"urn:swift:xsd:camt.042.001.04 CshInBrkdwnDtls,omitempty"`
	AddtlBal         FundBalance1                      `xml:"urn:swift:xsd:camt.042.001.04 AddtlBal,omitempty"`
}

type CashOutForecast5 struct {
	CshSttlmDt       ISODate                           `xml:"urn:swift:xsd:camt.042.001.04 CshSttlmDt"`
	SubTtlAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:camt.042.001.04 SubTtlAmt,omitempty"`
	SubTtlUnitsNb    FinancialInstrumentQuantity1      `xml:"urn:swift:xsd:camt.042.001.04 SubTtlUnitsNb,omitempty"`
	XcptnlCshFlowInd bool                              `xml:"urn:swift:xsd:camt.042.001.04 XcptnlCshFlowInd,omitempty"`
	CshOutBrkdwnDtls []FundCashOutBreakdown3           `xml:"urn:swift:xsd:camt.042.001.04 CshOutBrkdwnDtls,omitempty"`
	AddtlBal         FundBalance1                      `xml:"urn:swift:xsd:camt.042.001.04 AddtlBal,omitempty"`
}

type Charge26 struct {
	Tp       ChargeType4Choice   `xml:"urn:swift:xsd:camt.042.001.04 Tp"`
	ChrgApld AmountOrRate3Choice `xml:"urn:swift:xsd:camt.042.001.04 ChrgApld"`
}

// May be one of BEND, DISC, FEND, POST, REGF, SHIP, SPCN, TRAN
type ChargeType12Code string

type ChargeType4Choice struct {
	Cd    ChargeType12Code        `xml:"urn:swift:xsd:camt.042.001.04 Cd,omitempty"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:camt.042.001.04 Prtry,omitempty"`
}

type Commission21 struct {
	ComssnTp   CommissionType5Choice `xml:"urn:swift:xsd:camt.042.001.04 ComssnTp"`
	ComssnApld AmountOrRate3Choice   `xml:"urn:swift:xsd:camt.042.001.04 ComssnApld"`
}

type CommissionType5Choice struct {
	Cd    CommissionType6Code     `xml:"urn:swift:xsd:camt.042.001.04 Cd,omitempty"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:camt.042.001.04 Prtry,omitempty"`
}

// May be one of FEND, BEND, CDPL
type CommissionType6Code string

// May be no more than 35 items long
type ConsolidatedTapeAssociationIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyDesignation1 struct {
	CcyDsgnt CurrencyDesignation1Code `xml:"urn:swift:xsd:camt.042.001.04 CcyDsgnt,omitempty"`
	Lctn     CountryCode              `xml:"urn:swift:xsd:camt.042.001.04 Lctn,omitempty"`
	AddtlInf Max350Text               `xml:"urn:swift:xsd:camt.042.001.04 AddtlInf,omitempty"`
}

// May be one of ONSH, OFFS
type CurrencyDesignation1Code string

type DataFormat2Choice struct {
	Strd  GenericIdentification1 `xml:"urn:swift:xsd:camt.042.001.04 Strd,omitempty"`
	Ustrd Max140Text             `xml:"urn:swift:xsd:camt.042.001.04 Ustrd,omitempty"`
}

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:swift:xsd:camt.042.001.04 Dt,omitempty"`
	DtTm ISODateTime `xml:"urn:swift:xsd:camt.042.001.04 DtTm,omitempty"`
}

// May be one of DIST, ACCU
type DistributionPolicy1Code string

type Document struct {
	FndDtldEstmtdCshFcstRpt FundDetailedEstimatedCashForecastReportV04 `xml:"urn:swift:xsd:camt.042.001.04 FndDtldEstmtdCshFcstRpt"`
}

type EstimatedFundCashForecast5 struct {
	Id                        Max35Text                           `xml:"urn:swift:xsd:camt.042.001.04 Id"`
	TradDtTm                  DateAndDateTimeChoice               `xml:"urn:swift:xsd:camt.042.001.04 TradDtTm"`
	PrvsTradDtTm              DateAndDateTimeChoice               `xml:"urn:swift:xsd:camt.042.001.04 PrvsTradDtTm,omitempty"`
	FinInstrmDtls             FinancialInstrument9                `xml:"urn:swift:xsd:camt.042.001.04 FinInstrmDtls"`
	EstmtdTtlNAV              []ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:camt.042.001.04 EstmtdTtlNAV,omitempty"`
	PrvsTtlNAV                []ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:camt.042.001.04 PrvsTtlNAV,omitempty"`
	EstmtdTtlUnitsNb          FinancialInstrumentQuantity1        `xml:"urn:swift:xsd:camt.042.001.04 EstmtdTtlUnitsNb,omitempty"`
	PrvsTtlUnitsNb            FinancialInstrumentQuantity1        `xml:"urn:swift:xsd:camt.042.001.04 PrvsTtlUnitsNb,omitempty"`
	EstmtdTtlNAVChngRate      float64                             `xml:"urn:swift:xsd:camt.042.001.04 EstmtdTtlNAVChngRate,omitempty"`
	InvstmtCcy                []ActiveOrHistoricCurrencyCode      `xml:"urn:swift:xsd:camt.042.001.04 InvstmtCcy,omitempty"`
	CcySts                    CurrencyDesignation1                `xml:"urn:swift:xsd:camt.042.001.04 CcySts,omitempty"`
	XcptnlNetCshFlowInd       bool                                `xml:"urn:swift:xsd:camt.042.001.04 XcptnlNetCshFlowInd"`
	Pric                      UnitPrice19                         `xml:"urn:swift:xsd:camt.042.001.04 Pric,omitempty"`
	FXRate                    ForeignExchangeTerms19              `xml:"urn:swift:xsd:camt.042.001.04 FXRate,omitempty"`
	EstmtdPctgOfShrClssTtlNAV float64                             `xml:"urn:swift:xsd:camt.042.001.04 EstmtdPctgOfShrClssTtlNAV,omitempty"`
	BrkdwnByPty               []BreakdownByParty3                 `xml:"urn:swift:xsd:camt.042.001.04 BrkdwnByPty,omitempty"`
	BrkdwnByCtry              []BreakdownByCountry2               `xml:"urn:swift:xsd:camt.042.001.04 BrkdwnByCtry,omitempty"`
	BrkdwnByCcy               []BreakdownByCurrency2              `xml:"urn:swift:xsd:camt.042.001.04 BrkdwnByCcy,omitempty"`
	BrkdwnByUsrDfndParam      []BreakdownByUserDefinedParameter3  `xml:"urn:swift:xsd:camt.042.001.04 BrkdwnByUsrDfndParam,omitempty"`
	EstmtdNetCshFcstDtls      []NetCashForecast4                  `xml:"urn:swift:xsd:camt.042.001.04 EstmtdNetCshFcstDtls,omitempty"`
}

// May be no more than 12 items long
type EuroclearClearstreamIdentifier string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:swift:xsd:camt.042.001.04 PlcAndNm"`
	Txt      Max350Text `xml:"urn:swift:xsd:camt.042.001.04 Txt"`
}

type FinancialInstrument9 struct {
	Id          SecurityIdentification3Choice `xml:"urn:swift:xsd:camt.042.001.04 Id"`
	Nm          Max350Text                    `xml:"urn:swift:xsd:camt.042.001.04 Nm,omitempty"`
	SplmtryId   Max35Text                     `xml:"urn:swift:xsd:camt.042.001.04 SplmtryId,omitempty"`
	ReqdNAVCcy  ActiveOrHistoricCurrencyCode  `xml:"urn:swift:xsd:camt.042.001.04 ReqdNAVCcy,omitempty"`
	ClssTp      Max35Text                     `xml:"urn:swift:xsd:camt.042.001.04 ClssTp,omitempty"`
	SctiesForm  FormOfSecurity1Code           `xml:"urn:swift:xsd:camt.042.001.04 SctiesForm,omitempty"`
	DstrbtnPlcy DistributionPolicy1Code       `xml:"urn:swift:xsd:camt.042.001.04 DstrbtnPlcy,omitempty"`
	DualFndInd  bool                          `xml:"urn:swift:xsd:camt.042.001.04 DualFndInd"`
}

type FinancialInstrumentQuantity1 struct {
	Unit float64 `xml:"urn:swift:xsd:camt.042.001.04 Unit"`
}

// May be one of INCG, OUTG
type FlowDirectionType1Code string

type ForeignExchangeTerms19 struct {
	UnitCcy  ActiveCurrencyCode `xml:"urn:swift:xsd:camt.042.001.04 UnitCcy"`
	QtdCcy   ActiveCurrencyCode `xml:"urn:swift:xsd:camt.042.001.04 QtdCcy"`
	XchgRate float64            `xml:"urn:swift:xsd:camt.042.001.04 XchgRate"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type Fund3 struct {
	Nm                    Max350Text                        `xml:"urn:swift:xsd:camt.042.001.04 Nm,omitempty"`
	LglNttyIdr            LEIIdentifier                     `xml:"urn:swift:xsd:camt.042.001.04 LglNttyIdr,omitempty"`
	Id                    OtherIdentification4              `xml:"urn:swift:xsd:camt.042.001.04 Id,omitempty"`
	Ccy                   ActiveOrHistoricCurrencyCode      `xml:"urn:swift:xsd:camt.042.001.04 Ccy,omitempty"`
	EstmtdTtlNAV          ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:camt.042.001.04 EstmtdTtlNAV,omitempty"`
	PrvsTtlNAV            ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:camt.042.001.04 PrvsTtlNAV,omitempty"`
	EstmtdTtlUnitsNb      FinancialInstrumentQuantity1      `xml:"urn:swift:xsd:camt.042.001.04 EstmtdTtlUnitsNb,omitempty"`
	PrvsTtlUnitsNb        FinancialInstrumentQuantity1      `xml:"urn:swift:xsd:camt.042.001.04 PrvsTtlUnitsNb,omitempty"`
	EstmtdPctgOfFndTtlNAV float64                           `xml:"urn:swift:xsd:camt.042.001.04 EstmtdPctgOfFndTtlNAV,omitempty"`
}

type FundBalance1 struct {
	TtlUnitsFrUnitOrdrs FinancialInstrumentQuantity1      `xml:"urn:swift:xsd:camt.042.001.04 TtlUnitsFrUnitOrdrs,omitempty"`
	TtlUnitsFrCshOrdrs  FinancialInstrumentQuantity1      `xml:"urn:swift:xsd:camt.042.001.04 TtlUnitsFrCshOrdrs,omitempty"`
	TtlCshFrUnitOrdrs   ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:camt.042.001.04 TtlCshFrUnitOrdrs,omitempty"`
	TtlCshFrCshOrdrs    ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:camt.042.001.04 TtlCshFrCshOrdrs,omitempty"`
}

type FundCashInBreakdown3 struct {
	Amt              ActiveOrHistoricCurrencyAndAmount      `xml:"urn:swift:xsd:camt.042.001.04 Amt,omitempty"`
	UnitsNb          FinancialInstrumentQuantity1           `xml:"urn:swift:xsd:camt.042.001.04 UnitsNb,omitempty"`
	NewAmtInd        bool                                   `xml:"urn:swift:xsd:camt.042.001.04 NewAmtInd,omitempty"`
	InvstmtFndTxInTp InvestmentFundTransactionInType1Choice `xml:"urn:swift:xsd:camt.042.001.04 InvstmtFndTxInTp"`
	OrgnlOrdrQtyTp   QuantityType1Choice                    `xml:"urn:swift:xsd:camt.042.001.04 OrgnlOrdrQtyTp"`
	ChrgDtls         []Charge26                             `xml:"urn:swift:xsd:camt.042.001.04 ChrgDtls,omitempty"`
	ComssnDtls       []Commission21                         `xml:"urn:swift:xsd:camt.042.001.04 ComssnDtls,omitempty"`
	SttlmCcy         ActiveCurrencyCode                     `xml:"urn:swift:xsd:camt.042.001.04 SttlmCcy,omitempty"`
}

type FundCashOutBreakdown3 struct {
	Amt               ActiveOrHistoricCurrencyAndAmount       `xml:"urn:swift:xsd:camt.042.001.04 Amt,omitempty"`
	UnitsNb           FinancialInstrumentQuantity1            `xml:"urn:swift:xsd:camt.042.001.04 UnitsNb,omitempty"`
	NewAmtInd         bool                                    `xml:"urn:swift:xsd:camt.042.001.04 NewAmtInd,omitempty"`
	InvstmtFndTxOutTp InvestmentFundTransactionOutType1Choice `xml:"urn:swift:xsd:camt.042.001.04 InvstmtFndTxOutTp"`
	OrgnlOrdrQtyTp    QuantityType1Choice                     `xml:"urn:swift:xsd:camt.042.001.04 OrgnlOrdrQtyTp"`
	ChrgDtls          []Charge26                              `xml:"urn:swift:xsd:camt.042.001.04 ChrgDtls,omitempty"`
	ComssnDtls        []Commission21                          `xml:"urn:swift:xsd:camt.042.001.04 ComssnDtls,omitempty"`
	SttlmCcy          ActiveCurrencyCode                      `xml:"urn:swift:xsd:camt.042.001.04 SttlmCcy,omitempty"`
}

type FundDetailedEstimatedCashForecastReportV04 struct {
	MsgId                MessageIdentification1       `xml:"urn:swift:xsd:camt.042.001.04 MsgId"`
	PoolRef              AdditionalReference3         `xml:"urn:swift:xsd:camt.042.001.04 PoolRef,omitempty"`
	PrvsRef              []AdditionalReference3       `xml:"urn:swift:xsd:camt.042.001.04 PrvsRef,omitempty"`
	RltdRef              []AdditionalReference3       `xml:"urn:swift:xsd:camt.042.001.04 RltdRef,omitempty"`
	MsgPgntn             Pagination                   `xml:"urn:swift:xsd:camt.042.001.04 MsgPgntn"`
	FndOrSubFndDtls      Fund3                        `xml:"urn:swift:xsd:camt.042.001.04 FndOrSubFndDtls,omitempty"`
	EstmtdFndCshFcstDtls []EstimatedFundCashForecast5 `xml:"urn:swift:xsd:camt.042.001.04 EstmtdFndCshFcstDtls"`
	CnsltdNetCshFcst     NetCashForecast3             `xml:"urn:swift:xsd:camt.042.001.04 CnsltdNetCshFcst,omitempty"`
	Xtnsn                []Extension1                 `xml:"urn:swift:xsd:camt.042.001.04 Xtnsn,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:swift:xsd:camt.042.001.04 Id"`
	SchmeNm Max35Text `xml:"urn:swift:xsd:camt.042.001.04 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:swift:xsd:camt.042.001.04 Issr,omitempty"`
}

type GenericIdentification47 struct {
	Id      Exact4AlphaNumericText `xml:"urn:swift:xsd:camt.042.001.04 Id"`
	Issr    Max4AlphaNumericText   `xml:"urn:swift:xsd:camt.042.001.04 Issr"`
	SchmeNm Max4AlphaNumericText   `xml:"urn:swift:xsd:camt.042.001.04 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

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

type IdentificationSource5Choice struct {
	DmstIdSrc  CountryCode `xml:"urn:swift:xsd:camt.042.001.04 DmstIdSrc,omitempty"`
	PrtryIdSrc Max35Text   `xml:"urn:swift:xsd:camt.042.001.04 PrtryIdSrc,omitempty"`
}

type InvestmentAccount42 struct {
	AcctId   AccountIdentification1     `xml:"urn:swift:xsd:camt.042.001.04 AcctId,omitempty"`
	OwnrId   PartyIdentification2Choice `xml:"urn:swift:xsd:camt.042.001.04 OwnrId,omitempty"`
	AcctSvcr PartyIdentification2Choice `xml:"urn:swift:xsd:camt.042.001.04 AcctSvcr,omitempty"`
}

type InvestmentFundTransactionInType1Choice struct {
	Cd    InvestmentFundTransactionInType1Code `xml:"urn:swift:xsd:camt.042.001.04 Cd,omitempty"`
	Prtry GenericIdentification47              `xml:"urn:swift:xsd:camt.042.001.04 Prtry,omitempty"`
}

// May be one of SUBS, SWII, INSP, CROI, RDIV
type InvestmentFundTransactionInType1Code string

type InvestmentFundTransactionOutType1Choice struct {
	Cd    InvestmentFundTransactionOutType1Code `xml:"urn:swift:xsd:camt.042.001.04 Cd,omitempty"`
	Prtry GenericIdentification47               `xml:"urn:swift:xsd:camt.042.001.04 Prtry,omitempty"`
}

// May be one of REDM, SWIO, INSP, CROO
type InvestmentFundTransactionOutType1Code string

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// May be no more than 70 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:swift:xsd:camt.042.001.04 Id"`
	CreDtTm ISODateTime `xml:"urn:swift:xsd:camt.042.001.04 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:swift:xsd:camt.042.001.04 Nm"`
	Adr PostalAddress1 `xml:"urn:swift:xsd:camt.042.001.04 Adr,omitempty"`
}

type NetCashForecast3 struct {
	NetAmt     ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:camt.042.001.04 NetAmt,omitempty"`
	NetUnitsNb FinancialInstrumentQuantity1      `xml:"urn:swift:xsd:camt.042.001.04 NetUnitsNb,omitempty"`
	FlowDrctn  FlowDirectionType1Code            `xml:"urn:swift:xsd:camt.042.001.04 FlowDrctn"`
}

type NetCashForecast4 struct {
	CshSttlmDt ISODate                           `xml:"urn:swift:xsd:camt.042.001.04 CshSttlmDt"`
	NetAmt     ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:camt.042.001.04 NetAmt,omitempty"`
	NetUnitsNb FinancialInstrumentQuantity1      `xml:"urn:swift:xsd:camt.042.001.04 NetUnitsNb,omitempty"`
	FlowDrctn  FlowDirectionType1Code            `xml:"urn:swift:xsd:camt.042.001.04 FlowDrctn"`
	AddtlBal   FundBalance1                      `xml:"urn:swift:xsd:camt.042.001.04 AddtlBal,omitempty"`
}

// May be one of UNIT, CASH
type OrderQuantityType2Code string

type OtherIdentification4 struct {
	Id Max35Text                   `xml:"urn:swift:xsd:camt.042.001.04 Id"`
	Tp IdentificationSource5Choice `xml:"urn:swift:xsd:camt.042.001.04 Tp"`
}

type Pagination struct {
	PgNb      Max5NumericText `xml:"urn:swift:xsd:camt.042.001.04 PgNb"`
	LastPgInd bool            `xml:"urn:swift:xsd:camt.042.001.04 LastPgInd"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:swift:xsd:camt.042.001.04 BICOrBEI,omitempty"`
	PrtryId  GenericIdentification1 `xml:"urn:swift:xsd:camt.042.001.04 PrtryId,omitempty"`
	NmAndAdr NameAndAddress5        `xml:"urn:swift:xsd:camt.042.001.04 NmAndAdr,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:swift:xsd:camt.042.001.04 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:swift:xsd:camt.042.001.04 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:swift:xsd:camt.042.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:swift:xsd:camt.042.001.04 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:swift:xsd:camt.042.001.04 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:swift:xsd:camt.042.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:swift:xsd:camt.042.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:swift:xsd:camt.042.001.04 Ctry"`
}

type PriceValue1 struct {
	Amt ActiveCurrencyAnd13DecimalAmount `xml:"urn:swift:xsd:camt.042.001.04 Amt"`
}

type QuantityType1Choice struct {
	Cd    OrderQuantityType2Code  `xml:"urn:swift:xsd:camt.042.001.04 Cd,omitempty"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:camt.042.001.04 Prtry,omitempty"`
}

// May be no more than 35 items long
type RICIdentifier string

type SecurityIdentification3Choice struct {
	ISIN        ISINIdentifier                        `xml:"urn:swift:xsd:camt.042.001.04 ISIN,omitempty"`
	SEDOL       string                                `xml:"urn:swift:xsd:camt.042.001.04 SEDOL,omitempty"`
	CUSIP       string                                `xml:"urn:swift:xsd:camt.042.001.04 CUSIP,omitempty"`
	RIC         RICIdentifier                         `xml:"urn:swift:xsd:camt.042.001.04 RIC,omitempty"`
	TckrSymb    TickerIdentifier                      `xml:"urn:swift:xsd:camt.042.001.04 TckrSymb,omitempty"`
	Blmbrg      BloombergIdentifier                   `xml:"urn:swift:xsd:camt.042.001.04 Blmbrg,omitempty"`
	CTA         ConsolidatedTapeAssociationIdentifier `xml:"urn:swift:xsd:camt.042.001.04 CTA,omitempty"`
	QUICK       string                                `xml:"urn:swift:xsd:camt.042.001.04 QUICK,omitempty"`
	Wrtppr      string                                `xml:"urn:swift:xsd:camt.042.001.04 Wrtppr,omitempty"`
	Dtch        string                                `xml:"urn:swift:xsd:camt.042.001.04 Dtch,omitempty"`
	Vlrn        string                                `xml:"urn:swift:xsd:camt.042.001.04 Vlrn,omitempty"`
	SCVM        string                                `xml:"urn:swift:xsd:camt.042.001.04 SCVM,omitempty"`
	Belgn       string                                `xml:"urn:swift:xsd:camt.042.001.04 Belgn,omitempty"`
	Cmon        EuroclearClearstreamIdentifier        `xml:"urn:swift:xsd:camt.042.001.04 Cmon,omitempty"`
	OthrPrtryId AlternateSecurityIdentification1      `xml:"urn:swift:xsd:camt.042.001.04 OthrPrtryId,omitempty"`
}

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:swift:xsd:camt.042.001.04 Id"`
}

// May be no more than 35 items long
type TickerIdentifier string

// May be one of BIDE, OFFR, NAVL, CREA, CANC, INTE, SWNG, MIDD, RINV, SWIC, DDVR, ACTU
type TypeOfPrice10Code string

type UnitPrice19 struct {
	PricTp UnitPriceType2Choice `xml:"urn:swift:xsd:camt.042.001.04 PricTp"`
	Val    PriceValue1          `xml:"urn:swift:xsd:camt.042.001.04 Val"`
}

type UnitPriceType2Choice struct {
	Cd    TypeOfPrice10Code       `xml:"urn:swift:xsd:camt.042.001.04 Cd,omitempty"`
	Prtry GenericIdentification47 `xml:"urn:swift:xsd:camt.042.001.04 Prtry,omitempty"`
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
