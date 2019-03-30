# Setup golang

- Download go
- Put it in /home/rockyj/Workspace/2-Personal/91-Go/
- Run - set -gx GOPATH /home/rockyj/Workspace/2-Personal/91-Go/go/
- Run - set -gx PATH $PATH /home/rockyj/Workspace/2-Personal/91-Go/go/bin/
- Run - mkdir src
- Run - mkdir src/deployer
- Run - curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
- Run - set -gx GOPATH /home/rockyj/Workspace/2-Personal/91-Go/
- Run - cd src/deployer/
- Run - dep init
- Create - vi run.go
- Run - go build
- Run - env DEPLOYER_USER="foo" DEPLOYER_PASSWORD="bar123" ./deployer
