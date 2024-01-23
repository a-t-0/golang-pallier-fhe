# Golang implementation of Paillier FHE scheme

This is used by myself to learn, not for production.

## Prerequisites

```sh
sudo apt install golang-go
pre-commit install

# Install errcheck
#go get -u github.com/kisielk/errcheck # Did not resolve issue
sudo snap install errcheck # Worked.

# Install staticcheck
go install honnef.co/go/tools/cmd/staticcheck@latest

# Install goimports
go install golang.org/x/tools/cmd/goimports@latest

# Install gocyclo
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

# Install gocritic
go install -v github.com/go-critic/go-critic/cmd/gocritic@latest

# Optional (if it says ... gostatic not found etc):
export GOPATH="$HOME/go"
export PATH=$PATH:$GOPATH/bin

```

## Run code

```sh
go build -o out
./out
```

## Go=Golang resources

[handbook](https://www.freecodecamp.org/news/learn-golang-handbook/)

[pre-commit](https://github.com/dnephin/pre-commit-golang)

[Template repo](https://github.com/hashicorp/consul/tree/b37fe80eeea2b8df8ecc08ccde863ba6b7eba67d)

[pre-commit example](https://github.com/instill-ai/vdp/blob/8acdbee9885b38b5932e94ec83e7948338a40f5b/.pre-commit-configyaml#L9)

## FHE resources

- Explanations:

[Wiki](https://en.wikipedia.org/wiki/Paillier_cryptosystem)

[blog-on-Pailier-cryptosystem](https://blog.openmined.org/the-paillier-cryptosystem/)

[Verification](https://www.linksight.nl/en/content/homomorphic-encryption/)

- Python:

[e-voting](https://github.com/BenLanders/Homomorphic-E-voting-System/blob/main/paillier_e_voting.py)

[fhe](https://github.com/happydracula/Pallier-Homomorphic-Encryption/blob/main/pallier.py)

- Go:

[fhe](https://github.com/niclabs/tcecdsa/blob/58c70fa14b8e2b2eed238c623cda6d475678f560/l2fhe/l2fhe.go#L12)

[fabmorph](https://github.com/wsmelojr/aulacripto/blob/921d8c1dfbba20c478da66bf636a4ad430d69ec8/hyperledger/fabmorph/fabmorph.go#L232)

[paillier](https://github.com/IPSE-TEAM/go-supernodepdp/blob/807cd751bc9e47d88a43883e932e48a7bf854ccc/paillier.go#L12)
