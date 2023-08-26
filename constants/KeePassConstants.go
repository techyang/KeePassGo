package constants

const KEEPASS_DB_EXT = ".kdbx"
const KEEPASS_DB_DEFAULT_PASSWORD = ""

// 定义一个枚举类型
type Browser string

const (
	Browser_Default          Browser = Browser("default")
	Browser_Chrome           Browser = Browser("chrome")
	Browser_Firefox          Browser = Browser("firefox")
	Browser_Safari           Browser = Browser("safari")
	Browser_Edge             Browser = Browser("edge")
	Browser_InternetExplorer Browser = Browser("iexplore")
)
