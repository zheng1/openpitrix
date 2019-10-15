module openpitrix.io/openpitrix

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Masterminds/semver v1.5.0
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/aws/aws-sdk-go v1.25.12
	github.com/bitly/go-simplejson v0.5.0
	github.com/chai2010/jsonmap v1.0.0
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/cyphar/filepath-securejoin v0.2.2 // indirect
	github.com/disintegration/imaging v1.6.1
	github.com/fatih/camelcase v1.0.0
	github.com/fatih/structs v1.1.0
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/runtime v0.19.6
	github.com/go-openapi/spec v0.19.3
	github.com/go-openapi/strfmt v0.19.3
	github.com/go-openapi/swag v0.19.5
	github.com/go-openapi/validate v0.19.3
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocraft/dbr v0.0.0-20190714181702-8114670a83bd
	github.com/golang/protobuf v1.3.2
	github.com/google/gops v0.3.6
	github.com/gorilla/websocket v1.4.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/koding/multiconfig v0.0.0-20171124222453-69c27309b2d7
	github.com/kubernetes/helm v2.14.3+incompatible
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/robfig/cron v1.2.0
	github.com/sony/sonyflake v1.0.0
	github.com/speps/go-hashids v2.0.0+incompatible
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.4.0
	github.com/ugorji/go/codec v1.1.7 // indirect
	github.com/urfave/cli v1.22.1
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.1.0
	go.etcd.io/etcd v3.3.17+incompatible
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	golang.org/x/net v0.0.0-20191014212845-da9a3fd4c582
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/tools v0.0.0-20191014205221-18e3458ac98b
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03
	google.golang.org/grpc v1.24.0
	gopkg.in/square/go-jose.v2 v2.3.1
	gopkg.in/yaml.v2 v2.2.4
	k8s.io/apimachinery v0.0.0-20191014065749-fb3eea214746 // indirect
	k8s.io/client-go v11.0.0+incompatible // indirect
	k8s.io/helm v2.14.3+incompatible // indirect
	kubesphere.io/im v0.1.0
	openpitrix.io/iam v0.1.0
	openpitrix.io/notification v0.2.2
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace k8s.io/kubernetes => k8s.io/kubernetes v1.14.7

replace k8s.io/api => k8s.io/api v0.0.0-20190816222004-e3a6b8045b0b

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190918224502-6154570c2037

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190816221834-a9f1d8a9c101

replace k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190918223255-26459790ef01 // indirect

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190918224932-e56234cc6b3d // indirect

replace k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190918222721-c0e3722d5cf0+incompatible

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190918225840-7f3416179ad8 // indirect
