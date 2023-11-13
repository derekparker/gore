// This file is part of GoRE.
//
// Copyright (C) 2019-2021 GoRE Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2023-11-13 19:12:33.785736 +0000 UTC

package gore

var stdPkgs = map[string]struct{}{
	"archive": {},
	"archive/tar": {},
	"archive/tar/testdata": {},
	"archive/zip": {},
	"archive/zip/testdata": {},
	"arena": {},
	"bufio": {},
	"builtin": {},
	"bytes": {},
	"cmp": {},
	"compress": {},
	"compress/bzip2": {},
	"compress/bzip2/testdata": {},
	"compress/flate": {},
	"compress/flate/testdata": {},
	"compress/gzip": {},
	"compress/gzip/testdata": {},
	"compress/lzw": {},
	"compress/testdata": {},
	"compress/zlib": {},
	"container": {},
	"container/heap": {},
	"container/list": {},
	"container/ring": {},
	"context": {},
	"crypto": {},
	"crypto/aes": {},
	"crypto/boring": {},
	"crypto/cipher": {},
	"crypto/des": {},
	"crypto/dsa": {},
	"crypto/ecdh": {},
	"crypto/ecdsa": {},
	"crypto/ecdsa/testdata": {},
	"crypto/ed25519": {},
	"crypto/ed25519/testdata": {},
	"crypto/elliptic": {},
	"crypto/hmac": {},
	"crypto/internal": {},
	"crypto/internal/alias": {},
	"crypto/internal/bigmod": {},
	"crypto/internal/bigmod/_asm": {},
	"crypto/internal/boring": {},
	"crypto/internal/boring/bbig": {},
	"crypto/internal/boring/bcache": {},
	"crypto/internal/boring/fipstls": {},
	"crypto/internal/boring/sig": {},
	"crypto/internal/boring/syso": {},
	"crypto/internal/edwards25519": {},
	"crypto/internal/edwards25519/field": {},
	"crypto/internal/edwards25519/field/_asm": {},
	"crypto/internal/nistec": {},
	"crypto/internal/nistec/fiat": {},
	"crypto/internal/randutil": {},
	"crypto/md5": {},
	"crypto/rand": {},
	"crypto/rc4": {},
	"crypto/rsa": {},
	"crypto/rsa/testdata": {},
	"crypto/sha1": {},
	"crypto/sha256": {},
	"crypto/sha512": {},
	"crypto/subtle": {},
	"crypto/tls": {},
	"crypto/tls/fipsonly": {},
	"crypto/tls/testdata": {},
	"crypto/x509": {},
	"crypto/x509/internal": {},
	"crypto/x509/internal/macos": {},
	"crypto/x509/pkix": {},
	"crypto/x509/testdata": {},
	"database": {},
	"database/sql": {},
	"database/sql/driver": {},
	"debug": {},
	"debug/buildinfo": {},
	"debug/dwarf": {},
	"debug/dwarf/testdata": {},
	"debug/elf": {},
	"debug/elf/testdata": {},
	"debug/gosym": {},
	"debug/gosym/testdata": {},
	"debug/macho": {},
	"debug/macho/testdata": {},
	"debug/pe": {},
	"debug/pe/testdata": {},
	"debug/plan9obj": {},
	"debug/plan9obj/testdata": {},
	"embed": {},
	"embed/internal": {},
	"embed/internal/embedtest": {},
	"embed/internal/embedtest/testdata": {},
	"embed/internal/embedtest/testdata/-not-hidden": {},
	"embed/internal/embedtest/testdata/.hidden": {},
	"embed/internal/embedtest/testdata/.hidden/.more": {},
	"embed/internal/embedtest/testdata/.hidden/_more": {},
	"embed/internal/embedtest/testdata/.hidden/more": {},
	"embed/internal/embedtest/testdata/_hidden": {},
	"embed/internal/embedtest/testdata/i": {},
	"embed/internal/embedtest/testdata/i/j": {},
	"embed/internal/embedtest/testdata/i/j/k": {},
	"encoding": {},
	"encoding/ascii85": {},
	"encoding/asn1": {},
	"encoding/base32": {},
	"encoding/base64": {},
	"encoding/binary": {},
	"encoding/csv": {},
	"encoding/gob": {},
	"encoding/hex": {},
	"encoding/json": {},
	"encoding/json/testdata": {},
	"encoding/pem": {},
	"encoding/xml": {},
	"errors": {},
	"expvar": {},
	"flag": {},
	"fmt": {},
	"go": {},
	"go/ast": {},
	"go/build": {},
	"go/build/constraint": {},
	"go/build/testdata": {},
	"go/build/testdata/alltags": {},
	"go/build/testdata/bads": {},
	"go/build/testdata/cgo_disabled": {},
	"go/build/testdata/directives": {},
	"go/build/testdata/doc": {},
	"go/build/testdata/empty": {},
	"go/build/testdata/multi": {},
	"go/build/testdata/non_source_tags": {},
	"go/build/testdata/other": {},
	"go/build/testdata/other/file": {},
	"go/build/testdata/withvendor": {},
	"go/build/testdata/withvendor/src": {},
	"go/build/testdata/withvendor/src/a": {},
	"go/build/testdata/withvendor/src/a/b": {},
	"go/build/testdata/withvendor/src/a/vendor": {},
	"go/build/testdata/withvendor/src/a/vendor/c": {},
	"go/build/testdata/withvendor/src/a/vendor/c/d": {},
	"go/constant": {},
	"go/doc": {},
	"go/doc/comment": {},
	"go/doc/comment/testdata": {},
	"go/doc/testdata": {},
	"go/doc/testdata/examples": {},
	"go/doc/testdata/pkgdoc": {},
	"go/format": {},
	"go/importer": {},
	"go/internal": {},
	"go/internal/gccgoimporter": {},
	"go/internal/gccgoimporter/testdata": {},
	"go/internal/gcimporter": {},
	"go/internal/gcimporter/testdata": {},
	"go/internal/gcimporter/testdata/versions": {},
	"go/internal/srcimporter": {},
	"go/internal/srcimporter/testdata": {},
	"go/internal/srcimporter/testdata/issue20855": {},
	"go/internal/srcimporter/testdata/issue23092": {},
	"go/internal/srcimporter/testdata/issue24392": {},
	"go/internal/typeparams": {},
	"go/parser": {},
	"go/parser/testdata": {},
	"go/parser/testdata/goversion": {},
	"go/parser/testdata/issue42951": {},
	"go/parser/testdata/issue42951/not_a_file.go": {},
	"go/parser/testdata/resolution": {},
	"go/printer": {},
	"go/printer/testdata": {},
	"go/scanner": {},
	"go/token": {},
	"go/types": {},
	"go/types/testdata": {},
	"go/types/testdata/local": {},
	"go/version": {},
	"hash": {},
	"hash/adler32": {},
	"hash/crc32": {},
	"hash/crc64": {},
	"hash/fnv": {},
	"hash/maphash": {},
	"html": {},
	"html/template": {},
	"html/template/testdata": {},
	"image": {},
	"image/color": {},
	"image/color/palette": {},
	"image/draw": {},
	"image/gif": {},
	"image/internal": {},
	"image/internal/imageutil": {},
	"image/jpeg": {},
	"image/png": {},
	"image/png/testdata": {},
	"image/png/testdata/pngsuite": {},
	"image/testdata": {},
	"index": {},
	"index/suffixarray": {},
	"internal": {},
	"internal/abi": {},
	"internal/abi/testdata": {},
	"internal/bisect": {},
	"internal/buildcfg": {},
	"internal/bytealg": {},
	"internal/cfg": {},
	"internal/coverage": {},
	"internal/coverage/calloc": {},
	"internal/coverage/cformat": {},
	"internal/coverage/cmerge": {},
	"internal/coverage/decodecounter": {},
	"internal/coverage/decodemeta": {},
	"internal/coverage/encodecounter": {},
	"internal/coverage/encodemeta": {},
	"internal/coverage/pods": {},
	"internal/coverage/rtcov": {},
	"internal/coverage/slicereader": {},
	"internal/coverage/slicewriter": {},
	"internal/coverage/stringtab": {},
	"internal/coverage/test": {},
	"internal/coverage/uleb128": {},
	"internal/cpu": {},
	"internal/dag": {},
	"internal/diff": {},
	"internal/diff/testdata": {},
	"internal/fmtsort": {},
	"internal/fuzz": {},
	"internal/goarch": {},
	"internal/godebug": {},
	"internal/godebugs": {},
	"internal/goexperiment": {},
	"internal/goos": {},
	"internal/goroot": {},
	"internal/gover": {},
	"internal/goversion": {},
	"internal/intern": {},
	"internal/itoa": {},
	"internal/lazyregexp": {},
	"internal/lazytemplate": {},
	"internal/nettrace": {},
	"internal/obscuretestdata": {},
	"internal/oserror": {},
	"internal/pkgbits": {},
	"internal/platform": {},
	"internal/poll": {},
	"internal/profile": {},
	"internal/race": {},
	"internal/reflectlite": {},
	"internal/safefilepath": {},
	"internal/saferio": {},
	"internal/singleflight": {},
	"internal/syscall": {},
	"internal/syscall/execenv": {},
	"internal/syscall/unix": {},
	"internal/syscall/windows": {},
	"internal/syscall/windows/registry": {},
	"internal/syscall/windows/sysdll": {},
	"internal/sysinfo": {},
	"internal/testenv": {},
	"internal/testlog": {},
	"internal/testpty": {},
	"internal/trace": {},
	"internal/trace/testdata": {},
	"internal/trace/v2": {},
	"internal/trace/v2/event": {},
	"internal/trace/v2/event/go122": {},
	"internal/trace/v2/internal": {},
	"internal/trace/v2/internal/testgen": {},
	"internal/trace/v2/internal/testgen/go122": {},
	"internal/trace/v2/raw": {},
	"internal/trace/v2/testdata": {},
	"internal/trace/v2/testdata/cmd": {},
	"internal/trace/v2/testdata/cmd/gotraceraw": {},
	"internal/trace/v2/testdata/cmd/gotracevalidate": {},
	"internal/trace/v2/testdata/generators": {},
	"internal/trace/v2/testdata/testprog": {},
	"internal/trace/v2/testdata/tests": {},
	"internal/trace/v2/testtrace": {},
	"internal/trace/v2/version": {},
	"internal/txtar": {},
	"internal/types": {},
	"internal/types/errors": {},
	"internal/types/testdata": {},
	"internal/types/testdata/check": {},
	"internal/types/testdata/check/decls2": {},
	"internal/types/testdata/check/importdecl0": {},
	"internal/types/testdata/check/importdecl1": {},
	"internal/types/testdata/check/issue25008": {},
	"internal/types/testdata/examples": {},
	"internal/types/testdata/fixedbugs": {},
	"internal/types/testdata/spec": {},
	"internal/unsafeheader": {},
	"internal/xcoff": {},
	"internal/xcoff/testdata": {},
	"internal/zstd": {},
	"internal/zstd/testdata": {},
	"io": {},
	"io/fs": {},
	"io/ioutil": {},
	"io/ioutil/testdata": {},
	"log": {},
	"log/internal": {},
	"log/slog": {},
	"log/slog/internal": {},
	"log/slog/internal/benchmarks": {},
	"log/slog/internal/buffer": {},
	"log/slog/internal/slogtest": {},
	"log/syslog": {},
	"maps": {},
	"math": {},
	"math/big": {},
	"math/bits": {},
	"math/cmplx": {},
	"math/rand": {},
	"math/rand/v2": {},
	"mime": {},
	"mime/multipart": {},
	"mime/multipart/testdata": {},
	"mime/quotedprintable": {},
	"mime/testdata": {},
	"net": {},
	"net/http": {},
	"net/http/cgi": {},
	"net/http/cookiejar": {},
	"net/http/fcgi": {},
	"net/http/httptest": {},
	"net/http/httptrace": {},
	"net/http/httputil": {},
	"net/http/internal": {},
	"net/http/internal/ascii": {},
	"net/http/internal/testcert": {},
	"net/http/pprof": {},
	"net/http/testdata": {},
	"net/internal": {},
	"net/internal/socktest": {},
	"net/mail": {},
	"net/netip": {},
	"net/rpc": {},
	"net/rpc/jsonrpc": {},
	"net/smtp": {},
	"net/testdata": {},
	"net/textproto": {},
	"net/url": {},
	"os": {},
	"os/exec": {},
	"os/exec/internal": {},
	"os/exec/internal/fdtest": {},
	"os/signal": {},
	"os/testdata": {},
	"os/testdata/dirfs": {},
	"os/testdata/dirfs/dir": {},
	"os/testdata/issue37161": {},
	"os/user": {},
	"path": {},
	"path/filepath": {},
	"plugin": {},
	"reflect": {},
	"reflect/internal": {},
	"reflect/internal/example1": {},
	"reflect/internal/example2": {},
	"regexp": {},
	"regexp/syntax": {},
	"regexp/testdata": {},
	"runtime": {},
	"runtime/asan": {},
	"runtime/cgo": {},
	"runtime/coverage": {},
	"runtime/coverage/testdata": {},
	"runtime/coverage/testdata/issue56006": {},
	"runtime/coverage/testdata/issue59563": {},
	"runtime/debug": {},
	"runtime/debug/testdata": {},
	"runtime/debug/testdata/fuzz": {},
	"runtime/debug/testdata/fuzz/FuzzParseBuildInfoRoundTrip": {},
	"runtime/internal": {},
	"runtime/internal/atomic": {},
	"runtime/internal/math": {},
	"runtime/internal/startlinetest": {},
	"runtime/internal/sys": {},
	"runtime/internal/syscall": {},
	"runtime/internal/wasitest": {},
	"runtime/internal/wasitest/testdata": {},
	"runtime/metrics": {},
	"runtime/msan": {},
	"runtime/pprof": {},
	"runtime/pprof/testdata": {},
	"runtime/pprof/testdata/mappingtest": {},
	"runtime/race": {},
	"runtime/race/internal": {},
	"runtime/race/internal/amd64v1": {},
	"runtime/race/internal/amd64v3": {},
	"runtime/race/testdata": {},
	"runtime/testdata": {},
	"runtime/testdata/testexithooks": {},
	"runtime/testdata/testfaketime": {},
	"runtime/testdata/testfds": {},
	"runtime/testdata/testprog": {},
	"runtime/testdata/testprogcgo": {},
	"runtime/testdata/testprogcgo/windows": {},
	"runtime/testdata/testprognet": {},
	"runtime/testdata/testsuid": {},
	"runtime/testdata/testwinlib": {},
	"runtime/testdata/testwinlibsignal": {},
	"runtime/testdata/testwinlibthrow": {},
	"runtime/testdata/testwinsignal": {},
	"runtime/testdata/testwintls": {},
	"runtime/trace": {},
	"slices": {},
	"sort": {},
	"strconv": {},
	"strconv/testdata": {},
	"strings": {},
	"sync": {},
	"sync/atomic": {},
	"syscall": {},
	"syscall/js": {},
	"testdata": {},
	"testing": {},
	"testing/fstest": {},
	"testing/internal": {},
	"testing/internal/testdeps": {},
	"testing/iotest": {},
	"testing/quick": {},
	"testing/slogtest": {},
	"text": {},
	"text/scanner": {},
	"text/tabwriter": {},
	"text/template": {},
	"text/template/parse": {},
	"text/template/testdata": {},
	"time": {},
	"time/testdata": {},
	"time/tzdata": {},
	"unicode": {},
	"unicode/utf16": {},
	"unicode/utf8": {},
	"unsafe": {},
	"vendor": {},
	"vendor/golang.org": {},
	"vendor/golang.org/x": {},
	"vendor/golang.org/x/crypto": {},
	"vendor/golang.org/x/crypto/chacha20": {},
	"vendor/golang.org/x/crypto/chacha20poly1305": {},
	"vendor/golang.org/x/crypto/cryptobyte": {},
	"vendor/golang.org/x/crypto/cryptobyte/asn1": {},
	"vendor/golang.org/x/crypto/hkdf": {},
	"vendor/golang.org/x/crypto/internal": {},
	"vendor/golang.org/x/crypto/internal/alias": {},
	"vendor/golang.org/x/crypto/internal/poly1305": {},
	"vendor/golang.org/x/net": {},
	"vendor/golang.org/x/net/dns": {},
	"vendor/golang.org/x/net/dns/dnsmessage": {},
	"vendor/golang.org/x/net/http": {},
	"vendor/golang.org/x/net/http/httpguts": {},
	"vendor/golang.org/x/net/http/httpproxy": {},
	"vendor/golang.org/x/net/http2": {},
	"vendor/golang.org/x/net/http2/hpack": {},
	"vendor/golang.org/x/net/idna": {},
	"vendor/golang.org/x/net/lif": {},
	"vendor/golang.org/x/net/nettest": {},
	"vendor/golang.org/x/net/route": {},
	"vendor/golang.org/x/sys": {},
	"vendor/golang.org/x/sys/cpu": {},
	"vendor/golang.org/x/text": {},
	"vendor/golang.org/x/text/secure": {},
	"vendor/golang.org/x/text/secure/bidirule": {},
	"vendor/golang.org/x/text/transform": {},
	"vendor/golang.org/x/text/unicode": {},
	"vendor/golang.org/x/text/unicode/bidi": {},
	"vendor/golang.org/x/text/unicode/norm": {},
}
