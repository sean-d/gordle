package gordle

var fiveLetterWords = []string{
	"ABACK", "ABASE", "ABATE",
	"ABBEY", "ABBOT", "ABHOR",
	"ABIDE", "ABLED", "ABODE",
	"ABORT", "ABOUT", "ABOVE",
	"ABUSE", "ABYSS", "ACORN",
	"ACRID", "ACTOR", "ACUTE",
	"ADAGE", "ADAPT", "ADEPT",
	"ADMIN", "ADMIT", "ADOBE",
	"ADOPT", "ADORE", "ADORN",
	"ADULT", "AFFIX", "AFIRE",
	"AFOOT", "AFOUL", "AFTER",
	"AGAIN", "AGAPE", "AGATE",
	"AGENT", "AGILE", "AGING",
	"AGLOW", "AGONY", "AGORA",
	"AGREE", "AHEAD", "AIDER",
	"AISLE", "ALARM", "ALBUM",
	"ALERT", "ALGAE", "ALIBI",
	"ALIEN", "ALIGN", "ALIKE",
	"ALIVE", "ALLAY", "ALLEY",
	"ALLOT", "ALLOW", "ALLOY",
	"ALOFT", "ALONE", "ALONG",
	"ALOOF", "ALOUD", "ALPHA",
	"ALTAR", "ALTER", "AMASS",
	"AMAZE", "AMBER", "AMBLE",
	"AMEND", "AMISS", "AMITY",
	"AMONG", "AMPLE", "AMPLY",
	"AMUSE", "ANGEL", "ANGER",
	"ANGLE", "ANGRY", "ANGST",
	"ANIME", "ANKLE", "ANNEX",
	"ANNOY", "ANNUL", "ANODE",
	"ANTIC", "ANVIL", "AORTA",
	"APART", "APHID", "APING",
	"APNEA", "APPLE", "APPLY",
	"APRON", "APTLY", "ARBOR",
	"ARDOR", "ARENA", "ARGUE",
	"ARISE", "ARMOR", "AROMA",
	"AROSE", "ARRAY", "ARROW",
	"ARSON", "ARTSY", "ASCOT",
	"ASHEN", "ASIAN", "ASIDE",
	"ASKEW", "ASSAY", "ASSET",
	"ATOLL", "ATONE", "ATTIC",
	"AUDIO", "AUDIT", "AUGUR",
	"AUNTY", "AVAIL", "AVERT",
	"AVIAN", "AVOID", "AWAIT",
	"AWAKE", "AWARD", "AWARE",
	"AWASH", "AWFUL", "AWOKE",
	"AXIAL", "AXIOM", "AXION",
	"AZURE", "BACON", "BADGE",
	"BADLY", "BAGEL", "BAGGY",
	"BAKER", "BALER", "BALMY",
	"BANAL", "BANJO", "BARGE",
	"BARON", "BASAL", "BASIC",
	"BASIL", "BASIN", "BASIS",
	"BASTE", "BATCH", "BATHE",
	"BATON", "BATTY", "BAWDY",
	"BAYOU", "BEACH", "BEADY",
	"BEARD", "BEAST", "BEECH",
	"BEEFY", "BEFIT", "BEGAN",
	"BEGAT", "BEGET", "BEGIN",
	"BEGUN", "BEING", "BELCH",
	"BELIE", "BELLE", "BELLY",
	"BELOW", "BENCH", "BERET",
	"BERRY", "BERTH", "BESET",
	"BETEL", "BEVEL", "BEZEL",
	"BIBLE", "BICEP", "BIDDY",
	"BIGOT", "BILGE", "BILLY",
}

var sixLetterWords = []string{
	"ABROAD", "CASUAL", "AROUND", "COUPLE",
	"ACCEPT", "CAUGHT", "ARRIVE", "COURSE",
	"ACCESS", "CENTRE", "ARTIST", "COVERS",
	"ACROSS", "CENTUM", "ASPECT", "CREATE",
	"ACTING", "CHANCE", "ASSESS", "CREDIT",
	"ACTION", "CHANGE", "ASSIST", "CRISIS",
	"ACTIVE", "CHARGE", "ASSUME", "CUSTOM",
	"ACTUAL", "CHOICE", "ATTACK", "DAMAGE",
	"ADVICE", "CHOOSE", "ATTEND", "DANGER",
	"ADVISE", "CHOSEN", "AUGUST", "DEALER",
	"AFFECT", "CHURCH", "AUTHOR", "DEBATE",
	"AFFORD", "CIRCLE", "AVENUE", "DECADE",
	"AFRAID", "CLIENT", "BACKED", "DECIDE",
	"AGENCY", "CLOSED", "BARELY", "DEFEAT",
	"AGENDA", "CLOSER", "BATTLE", "DEFEND",
	"ALMOST", "COFFEE", "BEAUTY", "DEFINE",
	"ALWAYS", "COLUMN", "BECAME", "DEGREE",
	"AMOUNT", "COMBAT", "BECOME", "DEMAND",
	"ANIMAL", "COMING", "BEFORE", "DEPEND",
	"ANNUAL", "COMMON", "BEHALF", "DEPUTY",
	"ANSWER", "COMPLY", "BEHIND", "DESERT",
	"ANYONE", "COPPER", "BELIEF", "DESIGN",
	"ANYWAY", "CORNER", "BELONG", "DESIRE",
	"APPEAL", "COSTLY", "BEAKER", "DETAIL",
	"APPEAR", "COUNTY", "BETTER", "DETECT",
	"BEYOND", "BUDGET", "DURING", "DEVICE",
	"BISHOP", "BURDEN", "EATING", "DIFFER",
	"BORDER", "BUREAU", "EDITOR", "DINNER",
	"BOTTLE", "BUTTON", "EFFECT", "DIRECT",
	"BOTTOM", "CAMERA", "EFFORT", "DOCTOR",
	"BOUGHT", "CANCER", "EIGHTH", "DOLLAR",
	"BRANCH", "CACTUS", "EITHER", "DOMAIN",
	"BREATH", "CARBON", "ELEVEN", "DOUBLE",
	"BRIDGE", "CAREER", "EMERGE", "DRIVEN",
	"BRIGHT", "CASTLE", "DRIVER", "DRIVER",
}

var sevenLetterWords = []string{
	  "ABILITY", "BACKING", "CABINET",
    "ABSENCE", "BALANCE", "CALIBRE",
    "ACADEMY", "BANKING", "CALLING",
    "ACCOUNT", "BARRIER", "CAPABLE",
    "ACCUSED", "BATTERY", "CAPITAL",
    "ACHIEVE", "BEARING", "CAPTAIN",
    "ACQUIRE", "BEATING", "CAPTION",
    "ADDRESS", "BECAUSE", "CAPTURE",
    "ADVANCE", "BEDROOM", "CAREFUL",
    "ADVERSE", "BELIEVE", "CARRIER",
    "ADVISED", "BENEATH", "CAUTION",
    "ADVISER", "BENEFIT", "CEILING",
    "AGAINST", "BESIDES", "CENTRAL",
    "AIRLINE", "BETWEEN", "CENTRIC",
    "AIRPORT", "BILLION", "CENTURY",
    "ALCOHOL", "BINDING", "CERTAIN",
    "ALLEGED", "BROTHER", "CHAMBER",
    "ALREADY", "BROUGHT", "CHANNEL",
    "ANALYST", "BURNING", "CHAPTER",
    "ANCIENT", "DEALING", "CHARITY",
    "ANOTHER", "DECIDED", "CHARLIE",
    "ANXIETY", "DECLINE", "CHARTER",
    "ANXIOUS", "DEFAULT", "CHECKED",
    "ANYBODY", "DEFENCE", "CHICKEN",
    "APPLIED", "DEFICIT", "CHRONIC",
    "ARRANGE", "DELIVER", "CIRCUIT",
    "ARRIVAL", "DENSITY", "CLASSES",
    "ARTICLE", "DEPOSIT", "CLASSIC",
    "ASSAULT", "DESKTOP", "CLIMATE",
    "ASSUMED", "DESPITE", "CLOSING",
    "ASSURED", "DESTROY", "CLOSURE",
    "ATTEMPT", "DEVELOP", "CLOTHES",
    "ATTRACT", "DEVOTED", "COLLECT",
    "AUCTION", "DIAMOND", "COLLEGE",
    "AVERAGE", "DIGITAL", "COMBINE",
    "EASTERN", "DISCUSS", "COMFORT",
    "ECONOMY", "DISEASE", "COMMAND",
    "EDITION", "DISPLAY", "COMMENT",
    "ELDERLY", "DISPUTE", "COMPACT",
    "ELEMENT", "DISTANT", "COMPANY",
    "ENGAGED", "DIVERSE", "COMPARE",
    "ENHANCE", "DIVIDED", "COMPETE",
    "ESSENCE", "DRAWING", "COMPLEX",
    "EVENING", "DRIVING", "CONCEPT",
    "EVIDENT", "DYNAMIC", "CONCERN",
    "EXACTLY", "FACTORY", "CONCERT",
    "EXAMINE", "FACULTY", "CONDUCT",
    "EXAMPLE", "FAILING", "CONFIRM",
    "EXCITED", "FAILURE", "CONNECT",
    "EXCLUDE", "FASHION", "CONSENT",
    "EXHIBIT", "FEATURE", "CONSIST",
    "EXPENSE", "FEDERAL", "CONTACT",
    "EXPLAIN", "FEELING", "CONTAIN",
    "EXPLORE", "FICTION", "CONTENT",
    "EXPRESS", "FIFTEEN", "CONTEST",
    "EXTREME", "FILLING", "CONTEXT",
    "GALLERY", "FINANCE", "CONTROL",
    "GATEWAY", "FINDING", "CONVERT",
    "GENERAL", "FISHING", "CORRECT",
    "GENETIC", "FITNESS", "COUNCIL",
    "GENUINE", "FOREIGN", "COUNSEL",
    "GIGABIT", "FOREVER", "COUNTER",
    "GREATER", "FORMULA", "COUNTRY",
    "HANGING", "FORTUNE", "CRUCIAL",
    "HEADING", "FORWARD", "CRYSTAL",
    "HEALTHY", "FOUNDER", "CULTURE",
    "HEARING", "FREEDOM", "CURRENT",
    "HEAVILY", "FURTHER", "CUTTING",
    "HELPFUL", "ILLEGAL", "JOINTLY",
    "HELPING", "ILLNESS", "JOURNAL",
    "HERSELF", "IMAGINE", "JOURNEY",
    "HIGHWAY", "IMAGING", "JUSTICE",
    "HIMSELF", "IMPROVE", "JUSTIFY",
    "HISTORY", "INCLUDE", "KEEPING",
    "HOLDING", "INITIAL", "KILLING",
    "HOLIDAY", "INQUIRY", "KINGDOM",
    "HOUSING", "INSIGHT", "KITCHEN",
    "HOWEVER", "INSTALL", "KNOWING",
    "HUNDRED", "INSTANT", "MACHINE",
    "HUSBAND", "INSTEAD", "MANAGER",
    "LANDING", "INTENSE", "MARRIED",
    "LARGELY", "INTERIM", "MASSIVE",
    "LASTING", "INVOLVE", "MAXIMUM",
    "LEADING", "NATURAL", "MEANING",
    "LEARNED", "NEITHER", "MEASURE",
    "LEISURE", "NERVOUS", "MEDICAL",
    "LIBERAL", "NETWORK", "MEETING",
    "LIBERTY", "NEUTRAL", "MENTION",
    "LIBRARY", "NOTABLE", "MESSAGE",
    "LICENSE", "NOTHING", "MILLION",
    "LIMITED", "NOWHERE", "MINERAL",
    "LISTING", "NUCLEAR", "MINIMAL",
    "LOGICAL", "NURSING", "MINIMUM",
    "LOYALTY", "PACIFIC", "MISSING",
    "OBVIOUS", "PACKAGE", "MISSION",
    "OFFENCE", "PAINTED", "MISTAKE",
    "OFFICER", "PARKING", "MIXTURE",
    "ONGOING", "PARTIAL", "MONITOR",
    "OPENING", "PARTNER", "MONTHLY",
    "OPERATE", "PASSAGE", "MORNING",
    "OPINION", "PASSING", "MUSICAL",
    "OPTICAL", "PASSION", "MYSTERY",
    "ORGANIC", "PASSIVE", "PORTION",
    "OUTCOME", "PATIENT", "POVERTY",
    "OUTDOOR", "PATTERN", "PRECISE",
    "OUTLOOK", "PAYABLE", "PREDICT",
    "OUTSIDE", "PAYMENT", "PREMIER",
    "OVERALL", "PENALTY", "PREMIUM",
    "PROUDLY", "PENDING", "PREPARE",
    "PROJECT", "PENSION", "PRESENT",
    "PROMISE", "PEALING", "PREVENT",
    "PROMOTE", "PERFECT", "PRIMARY",
    "PROTECT", "PERFORM", "PRINTER",
    "PROTEIN", "PERHAPS", "PRIVACY",
    "PROTEST", "PHOENIX", "PRIVATE",
    "PROVIDE", "PICKING", "PROBLEM",
}

var eightLetterWords = []string{
	"ABSOLUTE", "BACHELOR", "COMPUTER",
	"ABSTRACT", "BACTERIA", "CONCLUDE",
	"ACADEMIC", "BASEBALL", "CONCRETE",
	"ACCEPTED", "BATHROOM", "CONFLICT",
	"ACCIDENT", "BECOMING", "CONFUSED",
	"ACCURACY", "BIRTHDAY", "CONGRESS",
	"ACCURATE", "BOUNDARY", "CONSIDER",
	"ACHIEVED", "BREAKING", "CONSTANT",
	"ACQUIRED", "BREEDING", "CONSUMER",
	"ACTIVITY", "BUILDING", "CONTINUE",
	"ACTUALLY", "BULLETIN", "CONTRACT",
	"ADDITION", "BUSINESS", "CONTRARY",
	"ADEQUATE", "CALENDAR", "CONTRAST",
	"ADJACENT", "CAMPAIGN", "CONVINCE",
	"ADJUSTED", "CAPACITY", "CORRIDOR",
	"ADVANCED", "CASHMERE", "COVERAGE",
	"ADVISORY", "CASUALTY", "COVERING",
	"ADVOCATE", "CATCHING", "CREATION",
	"AFFECTED", "CATEGORY", "CREATIVE",
	"AIRCRAFT", "CATHOLIC", "CRIMINAL",
	"ALLIANCE", "CAUTIOUS", "CRITICAL",
	"ALTHOUGH", "CELLULAR", "CROSSING",
	"ALUMINUM", "CEREMONY", "CULTURAL",
	"ANALYSIS", "CHAIRMAN", "CURRENCY",
	"ANNOUNCE", "CHAMPION", "CUSTOMER",
	"ANYTHING", "CHEMICAL", "DATABASE",
	"ANYWHERE", "CHILDREN", "DAUGHTER",
	"APPARENT", "CIRCULAR", "DAYLIGHT",
	"APPENDIX", "CIVILIAN", "DEADLINE",
	"APPROACH", "CLEARING", "DECIDING",
	"APPROVAL", "CLINICAL", "DECISION",
	"ARGUMENT", "CLOTHING", "DECREASE",
	"ARTISTIC", "COLLAPSE", "DEFERRED",
	"ASSEMBLY", "COLONIAL", "DEFINITE",
	"ASSUMING", "COLORFUL", "DELICATE",
	"ATHLETIC", "COMMENCE", "DELIVERY",
	"ATTACHED", "COMMERCE", "DESCRIBE",
	"ATTITUDE", "COMPLAIN", "DESIGNER",
	"ATTORNEY", "COMPLETE", "DETAILED",
	"AUDIENCE", "COMPOSED", "DIABETES",
	"AUTONOMY", "COMPOUND", "DIALOGUE",
	"AVIATION", "COMPRISE", "DIAMETER",
	"DOUBTFUL", "EMERGING", "DIRECTLY",
	"DRAMATIC", "EMPHASIS", "DIRECTOR",
	"DRESSING", "EMPLOYEE", "DISABLED",
	"DROPPING", "ENDEAVOR", "DISASTER",
	"DURATION", "ENGAGING", "DISCLOSE",
	"DYNAMICS", "ENGINEER", "DISCOUNT",
	"EARNINGS", "ENORMOUS", "DISCOVER",
	"ECONOMIC", "ENTIRELY", "DISORDER",
	"EDUCATED", "ENTRANCE", "DISPOSAL",
	"EFFICACY", "ENVELOPE", "DISTANCE",
	"EIGHTEEN", "EQUALITY", "DISTINCT",
	"ELECTION", "EQUATION", "DISTRICT",
	"ELECTRIC", "ESTIMATE", "DIVIDEND",
	"ELIGIBLE", "EVALUATE", "DIVISION",
	"EXTERNAL", "EVENTUAL", "DOCTRINE",
	"FACILITY", "EVERYDAY", "DOCUMENT",
	"FAMILIAR", "EVERYONE", "DOMESTIC",
	"FEATURED", "EVIDENCE", "DOMINANT",
	"FEEDBACK", "EXCHANGE", "DOMINATE",
	"FESTIVAL", "EXCITING", "FLOATING",
	"FINISHED", "EXERCISE", "FOOTBALL",
	"FIREWALL", "EXPLICIT", "FOOTHILL",
	"FLAGSHIP", "EXPOSURE", "FORECAST",
	"FLEXIBLE", "EXTENDED", "FOREMOST",
	"FUNCTION", "GUIDANCE", "FORMERLY",
	"GENERATE", "HANDLING", "FOURTEEN",
	"GENEROUS", "HARDWARE", "FRACTION",
	"GENOMICS", "HERITAGE", "FRACTURE",
	"GOODWILL", "HIGHLAND", "FREQUENT",
	"GOVERNOR", "HISTORIC", "FRIENDLY",
	"GRADUATE", "HOMELESS", "FRONTIER",
	"GRAPHICS", "HOMEPAGE", "IMPERIAL",
	"GRATEFUL", "HOSPITAL", "INCIDENT",
	"GUARDIAN", "HUMANITY", "INCLUDED",
	"INITIATE", "IDENTIFY", "INCREASE",
	"INNOCENT", "IDENTITY", "INDICATE",
	"INSPIRED", "IDEOLOGY", "INDIRECT",
	"INSTANCE", "INTIMATE", "INDUSTRY",
	"INTEGRAL", "INTRANET", "INFORMAL",
	"INTENDED", "INVASION", "INFORMED",
	"INTERACT", "INVOLVED", "INHERENT",
	"INTEREST", "ISOLATED", "KEYBOARD",
	"INTERIOR", "JUDGMENT", "LANDLORD",
	"INTERNAL", "JUDICIAL", "LANGUAGE",
	"INTERVAL", "JUNCTION", "LAUGHTER",
	"MATERIAL", "MINORITY", "LEARNING",
	"MATURITY", "MOBILITY", "LEVERAGE",
	"MAXIMIZE", "MODELING", "LIFETIME",
	"MEANTIME", "MODERATE", "LIGHTING",
	"MEASURED", "MOMENTUM", "LIKEWISE",
	"MEDICINE", "MONETARY", "LIMITING",
	"MEDIEVAL", "MOREOVER", "LITERARY",
	"MEMORIAL", "MORTGAGE", "LOCATION",
	"MERCHANT", "MOUNTAIN", "MAGAZINE",
	"MIDNIGHT", "MOUNTING", "MAGNETIC",
	"MILITARY", "MOVEMENT", "MAINTAIN",
	"MINIMIZE", "MULTIPLE", "MAJORITY",
	"MINISTER", "NATIONAL", "MARGINAL",
	"MINISTRY", "NEGATIVE", "MARRIAGE",
	"ORIGINAL", "NINETEEN", "PRESERVE",
	"OVERCOME", "NORTHERN", "PRESSING",
	"OVERHEAD", "NOTEBOOK", "PRESSURE",
	"OVERSEAS", "NUMEROUS", "PREVIOUS",
	"OVERVIEW", "OBSERVER", "PRINCESS",
	"PAINTING", "OCCASION", "PRINTING",
	"PARALLEL", "OFFERING", "PRIORITY",
	"PARENTAL", "OFFICIAL", "PROBABLE",
	"PATENTED", "OFFSHORE", "PROBABLY",
	"PATIENCE", "OPERATOR", "PRODUCER",
	"PEACEFUL", "OPPONENT", "PROFOUND",
}