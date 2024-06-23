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



