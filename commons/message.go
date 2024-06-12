package commons

import "github.com/saddam-satria/posq-be/domains"

var ACCESS_DENIED map[domains.Lang]string = map[domains.Lang]string{
	"id": "akses ditolak",
	"en": "access denied",
}
var NOT_FOUND map[domains.Lang]string = map[domains.Lang]string{
	"id": "data tidak ditemukan",
	"en": "data not found",
}

var BAD_REQUEST map[domains.Lang]string = map[domains.Lang]string{
	"id": "body tidak boleh kosong",
	"en": "body required",
}

var SUCCESS map[domains.Lang]string = map[domains.Lang]string{
	"id": "berhasil",
	"en": "success",
}
