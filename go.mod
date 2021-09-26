module cj.com/gobot

go 1.17

require (
	github.com/256dpi/lungo v0.2.10
	github.com/cjkao/Rocket.Chat.Go.SDK v1.0.0
	github.com/docker/go-connections v0.4.0
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.7.1
// gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
// gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
)

require (
	github.com/256dpi/btree v0.0.0-20200517182607-63d76dfb3721 // indirect
	github.com/apex/log v1.9.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/gopackage/ddp v0.0.0-20170117053602-652027933df4
	github.com/klauspost/compress v1.9.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/sony/sonyflake v1.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require github.com/Jeffail/gabs/v2 v2.6.1 // indirect

replace github.com/cjkao/Rocket.Chat.Go.SDK => ../Rocket.Chat.Go.SDK-Rui

replace github.com/cjkao/Rocket.Chat.Go.SDK-Rui/realtime => ../Rocket.Chat.Go.SDK-Rui/realtime

replace github.com/cjkao/Rocket.Chat.Go.SDK-Rui/models => ../Rocket.Chat.Go.SDK-Rui/models

replace github.com/gopackage/ddp => ../ddp
