// Code generated by download. DO NOT EDIT.

package iso20022_colr_023_001_01

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

type AllocationStatus1Choice struct {
	FullyAllctd ProprietaryReason4          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 FullyAllctd,omitempty"`
	PrtlyAllctd ProprietaryReason4          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PrtlyAllctd,omitempty"`
	Prtry       ProprietaryStatusAndReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Prtry,omitempty"`
}

type AlternatePartyIdentification7 struct {
	IdTp    IdentificationType42Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 IdTp"`
	Ctry    CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Ctry"`
	AltrnId Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AltrnId"`
}

type AmountAndDirection44 struct {
	Amt                 ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms23            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 FXDtls,omitempty"`
}

type AmountAndDirection49 struct {
	Amt                 ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Amt"`
	CdtDbtInd           CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CdtDbtInd,omitempty"`
	OrgnlCcyAndOrdrdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 OrgnlCcyAndOrdrdAmt,omitempty"`
	FXDtls              ForeignExchangeTerms23            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 FXDtls,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type BlockChainAddressWallet3 struct {
	Id Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Nm,omitempty"`
}

type CashAccountIdentification5Choice struct {
	IBAN  IBAN2007Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 IBAN,omitempty"`
	Prtry Max34Text          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Prtry,omitempty"`
}

type CashMovement7 struct {
	CshMvmnt                   CreditDebit3Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CshMvmnt"`
	CshAmt                     ActiveCurrencyAndAmount          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CshAmt"`
	CshAcct                    CashAccountIdentification5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CshAcct,omitempty"`
	MvmntSts                   ProprietaryStatusAndReason6      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 MvmntSts,omitempty"`
	CollMvmnt                  bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CollMvmnt"`
	CshMvmntApprvd             bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CshMvmntApprvd,omitempty"`
	PosTp                      bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PosTp,omitempty"`
	ClntCshMvmntId             Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 ClntCshMvmntId,omitempty"`
	TrptyAgtSvcPrvdrCshMvmntId Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 TrptyAgtSvcPrvdrCshMvmntId,omitempty"`
}

type ClosingDate4Choice struct {
	Dt DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Dt,omitempty"`
	Cd Date3Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Cd,omitempty"`
}

type CollateralAmount14 struct {
	Tx       AmountAndDirection49 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Tx,omitempty"`
	Termntn  AmountAndDirection49 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Termntn,omitempty"`
	Acrd     AmountAndDirection49 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Acrd,omitempty"`
	ValSght  AmountAndDirection49 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 ValSght,omitempty"`
	UdsptdTx AmountAndDirection49 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 UdsptdTx,omitempty"`
}

type CollateralAmount5 struct {
	ReqrdMrgn  AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 ReqrdMrgn,omitempty"`
	Collsd     AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Collsd,omitempty"`
	RmngCollsd AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 RmngCollsd,omitempty"`
	Sttld      AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Sttld,omitempty"`
	RmngSttlm  AmountAndDirection44 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 RmngSttlm,omitempty"`
}

type CollateralDate2 struct {
	TradDt      ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 TradDt,omitempty"`
	ReqdExctnDt DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 ReqdExctnDt,omitempty"`
	SttlmDt     ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SttlmDt,omitempty"`
}

type CollateralParameters13 struct {
	CollInstrTp  CollateralTransactionType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CollInstrTp"`
	XpsrTp       ExposureType23Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 XpsrTp"`
	CollSd       CollateralRole1Code              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CollSd"`
	Prty         GenericIdentification30          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Prty,omitempty"`
	AutomtcAllcn bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AutomtcAllcn,omitempty"`
	CollApprvd   bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CollApprvd,omitempty"`
	SttlmApprvd  bool                             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SttlmApprvd,omitempty"`
	CollAmt      CollateralAmount5                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CollAmt,omitempty"`
}

type CollateralParties8 struct {
	PtyA     PartyIdentificationAndAccount202 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PtyA"`
	ClntPtyA PartyIdentificationAndAccount193 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 ClntPtyA,omitempty"`
	PtyB     PartyIdentificationAndAccount203 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PtyB"`
	ClntPtyB PartyIdentificationAndAccount193 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 ClntPtyB,omitempty"`
	TrptyAgt PartyIdentification136           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 TrptyAgt,omitempty"`
}

// May be one of GIVE, TAKE
type CollateralRole1Code string

type CollateralStatus3Choice struct {
	Pdg   []ProprietaryReason4          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Pdg,omitempty"`
	Prtry []ProprietaryStatusAndReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Prtry,omitempty"`
}

type CollateralTransactionType1Choice struct {
	Cd    CollateralTransactionType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Cd,omitempty"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Prtry,omitempty"`
}

// May be one of AADJ, CDTA, CADJ, DADJ, DBVT, INIT, MADJ, PADJ, RATA, TERM
type CollateralTransactionType1Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebit3Code string

// May be one of CRDT, DBIT
type CreditDebitCode string

type Date3Choice struct {
	Cd    DateType2Code           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Prtry,omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Dt,omitempty"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 DtTm,omitempty"`
}

// May be one of OPEN
type DateType2Code string

type DealTransactionDetails7 struct {
	ClsgDt      ClosingDate4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 ClsgDt"`
	DealDtlsAmt CollateralAmount14 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 DealDtlsAmt,omitempty"`
}

type Document struct {
	TrptyCollStsAdvc TripartyCollateralStatusAdviceV01 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 TrptyCollStsAdvc"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be one of BFWD, PAYM, CBCO, COMM, CRDS, CRTL, CRSP, CCIR, CRPR, EQPT, EQUS, EXTD, EXPT, FIXI, FORX, FORW, FUTR, OPTN, LIQU, OTCD, RVPO, SLOA, SBSC, SCRP, SLEB, SCIR, SCIE, SWPT, TBAS, TRCP, UDMS, CCPC, EQUI, TRBD, REPO, SHSL, MGLD
type ExposureType14Code string

type ExposureType23Choice struct {
	Cd    ExposureType14Code      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Prtry,omitempty"`
}

// May be no more than 4 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity33Choice struct {
	Unit        float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Unit,omitempty"`
	FaceAmt     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 FaceAmt,omitempty"`
	AmtsdVal    float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AmtsdVal,omitempty"`
	DgtlTknUnit float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 DgtlTknUnit,omitempty"`
}

type ForeignExchangeTerms23 struct {
	UnitCcy  ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 UnitCcy"`
	QtdCcy   ActiveCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 QtdCcy"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 XchgRate"`
	RsltgAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 RsltgAmt"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Cd,omitempty"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Prtry,omitempty"`
}

type IdentificationType42Choice struct {
	Cd    TypeOfIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Cd,omitempty"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Prtry,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 210 items long
type Max210Text string

// May be no more than 34 items long
type Max34Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be no more than 52 items long
type Max52Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// May be no more than 70 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Adr,omitempty"`
}

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Tp"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PgNb"`
	LastPgInd bool            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 LastPgInd"`
}

type PartyIdentification120Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AnyBIC,omitempty"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PrtryId,omitempty"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 NmAndAdr,omitempty"`
}

type PartyIdentification136 struct {
	Id  PartyIdentification120Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Id"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 LEI,omitempty"`
}

type PartyIdentificationAndAccount193 struct {
	Id      PartyIdentification120Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Id"`
	LEI     LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 LEI,omitempty"`
	AltrnId AlternatePartyIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AltrnId,omitempty"`
}

type PartyIdentificationAndAccount202 struct {
	Id                 PartyIdentification120Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Id"`
	LEI                LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 LEI,omitempty"`
	AltrnId            AlternatePartyIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AltrnId,omitempty"`
	SfkpgAcct          SecuritiesAccount19           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SfkpgAcct,omitempty"`
	BlckChainAdrOrWllt BlockChainAddressWallet3      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 BlckChainAdrOrWllt,omitempty"`
	AcctOwnr           PartyIdentification136        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AcctOwnr,omitempty"`
	PtyCpcty           TradingPartyCapacity5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PtyCpcty,omitempty"`
}

type PartyIdentificationAndAccount203 struct {
	Id                 PartyIdentification120Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Id"`
	LEI                LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 LEI,omitempty"`
	AltrnId            AlternatePartyIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AltrnId,omitempty"`
	SfkpgAcct          SecuritiesAccount19           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SfkpgAcct,omitempty"`
	BlckChainAdrOrWllt BlockChainAddressWallet3      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 BlckChainAdrOrWllt,omitempty"`
	PtyCpcty           TradingPartyCapacity5Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PtyCpcty,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Ctry"`
}

type ProprietaryReason4 struct {
	Rsn         GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Rsn,omitempty"`
	AddtlRsnInf Max210Text              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AddtlRsnInf,omitempty"`
}

type ProprietaryStatusAndReason6 struct {
	PrtrySts GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PrtrySts"`
	PrtryRsn []ProprietaryReason4    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PrtryRsn,omitempty"`
}

type Quantity51Choice struct {
	Qty             FinancialInstrumentQuantity33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Qty,omitempty"`
	OrgnlAndCurFace OriginalAndCurrentQuantities1       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 OrgnlAndCurFace,omitempty"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type SecuritiesAccount19 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Nm,omitempty"`
}

type SecuritiesMovement8 struct {
	SctiesMvmntTp                 ReceiveDelivery1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SctiesMvmntTp"`
	FinInstrmId                   SecurityIdentification19        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 FinInstrmId"`
	SctiesQty                     Quantity51Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SctiesQty"`
	MvmntSts                      SecuritiesMovementStatus1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 MvmntSts,omitempty"`
	CollMvmnt                     bool                            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CollMvmnt"`
	SctiesMvmntsApprvd            bool                            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SctiesMvmntsApprvd,omitempty"`
	PosTp                         bool                            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PosTp,omitempty"`
	SfkpgAcct                     SecuritiesAccount19             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SfkpgAcct,omitempty"`
	BlckChainAdrOrWllt            BlockChainAddressWallet3        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 BlckChainAdrOrWllt,omitempty"`
	ClntSctiesMvmntId             Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 ClntSctiesMvmntId,omitempty"`
	TrptyAgtSvcPrvdrSctiesMvmntId Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 TrptyAgtSvcPrvdrSctiesMvmntId,omitempty"`
	MrgndVal                      AmountAndDirection44            `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 MrgndVal,omitempty"`
}

type SecuritiesMovementStatus1Choice struct {
	Amt        ProprietaryReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Amt,omitempty"`
	Csh        ProprietaryReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Csh,omitempty"`
	Ccy        ProprietaryReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Ccy,omitempty"`
	Excld      ProprietaryReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Excld,omitempty"`
	Futr       ProprietaryReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Futr,omitempty"`
	Pdg        ProprietaryReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Pdg,omitempty"`
	MnlyAccptd ProprietaryReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 MnlyAccptd,omitempty"`
	Elgblty    ProprietaryReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Elgblty,omitempty"`
	Tax        ProprietaryReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Tax,omitempty"`
	Wait       ProprietaryReason4 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Wait,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Desc,omitempty"`
}

type SettlementStatus27Choice struct {
	PrtlSttlm []ProprietaryReason4        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PrtlSttlm,omitempty"`
	Sttld     []ProprietaryReason4        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Sttld,omitempty"`
	Usttld    []ProprietaryReason4        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Usttld,omitempty"`
	Prtry     ProprietaryStatusAndReason6 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Prtry,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of AGEN, PRIN
type TradingCapacity7Code string

type TradingPartyCapacity5Choice struct {
	Cd    TradingCapacity7Code    `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Prtry,omitempty"`
}

type TransactionIdentifications46 struct {
	ClntCollInstrId             Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 ClntCollInstrId"`
	ClntCollTxId                Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 ClntCollTxId,omitempty"`
	TrptyAgtSvcPrvdrCollInstrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 TrptyAgtSvcPrvdrCollInstrId,omitempty"`
	TrptyAgtSvcPrvdrCollTxId    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 TrptyAgtSvcPrvdrCollTxId,omitempty"`
	CtrPtyCollTxId              Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CtrPtyCollTxId,omitempty"`
	CmonTxId                    Max52Text `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CmonTxId,omitempty"`
}

type TripartyCollateralStatusAdviceV01 struct {
	TxInstrId   TransactionIdentifications46 `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 TxInstrId"`
	Pgntn       Pagination1                  `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 Pgntn"`
	AllcnSts    AllocationStatus1Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 AllcnSts,omitempty"`
	SttlmSts    SettlementStatus27Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SttlmSts,omitempty"`
	CollSts     CollateralStatus3Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CollSts,omitempty"`
	GnlParams   CollateralParameters13       `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 GnlParams"`
	CollPties   CollateralParties8           `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CollPties"`
	DealTxDtls  DealTransactionDetails7      `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 DealTxDtls"`
	DealTxDt    CollateralDate2              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 DealTxDt"`
	SctiesMvmnt []SecuritiesMovement8        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SctiesMvmnt,omitempty"`
	CshMvmnt    []CashMovement7              `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 CshMvmnt,omitempty"`
	SplmtryData []SupplementaryData1         `xml:"urn:iso:std:iso:20022:tech:xsd:colr.023.001.01 SplmtryData,omitempty"`
}

// May be one of ARNU, CCPT, CHTY, CORP, DRLC, FIIN, TXID
type TypeOfIdentification1Code string

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
