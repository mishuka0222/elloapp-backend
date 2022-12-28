// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package mtproto

import (
	"strings"
)

var (
	mimeToExtension map[string]string
	extensionToMime map[string]string
)

func init() {
	mimeToExtension = make(map[string]string)
	extensionToMime = make(map[string]string)

	mimeToExtension["application/andrew-inset"] = "ez"
	extensionToMime["ez"] = "application/andrew-inset"

	mimeToExtension["application/dsptype"] = "tsp"
	extensionToMime["tsp"] = "application/dsptype"

	mimeToExtension["application/futuresplash"] = "spl"
	extensionToMime["spl"] = "application/futuresplash"

	mimeToExtension["application/hta"] = "hta"
	extensionToMime["hta"] = "application/hta"

	mimeToExtension["application/mac-binhex40"] = "hqx"
	extensionToMime["hqx"] = "application/mac-binhex40"

	mimeToExtension["application/mac-compactpro"] = "cpt"
	extensionToMime["cpt"] = "application/mac-compactpro"

	mimeToExtension["application/mathematica"] = "nb"
	extensionToMime["nb"] = "application/mathematica"

	mimeToExtension["application/msaccess"] = "mdb"
	extensionToMime["mdb"] = "application/msaccess"

	mimeToExtension["application/oda"] = "oda"
	extensionToMime["oda"] = "application/oda"

	mimeToExtension["application/ogg"] = "ogg"
	extensionToMime["ogg"] = "application/ogg"

	mimeToExtension["application/pdf"] = "pdf"
	extensionToMime["pdf"] = "application/pdf"

	mimeToExtension["application/com.adobe.pdf"] = "pdf"

	mimeToExtension["application/pgp-keys"] = "key"
	extensionToMime["key"] = "application/pgp-keys"

	mimeToExtension["application/pgp-signature"] = "pgp"
	extensionToMime["pgp"] = "application/pgp-signature"

	mimeToExtension["application/pics-rules"] = "prf"
	extensionToMime["prf"] = "application/pics-rules"

	mimeToExtension["application/rar"] = "rar"
	extensionToMime["rar"] = "application/rar"

	mimeToExtension["application/rdf+xml"] = "rdf"
	extensionToMime["rdf"] = "application/rdf+xml"

	mimeToExtension["application/rss+xml"] = "rss"
	extensionToMime["rss"] = "application/rss+xml"

	mimeToExtension["application/zip"] = "zip"
	extensionToMime["zip"] = "application/zip"

	mimeToExtension["application/vnd.android.package-archive"] = "apk"
	extensionToMime["apk"] = "application/vnd.android.package-archive"

	mimeToExtension["application/vnd.cinderella"] = "cdy"
	extensionToMime["cdy"] = "application/vnd.cinderella"

	mimeToExtension["application/vnd.ms-pki.stl"] = "stl"
	extensionToMime["stl"] = "application/vnd.ms-pki.stl"

	mimeToExtension["application/vnd.oasis.opendocument.database"] = "odb"
	extensionToMime["odb"] = "application/vnd.oasis.opendocument.database"

	mimeToExtension["application/vnd.oasis.opendocument.formula"] = "odf"
	extensionToMime["odf"] = "application/vnd.oasis.opendocument.formula"

	mimeToExtension["application/vnd.oasis.opendocument.graphics"] = "odg"
	extensionToMime["odg"] = "application/vnd.oasis.opendocument.graphics"

	mimeToExtension["application/vnd.oasis.opendocument.graphics-template"] = "otg"
	extensionToMime["otg"] = "application/vnd.oasis.opendocument.graphics-template"

	mimeToExtension["application/vnd.oasis.opendocument.image"] = "odi"
	extensionToMime["odi"] = "application/vnd.oasis.opendocument.image"

	mimeToExtension["application/vnd.oasis.opendocument.spreadsheet"] = "ods"
	extensionToMime["ods"] = "application/vnd.oasis.opendocument.spreadsheet"

	mimeToExtension["application/vnd.oasis.opendocument.spreadsheet-template"] = "ots"
	extensionToMime["ots"] = "application/vnd.oasis.opendocument.spreadsheet-template"

	mimeToExtension["application/vnd.oasis.opendocument.text"] = "odt"
	extensionToMime["odt"] = "application/vnd.oasis.opendocument.text"

	mimeToExtension["application/vnd.oasis.opendocument.text-master"] = "odm"
	extensionToMime["odm"] = "application/vnd.oasis.opendocument.text-master"

	mimeToExtension["application/vnd.oasis.opendocument.text-template"] = "ott"
	extensionToMime["ott"] = "application/vnd.oasis.opendocument.text-template"

	mimeToExtension["application/vnd.oasis.opendocument.text-web"] = "oth"
	extensionToMime["oth"] = "application/vnd.oasis.opendocument.text-web"

	mimeToExtension["application/msword"] = "doc"
	extensionToMime["doc"] = "application/msword"

	mimeToExtension["application/msword"] = "dot"
	extensionToMime["dot"] = "application/msword"

	mimeToExtension["application/vnd.openxmlformats-officedocument.wordprocessingml.document"] = "docx"
	extensionToMime["docx"] = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"

	mimeToExtension["application/vnd.openxmlformats-officedocument.wordprocessingml.template"] = "dotx"
	extensionToMime["dotx"] = "application/vnd.openxmlformats-officedocument.wordprocessingml.template"

	mimeToExtension["application/vnd.ms-excel"] = "xls"
	extensionToMime["xls"] = "application/vnd.ms-excel"

	mimeToExtension["application/vnd.ms-excel"] = "xlt"
	extensionToMime["xlt"] = "application/vnd.ms-excel"

	mimeToExtension["application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"] = "xlsx"
	extensionToMime["xlsx"] = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

	mimeToExtension["application/vnd.openxmlformats-officedocument.spreadsheetml.template"] = "xltx"
	extensionToMime["xltx"] = "application/vnd.openxmlformats-officedocument.spreadsheetml.template"

	mimeToExtension["application/vnd.ms-powerpoint"] = "ppt"
	extensionToMime["ppt"] = "application/vnd.ms-powerpoint"

	mimeToExtension["application/vnd.ms-powerpoint"] = "pot"
	extensionToMime["pot"] = "application/vnd.ms-powerpoint"

	mimeToExtension["application/vnd.ms-powerpoint"] = "pps"
	extensionToMime["pps"] = "application/vnd.ms-powerpoint"

	mimeToExtension["application/vnd.openxmlformats-officedocument.presentationml.presentation"] = "pptx"
	extensionToMime["pptx"] = "application/vnd.openxmlformats-officedocument.presentationml.presentation"

	mimeToExtension["application/vnd.openxmlformats-officedocument.presentationml.template"] = "potx"
	extensionToMime["potx"] = "application/vnd.openxmlformats-officedocument.presentationml.template"

	mimeToExtension["application/vnd.openxmlformats-officedocument.presentationml.slideshow"] = "ppsx"
	extensionToMime["ppsx"] = "application/vnd.openxmlformats-officedocument.presentationml.slideshow"

	mimeToExtension["application/vnd.rim.cod"] = "cod"
	extensionToMime["cod"] = "application/vnd.rim.cod"

	mimeToExtension["application/vnd.smaf"] = "mmf"
	extensionToMime["mmf"] = "application/vnd.smaf"

	mimeToExtension["application/vnd.stardivision.calc"] = "sdc"
	extensionToMime["sdc"] = "application/vnd.stardivision.calc"

	mimeToExtension["application/vnd.stardivision.draw"] = "sda"
	extensionToMime["sda"] = "application/vnd.stardivision.draw"

	mimeToExtension["application/vnd.stardivision.impress"] = "sdd"
	extensionToMime["sdd"] = "application/vnd.stardivision.impress"

	mimeToExtension["application/vnd.stardivision.impress"] = "sdp"
	extensionToMime["sdp"] = "application/vnd.stardivision.impress"

	mimeToExtension["application/vnd.stardivision.math"] = "smf"
	extensionToMime["smf"] = "application/vnd.stardivision.math"

	mimeToExtension["application/vnd.stardivision.writer"] = "sdw"
	extensionToMime["sdw"] = "application/vnd.stardivision.writer"

	mimeToExtension["application/vnd.stardivision.writer"] = "vor"
	extensionToMime["vor"] = "application/vnd.stardivision.writer"

	mimeToExtension["application/vnd.stardivision.writer-global"] = "sgl"
	extensionToMime["sgl"] = "application/vnd.stardivision.writer-global"

	mimeToExtension["application/vnd.sun.xml.calc"] = "sxc"
	extensionToMime["sxc"] = "application/vnd.sun.xml.calc"

	mimeToExtension["application/vnd.sun.xml.calc.template"] = "stc"
	extensionToMime["stc"] = "application/vnd.sun.xml.calc.template"

	mimeToExtension["application/vnd.sun.xml.draw"] = "sxd"
	extensionToMime["sxd"] = "application/vnd.sun.xml.draw"

	mimeToExtension["application/vnd.sun.xml.draw.template"] = "std"
	extensionToMime["std"] = "application/vnd.sun.xml.draw.template"

	mimeToExtension["application/vnd.sun.xml.impress"] = "sxi"
	extensionToMime["sxi"] = "application/vnd.sun.xml.impress"

	mimeToExtension["application/vnd.sun.xml.impress.template"] = "sti"
	extensionToMime["sti"] = "application/vnd.sun.xml.impress.template"

	mimeToExtension["application/vnd.sun.xml.math"] = "sxm"
	extensionToMime["sxm"] = "application/vnd.sun.xml.math"

	mimeToExtension["application/vnd.sun.xml.writer"] = "sxw"
	extensionToMime["sxw"] = "application/vnd.sun.xml.writer"

	mimeToExtension["application/vnd.sun.xml.writer.global"] = "sxg"
	extensionToMime["sxg"] = "application/vnd.sun.xml.writer.global"

	mimeToExtension["application/vnd.sun.xml.writer.template"] = "stw"
	extensionToMime["stw"] = "application/vnd.sun.xml.writer.template"

	mimeToExtension["application/vnd.visio"] = "vsd"
	extensionToMime["vsd"] = "application/vnd.visio"

	mimeToExtension["application/x-abiword"] = "abw"
	extensionToMime["abw"] = "application/x-abiword"

	mimeToExtension["application/x-apple-diskimage"] = "dmg"
	extensionToMime["dmg"] = "application/x-apple-diskimage"

	mimeToExtension["application/x-bcpio"] = "bcpio"
	extensionToMime["bcpio"] = "application/x-bcpio"

	mimeToExtension["application/x-bittorrent"] = "torrent"
	extensionToMime["torrent"] = "application/x-bittorrent"

	mimeToExtension["application/x-cdf"] = "cdf"
	extensionToMime["cdf"] = "application/x-cdf"

	mimeToExtension["application/x-cdlink"] = "vcd"
	extensionToMime["vcd"] = "application/x-cdlink"

	mimeToExtension["application/x-chess-pgn"] = "pgn"
	extensionToMime["pgn"] = "application/x-chess-pgn"

	mimeToExtension["application/x-cpio"] = "cpio"
	extensionToMime["cpio"] = "application/x-cpio"

	mimeToExtension["application/x-debian-package"] = "deb"
	extensionToMime["deb"] = "application/x-debian-package"

	mimeToExtension["application/x-debian-package"] = "udeb"
	extensionToMime["udeb"] = "application/x-debian-package"

	mimeToExtension["application/x-director"] = "dcr"
	extensionToMime["dcr"] = "application/x-director"

	mimeToExtension["application/x-director"] = "dir"
	extensionToMime["dir"] = "application/x-director"

	mimeToExtension["application/x-director"] = "dxr"
	extensionToMime["dxr"] = "application/x-director"

	mimeToExtension["application/x-dms"] = "dms"
	extensionToMime["dms"] = "application/x-dms"

	mimeToExtension["application/x-doom"] = "wad"
	extensionToMime["wad"] = "application/x-doom"

	mimeToExtension["application/x-dvi"] = "dvi"
	extensionToMime["dvi"] = "application/x-dvi"

	mimeToExtension["application/x-flac"] = "flac"
	extensionToMime["flac"] = "application/x-flac"

	mimeToExtension["application/x-font"] = "pfa"
	extensionToMime["pfa"] = "application/x-font"

	mimeToExtension["application/x-font"] = "pfb"
	extensionToMime["pfb"] = "application/x-font"

	mimeToExtension["application/x-font"] = "gsf"
	extensionToMime["gsf"] = "application/x-font"

	mimeToExtension["application/x-font"] = "pcf"
	extensionToMime["pcf"] = "application/x-font"

	mimeToExtension["application/x-font"] = "pcf.Z"
	extensionToMime["pcf.Z"] = "application/x-font"

	mimeToExtension["application/x-freemind"] = "mm"
	extensionToMime["mm"] = "application/x-freemind"

	mimeToExtension["application/x-futuresplash"] = "spl"
	extensionToMime["spl"] = "application/x-futuresplash"

	mimeToExtension["application/x-gnumeric"] = "gnumeric"
	extensionToMime["gnumeric"] = "application/x-gnumeric"

	mimeToExtension["application/x-go-sgf"] = "sgf"
	extensionToMime["sgf"] = "application/x-go-sgf"

	mimeToExtension["application/x-graphing-calculator"] = "gcf"
	extensionToMime["gcf"] = "application/x-graphing-calculator"

	mimeToExtension["application/x-gtar"] = "gtar"
	extensionToMime["gtar"] = "application/x-gtar"

	mimeToExtension["application/x-gtar"] = "tgz"
	extensionToMime["tgz"] = "application/x-gtar"

	mimeToExtension["application/x-gtar"] = "taz"
	extensionToMime["taz"] = "application/x-gtar"

	mimeToExtension["application/x-hdf"] = "hdf"
	extensionToMime["hdf"] = "application/x-hdf"

	mimeToExtension["application/x-ica"] = "ica"
	extensionToMime["ica"] = "application/x-ica"

	mimeToExtension["application/x-internet-signup"] = "ins"
	extensionToMime["ins"] = "application/x-internet-signup"

	mimeToExtension["application/x-internet-signup"] = "isp"
	extensionToMime["isp"] = "application/x-internet-signup"

	mimeToExtension["application/x-iphone"] = "iii"
	extensionToMime["iii"] = "application/x-iphone"

	mimeToExtension["application/x-iso9660-image"] = "iso"
	extensionToMime["iso"] = "application/x-iso9660-image"

	mimeToExtension["application/x-jmol"] = "jmz"
	extensionToMime["jmz"] = "application/x-jmol"

	mimeToExtension["application/x-kchart"] = "chrt"
	extensionToMime["chrt"] = "application/x-kchart"

	mimeToExtension["application/x-killustrator"] = "kil"
	extensionToMime["kil"] = "application/x-killustrator"

	mimeToExtension["application/x-koan"] = "skp"
	extensionToMime["skp"] = "application/x-koan"

	mimeToExtension["application/x-koan"] = "skd"
	extensionToMime["skd"] = "application/x-koan"

	mimeToExtension["application/x-koan"] = "skt"
	extensionToMime["skt"] = "application/x-koan"

	mimeToExtension["application/x-koan"] = "skm"
	extensionToMime["skm"] = "application/x-koan"

	mimeToExtension["application/x-kpresenter"] = "kpr"
	extensionToMime["kpr"] = "application/x-kpresenter"

	mimeToExtension["application/x-kpresenter"] = "kpt"
	extensionToMime["kpt"] = "application/x-kpresenter"

	mimeToExtension["application/x-kspread"] = "ksp"
	extensionToMime["ksp"] = "application/x-kspread"

	mimeToExtension["application/x-kword"] = "kwd"
	extensionToMime["kwd"] = "application/x-kword"

	mimeToExtension["application/x-kword"] = "kwt"
	extensionToMime["kwt"] = "application/x-kword"

	mimeToExtension["application/x-latex"] = "latex"
	extensionToMime["latex"] = "application/x-latex"

	mimeToExtension["application/x-lha"] = "lha"
	extensionToMime["lha"] = "application/x-lha"

	mimeToExtension["application/x-lzh"] = "lzh"
	extensionToMime["lzh"] = "application/x-lzh"

	mimeToExtension["application/x-lzx"] = "lzx"
	extensionToMime["lzx"] = "application/x-lzx"

	mimeToExtension["application/x-maker"] = "frm"
	extensionToMime["frm"] = "application/x-maker"

	mimeToExtension["application/x-maker"] = "maker"
	extensionToMime["maker"] = "application/x-maker"

	mimeToExtension["application/x-maker"] = "frame"
	extensionToMime["frame"] = "application/x-maker"

	mimeToExtension["application/x-maker"] = "fb"
	extensionToMime["fb"] = "application/x-maker"

	mimeToExtension["application/x-maker"] = "book"
	extensionToMime["book"] = "application/x-maker"

	mimeToExtension["application/x-maker"] = "fbdoc"
	extensionToMime["fbdoc"] = "application/x-maker"

	mimeToExtension["application/x-mif"] = "mif"
	extensionToMime["mif"] = "application/x-mif"

	mimeToExtension["application/x-ms-wmd"] = "wmd"
	extensionToMime["wmd"] = "application/x-ms-wmd"

	mimeToExtension["application/x-ms-wmz"] = "wmz"
	extensionToMime["wmz"] = "application/x-ms-wmz"

	mimeToExtension["application/x-msi"] = "msi"
	extensionToMime["msi"] = "application/x-msi"

	mimeToExtension["application/x-ns-proxy-autoconfig"] = "pac"
	extensionToMime["pac"] = "application/x-ns-proxy-autoconfig"

	mimeToExtension["application/x-nwc"] = "nwc"
	extensionToMime["nwc"] = "application/x-nwc"

	mimeToExtension["application/x-object"] = "o"
	extensionToMime["o"] = "application/x-object"

	mimeToExtension["application/x-oz-application"] = "oza"
	extensionToMime["oza"] = "application/x-oz-application"

	mimeToExtension["application/x-pkcs12"] = "p12"
	extensionToMime["p12"] = "application/x-pkcs12"

	mimeToExtension["application/x-pkcs7-certreqresp"] = "p7r"
	extensionToMime["p7r"] = "application/x-pkcs7-certreqresp"

	mimeToExtension["application/x-pkcs7-crl"] = "crl"
	extensionToMime["crl"] = "application/x-pkcs7-crl"

	mimeToExtension["application/x-quicktimeplayer"] = "qtl"
	extensionToMime["qtl"] = "application/x-quicktimeplayer"

	mimeToExtension["application/x-shar"] = "shar"
	extensionToMime["shar"] = "application/x-shar"

	mimeToExtension["application/x-shockwave-flash"] = "swf"
	extensionToMime["swf"] = "application/x-shockwave-flash"

	mimeToExtension["application/x-stuffit"] = "sit"
	extensionToMime["sit"] = "application/x-stuffit"

	mimeToExtension["application/x-sv4cpio"] = "sv4cpio"
	extensionToMime["sv4cpio"] = "application/x-sv4cpio"

	mimeToExtension["application/x-sv4crc"] = "sv4crc"
	extensionToMime["sv4crc"] = "application/x-sv4crc"

	mimeToExtension["application/x-tar"] = "tar"
	extensionToMime["tar"] = "application/x-tar"

	mimeToExtension["application/x-texinfo"] = "texinfo"
	extensionToMime["texinfo"] = "application/x-texinfo"

	mimeToExtension["application/x-texinfo"] = "texi"
	extensionToMime["texi"] = "application/x-texinfo"

	mimeToExtension["application/x-troff"] = "t"
	extensionToMime["t"] = "application/x-troff"

	mimeToExtension["application/x-troff"] = "roff"
	extensionToMime["roff"] = "application/x-troff"

	mimeToExtension["application/x-troff-man"] = "man"
	extensionToMime["man"] = "application/x-troff-man"

	mimeToExtension["application/x-ustar"] = "ustar"
	extensionToMime["ustar"] = "application/x-ustar"

	mimeToExtension["application/x-wais-source"] = "src"
	extensionToMime["src"] = "application/x-wais-source"

	mimeToExtension["application/x-wingz"] = "wz"
	extensionToMime["wz"] = "application/x-wingz"

	mimeToExtension["application/x-webarchive"] = "webarchive"
	extensionToMime["webarchive"] = "application/x-webarchive"

	mimeToExtension["application/x-x509-ca-cert"] = "crt"
	extensionToMime["crt"] = "application/x-x509-ca-cert"

	mimeToExtension["application/x-x509-user-cert"] = "crt"
	extensionToMime["crt"] = "application/x-x509-user-cert"

	mimeToExtension["application/x-xcf"] = "xcf"
	extensionToMime["xcf"] = "application/x-xcf"

	mimeToExtension["application/x-xfig"] = "fig"
	extensionToMime["fig"] = "application/x-xfig"

	mimeToExtension["application/xhtml+xml"] = "xhtml"
	extensionToMime["xhtml"] = "application/xhtml+xml"

	mimeToExtension["audio/3gpp"] = "3gpp"
	extensionToMime["3gpp"] = "audio/3gpp"

	mimeToExtension["audio/basic"] = "snd"
	extensionToMime["snd"] = "audio/basic"

	mimeToExtension["audio/midi"] = "mid"
	extensionToMime["mid"] = "audio/midi"

	mimeToExtension["audio/midi"] = "midi"
	extensionToMime["midi"] = "audio/midi"

	mimeToExtension["audio/midi"] = "kar"
	extensionToMime["kar"] = "audio/midi"

	mimeToExtension["audio/mpeg"] = "mpga"
	extensionToMime["mpga"] = "audio/mpeg"

	mimeToExtension["audio/mpeg"] = "mpega"
	extensionToMime["mpega"] = "audio/mpeg"

	mimeToExtension["audio/mpeg"] = "mp2"
	extensionToMime["mp2"] = "audio/mpeg"

	mimeToExtension["audio/mpeg"] = "mp3"
	extensionToMime["mp3"] = "audio/mpeg"

	mimeToExtension["audio/mpeg"] = "m4a"
	extensionToMime["m4a"] = "audio/mpeg"

	mimeToExtension["audio/mpegurl"] = "m3u"
	extensionToMime["m3u"] = "audio/mpegurl"

	mimeToExtension["audio/prs.sid"] = "sid"
	extensionToMime["sid"] = "audio/prs.sid"

	mimeToExtension["audio/x-aiff"] = "aif"
	extensionToMime["aif"] = "audio/x-aiff"

	mimeToExtension["audio/x-aiff"] = "aiff"
	extensionToMime["aiff"] = "audio/x-aiff"

	mimeToExtension["audio/x-aiff"] = "aifc"
	extensionToMime["aifc"] = "audio/x-aiff"

	mimeToExtension["audio/x-gsm"] = "gsm"
	extensionToMime["gsm"] = "audio/x-gsm"

	mimeToExtension["audio/x-mpegurl"] = "m3u"
	extensionToMime["m3u"] = "audio/x-mpegurl"

	mimeToExtension["audio/x-ms-wma"] = "wma"
	extensionToMime["wma"] = "audio/x-ms-wma"

	mimeToExtension["audio/x-ms-wax"] = "wax"
	extensionToMime["wax"] = "audio/x-ms-wax"

	mimeToExtension["audio/x-pn-realaudio"] = "ra"
	extensionToMime["ra"] = "audio/x-pn-realaudio"

	mimeToExtension["audio/x-pn-realaudio"] = "rm"
	extensionToMime["rm"] = "audio/x-pn-realaudio"

	mimeToExtension["audio/x-pn-realaudio"] = "ram"
	extensionToMime["ram"] = "audio/x-pn-realaudio"

	mimeToExtension["audio/x-realaudio"] = "ra"
	extensionToMime["ra"] = "audio/x-realaudio"

	mimeToExtension["audio/x-scpls"] = "pls"
	extensionToMime["pls"] = "audio/x-scpls"

	mimeToExtension["audio/x-sd2"] = "sd2"
	extensionToMime["sd2"] = "audio/x-sd2"

	mimeToExtension["audio/x-wav"] = "wav"
	extensionToMime["wav"] = "audio/x-wav"

	mimeToExtension["image/bmp"] = "bmp"
	extensionToMime["bmp"] = "image/bmp"

	mimeToExtension["image/gif"] = "gif"
	extensionToMime["gif"] = "image/gif"

	mimeToExtension["image/ico"] = "cur"
	extensionToMime["cur"] = "image/ico"

	mimeToExtension["image/ico"] = "ico"
	extensionToMime["ico"] = "image/ico"

	mimeToExtension["image/ief"] = "ief"
	extensionToMime["ief"] = "image/ief"

	mimeToExtension["image/jpeg"] = "jpeg"
	extensionToMime["jpeg"] = "image/jpeg"

	mimeToExtension["image/jpeg"] = "jpg"
	extensionToMime["jpg"] = "image/jpeg"

	mimeToExtension["image/jpeg"] = "jpe"
	extensionToMime["jpe"] = "image/jpeg"

	mimeToExtension["image/pcx"] = "pcx"
	extensionToMime["pcx"] = "image/pcx"

	mimeToExtension["image/png"] = "png"
	extensionToMime["png"] = "image/png"

	mimeToExtension["image/svg+xml"] = "svg"
	extensionToMime["svg"] = "image/svg+xml"

	mimeToExtension["image/svg+xml"] = "svgz"
	extensionToMime["svgz"] = "image/svg+xml"

	mimeToExtension["image/tiff"] = "tiff"
	extensionToMime["tiff"] = "image/tiff"

	mimeToExtension["image/tiff"] = "tif"
	extensionToMime["tif"] = "image/tiff"

	mimeToExtension["image/vnd.djvu"] = "djvu"
	extensionToMime["djvu"] = "image/vnd.djvu"

	mimeToExtension["image/vnd.djvu"] = "djv"
	extensionToMime["djv"] = "image/vnd.djvu"

	mimeToExtension["image/vnd.wap.wbmp"] = "wbmp"
	extensionToMime["wbmp"] = "image/vnd.wap.wbmp"

	mimeToExtension["image/x-cmu-raster"] = "ras"
	extensionToMime["ras"] = "image/x-cmu-raster"

	mimeToExtension["image/x-coreldraw"] = "cdr"
	extensionToMime["cdr"] = "image/x-coreldraw"

	mimeToExtension["image/x-coreldrawpattern"] = "pat"
	extensionToMime["pat"] = "image/x-coreldrawpattern"

	mimeToExtension["image/x-coreldrawtemplate"] = "cdt"
	extensionToMime["cdt"] = "image/x-coreldrawtemplate"

	mimeToExtension["image/x-corelphotopaint"] = "cpt"
	extensionToMime["cpt"] = "image/x-corelphotopaint"

	mimeToExtension["image/x-icon"] = "ico"
	extensionToMime["ico"] = "image/x-icon"

	mimeToExtension["image/x-jg"] = "art"
	extensionToMime["art"] = "image/x-jg"

	mimeToExtension["image/x-jng"] = "jng"
	extensionToMime["jng"] = "image/x-jng"

	mimeToExtension["image/x-ms-bmp"] = "bmp"
	extensionToMime["bmp"] = "image/x-ms-bmp"

	mimeToExtension["image/x-photoshop"] = "psd"
	extensionToMime["psd"] = "image/x-photoshop"

	mimeToExtension["image/x-portable-anymap"] = "pnm"
	extensionToMime["pnm"] = "image/x-portable-anymap"

	mimeToExtension["image/x-portable-bitmap"] = "pbm"
	extensionToMime["pbm"] = "image/x-portable-bitmap"

	mimeToExtension["image/x-portable-graymap"] = "pgm"
	extensionToMime["pgm"] = "image/x-portable-graymap"

	mimeToExtension["image/x-portable-pixmap"] = "ppm"
	extensionToMime["ppm"] = "image/x-portable-pixmap"

	mimeToExtension["image/x-rgb"] = "rgb"
	extensionToMime["rgb"] = "image/x-rgb"

	mimeToExtension["image/x-xbitmap"] = "xbm"
	extensionToMime["xbm"] = "image/x-xbitmap"

	mimeToExtension["image/x-xpixmap"] = "xpm"
	extensionToMime["xpm"] = "image/x-xpixmap"

	mimeToExtension["image/x-xwindowdump"] = "xwd"
	extensionToMime["xwd"] = "image/x-xwindowdump"

	mimeToExtension["model/iges"] = "igs"
	extensionToMime["igs"] = "model/iges"

	mimeToExtension["model/iges"] = "iges"
	extensionToMime["iges"] = "model/iges"

	mimeToExtension["model/mesh"] = "msh"
	extensionToMime["msh"] = "model/mesh"

	mimeToExtension["model/mesh"] = "mesh"
	extensionToMime["mesh"] = "model/mesh"

	mimeToExtension["model/mesh"] = "silo"
	extensionToMime["silo"] = "model/mesh"

	mimeToExtension["text/calendar"] = "ics"
	extensionToMime["ics"] = "text/calendar"

	mimeToExtension["text/calendar"] = "icz"
	extensionToMime["icz"] = "text/calendar"

	mimeToExtension["text/comma-separated-values"] = "csv"
	extensionToMime["csv"] = "text/comma-separated-values"

	mimeToExtension["text/css"] = "css"
	extensionToMime["css"] = "text/css"

	mimeToExtension["text/html"] = "htm"
	extensionToMime["htm"] = "text/html"

	mimeToExtension["text/html"] = "html"
	extensionToMime["html"] = "text/html"

	mimeToExtension["text/h323"] = "323"
	extensionToMime["323"] = "text/h323"

	mimeToExtension["text/iuls"] = "uls"
	extensionToMime["uls"] = "text/iuls"

	mimeToExtension["text/mathml"] = "mml"
	extensionToMime["mml"] = "text/mathml"

	mimeToExtension["text/plain"] = "txt"
	extensionToMime["txt"] = "text/plain"

	mimeToExtension["text/plain"] = "asc"
	extensionToMime["asc"] = "text/plain"

	mimeToExtension["text/plain"] = "text"
	extensionToMime["text"] = "text/plain"

	mimeToExtension["text/plain"] = "diff"
	extensionToMime["diff"] = "text/plain"

	mimeToExtension["text/plain"] = "po"
	extensionToMime["po"] = "text/plain"
	// reserve "pot" for vnd.ms-powerpoint

	mimeToExtension["text/richtext"] = "rtx"
	extensionToMime["rtx"] = "text/richtext"

	mimeToExtension["text/rtf"] = "rtf"
	extensionToMime["rtf"] = "text/rtf"

	mimeToExtension["text/texmacs"] = "ts"
	extensionToMime["ts"] = "text/texmacs"

	mimeToExtension["text/text"] = "phps"
	extensionToMime["phps"] = "text/text"

	mimeToExtension["text/tab-separated-values"] = "tsv"
	extensionToMime["tsv"] = "text/tab-separated-values"

	mimeToExtension["text/xml"] = "xml"
	extensionToMime["xml"] = "text/xml"

	mimeToExtension["text/x-bibtex"] = "bib"
	extensionToMime["bib"] = "text/x-bibtex"

	mimeToExtension["text/x-boo"] = "boo"
	extensionToMime["boo"] = "text/x-boo"

	mimeToExtension["text/x-c++hdr"] = "h++"
	extensionToMime["h++"] = "text/x-c++hdr"

	mimeToExtension["text/x-c++hdr"] = "hpp"
	extensionToMime["hpp"] = "text/x-c++hdr"

	mimeToExtension["text/x-c++hdr"] = "hxx"
	extensionToMime["hxx"] = "text/x-c++hdr"

	mimeToExtension["text/x-c++hdr"] = "hh"
	extensionToMime["hh"] = "text/x-c++hdr"

	mimeToExtension["text/x-c++src"] = "c++"
	extensionToMime["c++"] = "text/x-c++src"

	mimeToExtension["text/x-c++src"] = "cpp"
	extensionToMime["cpp"] = "text/x-c++src"

	mimeToExtension["text/x-c++src"] = "cxx"
	extensionToMime["cxx"] = "text/x-c++src"

	mimeToExtension["text/x-chdr"] = "h"
	extensionToMime["h"] = "text/x-chdr"

	mimeToExtension["text/x-component"] = "htc"
	extensionToMime["htc"] = "text/x-component"

	mimeToExtension["text/x-csh"] = "csh"
	extensionToMime["csh"] = "text/x-csh"

	mimeToExtension["text/x-csrc"] = "c"
	extensionToMime["c"] = "text/x-csrc"

	mimeToExtension["text/x-dsrc"] = "d"
	extensionToMime["d"] = "text/x-dsrc"

	mimeToExtension["text/x-haskell"] = "hs"
	extensionToMime["hs"] = "text/x-haskell"

	mimeToExtension["text/x-java"] = "java"
	extensionToMime["java"] = "text/x-java"

	mimeToExtension["text/x-literate-haskell"] = "lhs"
	extensionToMime["lhs"] = "text/x-literate-haskell"

	mimeToExtension["text/x-moc"] = "moc"
	extensionToMime["moc"] = "text/x-moc"

	mimeToExtension["text/x-pascal"] = "p"
	extensionToMime["p"] = "text/x-pascal"

	mimeToExtension["text/x-pascal"] = "pas"
	extensionToMime["pas"] = "text/x-pascal"

	mimeToExtension["text/x-pcs-gcd"] = "gcd"
	extensionToMime["gcd"] = "text/x-pcs-gcd"

	mimeToExtension["text/x-setext"] = "etx"
	extensionToMime["etx"] = "text/x-setext"

	mimeToExtension["text/x-tcl"] = "tcl"
	extensionToMime["tcl"] = "text/x-tcl"

	mimeToExtension["text/x-tex"] = "tex"
	extensionToMime["tex"] = "text/x-tex"

	mimeToExtension["text/x-tex"] = "ltx"
	extensionToMime["ltx"] = "text/x-tex"

	mimeToExtension["text/x-tex"] = "sty"
	extensionToMime["sty"] = "text/x-tex"

	mimeToExtension["text/x-tex"] = "cls"
	extensionToMime["cls"] = "text/x-tex"

	mimeToExtension["text/x-vcalendar"] = "vcs"
	extensionToMime["vcs"] = "text/x-vcalendar"

	mimeToExtension["text/x-vcard"] = "vcf"
	extensionToMime["vcf"] = "text/x-vcard"

	mimeToExtension["video/3gpp"] = "3gpp"
	extensionToMime["3gpp"] = "video/3gpp"

	mimeToExtension["video/3gpp"] = "3gp"
	extensionToMime["3gp"] = "video/3gpp"

	mimeToExtension["video/3gpp"] = "3g2"
	extensionToMime["3g2"] = "video/3gpp"

	mimeToExtension["video/dl"] = "dl"
	extensionToMime["dl"] = "video/dl"

	mimeToExtension["video/dv"] = "dif"
	extensionToMime["dif"] = "video/dv"

	mimeToExtension["video/dv"] = "dv"
	extensionToMime["dv"] = "video/dv"

	mimeToExtension["video/fli"] = "fli"
	extensionToMime["fli"] = "video/fli"

	mimeToExtension["video/m4v"] = "m4v"
	extensionToMime["m4v"] = "video/m4v"

	mimeToExtension["video/mpeg"] = "mpeg"
	extensionToMime["mpeg"] = "video/mpeg"

	mimeToExtension["video/mpeg"] = "mpg"
	extensionToMime["mpg"] = "video/mpeg"

	mimeToExtension["video/mpeg"] = "mpe"
	extensionToMime["mpe"] = "video/mpeg"

	mimeToExtension["video/mp4"] = "mp4"
	extensionToMime["mp4"] = "video/mp4"

	mimeToExtension["video/mpeg"] = "VOB"
	extensionToMime["VOB"] = "video/mpeg"

	mimeToExtension["video/quicktime"] = "qt"
	extensionToMime["qt"] = "video/quicktime"

	mimeToExtension["video/quicktime"] = "mov"
	extensionToMime["mov"] = "video/quicktime"

	mimeToExtension["video/vnd.mpegurl"] = "mxu"
	extensionToMime["mxu"] = "video/vnd.mpegurl"

	mimeToExtension["video/x-la-asf"] = "lsf"
	extensionToMime["lsf"] = "video/x-la-asf"

	mimeToExtension["video/x-la-asf"] = "lsx"
	extensionToMime["lsx"] = "video/x-la-asf"

	mimeToExtension["video/x-mng"] = "mng"
	extensionToMime["mng"] = "video/x-mng"

	mimeToExtension["video/x-ms-asf"] = "asf"
	extensionToMime["asf"] = "video/x-ms-asf"

	mimeToExtension["video/x-ms-asf"] = "asx"
	extensionToMime["asx"] = "video/x-ms-asf"

	mimeToExtension["video/x-ms-wm"] = "wm"
	extensionToMime["wm"] = "video/x-ms-wm"

	mimeToExtension["video/x-ms-wmv"] = "wmv"
	extensionToMime["wmv"] = "video/x-ms-wmv"

	mimeToExtension["video/x-ms-wmx"] = "wmx"
	extensionToMime["wmx"] = "video/x-ms-wmx"

	mimeToExtension["video/x-ms-wvx"] = "wvx"
	extensionToMime["wvx"] = "video/x-ms-wvx"

	mimeToExtension["video/x-msvideo"] = "avi"
	extensionToMime["avi"] = "video/x-msvideo"

	mimeToExtension["video/x-sgi-movie"] = "movie"
	extensionToMime["movie"] = "video/x-sgi-movie"

	mimeToExtension["x-conference/x-cooltalk"] = "ice"
	extensionToMime["ice"] = "x-conference/x-cooltalk"

	mimeToExtension["x-epoc/x-sisx-app"] = "sisx"
	extensionToMime["sisx"] = "x-epoc/x-sisx-app"

	mimeToExtension["application/epub+zip"] = "epub"
	extensionToMime["epub"] = "application/epub+zip"

	mimeToExtension["text/swift"] = "swift"
	extensionToMime["swift"] = "text/swift"
}

func GuessMimeTypeByFileExtension(ext string) string {
	if strings.HasPrefix(ext, ".") {
		ext = ext[1:]
	}
	if mimeType, ok := extensionToMime[ext]; ok {
		return mimeType
	}

	return "application/binary"
}

// IsMimeAcceptedForPhotoVideoAlbum
/*
bool IsMimeAcceptedForPhotoVideoAlbum(const QString &mime) {
	return (mime == u"image/jpeg"_q)
		|| (mime == u"image/png"_q)
		|| (mime == u"video/mp4"_q)
		|| (mime == u"video/quicktime"_q);
}
*/
func IsMimeAcceptedForPhotoVideoAlbum(mime string) bool {
	return mime == "image/jpeg" ||
		mime == "image/png" ||
		mime == "video/mp4" ||
		mime == "video/quicktime"
}

// FileIsImage
/*
bool FileIsImage(const QString &name, const QString &mime) {
	QString lowermime = mime.toLower(), namelower = name.toLower();
	if (lowermime.startsWith(qstr("image/"))) {
		return true;
	} else if (namelower.endsWith(qstr(".bmp"))
		|| namelower.endsWith(qstr(".jpg"))
		|| namelower.endsWith(qstr(".jpeg"))
		|| namelower.endsWith(qstr(".gif"))
		|| namelower.endsWith(qstr(".webp"))
		|| namelower.endsWith(qstr(".tga"))
		|| namelower.endsWith(qstr(".tiff"))
		|| namelower.endsWith(qstr(".tif"))
		|| namelower.endsWith(qstr(".psd"))
		|| namelower.endsWith(qstr(".png"))) {
		return true;
	}
	return false;
}
*/
func FileIsImage(name, mime string) bool {
	lowermime := strings.ToLower(mime)
	namelower := strings.ToLower(name)
	if strings.HasPrefix(lowermime, "image/") {
		return true
	} else if strings.HasSuffix(namelower, ".bmp") ||
		strings.HasSuffix(namelower, ".jpg") ||
		strings.HasSuffix(namelower, ".jpeg") ||
		strings.HasSuffix(namelower, ".gif") ||
		strings.HasSuffix(namelower, ".webp") ||
		strings.HasSuffix(namelower, ".tga") ||
		strings.HasSuffix(namelower, ".tiff") ||
		strings.HasSuffix(namelower, ".tif") ||
		strings.HasSuffix(namelower, ".psd") ||
		strings.HasSuffix(namelower, ".png") {
		return true
	}
	return false
}
