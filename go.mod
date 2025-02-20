module gerrit.o-ran-sc.org/r/ric-plt/appmgr

go 1.18

replace gerrit.o-ran-sc.org/r/ric-plt/sdlgo => gerrit.o-ran-sc.org/r/ric-plt/sdlgo.git v0.8.0

replace gerrit.o-ran-sc.org/r/com/golog => gerrit.o-ran-sc.org/r/com/golog.git v0.0.2

require (
	gerrit.o-ran-sc.org/r/com/golog v0.0.2
	gerrit.o-ran-sc.org/r/ric-plt/sdlgo v0.0.0-00010101000000-000000000000
	github.com/fsnotify/fsnotify v1.4.9
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/loads v0.19.4
	github.com/go-openapi/runtime v0.19.7
	github.com/go-openapi/spec v0.19.4
	github.com/go-openapi/strfmt v0.19.3
	github.com/go-openapi/swag v0.19.5
	github.com/go-openapi/validate v0.19.4
	github.com/jessevdk/go-flags v1.4.0
	github.com/orcaman/concurrent-map v0.0.0-20190314100340-2693aad1ed75
	github.com/segmentio/ksuid v1.0.2
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.6.1
	github.com/valyala/fastjson v1.4.1
	github.com/xeipuuv/gojsonschema v1.1.0
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7
)

require (
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	go.mongodb.org/mongo-driver v1.5.1 // indirect
	golang.org/x/tools v0.0.0-20190617190820-da514acc4774 // indirect
	k8s.io/utils v0.0.0-20210527160623-6fdb442a123b // indirect
)
