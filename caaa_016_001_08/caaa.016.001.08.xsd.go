// Code generated by download. DO NOT EDIT.

package iso20022_caaa_016_001_08

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AdditionalData1 struct {
	Tp  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Tp,omitempty"`
	Val Max2048Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Val,omitempty"`
}

type AdditionalFee1 struct {
	Tp         TypeOfAmount10Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Tp"`
	OthrTp     Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrTp,omitempty"`
	FeePrgm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 FeePrgm,omitempty"`
	FeeDscrptr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 FeeDscrptr,omitempty"`
	Amt        FeeAmount2         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Amt"`
	Labl       Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Labl,omitempty"`
}

type AdditionalInformation22 struct {
	Rcpt PartyType19Code      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Rcpt,omitempty"`
	Trgt []UserInterface8Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Trgt,omitempty"`
	Frmt OutputFormat4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Frmt,omitempty"`
	Tp   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Tp,omitempty"`
	Val  Max20KText           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Val"`
}

type Address1 struct {
	AdrLine1       Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AdrLine1,omitempty"`
	AdrLine2       Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AdrLine2,omitempty"`
	StrtNm         Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 StrtNm,omitempty"`
	BldgNb         Max16Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 BldgNb,omitempty"`
	PstlCd         Max16Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PstlCd,omitempty"`
	TwnNm          Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 TwnNm,omitempty"`
	CtrySubDvsnMnr Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CtrySubDvsnMnr,omitempty"`
	CtrySubDvsnMjr Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CtrySubDvsnMjr,omitempty"`
	Ctry           Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Ctry,omitempty"`
}

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of HS25, HS38, HS51
type Algorithm20Code string

// May be one of EA2C, E3DC, EA9C, EA5C, EA2R, EA9R, EA5R, E3DR, E36C, E36R, SD5C
type Algorithm23Code string

// May be one of HS25, HS38, HS51
type Algorithm5Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification25 struct {
	Algo  Algorithm23Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Param,omitempty"`
}

type AlgorithmIdentification26 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Algo"`
	Param Algorithm5Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Param,omitempty"`
}

type AlgorithmIdentification27 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Algo"`
	Param Parameter13    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Param,omitempty"`
}

type AlgorithmIdentification28 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Algo"`
	Param Parameter14     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

// May be one of APKI, ADVF, ARNB, ARPC, ARQC, ATCC, BTHD, CHSA, CHDN, CUID, DRVI, DRLN, EMAL, EMIN, EMRN, IDCN, MANU, NVSC, FBIG, FBIO, OLDA, OLDS, OFPE, FCPN, OTPW, NBIG, NPIN, OCHI, OTHN, OTHP, PPSG, PSVE, PASN, PSWD, TOKP, PKIS, PLOB, PCDV, SCRT, SCNL, CSEC, SHAF, SHAT, CPSG, SSNB, TXIN, TOKA, CDHI, TOKN, QWAC, PHOM, PWOR, THDS, ADDB, ADDS, CSCV, CRYP, BIOM, MOBL, FPIN
type AuthenticationMethod11Code string

type AuthorisationStatus1 struct {
	AuthstnInd      bool            `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AuthstnInd,omitempty"`
	AuthstnNtty     PartyType26Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AuthstnNtty,omitempty"`
	OthrAuthstnNtty Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrAuthstnNtty,omitempty"`
}

type BatchManagementInformation1 struct {
	ColltnId         Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 ColltnId,omitempty"`
	BtchId           Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 BtchId"`
	MsgSeqNb         Max15NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MsgSeqNb,omitempty"`
	MsgChcksmInptVal Max140Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MsgChcksmInptVal,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CardData3 struct {
	PAN           Max19NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PAN,omitempty"`
	PrtctdPANInd  bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PrtctdPANInd,omitempty"`
	CardSeqNb     Min2Max3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CardSeqNb,omitempty"`
	FctvDt        Max10Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 FctvDt,omitempty"`
	XpryDt        Exact4NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 XpryDt,omitempty"`
	PmtAcctRef    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PmtAcctRef,omitempty"`
	PANRefIdr     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PANRefIdr,omitempty"`
	PANAcctRg     Max19NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PANAcctRg,omitempty"`
	CardCtryCd    ISO3NumericCountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CardCtryCd,omitempty"`
	CardPdctTp    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CardPdctTp,omitempty"`
	CardPdctSubTp Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CardPdctSubTp,omitempty"`
	CardPrtflIdr  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CardPrtflIdr,omitempty"`
	AddtlCardData Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AddtlCardData,omitempty"`
}

type CardNotReceivedDetails1 struct {
	DtOfCardMld    ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 DtOfCardMld"`
	MlngAdr        Address1                  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MlngAdr,omitempty"`
	MlngAdrUstrd   Max256Text                `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MlngAdrUstrd,omitempty"`
	MldFrPstlCd    Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MldFrPstlCd"`
	VldFr          ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 VldFr,omitempty"`
	CardSctyCdInd  bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CardSctyCdInd,omitempty"`
	CardSctyCpblty []CardSecurityCapability1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CardSctyCpblty,omitempty"`
}

type CardProgrammeMode1 struct {
	Tp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Tp,omitempty"`
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Id"`
}

type CardSecurityCapability1 struct {
	Cpblty     CardSecurityCapability1Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Cpblty"`
	OthrCpblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrCpblty,omitempty"`
}

// May be one of ICCD, MWOS, MSWS, OTHN, OTHP, OLPN
type CardSecurityCapability1Code string

type Cardholder15 struct {
	CrdhldrNm CardholderName1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CrdhldrNm,omitempty"`
	Id        []Credentials1  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Id,omitempty"`
	Adr       Address1        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Adr,omitempty"`
	CtctInf   Contact1        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CtctInf,omitempty"`
	DtOfBirth ISODate         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 DtOfBirth,omitempty"`
}

type CardholderName1 struct {
	Nm         Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Nm,omitempty"`
	GvnNm      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 GvnNm,omitempty"`
	MddlInitls Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MddlInitls,omitempty"`
	LastNm     Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 LastNm,omitempty"`
}

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 RltvDstngshdNm"`
}

type Contact1 struct {
	Nm            Max70Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Nm,omitempty"`
	HomePhneNb    PhoneNumber       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 HomePhneNb,omitempty"`
	BizPhneNb     PhoneNumber       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 BizPhneNb,omitempty"`
	MobPhneNb     PhoneNumber       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MobPhneNb,omitempty"`
	OthrPhneNb    PhoneNumber       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrPhneNb,omitempty"`
	PrsnlEmailAdr Max256Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PrsnlEmailAdr,omitempty"`
	BizEmailAdr   Max256Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 BizEmailAdr,omitempty"`
	OthrEmailAdr  Max256Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrEmailAdr,omitempty"`
	Lang          ISO2ALanguageCode `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Lang,omitempty"`
}

type ContentInformationType20 struct {
	MACData MACData1          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MACData"`
	MAC     Max8HexBinaryText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MAC"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// May be one of EVLP, IFSE
type ContentType3Code string

type Context8 struct {
	TxCntxt TransactionContext5 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 TxCntxt,omitempty"`
}

type Credentials1 struct {
	IdCd     Identification2Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 IdCd"`
	OthrIdCd Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrIdCd,omitempty"`
	IdVal    Max70Text           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 IdVal"`
}

type DisputeData2 struct {
	PresntmntCycl Exact1NumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PresntmntCycl,omitempty"`
	DsptCond      Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 DsptCond,omitempty"`
	DsptRef       []DisputeReference1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 DsptRef,omitempty"`
}

type DisputeIdentification1 struct {
	Tp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Tp,omitempty"`
	Id Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Id"`
}

type DisputeReference1 struct {
	AssgnrNtty     PartyType32Code          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AssgnrNtty,omitempty"`
	OthrAssgnrNtty Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrAssgnrNtty,omitempty"`
	DsptId         []DisputeIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 DsptId"`
}

type Document struct {
	FrdRptgInitn FraudReportingInitiationV01 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 FrdRptgInitn"`
}

type EncryptedContent5 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification25 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CnttNcrptnAlgo"`
	NcrptdDataElmt []EncryptedDataElement1   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 NcrptdDataElmt"`
}

type EncryptedData1 struct {
	Ctrl           Exact1HexBinaryText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Ctrl,omitempty"`
	KeySetIdr      Max8NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeySetIdr,omitempty"`
	DrvdInf        Max32HexBinaryText      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 DrvdInf,omitempty"`
	Algo           Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Algo,omitempty"`
	KeyLngth       Max4NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyLngth,omitempty"`
	KeyPrtcn       Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyPrtcn,omitempty"`
	KeyIndx        Max5NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyIndx,omitempty"`
	PddgMtd        Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PddgMtd,omitempty"`
	NcrptdDataFrmt Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 NcrptdDataFrmt,omitempty"`
	NcrptdDataElmt []EncryptedDataElement1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 NcrptdDataElmt"`
}

type EncryptedData1Choice struct {
	BinryData   Max100KBinary `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 BinryData,omitempty"`
	HexBinryVal string        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 HexBinryVal,omitempty"`
}

type EncryptedDataElement1 struct {
	Id                   ExternalEncryptedElementIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Id,omitempty"`
	OthrId               Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrId,omitempty"`
	NcrptdData           EncryptedData1Choice                        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 NcrptdData"`
	ClearTxtDataFrmt     EncryptedDataFormat1Code                    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 ClearTxtDataFrmt,omitempty"`
	OthrClearTxtDataFrmt Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrClearTxtDataFrmt,omitempty"`
}

// May be one of ASCI, BINF, EBCD, HEXF, OTHN, OTHP
type EncryptedDataFormat1Code string

// May be one of TR34, TR31, CTCE, CBCE
type EncryptionFormat3Code string

type EnvelopedData6 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Vrsn,omitempty"`
	Rcpt       []Recipient7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Rcpt"`
	NcrptdCntt EncryptedContent5  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 NcrptdCntt,omitempty"`
}

type Environment10 struct {
	Acqrr   PartyIdentification197 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Acqrr,omitempty"`
	Sndr    PartyIdentification197 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Sndr,omitempty"`
	Rcvr    PartyIdentification197 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Rcvr,omitempty"`
	Card    CardData3              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Card,omitempty"`
	Crdhldr Cardholder15           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Crdhldr,omitempty"`
	Tkn     Token1                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Tkn,omitempty"`
}

// Must match the pattern ([0-9A-F][0-9A-F]){1}
type Exact1HexBinaryText string

// Must match the pattern [0-9]
type Exact1NumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// Must match the pattern ([0-9A-F][0-9A-F]){1,3}
type ExternalEncryptedElementIdentification1Code string

type FeeAmount2 struct {
	Amt      float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Amt"`
	Ccy      ISO3NumericCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Ccy,omitempty"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 XchgRate,omitempty"`
	QtnDt    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 QtnDt,omitempty"`
	Sgn      bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Sgn,omitempty"`
}

type FraudCaseDetails1 struct {
	MktSgmt  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MktSgmt,omitempty"`
	LctrNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 LctrNb,omitempty"`
	CaseRef  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CaseRef,omitempty"`
	ArrstInd bool      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 ArrstInd,omitempty"`
}

// May be one of DUPL, CLSE, NEWF, OTHN, OTHP, REOP, UPDT
type FraudReportingAction1Code string

type FraudReportingInitiation1 struct {
	Envt        Environment10        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Envt"`
	Cntxt       Context8             `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Cntxt,omitempty"`
	Tx          Transaction79        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Tx"`
	PrtctdData  []ProtectedData1     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PrtctdData,omitempty"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 SplmtryData,omitempty"`
}

type FraudReportingInitiationV01 struct {
	Hdr      Header48                  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Hdr"`
	Body     FraudReportingInitiation1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Body"`
	SctyTrlr ContentInformationType20  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 SctyTrlr,omitempty"`
}

// May be one of ACTO, CWUI, CRNT, FRAC, FRAP, CWKA, CRDL, MISC, OTHN, OTHP, CRDS, CNPA, MUFD, COSN
type FraudType1Code string

type FraudulentTransactionData1 struct {
	AuthstnSts  AuthorisationStatus1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AuthstnSts,omitempty"`
	DsptDtls    DisputeData2         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 DsptDtls,omitempty"`
	MsgRsn      []Exact4NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MsgRsn,omitempty"`
	AltrnMsgRsn []Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AltrnMsgRsn,omitempty"`
	FrdlntMsg   Max100KBinary        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 FrdlntMsg,omitempty"`
}

type GenericIdentification172 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Id"`
	Tp     PartyType17Code   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Tp,omitempty"`
	OthrTp Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrTp,omitempty"`
	Assgnr PartyType18Code   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Assgnr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 ShrtNm,omitempty"`
}

type Header48 struct {
	MsgFctn        MessageFunction29Code       `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MsgFctn"`
	PrtcolVrsn     Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PrtcolVrsn"`
	XchgId         Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 XchgId,omitempty"`
	ReTrnsmssnCntr Max3NumericText             `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 ReTrnsmssnCntr,omitempty"`
	CreDtTm        ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CreDtTm"`
	BtchMgmtInf    BatchManagementInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 BtchMgmtInf,omitempty"`
	InitgPty       GenericIdentification172    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 InitgPty"`
	RcptPty        GenericIdentification172    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 RcptPty,omitempty"`
	TracData       []AdditionalData1           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 TracData,omitempty"`
	Tracblt        []Traceability7             `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Tracblt,omitempty"`
}

// Must match the pattern [a-z]{2,2}
type ISO2ALanguageCode string

// Must match the pattern [0-9]{3,3}
type ISO3NumericCountryCode string

// Must match the pattern [0-9]{3,3}
type ISO3NumericCurrencyCode string

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

// May be one of DRID, NTID, PASS, SSYN, ARNB, OTHP, OTHN, EMAL, PHNB
type Identification2Code string

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 SrlNb"`
}

type KEK6 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Vrsn,omitempty"`
	KEKId         KEKIdentifier6            `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification28 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 NcrptdKey,omitempty"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 DerivtnId,omitempty"`
}

type KEKIdentifier6 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyVrsn,omitempty"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 DerivtnId,omitempty"`
}

type KeyTransport6 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification27 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 NcrptdKey"`
}

type MACData1 struct {
	Ctrl         Exact1HexBinaryText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Ctrl"`
	KeySetIdr    Max8NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeySetIdr"`
	DrvdInf      Max32HexBinaryText  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 DrvdInf,omitempty"`
	Algo         Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Algo"`
	KeyLngth     Max4NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyLngth,omitempty"`
	KeyPrtcn     Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyPrtcn,omitempty"`
	KeyIndx      Max5NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyIndx,omitempty"`
	PddgMtd      Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PddgMtd,omitempty"`
	InitlstnVctr Max32HexBinaryText  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 InitlstnVctr,omitempty"`
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 10 items long
type Max10Text string

// Must match the pattern [0-9]{1,11}
type Max11NumericText string

type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 140 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// May be no more than 16 items long
type Max16Text string

// Must match the pattern [0-9]{1,19}
type Max19NumericText string

// May be no more than 2048 items long
type Max2048Text string

// May be no more than 20000 items long
type Max20KText string

// May be no more than 256 items long
type Max256Text string

// Must match the pattern [0-9]{1,2}
type Max2NumericText string

// Must match the pattern ([0-9A-F][0-9A-F]){1,32}
type Max32HexBinaryText string

// May be no more than 350 items long
type Max350Text string

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 35 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must match the pattern [0-9]{1,4}
type Max4NumericText string

type Max5000Binary []byte

func (t *Max5000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max5000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max500Binary []byte

func (t *Max500Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max500Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// May be no more than 70 items long
type Max70Text string

// Must match the pattern ([0-9A-F][0-9A-F]){1,8}
type Max8HexBinaryText string

// Must match the pattern [0-9]{1,8}
type Max8NumericText string

// May be one of ADVC, NOTI
type MessageFunction29Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be one of FLNM, MREF, OTHN, OTHP, SMSI, TEXT, URLI, HTML
type OutputFormat4Code string

type Parameter13 struct {
	DgstAlgo     Algorithm20Code           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification26 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 MskGnrtrAlgo,omitempty"`
}

type Parameter14 struct {
	NcrptnFrmt   EncryptionFormat3Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 BPddg,omitempty"`
}

type PartyIdentification197 struct {
	Id      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Id"`
	Assgnr  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Assgnr,omitempty"`
	Ctry    ISO3NumericCountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Ctry,omitempty"`
	ShrtNm  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 ShrtNm,omitempty"`
	AddtlId Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AddtlId,omitempty"`
}

// May be one of OTHN, OTHP, ACQR, ACQP, CISS, CISP, AGNT
type PartyType17Code string

// May be one of ACQR, CISS, CSCH, AGNT
type PartyType18Code string

// May be one of ACCP, ACQR, ACQP, CISS, CISP, AGNT, OTHN, OTHP
type PartyType19Code string

// May be one of ACCP, ACQR, ICCA, CISS, DLIS, AGNT, OTHN, OTHP
type PartyType26Code string

// May be one of ACQR, AGNT, ISUR, OTHN, OTHP
type PartyType32Code string

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type ProtectedData1 struct {
	CnttTp     ContentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CnttTp"`
	EnvlpdData EnvelopedData6   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 EnvlpdData,omitempty"`
	NcrptdData EncryptedData1   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 NcrptdData,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 IssrAndSrlNb,omitempty"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyIdr,omitempty"`
}

type Recipient7Choice struct {
	KeyTrnsprt KeyTransport6  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyTrnsprt,omitempty"`
	KEK        KEK6           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KEK,omitempty"`
	KeyIdr     KEKIdentifier6 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 KeyIdr,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AttrVal"`
}

type ReportedFraud1 struct {
	FrdTp           FraudType1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 FrdTp"`
	OthrFrdTp       Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrFrdTp,omitempty"`
	FrdRptgActn     FraudReportingAction1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 FrdRptgActn"`
	OthrFrdRptgActn Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrFrdRptgActn,omitempty"`
	RptgNtty        PartyType26Code              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 RptgNtty"`
	OthrRptgNtty    Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 OthrRptgNtty,omitempty"`
	CmprmsdCrdntl   []AuthenticationMethod11Code `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CmprmsdCrdntl,omitempty"`
	CrdhldrRptgDt   ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CrdhldrRptgDt,omitempty"`
	ConfRptgDt      ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 ConfRptgDt,omitempty"`
	SubmitrCaseRef  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 SubmitrCaseRef,omitempty"`
	FrdCaseDtls     FraudCaseDetails1            `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 FrdCaseDtls,omitempty"`
	FrdInvstgtnSts  Max256Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 FrdInvstgtnSts,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type Token1 struct {
	PmtTkn        Max19NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 PmtTkn,omitempty"`
	TknXpryDt     Exact4NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 TknXpryDt,omitempty"`
	TknRqstrId    Max11NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 TknRqstrId,omitempty"`
	TknAssrncData Max140Text        `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 TknAssrncData,omitempty"`
	TknAssrncMtd  Max2NumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 TknAssrncMtd,omitempty"`
	TknInittdInd  bool              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 TknInittdInd,omitempty"`
}

type Traceability7 struct {
	RlayId      GenericIdentification172 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 RlayId"`
	TracDtTmIn  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 TracDtTmIn,omitempty"`
	TracDtTmOut ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 TracDtTmOut,omitempty"`
}

type Transaction79 struct {
	FrdTxId         Max70Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 FrdTxId"`
	RptdFrd         ReportedFraud1             `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 RptdFrd"`
	FrdlntTxData    FraudulentTransactionData1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 FrdlntTxData"`
	CardNotRcvdDtls CardNotReceivedDetails1    `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CardNotRcvdDtls,omitempty"`
	CrdhldrCardNm   CardholderName1            `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CrdhldrCardNm,omitempty"`
	AddtlFees       []AdditionalFee1           `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AddtlFees,omitempty"`
	AddtlInf        []AdditionalInformation22  `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AddtlInf,omitempty"`
	AddtlData       []AdditionalData1          `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 AddtlData,omitempty"`
}

type TransactionContext5 struct {
	CardPrgrmmApld CardProgrammeMode1 `xml:"urn:iso:std:iso:20022:tech:xsd:cafr.001.001.01 CardPrgrmmApld,omitempty"`
}

// May be one of INTC, FEEP, OTHN, OTHP, FEEA
type TypeOfAmount10Code string

// May be one of DSPU, FILE, LOGF, OTHP, OTHN
type UserInterface8Code string

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
