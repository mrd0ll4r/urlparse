package urlparse_test
import (
	"testing"
	"github.com/mrd0ll4r/urlparse"
)

const (
	query = "https://www.subdomain.tracker.com:80/announce?compact=0&downloaded=1234&info_hash=01234567890123456789&key=peerKey&left=4321&no_peer_id=1&peer_id=-TEST01-6wfG2wk6wWLc&port=6881&trackerid=trackerId"
	path = "/announce?compact=0&downloaded=1234&info_hash=01234567890123456789&key=peerKey&left=4321&no_peer_id=1&peer_id=-TEST01-6wfG2wk6wWLc&port=6881&trackerid=trackerId"
)

func compare(X, Y []string) []string {
	m := make(map[string]int)

	for _, y := range Y {
		m[y]++
	}

	var ret []string
	for _, x := range X {
		if m[x] > 0 {
			m[x]--
			continue
		}
		ret = append(ret, x)
	}

	return ret
}


func BenchmarkParseStdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlparse.ParseStdlib(query)
	}
}

func BenchmarkParseStdlibNoUrldecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlparse.ParseNoUrldecode(query)
	}
}

func BenchmarkQueryNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlparse.New(query)
	}
}

func BenchmarkQueryNewSliced(b *testing.B){
	for i := 0; i < b.N; i++ {
		urlparse.NewSliced(query)
	}
}

func BenchmarkUrlparseStdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u, _ := urlparse.ParseFull(query)
		u.Query()
	}
}

func BenchmarkPathParseStdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlparse.ParseStdlib(path)
	}
}

func BenchmarkPathParseStdlibNoUrldecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlparse.ParseNoUrldecode(path)
	}
}

func BenchmarkPathQueryNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlparse.New(path)
	}
}

func BenchmarkPathQueryNewSliced(b *testing.B){
	for i := 0; i < b.N; i++ {
		urlparse.NewSliced(path)
	}
}

func BenchmarkPathUrlparseStdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u, _ := urlparse.ParseFull(path)
		u.Query()
	}
}