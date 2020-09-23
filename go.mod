module luban

go 1.15

require (
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/drone/drone v1.9.0
	github.com/drone/drone-yaml v1.2.4-0.20200326192514-6f4d6dfb39e4
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/cors v1.1.1
	github.com/go-chi/docgen v1.0.5
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.4.0
	github.com/zc2638/drone-control v0.0.0-20200923072241-3198b58992c9
	github.com/zc2638/gotool v0.0.0-20200528080342-200e82def869
	gopkg.in/yaml.v2 v2.2.8
	gorm.io/driver/mysql v1.0.1 // indirect
	gorm.io/driver/sqlite v1.1.1
	gorm.io/gorm v1.20.0
)

replace github.com/h2non/gock => gopkg.in/h2non/gock.v1 v1.0.15
