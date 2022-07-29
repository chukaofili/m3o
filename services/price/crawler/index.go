package crawler

var (
	Index = map[string]string{
		"AFN":      "Afghan Afghani",
		"ALL":      "Albanian Lek",
		"DZD":      "Algerian Dinar",
		"ALU":      "Aluminum",
		"AOA":      "Angolan Kwanza",
		"ARS":      "Argentine Peso",
		"AMD":      "Armenian Dram",
		"AWG":      "Aruban Florin",
		"AUD":      "Australian Dollar",
		"AZN":      "Azerbaijani Manat",
		"BSD":      "Bahamian Dollar",
		"BHD":      "Bahraini Dinar",
		"BDT":      "Bangladeshi Taka",
		"BBD":      "Barbadian Dollar",
		"BYN":      "Belarusian Ruble",
		"BYR":      "Belarusian Ruble",
		"BZD":      "Belize Dollar",
		"BMD":      "Bermudian Dollar",
		"BTN":      "Bhutanese Ngultrum",
		"BTC":      "Bitcoin",
		"BCH":      "Bitcoin Cash",
		"BOB":      "Bolivian Boliviano",
		"BAM":      "Bosnia and Herzegovina Convertible Mark",
		"BWP":      "Botswana Pula",
		"BRL":      "Brazilian Real",
		"BRENTOIL": "Brent Crude Oil",
		"GBP":      "British Pound",
		"BND":      "Brunei Dollar",
		"BGN":      "Bulgarian Lev",
		"BIF":      "Burundian Franc",
		"KHR":      "Cambodian Riel",
		"CAD":      "Canadian Dollar",
		"CVE":      "Cape Verdean Escudo",
		"ADA":      "Cardano",
		"KYD":      "Cayman Islands Dollar",
		"XAF":      "Central African Cfa Franc",
		"XPF":      "Cfp Franc",
		"LINK":     "Chainlink",
		"CLP":      "Chilean Peso",
		"CNY":      "Chinese Renminbi Yuan",
		"CNH":      "Chinese Renminbi Yuan Offshore",
		"LCO":      "Cobalt (Troy Ounce)",
		"COFFEE":   "Coffee",
		"COP":      "Colombian Peso",
		"KMF":      "Comorian Franc",
		"CDF":      "Congolese Franc",
		"XCU":      "Copper",
		"CORN":     "Corn",
		"CRC":      "Costa Rican Colón",
		"COTTON":   "Cotton",
		"HRK":      "Croatian Kuna",
		"CPO":      "Crude Palm Oil",
		"CUC":      "Cuban Convertible Peso",
		"CZK":      "Czech Koruna",
		"DKK":      "Danish Krone",
		"DJF":      "Djiboutian Franc",
		"DOP":      "Dominican Peso",
		"XCD":      "East Caribbean Dollar",
		"EGP":      "Egyptian Pound",
		"ERN":      "Eritrean Nakfa",
		"EEK":      "Estonian Kroon",
		"ETHANOL":  "Ethanol",
		"ETH":      "Ethereum",
		"ETB":      "Ethiopian Birr",
		"EUR":      "Euro",
		"FKP":      "Falkland Pound",
		"FJD":      "Fijian Dollar",
		"GMD":      "Gambian Dalasi",
		"GEL":      "Georgian Lari",
		"GHS":      "Ghanaian Cedi",
		"GIP":      "Gibraltar Pound",
		"XAU":      "Gold (Troy Ounce)",
		"GTQ":      "Guatemalan Quetzal",
		"GGP":      "Guernsey Pound",
		"GNF":      "Guinean Franc",
		"GYD":      "Guyanese Dollar",
		"HTG":      "Haitian Gourde",
		"HNL":      "Honduran Lempira",
		"HKD":      "Hong Kong Dollar",
		"HUF":      "Hungarian Forint",
		"ISK":      "Icelandic Króna",
		"INR":      "Indian Rupee",
		"IDR":      "Indonesian Rupiah",
		"IQD":      "Iraqi Dinar",
		"IRD":      "Iridium (Troy Ounce)",
		"IMP":      "Isle of Man Pound",
		"ILS":      "Israeli New Sheqel",
		"JMD":      "Jamaican Dollar",
		"JPY":      "Japanese Yen",
		"JEP":      "Jersey Pound",
		"JOD":      "Jordanian Dinar",
		"KZT":      "Kazakhstani Tenge",
		"KES":      "Kenyan Shilling",
		"KWD":      "Kuwaiti Dinar",
		"KGS":      "Kyrgyzstani Som",
		"LAK":      "Lao Kip",
		"LVL":      "Latvian Lats",
		"LBP":      "Lebanese Pound",
		"LSL":      "Lesotho Loti",
		"LRD":      "Liberian Dollar",
		"LYD":      "Libyan Dinar",
		"LTC":      "Litecoin",
		"LTL":      "Lithuanian Litas",
		"MOP":      "Macanese Pataca",
		"MKD":      "Macedonian Denar",
		"MGA":      "Malagasy Ariary",
		"MWK":      "Malawian Kwacha",
		"MYR":      "Malaysian Ringgit",
		"MVR":      "Maldivian Rufiyaa",
		"MTL":      "Maltese Lira",
		"MRO":      "Mauritanian Ouguiya",
		"MUR":      "Mauritian Rupee",
		"MXN":      "Mexican Peso",
		"MDL":      "Moldovan Leu",
		"MNT":      "Mongolian Tögrög",
		"MAD":      "Moroccan Dirham",
		"MZN":      "Mozambican Metical",
		"MMK":      "Myanmar Kyat",
		"NAD":      "Namibian Dollar",
		"NG":       "Natural Gas",
		"NPR":      "Nepalese Rupee",
		"ANG":      "Netherlands Antillean Gulden",
		"TWD":      "New Taiwan Dollar",
		"NZD":      "New Zealand Dollar",
		"NIO":      "Nicaraguan Córdoba",
		"NI":       "Nickel",
		"NGN":      "Nigerian Naira",
		"NOK":      "Norwegian Krone",
		"OMR":      "Omani Rial",
		"PKR":      "Pakistani Rupee",
		"XPD":      "Palladium (Troy Ounce)",
		"PAB":      "Panamanian Balboa",
		"PGK":      "Papua New Guinean Kina",
		"PYG":      "Paraguayan Guaraní",
		"PEN":      "Peruvian Sol",
		"PHP":      "Philippine Peso",
		"XPT":      "Platinum (Troy Ounce)",
		"PLN":      "Polish Złoty",
		"QAR":      "Qatari Riyal",
		"XRH":      "Rhodium (Troy Ounce)",
		"RICE":     "Rice",
		"XRP":      "Ripple",
		"XRP2":     "Ripple",
		"RON":      "Romanian Leu",
		"RUBBER":   "Rubber",
		"RUB":      "Russian Ruble",
		"RUTH":     "Ruthenium",
		"RWF":      "Rwandan Franc",
		"SHP":      "Saint Helenian Pound",
		"SVC":      "Salvadoran Colón",
		"WST":      "Samoan Tala",
		"STD":      "São Tomé and Príncipe Dobra",
		"SAR":      "Saudi Riyal",
		"RSD":      "Serbian Dinar",
		"SCR":      "Seychellois Rupee",
		"SLL":      "Sierra Leonean Leone",
		"XAG":      "Silver (Troy Ounce)",
		"SGD":      "Singapore Dollar",
		"SBD":      "Solomon Islands Dollar",
		"SOS":      "Somali Shilling",
		"ZAR":      "South African Rand",
		"KRW":      "South Korean Won",
		"SSP":      "South Sudanese Pound",
		"SOYBEAN":  "Soybeans",
		"XDR":      "Special Drawing Rights",
		"LKR":      "Sri Lankan Rupee",
		"XLM":      "Stellar",
		"SUGAR":    "Sugar",
		"SRD":      "Surinamese Dollar",
		"SZL":      "Swazi Lilangeni",
		"SEK":      "Swedish Krona",
		"CHF":      "Swiss Franc",
		"TJS":      "Tajikistani Somoni",
		"TZS":      "Tanzanian Shilling",
		"THB":      "Thai Baht",
		"TIN":      "Tin",
		"TOP":      "Tongan Paʻanga",
		"TTD":      "Trinidad and Tobago Dollar",
		"TND":      "Tunisian Dinar",
		"TRY":      "Turkish Lira",
		"TMT":      "Turkmenistani Manat",
		"UGX":      "Ugandan Shilling",
		"UAH":      "Ukrainian Hryvnia",
		"CLF":      "Unidad de Fomento",
		"UNI":      "Uniswap",
		"AED":      "United Arab Emirates Dirham",
		"USD":      "United States Dollar",
		"UYU":      "Uruguayan Peso",
		"UZS":      "Uzbekistan Som",
		"VUV":      "Vanuatu Vatu",
		"VEF":      "Venezuelan Bolívar",
		"VES":      "Venezuelan Bolívar Soberano",
		"VND":      "Vietnamese Đồng",
		"XOF":      "West African Cfa Franc",
		"WHEAT":    "Wheat",
		"WTIOIL":   "WTI Crude Oil",
		"YER":      "Yemeni Rial",
		"ZMK":      "Zambian Kwacha",
		"ZMW":      "Zambian Kwacha",
		"ZWL":      "Zimbabwean Dollar",
		"ZNC":      "Zinc",
	}
)