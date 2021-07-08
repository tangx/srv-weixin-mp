package confwxmp

import (
	"bytes"
	"sort"
	"strings"

	"github.com/tangx/srv-weixin-mp/pkg/utils/sha"
)

type Server struct {
	OpenID string `env:""`
	Token  string `env:""`
}

func (s *Server) sign(ts string, nonce string) string {
	_parts := []string{s.Token, ts, nonce}
	sort.Strings(_parts)

	buf := bytes.NewBuffer(nil)
	for _, str := range _parts {
		buf.WriteString(str)
	}

	sig := sha.Sha1(buf.String())

	return strings.ToLower(sig)
}

func (s *Server) IsSignMatch(ts string, nonce string, sign string) bool {
	sig := s.sign(ts, nonce)
	return sig == sign
}
