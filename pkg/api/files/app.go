package files

import (
	"net/http"

	"github.com/provokateurin/rain-cloud/pkg/common"
	"github.com/provokateurin/rain-cloud/pkg/registration"

	"github.com/go-chi/chi/v5"
	"golang.org/x/net/webdav"
)

//nolint:revive
type FilesRegistration struct{}

var _ registration.StrictServerInterface = (*FilesRegistration)(nil)

//nolint:gochecknoinits
func init() {
	chi.RegisterMethod("PROPFIND")
	chi.RegisterMethod("MKCOL")

	prefix := "/remote.php/dav/files/admin"
	handler := &webdav.Handler{
		Prefix:     prefix,
		FileSystem: webdav.NewMemFS(),
		LockSystem: webdav.NewMemLS(),
		Logger: func(request *http.Request, err error) {
			if err != nil {
				panic(err)
			}
		},
	}
	common.Router.Mount(prefix, handler)
	registration.HandlerFromMux(registration.NewStrictHandler(&FilesRegistration{}, nil), common.Router)
}
