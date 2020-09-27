package consts

const (
	// 当前支持平台
	PLATFORM_SAMH = "samh"
	PLATFORM_MKZ  = "mkz"
	PLATFORM_KMH  = "kmh"
)

const (
	// 运行环境
	RUN_MODE_DEV  = "dev"
	RUN_MODE_PROD = "prod"
)

const (
	SQLCON    = "#"
	SQLSEP    = ","
	SEP       = "|"
	BLANK     = " "
	EMPTY_STR = ""
	ZERO      = 0
	ONE       = 1
)

const (
	DEFAULT_UID = 0
	ALL_UID     = 0
	ALL_GID     = 0
)

const (
	DEFAULT_NEWER = 0
	NEWER_NEW_ID  = 1
	NEWER_OLD_ID  = 2
)

const (
	DEFAULT_GENDER   = 0
	GENDER_MALE_ID   = 1
	GENDER_FEMALE_ID = 2
)

const (
	TIMEFORMAT = "2006-01-02"

	ONE_DAY   = 1
	ONE_WEEK  = 7
	ONE_MONTH = 30
)

const (
	INT_MAX       = int(^uint(0) >> 1)
	INT_MIN       = ^INT_MAX
	UINT_MIN uint = 0
	UINT_MAX      = ^uint(0)
)
