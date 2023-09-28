package test

import (
	"github.com/nbio/st"
	"github.com/h2non/gock"
	"io/ioutil"
	"github.com/nycu-ucr/gonet/http"
	"testing"
)

func TestMatchURL(t *testing.T) {
	defer gock.Disable()

	gock.New("http://(.*).com").
		Reply(200).
		BodyString("foo foo")

	res, err := http.Get("http://foo.com")
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)
	body, _ := ioutil.ReadAll(res.Body)
	st.Expect(t, string(body), "foo foo")
}
