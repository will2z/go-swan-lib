package constants

const (
	EMPTY_STRING         = ""
	DEFAULT_SELECT_LIMIT = "100"

	STORAGE_SERVER_TYPE_WEB_SERVER  = "web server"
	STORAGE_SERVER_TYPE_IPFS_SERVER = "ipfs server"

	MARKET_VERSION_1 = "1.1"
	MARKET_VERSION_2 = "1.2"

	SWAN_API_STATUS_SUCCESS = "success"
	SWAN_API_STATUS_FAIL    = "fail"

	TASK_TYPE_VERIFIED = "verified"
	TASK_TYPE_REGULAR  = "regular"

	TASK_STATUS_ASSIGNED              = "Assigned"
	TASK_STATUS_DEAL_SENT             = "DealSent"
	TASK_STATUS_PROGRESS_WITH_FAILURE = "ProgressWithFailure"

	TASK_FAST_RETRIEVAL_NO  = 0
	TASK_FAST_RETRIEVAL_YES = 1

	TASK_BID_MODE_MANUAL = 0 // allocate miner manually after creating task
	TASK_BID_MODE_AUTO   = 1 // allocate miner by market matcher after creating task
	TASK_BID_MODE_NONE   = 2 // set miner when creating task

	TASK_IS_PUBLIC = 1

	CAR_FILE_STATUS_CREATED  = "Created"
	CAR_FILE_STATUS_ASSIGNED = "Assigned"

	OFFLINE_DEAL_STATUS_ASSIGNED = "Assigned"
	OFFLINE_DEAL_STATUS_CREATED  = "Created"

	EPOCH_PER_HOUR = 120

	PATH_TYPE_NOT_EXIST = 0 //this path not exists
	PATH_TYPE_FILE      = 1 //file
	PATH_TYPE_DIR       = 2 //directory
	PATH_TYPE_UNKNOWN   = 3 //unknown path type

	AuthorizationHeaderKey = "Authorization"

	TASK_SOURCE_ID_DEFAULT      = 0
	TASK_SOURCE_ID_SWAN         = 1
	TASK_SOURCE_ID_SWAN_CLIENT  = 2
	TASK_SOURCE_ID_SWAN_FS3     = 3
	TASK_SOURCE_ID_SWAN_PAYMENT = 4
	TASK_SOURCE_ID_OTHER        = 5

	LOTUS_PRICE_MULTIPLE_1E18 = 1e18
	LOTUS_PRICE_MULTIPLE_1E15 = 1e15
	LOTUS_PRICE_MULTIPLE_1E12 = 1e12
	LOTUS_PRICE_MULTIPLE_1E9  = 1e9
	LOTUS_PRICE_MULTIPLE_1E6  = 1e6
	LOTUS_PRICE_MULTIPLE_1E3  = 1e3

	WALLET_NON_VERIFIED_MESSAGE = "Not a Verified Client"

	LOTUS_AUTH_READ  = "read"
	LOTUS_AUTH_WRITE = "write"
	LOTUS_AUTH_SIGN  = "sign"
	LOTUS_AUTH_ADMIN = "admin"

	LOTUS_TRANSFER_TYPE_MANUAL = "manual"

	MAX_AUTO_BID_COPY_NUMBER = 8

	DURATION_DEFAULT = 1512000

	DURATION_MIN = 518400
	DURATION_MAX = 1555200

	HTTP_API_TIMEOUT_SECOND = 30

	WALLET_TYPE_256 = "secp256k1"
	WALLET_TYPE_BLS = "bls"
)

var ChainMap = map[string]string{
	"1":  "ETH",   // 1     Ethereum Mainnet
	"2":  "BNB",   // 56    Binance Smart Chain Mainnet
	"3":  "AVAX",  // 43114 Avalanche C-Chain
	"4":  "MATIC", // 137   Polygon Mainnet
	"5":  "FTM",   // 250   Fantom Opera
	"6":  "xDAI",  // 100   Gnosis Chain (formerly xDai)
	"7":  "IOTX",  // 4689  IoTeX Network Mainnet
	"8":  "ONE",   // 1666600000 Harmony Mainnet Shard 0
	"9":  "BOBA",  // 288   Boba Network
	"10": "FUSE",  // 122   Fuse Mainnet
	"11": "JEWEL", // 53935  DFK Chain
	"12": "EVMOS", // 9001   Evmos
	"13": "TUS",   // 73772  Swimmer Network
}
