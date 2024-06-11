package apis

type AuthHeader struct {
	Authorization string `reqHeader:"Authorization"`
}