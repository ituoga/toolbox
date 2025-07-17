package hotreload

import (
	"net/http"
	"sync"

	"github.com/starfederation/datastar-go/datastar"
)

var syncOnce sync.Once

func Handler(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	syncOnce.Do(func() {
		sse.ExecuteScript("window.location.reload(true)") // Force hard refresh
	})
	<-r.Context().Done()
}

const HTML = `<div data-on-load="@get('/hotreload', {retryMaxCount: 1000,retryInterval:20, retryMaxWaitMs:200})"></div>`
