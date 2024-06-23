## Using Make to Run Commands at a Glance
```# apt-get -y install make```

create Makefile

```# touch Makefile```

file structure:

```
command_name:
	command to run like cd ~

.PHONY: command_name

```

usage:

```# make command_name```

## Powerfull tool to compile sql codes to go ORM

```
# snap install sqlc
# sqlc version
# sqlc init
```

sqlc.yaml config refrence:

https://docs.sqlc.dev/en/stable/reference/config.html


After creating migration files and query files, run this command to generate compiled Go files!

```
# sqlc generate
```


## Installing Golang Viper for defining environment variables

https://github.com/spf13/viper

```
# go get github.com/spf13/viper
```


## Install Gomock for mocking database in test mode

https://github.com/uber-go/mock

```
# go install go.uber.org/mock/mockgen@latest
# mockgen -version
```
if mockgen not defiend:

find mockgen path in your system:

```
# find / -name mockgen
"/home/docker/go/bin/mockgen"

# cp /home/docker/go/bin/mockgen /usr/local/bin/
```

use mock in db/sqlc/store.go add this to import

```
_ "github.com/golang/mock/mockgen/model"
```

then:
```
# go get go.uber.org/mock/mockgen/model

# mockgen -destination db/mock/store.go github.com/farshan-dev/gin/db/sqlc Store

```
read ```# mockgen -help``` for more info! i use Reflect mode in above command.
