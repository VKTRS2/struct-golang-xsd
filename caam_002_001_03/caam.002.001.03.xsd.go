// Code generated by download. DO NOT EDIT.

package iso20022_caam_002_001_03

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type ATMCommand10 struct {
	Tp          ATMCommand6Code             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Tp"`
	Urgcy       TMSContactLevel2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Urgcy"`
	DtTm        ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 DtTm,omitempty"`
	CmdId       ATMCommandIdentification1   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 CmdId,omitempty"`
	Rsn         ATMCommandReason1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Rsn,omitempty"`
	TracRsn     []ATMCommandReason1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 TracRsn,omitempty"`
	AddtlRsnInf Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 AddtlRsnInf,omitempty"`
	CmdParams   ATMCommandParameters3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 CmdParams,omitempty"`
}

// May be one of ABAL, ASTS, CFGT, CCNT, DISC, KACT, KDAC, KDWL, KRMV, SCFU, SSCU, SSTU, SNDM, HKCG, HKRV, KCHG
type ATMCommand6Code string

type ATMCommandIdentification1 struct {
	Orgn Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Orgn,omitempty"`
	Ref  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Ref,omitempty"`
	Prcr Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Prcr,omitempty"`
}

type ATMCommandParameters1 struct {
	SrlNb      Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SrlNb,omitempty"`
	ReqrdCfgtn ATMSecurityConfiguration1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ReqrdCfgtn,omitempty"`
	ReqrdSts   ATMStatus2Code            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ReqrdSts,omitempty"`
}

type ATMCommandParameters3Choice struct {
	ATMReqrdGblSts  ATMStatus1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ATMReqrdGblSts,omitempty"`
	XpctdMsgFctn    MessageFunction8Code       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 XpctdMsgFctn,omitempty"`
	ReqrdCfgtnParam ATMConfigurationParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ReqrdCfgtnParam,omitempty"`
	ReqrdSctySchme  ATMSecurityScheme4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ReqrdSctySchme,omitempty"`
	SctyDvc         ATMCommandParameters1      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SctyDvc,omitempty"`
	Key             ATMConfigurationParameter2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Key,omitempty"`
}

// May be one of DIAG, MONI, SECU, SYNC, UPDT
type ATMCommandReason1Code string

type ATMConfigurationParameter1 struct {
	Tp   DataSetCategory7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Tp"`
	Vrsn Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Vrsn"`
}

type ATMConfigurationParameter2 struct {
	KeyCtgy   CryptographicKeyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KeyCtgy,omitempty"`
	HstChllng Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 HstChllng,omitempty"`
	Cert      []Max5000Binary           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Cert,omitempty"`
	KeyProps  []KEKIdentifier4          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KeyProps,omitempty"`
}

type ATMDeviceControl2 struct {
	Envt ATMEnvironment7 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Envt"`
	Cmd  []ATMCommand10  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Cmd,omitempty"`
}

type ATMDeviceControlV03 struct {
	Hdr              Header31                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Hdr"`
	PrtctdATMDvcCtrl ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 PrtctdATMDvcCtrl,omitempty"`
	ATMDvcCtrl       ATMDeviceControl2        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ATMDvcCtrl,omitempty"`
	SctyTrlr         ContentInformationType13 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SctyTrlr,omitempty"`
}

type ATMEnvironment7 struct {
	Acqrr    Acquirer7               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Acqrr,omitempty"`
	ATMMgr   Acquirer8               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ATMMgr,omitempty"`
	HstgNtty TerminalHosting1        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 HstgNtty,omitempty"`
	ATM      AutomatedTellerMachine3 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ATM"`
}

type ATMMessageFunction2 struct {
	Fctn     MessageFunction11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Fctn"`
	ATMSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ATMSvcCd,omitempty"`
	HstSvcCd Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 HstSvcCd,omitempty"`
}

type ATMSecurityConfiguration1 struct {
	Keys      ATMSecurityConfiguration2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Keys,omitempty"`
	Ncrptn    ATMSecurityConfiguration3 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Ncrptn,omitempty"`
	MACAlgo   []Algorithm12Code         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MACAlgo,omitempty"`
	DgstAlgo  []Algorithm11Code         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 DgstAlgo,omitempty"`
	DgtlSgntr ATMSecurityConfiguration4 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 DgtlSgntr,omitempty"`
	PIN       ATMSecurityConfiguration5 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 PIN,omitempty"`
	MsgPrtcn  []MessageProtection1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MsgPrtcn,omitempty"`
}

type ATMSecurityConfiguration2 struct {
	MaxSmmtrcKey    float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MaxSmmtrcKey,omitempty"`
	MaxAsmmtrcKey   float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MaxAsmmtrcKey,omitempty"`
	MaxRSAKeyLngth  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MaxRSAKeyLngth,omitempty"`
	MaxRootKeyLngth float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MaxRootKeyLngth,omitempty"`
}

type ATMSecurityConfiguration3 struct {
	AsmmtrcNcrptn        bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 AsmmtrcNcrptn,omitempty"`
	AsmmtrcKeyStdId      bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 AsmmtrcKeyStdId,omitempty"`
	AsmmtrcNcrptnAlgo    []Algorithm7Code        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 AsmmtrcNcrptnAlgo,omitempty"`
	SmmtrcTrnsprtKey     bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SmmtrcTrnsprtKey,omitempty"`
	SmmtrcTrnsprtKeyAlgo []Algorithm13Code       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SmmtrcTrnsprtKeyAlgo,omitempty"`
	SmmtrcNcrptnAlgo     []Algorithm15Code       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SmmtrcNcrptnAlgo,omitempty"`
	NcrptnFrmt           []EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 NcrptnFrmt,omitempty"`
}

type ATMSecurityConfiguration4 struct {
	MaxCerts      float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MaxCerts,omitempty"`
	MaxSgntrs     float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MaxSgntrs,omitempty"`
	DgtlSgntrAlgo []Algorithm14Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 DgtlSgntrAlgo,omitempty"`
}

type ATMSecurityConfiguration5 struct {
	PINFrmt          []PINFormat4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 PINFrmt,omitempty"`
	PINLngthCpblties float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 PINLngthCpblties,omitempty"`
}

// May be one of APPK, CERT, FRAN, DTCH, LUXG, MANU, PKIP, SIGN, TR34
type ATMSecurityScheme4Code string

// May be one of INSV, OUTS
type ATMStatus1Code string

// May be one of OPER, OUTS
type ATMStatus2Code string

type Acquirer7 struct {
	AcqrgInstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 AcqrgInstn,omitempty"`
	Brnch      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Brnch,omitempty"`
}

type Acquirer8 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Id"`
	ApplVrsn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ApplVrsn,omitempty"`
}

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of ERS2, ERS1, RPSS
type Algorithm14Code string

// May be one of EA2C, E3DC, EA9C, EA5C
type Algorithm15Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Param,omitempty"`
}

type AlgorithmIdentification16 struct {
	Algo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Algo"`
}

type AlgorithmIdentification17 struct {
	Algo  Algorithm14Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Algo"`
	Param Parameter8      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MAC"`
}

type AutomatedTellerMachine3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Id"`
	AddtlId Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 AddtlId,omitempty"`
	SeqNb   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SeqNb,omitempty"`
	Lctn    PostalAddress17 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Lctn,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 RltvDstngshdNm"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 EnvlpdData"`
}

type ContentInformationType13 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 AuthntcdData,omitempty"`
	SgndData     SignedData4        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SgndData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of APPL, DATA, DYNC, KENC, MACK, PINK, WRKG
type CryptographicKeyType4Code string

// May be one of ATMC, ATMP, APPR, CRAP, CPRC, OEXR, AMNT, LOCC, MNOC
type DataSetCategory7Code string

type Document struct {
	ATMDvcCtrl ATMDeviceControlV03 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ATMDvcCtrl"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 NcrptdCntt,omitempty"`
}

type GenericIdentification77 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Id"`
	Tp     PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Tp"`
	Issr   PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 ShrtNm,omitempty"`
}

type GeographicCoordinates1 struct {
	Lat  Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Lat"`
	Long Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Long"`
}

type GeographicLocation1Choice struct {
	GeogcCordints GeographicCoordinates1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 GeogcCordints,omitempty"`
	UTMCordints   UTMCoordinates1        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 UTMCordints,omitempty"`
}

type Header31 struct {
	MsgFctn    ATMMessageFunction2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MsgFctn"`
	PrtcolVrsn Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 PrtcolVrsn"`
	XchgId     Max3NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 XchgId"`
	CreDtTm    ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 CreDtTm"`
	InitgPty   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 InitgPty"`
	RcptPty    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 RcptPty,omitempty"`
	PrcStat    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 PrcStat,omitempty"`
	Tracblt    []Traceability4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 DerivtnId,omitempty"`
}

type KEKIdentifier4 struct {
	Nm      Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Nm,omitempty"`
	KeyId   Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KeyId,omitempty"`
	KeyVrsn Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KeyVrsn,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 NcrptdKey"`
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

type Max3000Binary []byte

func (t *Max3000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max3000Binary) MarshalText() ([]byte, error) {
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

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

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

// May be no more than 6 items long
type Max6Text string

// May be no more than 70 items long
type Max70Text string

// May be one of BALN, CMPA, CMPD, ACMD, DVCC, DIAQ, DIAP, GSTS, INQQ, INQP, KYAQ, KYAP, PINQ, PINP, RJAQ, RJAP, WITV, WITK, WITQ, WITP, INQC, H2AP, H2AQ, TMOP, CSEC, DSEC, SKSC, SSTS, DPSK, DPSV, DPSQ, DPSP, EXPK, EXPV, TRFQ, TRFP, RPTC
type MessageFunction11Code string

// May be one of BALN, GSTS, DSEC, INQC, KEYQ, SSTS
type MessageFunction8Code string

// May be one of EVLP, MACB, MACM, UNPR
type MessageProtection1Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// May be one of ANSI, BNCM, BKSY, DBLD, DBLC, ECI2, ECI3, EMVS, IBM3, ISO0, ISO1, ISO2, ISO3, ISO4, ISO5, VIS2, VIS3
type PINFormat4Code string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 BPddg,omitempty"`
}

type Parameter8 struct {
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 DgstAlgo"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 MskGnrtrAlgo"`
	SaltLngth    float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SaltLngth"`
	TrlrFld      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 TrlrFld,omitempty"`
}

// May be one of ACQR, ATMG, CISP, DLIS, HSTG, ITAG, OATM
type PartyType12Code string

type PostalAddress17 struct {
	AdrLine     []Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 BldgNb,omitempty"`
	PstCd       Max16Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 PstCd,omitempty"`
	TwnNm       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 TwnNm"`
	CtrySubDvsn []Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Ctry"`
	GLctn       GeographicLocation1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 GLctn,omitempty"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KeyTrnsprt,omitempty"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KEK,omitempty"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KeyIdr,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 IssrAndSrlNb,omitempty"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 KeyIdr,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 AttrVal"`
}

type SignedData4 struct {
	Vrsn        float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 NcpsltdCntt"`
	Cert        []Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Cert,omitempty"`
	Sgnr        []Signer3                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Sgnr"`
}

type Signer3 struct {
	Vrsn      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Vrsn,omitempty"`
	SgnrId    Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SgnrId,omitempty"`
	DgstAlgo  AlgorithmIdentification16 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 DgstAlgo"`
	SgntrAlgo AlgorithmIdentification17 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SgntrAlgo"`
	Sgntr     Max3000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Sgntr"`
}

// May be one of ASAP, CRIT, DTIM, ENCS
type TMSContactLevel2Code string

type TerminalHosting1 struct {
	Ctgy TransactionEnvironment3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Ctgy,omitempty"`
	Id   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 Id,omitempty"`
}

type Traceability4 struct {
	RlayId      GenericIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 RlayId"`
	SeqNb       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 SeqNb,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 TracDtTmOut"`
}

// May be one of BRCH, MERC, OTHR
type TransactionEnvironment3Code string

type UTMCoordinates1 struct {
	UTMZone    Max16Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 UTMZone"`
	UTMEstwrd  float64   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 UTMEstwrd"`
	UTMNrthwrd float64   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.002.001.03 UTMNrthwrd"`
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
