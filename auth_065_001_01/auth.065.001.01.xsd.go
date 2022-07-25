// Code generated by download. DO NOT EDIT.

package iso20022_auth_065_001_01

type BackTestingMethodology1 struct {
	RskMdlTp          ModelType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 RskMdlTp"`
	MdlCnfdncLvl      float64          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 MdlCnfdncLvl"`
	VartnMrgnCleanInd bool             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 VartnMrgnCleanInd"`
	Desc              Max2000Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 Desc,omitempty"`
}

type CCPBackTestingDefinitionReportV01 struct {
	Mthdlgy     []BackTestingMethodology1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 Mthdlgy"`
	SplmtryData []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 SplmtryData,omitempty"`
}

type Document struct {
	CCPBckTstgDefRpt CCPBackTestingDefinitionReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 CCPBckTstgDefRpt"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 SchmeNm,omitempty"`
}

// May be no more than 2000 items long
type Max2000Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

type ModelType1Choice struct {
	Cd    ModelType1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 Cd,omitempty"`
	Prtry GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 Prtry,omitempty"`
}

// May be one of EXPS, OTHR, ORIA, SPAN, VARI, SAMO
type ModelType1Code string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.065.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
