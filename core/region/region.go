package region

import (
	"sort"
)

const (
	AC  = "AC"
	AD  = "AD"
	AE  = "AE"
	AF  = "AF"
	AG  = "AG"
	AI  = "AI"
	AL  = "AL"
	AM  = "AM"
	AN  = "AN"
	AO  = "AO"
	AQ  = "AQ"
	AR  = "AR"
	AS  = "AS"
	AT  = "AT"
	AU  = "AU"
	AW  = "AW"
	AX  = "AX"
	AZ  = "AZ"
	BA  = "BA"
	BB  = "BB"
	BD  = "BD"
	BE  = "BE"
	BF  = "BF"
	BG  = "BG"
	BH  = "BH"
	BI  = "BI"
	BJ  = "BJ"
	BL  = "BL"
	BM  = "BM"
	BN  = "BN"
	BO  = "BO"
	BR  = "BR"
	BS  = "BS"
	BT  = "BT"
	BW  = "BW"
	BY  = "BY"
	BZ  = "BZ"
	CA  = "CA"
	CC  = "CC"
	CD  = "CD"
	CF  = "CF"
	CG  = "CG"
	CH  = "CH"
	CI  = "CI"
	CK  = "CK"
	CL  = "CL"
	CM  = "CM"
	CN  = "CN"
	CO  = "CO"
	CR  = "CR"
	CU  = "CU"
	CV  = "CV"
	CX  = "CX"
	CY  = "CY"
	CZ  = "CZ"
	DE  = "DE"
	DJ  = "DJ"
	DK  = "DK"
	DM  = "DM"
	DO  = "DO"
	DZ  = "DZ"
	EC  = "EC"
	EE  = "EE"
	EG  = "EG"
	ER  = "ER"
	ES  = "ES"
	ET  = "ET"
	FI  = "FI"
	FJ  = "FJ"
	FK  = "FK"
	FM  = "FM"
	FO  = "FO"
	FR  = "FR"
	GA  = "GA"
	GB  = "GB"
	GD  = "GD"
	GE  = "GE"
	GF  = "GF"
	GG  = "GG"
	GH  = "GH"
	GI  = "GI"
	GL  = "GL"
	GM  = "GM"
	GN  = "GN"
	GP  = "GP"
	GQ  = "GQ"
	GR  = "GR"
	GS  = "GS"
	GT  = "GT"
	GU  = "GU"
	GW  = "GW"
	GY  = "GY"
	HK  = "HK"
	HN  = "HN"
	HR  = "HR"
	HT  = "HT"
	HU  = "HU"
	ID  = "ID"
	IE  = "IE"
	IL  = "IL"
	IM  = "IM"
	IN  = "IN"
	IO  = "IO"
	IQ  = "IQ"
	IR  = "IR"
	IS  = "IS"
	IT  = "IT"
	JE  = "JE"
	JM  = "JM"
	JO  = "JO"
	JP  = "JP"
	KE  = "KE"
	KG  = "KG"
	KH  = "KH"
	KI  = "KI"
	KM  = "KM"
	KN  = "KN"
	KP  = "KP"
	KR  = "KR"
	KW  = "KW"
	KY  = "KY"
	KZ  = "KZ"
	LA  = "LA"
	LB  = "LB"
	LC  = "LC"
	LI  = "LI"
	LK  = "LK"
	LR  = "LR"
	LS  = "LS"
	LT  = "LT"
	LU  = "LU"
	LV  = "LV"
	LY  = "LY"
	MA  = "MA"
	MC  = "MC"
	MD  = "MD"
	ME  = "ME"
	MF  = "MF"
	MG  = "MG"
	MH  = "MH"
	MK  = "MK"
	ML  = "ML"
	MM  = "MM"
	MN  = "MN"
	MO  = "MO"
	MP  = "MP"
	MQ  = "MQ"
	MR  = "MR"
	MS  = "MS"
	MT  = "MT"
	MU  = "MU"
	MV  = "MV"
	MW  = "MW"
	MX  = "MX"
	MY  = "MY"
	MZ  = "MZ"
	NA  = "NA"
	NC  = "NC"
	NE  = "NE"
	NF  = "NF"
	NG  = "NG"
	NI  = "NI"
	NL  = "NL"
	NO  = "NO"
	NP  = "NP"
	NR  = "NR"
	NU  = "NU"
	NZ  = "NZ"
	OM  = "OM"
	PA  = "PA"
	PE  = "PE"
	PF  = "PF"
	PG  = "PG"
	PH  = "PH"
	PK  = "PK"
	PL  = "PL"
	PM  = "PM"
	PN  = "PN"
	PR  = "PR"
	PS  = "PS"
	PT  = "PT"
	PW  = "PW"
	PY  = "PY"
	QA  = "QA"
	RE  = "RE"
	RO  = "RO"
	RS  = "RS"
	RU  = "RU"
	RW  = "RW"
	SA  = "SA"
	SB  = "SB"
	SC  = "SC"
	SD  = "SD"
	SE  = "SE"
	SG  = "SG"
	SH  = "SH"
	SI  = "SI"
	SJ  = "SJ"
	SK  = "SK"
	SL  = "SL"
	SM  = "SM"
	SN  = "SN"
	SO  = "SO"
	SR  = "SR"
	SS  = "SS"
	ST  = "ST"
	SV  = "SV"
	SX  = "SX"
	SY  = "SY"
	SZ  = "SZ"
	TC  = "TC"
	TD  = "TD"
	TG  = "TG"
	TH  = "TH"
	TJ  = "TJ"
	TK  = "TK"
	TL  = "TL"
	TM  = "TM"
	TN  = "TN"
	TO  = "TO"
	TR  = "TR"
	TT  = "TT"
	TV  = "TV"
	TW  = "TW"
	TZ  = "TZ"
	UA  = "UA"
	UG  = "UG"
	UMI = "UMI"
	US  = "US"
	UY  = "UY"
	UZ  = "UZ"
	VA  = "VA"
	VC  = "VC"
	VE  = "VE"
	VG  = "VG"
	VI  = "VI"
	VN  = "VN"
	VU  = "VU"
	WF  = "WF"
	WS  = "WS"
	XK  = "XK"
	YE  = "YE"
	YT  = "YT"
	ZA  = "ZA"
	ZM  = "ZM"
	ZW  = "ZW"
)

type Region struct {
	Name           string
	Nationality    string
	ISO            string
	Currency       string
	CurrencySymbol string
	CC             string
	lengthAfterCC  int

	IDTypes           []string
	IDValidationFuncs map[string]IDValidationFunc
	sanitiseMobile    []SanitiseMobileFunc
}

// RegionList implements sort.Interface and sorts regions by Name
type RegionList []Region

func (a RegionList) Len() int { return len(a) }

func (a RegionList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a RegionList) Less(i, j int) bool { return a[i].Name < a[j].Name }

var (
	// SortedRegionList is a list of regions sorted by Name, this is because in most UI scenarios we
	// want to sort by name rather than ISO
	SortedRegionList RegionList = make([]Region, len(IsoRegionMap))
)

func init() {
	var i int
	for _, r := range IsoRegionMap {
		SortedRegionList[i] = r
		i++
	}

	sort.Sort(SortedRegionList)
}

var IsoRegionMap = map[string]Region{
	ZA: {
		Name:           "South Africa",
		Nationality:    "South African",
		ISO:            ZA,
		Currency:       "Rand",
		CurrencySymbol: "R",
		CC:             "27",
		lengthAfterCC:  9,
		IDTypes:        []string{ID_Type_NationalID, ID_Type_Passport},
		IDValidationFuncs: map[string]IDValidationFunc{
			ID_Type_NationalID: validateZaID,
			ID_Type_Passport:   validateZaPassport,
		},
		sanitiseMobile: []SanitiseMobileFunc{
			StripSpaces,
			StripCountryCode,
			StripLeadingZero,
			ValidateAfterCC,
			PrependCC,
		},
	},
	ZW: {
		Name:        "Zimbabwe",
		Nationality: "Zimbabwean",
		ISO:         ZW,
		CC:          "263",
	},
	AF: {
		Name:        "Afghanistan",
		Nationality: "Afghan",
		ISO:         AF,
		CC:          "93",
	},
	AX: {
		Name:        "Aland Islands",
		Nationality: "\u00c5land Island",
		ISO:         AX,
		CC:          "358",
	},
	AL: {
		Name:        "Albania",
		Nationality: "Albanian",
		ISO:         AL,
		CC:          "355",
	},
	DZ: {
		Name:        "Algeria",
		Nationality: "Algerian",
		ISO:         DZ,
		CC:          "213",
	},
	AS: {
		Name:        "American Samoa",
		Nationality: "American Samoan",
		ISO:         AS,
		CC:          "1684",
	},
	AD: {
		Name:        "Andorra",
		Nationality: "Andorran",
		ISO:         AD,
		CC:          "376",
	},
	AO: {
		Name:        "Angola",
		Nationality: "Angolan",
		ISO:         AO,
		CC:          "244",
	},
	AI: {
		Name:        "Anguilla",
		Nationality: "Anguillan",
		ISO:         AI,
		CC:          "1264",
	},
	AQ: {
		Name:        "Antarctica",
		Nationality: "Antarctic",
		ISO:         AQ,
		CC:          "672",
	},
	AG: {
		Name:        "Antigua and Barbuda",
		Nationality: "Antiguan or Barbudan",
		ISO:         AG,
		CC:          "1268",
	},
	AR: {
		Name:        "Argentina",
		Nationality: "Argentine",
		ISO:         AR,
		CC:          "54",
	},
	AM: {
		Name:        "Armenia",
		Nationality: "Armenian",
		ISO:         AM,
		CC:          "374",
	},
	AW: {
		Name:        "Aruba",
		Nationality: "Aruban",
		ISO:         AW,
		CC:          "297",
	},
	AC: {
		Name:        "Ascension Island",
		Nationality: "",
		ISO:         AC,
		CC:          "247",
	},
	AU: {
		Name:        "Australia",
		Nationality: "Australian",
		ISO:         AU,
		CC:          "61",
	},
	AT: {
		Name:        "Austria",
		Nationality: "Austrian",
		ISO:         AT,
		CC:          "43",
	},
	AZ: {
		Name:        "Azerbaijan",
		Nationality: "Azerbaijani, Azeri",
		ISO:         AZ,
		CC:          "994",
	},
	BS: {
		Name:        "Bahamas",
		Nationality: "Bahamian",
		ISO:         BS,
		CC:          "1242",
	},
	BH: {
		Name:        "Bahrain",
		Nationality: "Bahraini",
		ISO:         BH,
		CC:          "973",
	},
	BD: {
		Name:        "Bangladesh",
		Nationality: "Bangladeshi",
		ISO:         BD,
		CC:          "880",
	},
	BB: {
		Name:        "Barbados",
		Nationality: "Barbadian",
		ISO:         BB,
		CC:          "1246",
	},
	BY: {
		Name:        "Belarus",
		Nationality: "Belarusian",
		ISO:         BY,
		CC:          "375",
	},
	BE: {
		Name:        "Belgium",
		Nationality: "Belgian",
		ISO:         BE,
		CC:          "32",
	},
	BZ: {
		Name:        "Belize",
		Nationality: "Belizean",
		ISO:         BZ,
		CC:          "501",
	},
	BJ: {
		Name:        "Benin",
		Nationality: "Beninese, Beninois",
		ISO:         BJ,
		CC:          "229",
	},
	BM: {
		Name:        "Bermuda",
		Nationality: "Bermudian, Bermudan",
		ISO:         BM,
		CC:          "1441",
	},
	BT: {
		Name:        "Bhutan",
		Nationality: "Bhutanese",
		ISO:         BT,
		CC:          "975",
	},
	BO: {
		Name:        "Bolivia",
		Nationality: "Bolivian",
		ISO:         BO,
		CC:          "591",
	},
	BA: {
		Name:        "Bosnia and Herzegovina",
		Nationality: "Bosnian or Herzegovinian",
		ISO:         BA,
		CC:          "387",
	},
	BW: {
		Name:        "Botswana",
		Nationality: "Motswana, Botswanan",
		ISO:         BW,
		CC:          "267",
	},
	BR: {
		Name:        "Brazil",
		Nationality: "Brazilian",
		ISO:         BR,
		CC:          "55",
	},
	IO: {
		Name:        "British Indian Ocean Territory",
		Nationality: "BIOT",
		ISO:         IO,
		CC:          "246",
	},
	BN: {
		Name:        "Brunei Darussalam",
		Nationality: "Bruneian",
		ISO:         BN,
		CC:          "673",
	},
	BG: {
		Name:        "Bulgaria",
		Nationality: "Bulgarian",
		ISO:         BG,
		CC:          "359",
	},
	BF: {
		Name:        "Burkina Faso",
		Nationality: "Burkinab\u00e9",
		ISO:         BF,
		CC:          "226",
	},
	BI: {
		Name:        "Burundi",
		Nationality: "Burundian",
		ISO:         BI,
		CC:          "257",
	},
	KH: {
		Name:        "Cambodia",
		Nationality: "Cambodian",
		ISO:         KH,
		CC:          "855",
	},
	CM: {
		Name:        "Cameroon",
		Nationality: "Cameroonian",
		ISO:         CM,
		CC:          "237",
	},
	CA: {
		Name:        "Canada",
		Nationality: "Canadian",
		ISO:         CA,
		CC:          "1",
	},
	CV: {
		Name:        "Cape Verde",
		Nationality: "Cabo Verdean",
		ISO:         CV,
		CC:          "238",
	},
	KY: {
		Name:        "Cayman Islands",
		Nationality: "Caymanian",
		ISO:         KY,
		CC:          "1345",
	},
	CF: {
		Name:        "Central African Republic",
		Nationality: "Central African",
		ISO:         CF,
		CC:          "236",
	},
	TD: {
		Name:        "Chad",
		Nationality: "Chadian",
		ISO:         TD,
		CC:          "235",
	},
	CL: {
		Name:        "Chile",
		Nationality: "Chilean",
		ISO:         CL,
		CC:          "56",
	},
	CN: {
		Name:        "China",
		Nationality: "Chinese",
		ISO:         CN,
		CC:          "86",
	},
	CX: {
		Name:        "Christmas Island",
		Nationality: "Christmas Island",
		ISO:         CX,
		CC:          "61",
	},
	CC: {
		Name:        "Cocos (Keeling) Islands",
		Nationality: "Cocos Island",
		ISO:         CC,
		CC:          "61",
	},
	CO: {
		Name:        "Colombia",
		Nationality: "Colombian",
		ISO:         CO,
		CC:          "57",
	},
	KM: {
		Name:        "Comoros",
		Nationality: "Comoran, Comorian",
		ISO:         KM,
		CC:          "269",
	},
	CG: {
		Name:        "Congo",
		Nationality: "Congolese",
		ISO:         CG,
		CC:          "242",
	},
	CK: {
		Name:        "Cook Islands",
		Nationality: "Cook Island",
		ISO:         CK,
		CC:          "682",
	},
	CR: {
		Name:        "Costa Rica",
		Nationality: "Costa Rican",
		ISO:         CR,
		CC:          "506",
	},
	HR: {
		Name:        "Croatia",
		Nationality: "Croatian",
		ISO:         HR,
		CC:          "385",
	},
	CU: {
		Name:        "Cuba",
		Nationality: "Cuban",
		ISO:         CU,
		CC:          "53",
	},
	CY: {
		Name:        "Cyprus",
		Nationality: "Cypriot",
		ISO:         CY,
		CC:          "357",
	},
	CZ: {
		Name:        "Czech Republic",
		Nationality: "Czech",
		ISO:         CZ,
		CC:          "420",
	},
	CD: {
		Name:        "Democratic Republic of the Congo",
		Nationality: "Congolese",
		ISO:         CD,
		CC:          "243",
	},
	DK: {
		Name:        "Denmark",
		Nationality: "Danish",
		ISO:         DK,
		CC:          "45",
	},
	DJ: {
		Name:        "Djibouti",
		Nationality: "Djiboutian",
		ISO:         DJ,
		CC:          "253",
	},
	DM: {
		Name:        "Dominica",
		Nationality: "Dominican",
		ISO:         DM,
		CC:          "1767",
	},
	DO: {
		Name:        "Dominican Republic",
		Nationality: "Dominican",
		ISO:         DO,
		CC:          "1849",
	},
	EC: {
		Name:        "Ecuador",
		Nationality: "Ecuadorian",
		ISO:         EC,
		CC:          "593",
	},
	EG: {
		Name:        "Egypt",
		Nationality: "Egyptian",
		ISO:         EG,
		CC:          "20",
	},
	SV: {
		Name:        "El Salvador",
		Nationality: "Salvadoran",
		ISO:         SV,
		CC:          "503",
	},
	GQ: {
		Name:        "Equatorial Guinea",
		Nationality: "Equatorial Guinean, Equatoguinean",
		ISO:         GQ,
		CC:          "240",
	},
	ER: {
		Name:        "Eritrea",
		Nationality: "Eritrean",
		ISO:         ER,
		CC:          "291",
	},
	EE: {
		Name:        "Estonia",
		Nationality: "Estonian",
		ISO:         EE,
		CC:          "372",
	},
	SZ: {
		Name:        "Eswatini",
		Nationality: "Swazi",
		ISO:         SZ,
		CC:          "268",
	},
	ET: {
		Name:        "Ethiopia",
		Nationality: "Ethiopian",
		ISO:         ET,
		CC:          "251",
	},
	FK: {
		Name:        "Falkland Islands (Malvinas)",
		Nationality: "Falkland Island",
		ISO:         FK,
		CC:          "500",
	},
	FO: {
		Name:        "Faroe Islands",
		Nationality: "Faroese",
		ISO:         FO,
		CC:          "298",
	},
	FJ: {
		Name:        "Fiji",
		Nationality: "Fijian",
		ISO:         FJ,
		CC:          "679",
	},
	FI: {
		Name:        "Finland",
		Nationality: "Finnish",
		ISO:         FI,
		CC:          "358",
	},
	FR: {
		Name:        "France",
		Nationality: "French",
		ISO:         FR,
		CC:          "33",
	},
	GF: {
		Name:        "French Guiana",
		Nationality: "French Guianese",
		ISO:         GF,
		CC:          "594",
	},
	PF: {
		Name:        "French Polynesia",
		Nationality: "French Polynesian",
		ISO:         PF,
		CC:          "689",
	},
	GA: {
		Name:        "Gabon",
		Nationality: "Gabonese",
		ISO:         GA,
		CC:          "241",
	},
	GM: {
		Name:        "Gambia",
		Nationality: "Gambian",
		ISO:         GM,
		CC:          "220",
	},
	GE: {
		Name:        "Georgia",
		Nationality: "Georgian",
		ISO:         GE,
		CC:          "995",
	},
	DE: {
		Name:        "Germany",
		Nationality: "German",
		ISO:         DE,
		CC:          "49",
	},
	GH: {
		Name:        "Ghana",
		Nationality: "Ghanaian",
		ISO:         GH,
		CC:          "233",
	},
	GI: {
		Name:        "Gibraltar",
		Nationality: "Gibraltar",
		ISO:         GI,
		CC:          "350",
	},
	GR: {
		Name:        "Greece",
		Nationality: "Greek, Hellenic",
		ISO:         GR,
		CC:          "30",
	},
	GL: {
		Name:        "Greenland",
		Nationality: "Greenlandic",
		ISO:         GL,
		CC:          "299",
	},
	GD: {
		Name:        "Grenada",
		Nationality: "Grenadian",
		ISO:         GD,
		CC:          "1473",
	},
	GP: {
		Name:        "Guadeloupe",
		Nationality: "Guadeloupe",
		ISO:         GP,
		CC:          "590",
	},
	GU: {
		Name:        "Guam",
		Nationality: "Guamanian, Guambat",
		ISO:         GU,
		CC:          "1671",
	},
	GT: {
		Name:        "Guatemala",
		Nationality: "Guatemalan",
		ISO:         GT,
		CC:          "502",
	},
	GG: {
		Name:        "Guernsey",
		Nationality: "Channel Island",
		ISO:         GG,
		CC:          "44-1481",
	},
	GN: {
		Name:        "Guinea",
		Nationality: "Guinean",
		ISO:         GN,
		CC:          "224",
	},
	GW: {
		Name:        "Guinea-Bissau",
		Nationality: "Bissau-Guinean",
		ISO:         GW,
		CC:          "245",
	},
	GY: {
		Name:        "Guyana",
		Nationality: "Guyanese",
		ISO:         GY,
		CC:          "592",
	},
	HT: {
		Name:        "Haiti",
		Nationality: "Haitian",
		ISO:         HT,
		CC:          "509",
	},
	VA: {
		Name:        "Holy See (Vatican City State)",
		Nationality: "Vatican",
		ISO:         VA,
		CC:          "379",
	},
	HN: {
		Name:        "Honduras",
		Nationality: "Honduran",
		ISO:         HN,
		CC:          "504",
	},
	HK: {
		Name:        "Hong Kong",
		Nationality: "Hong Kong, Hong Kongese",
		ISO:         HK,
		CC:          "852",
	},
	HU: {
		Name:        "Hungary",
		Nationality: "Hungarian, Magyar",
		ISO:         HU,
		CC:          "36",
	},
	IS: {
		Name:        "Iceland",
		Nationality: "Icelandic",
		ISO:         IS,
		CC:          "354",
	},
	IN: {
		Name:        "India",
		Nationality: "Indian",
		ISO:         IN,
		CC:          "91",
	},
	ID: {
		Name:        "Indonesia",
		Nationality: "Indonesian",
		ISO:         ID,
		CC:          "62",
	},
	IR: {
		Name:        "Iran",
		Nationality: "Iranian, Persian",
		ISO:         IR,
		CC:          "98",
	},
	IQ: {
		Name:        "Iraq",
		Nationality: "Iraqi",
		ISO:         IQ,
		CC:          "964",
	},
	IE: {
		Name:        "Ireland",
		Nationality: "Irish",
		ISO:         IE,
		CC:          "353",
	},
	IM: {
		Name:        "Isle of Man",
		Nationality: "Manx",
		ISO:         IM,
		CC:          "44-1624",
	},
	IL: {
		Name:        "Israel",
		Nationality: "Israeli",
		ISO:         IL,
		CC:          "972",
	},
	IT: {
		Name:        "Italy",
		Nationality: "Italian",
		ISO:         IT,
		CC:          "39",
	},
	CI: {
		Name:        "Ivory Coast / Cote d'Ivoire",
		Nationality: "Ivorian",
		ISO:         CI,
		CC:          "225",
	},
	JM: {
		Name:        "Jamaica",
		Nationality: "Jamaican",
		ISO:         JM,
		CC:          "1876",
	},
	JP: {
		Name:        "Japan",
		Nationality: "Japanese",
		ISO:         JP,
		CC:          "81",
	},
	JE: {
		Name:        "Jersey",
		Nationality: "Channel Island",
		ISO:         JE,
		CC:          "44-1534",
	},
	JO: {
		Name:        "Jordan",
		Nationality: "Jordanian",
		ISO:         JO,
		CC:          "962",
	},
	KZ: {
		Name:        "Kazakhstan",
		Nationality: "Kazakhstani, Kazakh",
		ISO:         KZ,
		CC:          "77",
	},
	KE: {
		Name:        "Kenya",
		Nationality: "Kenyan",
		ISO:         KE,
		CC:          "254",
	},
	KI: {
		Name:        "Kiribati",
		Nationality: "I-Kiribati",
		ISO:         KI,
		CC:          "686",
	},
	KP: {
		Name:        "Korea, Democratic People's Republic of Korea",
		Nationality: "North Korean",
		ISO:         KP,
		CC:          "850",
	},
	KR: {
		Name:        "Korea, Republic of South Korea",
		Nationality: "South Korean",
		ISO:         KR,
		CC:          "82",
	},
	XK: {
		Name:        "Kosovo",
		Nationality: "",
		ISO:         XK,
		CC:          "383",
	},
	KW: {
		Name:        "Kuwait",
		Nationality: "Kuwaiti",
		ISO:         KW,
		CC:          "965",
	},
	KG: {
		Name:        "Kyrgyzstan",
		Nationality: "Kyrgyzstani, Kyrgyz, Kirgiz, Kirghiz",
		ISO:         KG,
		CC:          "996",
	},
	LA: {
		Name:        "Laos",
		Nationality: "Lao, Laotian",
		ISO:         LA,
		CC:          "856",
	},
	LV: {
		Name:        "Latvia",
		Nationality: "Latvian",
		ISO:         LV,
		CC:          "371",
	},
	LB: {
		Name:        "Lebanon",
		Nationality: "Lebanese",
		ISO:         LB,
		CC:          "961",
	},
	LS: {
		Name:        "Lesotho",
		Nationality: "Basotho",
		ISO:         LS,
		CC:          "266",
	},
	LR: {
		Name:        "Liberia",
		Nationality: "Liberian",
		ISO:         LR,
		CC:          "231",
	},
	LY: {
		Name:        "Libya",
		Nationality: "Libyan",
		ISO:         LY,
		CC:          "218",
	},
	LI: {
		Name:        "Liechtenstein",
		Nationality: "Liechtenstein",
		ISO:         LI,
		CC:          "423",
	},
	LT: {
		Name:        "Lithuania",
		Nationality: "Lithuanian",
		ISO:         LT,
		CC:          "370",
	},
	LU: {
		Name:        "Luxembourg",
		Nationality: "Luxembourg, Luxembourgish",
		ISO:         LU,
		CC:          "352",
	},
	MO: {
		Name:        "Macau",
		Nationality: "Macanese, Chinese",
		ISO:         MO,
		CC:          "853",
	},
	MG: {
		Name:        "Madagascar",
		Nationality: "Malagasy",
		ISO:         MG,
		CC:          "261",
	},
	MW: {
		Name:        "Malawi",
		Nationality: "Malawian",
		ISO:         MW,
		CC:          "265",
	},
	MY: {
		Name:        "Malaysia",
		Nationality: "Malaysian",
		ISO:         MY,
		CC:          "60",
	},
	MV: {
		Name:        "Maldives",
		Nationality: "Maldivian",
		ISO:         MV,
		CC:          "960",
	},
	ML: {
		Name:        "Mali",
		Nationality: "Malian, Malinese",
		ISO:         ML,
		CC:          "223",
	},
	MT: {
		Name:        "Malta",
		Nationality: "Maltese",
		ISO:         MT,
		CC:          "356",
	},
	MH: {
		Name:        "Marshall Islands",
		Nationality: "Marshallese",
		ISO:         MH,
		CC:          "692",
	},
	MQ: {
		Name:        "Martinique",
		Nationality: "Martiniquais, Martinican",
		ISO:         MQ,
		CC:          "596",
	},
	MR: {
		Name:        "Mauritania",
		Nationality: "Mauritanian",
		ISO:         MR,
		CC:          "222",
	},
	MU: {
		Name:        "Mauritius",
		Nationality: "Mauritian",
		ISO:         MU,
		CC:          "230",
	},
	YT: {
		Name:        "Mayotte",
		Nationality: "Mahoran",
		ISO:         YT,
		CC:          "262",
	},
	MX: {
		Name:        "Mexico",
		Nationality: "Mexican",
		ISO:         MX,
		CC:          "52",
	},
	FM: {
		Name:        "Micronesia, Federated States of Micronesia",
		Nationality: "Micronesian",
		ISO:         FM,
		CC:          "691",
	},
	MD: {
		Name:        "Moldova",
		Nationality: "Moldovan",
		ISO:         MD,
		CC:          "373",
	},
	MC: {
		Name:        "Monaco",
		Nationality: "Mon\u00e9gasque, Monacan",
		ISO:         MC,
		CC:          "377",
	},
	MN: {
		Name:        "Mongolia",
		Nationality: "Mongolian",
		ISO:         MN,
		CC:          "976",
	},
	ME: {
		Name:        "Montenegro",
		Nationality: "Montenegrin",
		ISO:         ME,
		CC:          "382",
	},
	MS: {
		Name:        "Montserrat",
		Nationality: "Montserratian",
		ISO:         MS,
		CC:          "1664",
	},
	MA: {
		Name:        "Morocco",
		Nationality: "Moroccan",
		ISO:         MA,
		CC:          "212",
	},
	MZ: {
		Name:        "Mozambique",
		Nationality: "Mozambican",
		ISO:         MZ,
		CC:          "258",
	},
	MM: {
		Name:        "Myanmar",
		Nationality: "Burmese",
		ISO:         MM,
		CC:          "95",
	},
	NA: {
		Name:        "Namibia",
		Nationality: "Namibian",
		ISO:         NA,
		CC:          "264",
	},
	NR: {
		Name:        "Nauru",
		Nationality: "Nauruan",
		ISO:         NR,
		CC:          "674",
	},
	NP: {
		Name:        "Nepal",
		Nationality: "Nepali, Nepalese",
		ISO:         NP,
		CC:          "977",
	},
	NL: {
		Name:        "Netherlands",
		Nationality: "Dutch, Netherlandic",
		ISO:         NL,
		CC:          "31",
	},
	AN: {
		Name:        "Netherlands Antilles",
		Nationality: "",
		ISO:         AN,
		CC:          "599",
	},
	NC: {
		Name:        "New Caledonia",
		Nationality: "New Caledonian",
		ISO:         NC,
		CC:          "687",
	},
	NZ: {
		Name:        "New Zealand",
		Nationality: "New Zealand, NZ",
		ISO:         NZ,
		CC:          "64",
	},
	NI: {
		Name:        "Nicaragua",
		Nationality: "Nicaraguan",
		ISO:         NI,
		CC:          "505",
	},
	NE: {
		Name:        "Niger",
		Nationality: "Nigerien",
		ISO:         NE,
		CC:          "227",
	},
	NG: {
		Name:        "Nigeria",
		Nationality: "Nigerian",
		ISO:         NG,
		CC:          "234",
	},
	NU: {
		Name:        "Niue",
		Nationality: "Niuean",
		ISO:         NU,
		CC:          "683",
	},
	NF: {
		Name:        "Norfolk Island",
		Nationality: "Norfolk Island",
		ISO:         NF,
		CC:          "672",
	},
	MK: {
		Name:        "North Macedonia",
		Nationality: "Macedonian",
		ISO:         MK,
		CC:          "389",
	},
	MP: {
		Name:        "Northern Mariana Islands",
		Nationality: "Northern Marianan",
		ISO:         MP,
		CC:          "1670",
	},
	NO: {
		Name:        "Norway",
		Nationality: "Norwegian",
		ISO:         NO,
		CC:          "47",
	},
	OM: {
		Name:        "Oman",
		Nationality: "Omani",
		ISO:         OM,
		CC:          "968",
	},
	PK: {
		Name:        "Pakistan",
		Nationality: "Pakistani",
		ISO:         PK,
		CC:          "92",
	},
	PW: {
		Name:        "Palau",
		Nationality: "Palauan",
		ISO:         PW,
		CC:          "680",
	},
	PS: {
		Name:        "Palestine",
		Nationality: "Palestinian",
		ISO:         PS,
		CC:          "970",
	},
	PA: {
		Name:        "Panama",
		Nationality: "Panamanian",
		ISO:         PA,
		CC:          "507",
	},
	PG: {
		Name:        "Papua New Guinea",
		Nationality: "Papua New Guinean, Papuan",
		ISO:         PG,
		CC:          "675",
	},
	PY: {
		Name:        "Paraguay",
		Nationality: "Paraguayan",
		ISO:         PY,
		CC:          "595",
	},
	PE: {
		Name:        "Peru",
		Nationality: "Peruvian",
		ISO:         PE,
		CC:          "51",
	},
	PH: {
		Name:        "Philippines",
		Nationality: "Philippine, Filipino",
		ISO:         PH,
		CC:          "63",
	},
	PN: {
		Name:        "Pitcairn",
		Nationality: "Pitcairn Island",
		ISO:         PN,
		CC:          "872",
	},
	PL: {
		Name:        "Poland",
		Nationality: "Polish",
		ISO:         PL,
		CC:          "48",
	},
	PT: {
		Name:        "Portugal",
		Nationality: "Portuguese",
		ISO:         PT,
		CC:          "351",
	},
	PR: {
		Name:        "Puerto Rico",
		Nationality: "Puerto Rican",
		ISO:         PR,
		CC:          "1939",
	},
	QA: {
		Name:        "Qatar",
		Nationality: "Qatari",
		ISO:         QA,
		CC:          "974",
	},
	RE: {
		Name:        "Reunion",
		Nationality: "R\u00e9unionese, R\u00e9unionnais",
		ISO:         RE,
		CC:          "262",
	},
	RO: {
		Name:        "Romania",
		Nationality: "Romanian",
		ISO:         RO,
		CC:          "40",
	},
	RU: {
		Name:        "Russia",
		Nationality: "Russian",
		ISO:         RU,
		CC:          "7",
	},
	RW: {
		Name:        "Rwanda",
		Nationality: "Rwandan",
		ISO:         RW,
		CC:          "250",
	},
	BL: {
		Name:        "Saint Barthelemy",
		Nationality: "Barth\u00e9lemois",
		ISO:         BL,
		CC:          "590",
	},
	SH: {
		Name:        "Saint Helena, Ascension and Tristan Da Cunha",
		Nationality: "Saint Helenian",
		ISO:         SH,
		CC:          "290",
	},
	KN: {
		Name:        "Saint Kitts and Nevis",
		Nationality: "Kittitian or Nevisian",
		ISO:         KN,
		CC:          "1869",
	},
	LC: {
		Name:        "Saint Lucia",
		Nationality: "Saint Lucian",
		ISO:         LC,
		CC:          "1758",
	},
	MF: {
		Name:        "Saint Martin",
		Nationality: "Saint-Martinoise",
		ISO:         MF,
		CC:          "590",
	},
	PM: {
		Name:        "Saint Pierre and Miquelon",
		Nationality: "Saint-Pierrais or Miquelonnais",
		ISO:         PM,
		CC:          "508",
	},
	VC: {
		Name:        "Saint Vincent and the Grenadines",
		Nationality: "Saint Vincentian, Vincentian",
		ISO:         VC,
		CC:          "1784",
	},
	WS: {
		Name:        "Samoa",
		Nationality: "Samoan",
		ISO:         WS,
		CC:          "685",
	},
	SM: {
		Name:        "San Marino",
		Nationality: "Sammarinese",
		ISO:         SM,
		CC:          "378",
	},
	ST: {
		Name:        "Sao Tome and Principe",
		Nationality: "S\u00e3o Tom\u00e9an",
		ISO:         ST,
		CC:          "239",
	},
	SA: {
		Name:        "Saudi Arabia",
		Nationality: "Saudi, Saudi Arabian",
		ISO:         SA,
		CC:          "966",
	},
	SN: {
		Name:        "Senegal",
		Nationality: "Senegalese",
		ISO:         SN,
		CC:          "221",
	},
	RS: {
		Name:        "Serbia",
		Nationality: "Serbian",
		ISO:         RS,
		CC:          "381",
	},
	SC: {
		Name:        "Seychelles",
		Nationality: "Seychellois",
		ISO:         SC,
		CC:          "248",
	},
	SL: {
		Name:        "Sierra Leone",
		Nationality: "Sierra Leonean",
		ISO:         SL,
		CC:          "232",
	},
	SG: {
		Name:        "Singapore",
		Nationality: "Singaporean",
		ISO:         SG,
		CC:          "65",
	},
	SX: {
		Name:        "Sint Maarten",
		Nationality: "Sint Maarten",
		ISO:         SX,
		CC:          "1721",
	},
	SK: {
		Name:        "Slovakia",
		Nationality: "Slovak",
		ISO:         SK,
		CC:          "421",
	},
	SI: {
		Name:        "Slovenia",
		Nationality: "Slovenian, Slovene",
		ISO:         SI,
		CC:          "386",
	},
	SB: {
		Name:        "Solomon Islands",
		Nationality: "Solomon Island",
		ISO:         SB,
		CC:          "677",
	},
	SO: {
		Name:        "Somalia",
		Nationality: "Somali, Somalian",
		ISO:         SO,
		CC:          "252",
	},
	GS: {
		Name:        "South Georgia and the South Sandwich Islands",
		Nationality: "South Georgia or South Sandwich Islands",
		ISO:         GS,
		CC:          "500",
	},
	SS: {
		Name:        "South Sudan",
		Nationality: "South Sudanese",
		ISO:         SS,
		CC:          "211",
	},
	ES: {
		Name:        "Spain",
		Nationality: "Spanish",
		ISO:         ES,
		CC:          "34",
	},
	LK: {
		Name:        "Sri Lanka",
		Nationality: "Sri Lankan",
		ISO:         LK,
		CC:          "94",
	},
	SD: {
		Name:        "Sudan",
		Nationality: "Sudanese",
		ISO:         SD,
		CC:          "249",
	},
	SR: {
		Name:        "Suriname",
		Nationality: "Surinamese",
		ISO:         SR,
		CC:          "597",
	},
	SJ: {
		Name:        "Svalbard and Jan Mayen",
		Nationality: "Svalbard",
		ISO:         SJ,
		CC:          "47",
	},
	SE: {
		Name:        "Sweden",
		Nationality: "Swedish",
		ISO:         SE,
		CC:          "46",
	},
	CH: {
		Name:        "Switzerland",
		Nationality: "Swiss",
		ISO:         CH,
		CC:          "41",
	},
	SY: {
		Name:        "Syrian Arab Republic",
		Nationality: "Syrian",
		ISO:         SY,
		CC:          "963",
	},
	TW: {
		Name:        "Taiwan",
		Nationality: "Chinese, Taiwanese",
		ISO:         TW,
		CC:          "886",
	},
	TJ: {
		Name:        "Tajikistan",
		Nationality: "Tajikistani",
		ISO:         TJ,
		CC:          "992",
	},
	TZ: {
		Name:        "Tanzania, United Republic of Tanzania",
		Nationality: "Tanzanian",
		ISO:         TZ,
		CC:          "255",
	},
	TH: {
		Name:        "Thailand",
		Nationality: "Thai",
		ISO:         TH,
		CC:          "66",
	},
	TL: {
		Name:        "Timor-Leste",
		Nationality: "Timorese",
		ISO:         TL,
		CC:          "670",
	},
	TG: {
		Name:        "Togo",
		Nationality: "Togolese",
		ISO:         TG,
		CC:          "228",
	},
	TK: {
		Name:        "Tokelau",
		Nationality: "Tokelauan",
		ISO:         TK,
		CC:          "690",
	},
	TO: {
		Name:        "Tonga",
		Nationality: "Tongan",
		ISO:         TO,
		CC:          "676",
	},
	TT: {
		Name:        "Trinidad and Tobago",
		Nationality: "Trinidadian or Tobagonian",
		ISO:         TT,
		CC:          "1868",
	},
	TN: {
		Name:        "Tunisia",
		Nationality: "Tunisian",
		ISO:         TN,
		CC:          "216",
	},
	TR: {
		Name:        "Turkey",
		Nationality: "Turkish",
		ISO:         TR,
		CC:          "90",
	},
	TM: {
		Name:        "Turkmenistan",
		Nationality: "Turkmen",
		ISO:         TM,
		CC:          "993",
	},
	TC: {
		Name:        "Turks and Caicos Islands",
		Nationality: "Turks and Caicos Island",
		ISO:         TC,
		CC:          "1649",
	},
	TV: {
		Name:        "Tuvalu",
		Nationality: "Tuvaluan",
		ISO:         TV,
		CC:          "688",
	},
	UG: {
		Name:        "Uganda",
		Nationality: "Ugandan",
		ISO:         UG,
		CC:          "256",
	},
	UA: {
		Name:        "Ukraine",
		Nationality: "Ukrainian",
		ISO:         UA,
		CC:          "380",
	},
	AE: {
		Name:        "United Arab Emirates",
		Nationality: "Emirati, Emirian, Emiri",
		ISO:         AE,
		CC:          "971",
	},
	GB: {
		Name:        "United Kingdom",
		Nationality: "British, UK",
		ISO:         GB,
		CC:          "44",
	},
	US: {
		Name:        "United States",
		Nationality: "American",
		ISO:         US,
		CC:          "1",
	},
	UMI: {
		Name:        "United States Minor Outlying Islands",
		Nationality: "",
		ISO:         UMI,
		CC:          "246",
	},
	UY: {
		Name:        "Uruguay",
		Nationality: "Uruguayan",
		ISO:         UY,
		CC:          "598",
	},
	UZ: {
		Name:        "Uzbekistan",
		Nationality: "Uzbekistani, Uzbek",
		ISO:         UZ,
		CC:          "998",
	},
	VU: {
		Name:        "Vanuatu",
		Nationality: "Ni-Vanuatu, Vanuatuan",
		ISO:         VU,
		CC:          "678",
	},
	VE: {
		Name:        "Venezuela, Bolivarian Republic of Venezuela",
		Nationality: "Venezuelan",
		ISO:         VE,
		CC:          "58",
	},
	VN: {
		Name:        "Vietnam",
		Nationality: "Vietnamese",
		ISO:         VN,
		CC:          "84",
	},
	VG: {
		Name:        "Virgin Islands, British",
		Nationality: "British Virgin Island",
		ISO:         VG,
		CC:          "1284",
	},
	VI: {
		Name:        "Virgin Islands, U.S.",
		Nationality: "U.S. Virgin Island",
		ISO:         VI,
		CC:          "1340",
	},
	WF: {
		Name:        "Wallis and Futuna",
		Nationality: "Wallis and Futuna, Wallisian or Futunan",
		ISO:         WF,
		CC:          "681",
	},
	YE: {
		Name:        "Yemen",
		Nationality: "Yemeni",
		ISO:         YE,
		CC:          "967",
	},
	ZM: {
		Name:        "Zambia",
		Nationality: "Zambian",
		ISO:         ZM,
		CC:          "260",
	},
}
