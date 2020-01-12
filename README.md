# zx

zx is a set of handy commands to make some daily tasks easier.

[![Circle CI](https://circleci.com/gh/martinusso/zx/tree/master.svg?style=shield&circle-token=:circle-token)](https://circleci.com/gh/martinusso/zx/tree/master)
[![Build Status](https://travis-ci.org/martinusso/zx.svg?branch=master)](https://travis-ci.org/martinusso/zx)
[![Coverage Status](https://coveralls.io/repos/github/martinusso/zx/badge.svg?branch=master)](https://coveralls.io/github/martinusso/zx?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/martinusso/zx)](https://goreportcard.com/report/github.com/martinusso/zx)

## Installing

### Linux

##### Prerequisites

- Check that Go is installed correctly
- `curl` or `wget` should be installed
- `git` should be installed
- Copying to the Clipboard requires `xclip` or `xsel` command to be installed

#### Basic Installation

**zx** is installed by running one of the following commands in your terminal.

##### via curl

```
sh -c "$(curl -fsSL https://raw.githubusercontent.com/martinusso/zx/master/install.sh)"
```
##### via wget

```
sh -c "$(wget -O- https://raw.githubusercontent.com/martinusso/zx/master/install.sh)"
```

As an alternative, you can first download the install script and run it afterwards:

```
wget https://raw.githubusercontent.com/martinusso/zx/master/install.sh
sh install.sh
```

#### Using git

Alternatively, you can "git clone" this repository to any directory and run install command.

```
git clone --depth 1 https://github.com/martinusso/zx.git
cd zx
make install
```
## Commands

### info

```
~ zx info
zx is a set of handy commands to make some daily tasks easier.
version: v0.3.0
```

### base64

Decode from Base64 or Encode to Base64.

_Copies automatically the output to the clipboard._

##### Usage:

Encode
```
~ zx base64 -e "zx is a set of handy commands to make some daily tasks easier"
enggaXMgYSBzZXQgb2YgaGFuZHkgY29tbWFuZHMgdG8gbWFrZSBzb21lIGRhaWx5IHRhc2tzIGVhc2llcg==
```

Decode
```
~ zx base64 enggaXMgYSBzZXQgb2YgaGFuZHkgY29tbWFuZHMgdG8gbWFrZSBzb21lIGRhaWx5IHRhc2tzIGVhc2llcg==
zx is a set of handy commands to make some daily tasks easier
```

##### Aliases:
  `base64`, `b64`

##### Flags:
```
  -d, --decode   Decode from Base64
  -e, --encode   Encode to Base64
  -h, --help     help for base64
```

### cnpj

Generate a valid CNPJ or Validate if pass a CNPJ as args

_Copies automatically the output to the clipboard._

##### Usage:

```
~ zx cnpj
72114610000127
```

```
~ zx cnpj 72114610000127
72114610000127 ➜ valid
```

```
~ zx cnpj 72114610000121
72114610000121 ➜ invalid
```

##### Flags:
```
  -h, --help   help for cnpj
```

### cpf

Generate a valid CPF or Validate if pass a CPF as args.

_Copies automatically the output to the clipboard._

##### Usage:

```
~ zx cpf
08507460003
```

```
~ zx cpf 08507460003
08507460003 ➜ valid
```

```
~ zx cpf 08507460001
08507460001 ➜ invalid
```

##### Flags:
```
  -h, --help   help for cpf
```

### exchange

List of foreign currency rates.

Reference exchange rate: US dollar (USD).

`BRL: ? (BRL 1 = USD ?), EUR: ? (EUR 1 = USD ?), GBP: ? (GBP 1 = USD ?)`

##### Usage:
```
~ zx exchange
BRL: 4.05 (USD 0.25), EUR: 0.90 (USD 1.12), GBP: 0.76 (USD 1.31)
```

##### Flags:
```
  -h, --help   help for exchange
```

### ip

IP geolocation.
It provides a simple IP to country, state and city mapping. Using db-ip.com/api.

_Copies automatically the output to the clipboard._

##### Usage:

```
~ zx ip 8.8.8.8
IP: 8.8.8.8
Continent: (SA) South America
Country: (BR) Brazil
City: Vitória (Espírito Santo)
```

##### Flags:
```
  -h, --help   help for ip
```


### jwt

Decode a JWT token.

This command doesn't validate the token, any well formed JWT can be decoded.

##### Usage:
```
~ zx jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c 
Header:
{
  "alg": "HS256",
  "typ": "JWT"
}

Payload:
{
  "iat": 1516239022,
  "name": "John Doe",
  "sub": "1234567890"
}
```

##### Flags:
```
  -h, --help   help for jwt
```

### minify

Minify removes whitespace, strips comments, combines files, and optimizes/shortens a few common programming patterns.
Supported media types: `css`, `html`, `js`, `json`, `svg`, `xml`.

To confirm the input, press `Ctrl+] ENTER`

_Copies automatically the output to the clipboard._


##### Usage:
```
~ zx minify css
Type (or paste) here:
body {
   overflow: hidden;
   background-color: #000000;
   background-image: url(images/bg.gif);
   background-repeat: no-repeat;
}
^]
body{overflow:hidden;background-color:#000;background-image:url(images/bg.gif);background-repeat:no-repeat}
```

##### Flags:
```
  -h, --help          help for minify
  -t, --type string   Supported media types: css, html, js, json, svg, xml
```

### password

Generate a random password.

_Copies automatically the output to the clipboard._


##### Usage:
```
~ zx password
0n&r82i$
```

Specifying the size:

```
~ zx password 16
F@+pqBwkSH1907=2

~ zx password -l=16
Hk^vO4N]#tuI83le
```

zx password has the `no|N` flag which allows you to remove character types in the password generation.

Password without lowercase letters:

```
~ zx password -Nl
F@+BSH1907=2
```

Password without uppercase letters:

```
~ zx password -Nu
@+pqwk1907=2
```

Password without numbers: 

```
~ zx password -Nn
F@+pqBwkSH=
```

Password without symbols:

```
~ zx password -Ns
FpqBwkSH19072
```

And you can even combine several of them

```
~ zx password -Nsu
pqwk19072
```

##### Flags:
```
  -h, --help         help for password
  -l, --length int   Password length
  -N, --no string    Password without [l]owercase, [u]ppercase, [n]umbers, [s]ymbols
```

### salt

Generate a random and unique salt and hash for passwords.

##### Usage:
```
~ zx salt
Password: S[H3/7l1
Hash: efb283bcde5fe42c177558a563c0b998f179ba312777aa36218465ffe35ca3ba115f7e51cff736ca7a783b7d27be2ec072b80a2f6c817139ab2b32429a4fcd03
Salt: 6e034879d4c3e97def89875d23b7a173a438b1611269168d5fcb548cb7676b7b703e82887508bc1224732ef04373b45154760982bd90d7e40a4a5057dafd7c05
```

```
~ zx salt MyPassword       
Password: MyPassword
Hash: c60bf872599be8327d40d2c29bf5abfec1c78da6870f9c6c83d8a1bb7069dd4fdfee535b5a24d658a1ff085830385a9242ff7f5a51ad9e60542e4b661f582783
Salt: 4c27d2a1ffc2fda6cb678e917679d39d7cfcaba1d9d207a0d8509c1c02946155071cdd2515aec0fe266e05d84ce1b2d13dd85ad8db0764b55e8e4024e90f2771
```

##### Flags:
```
  -h, --help              help for salt
  -p, --password string   Password
```

### uuid

Generate a random UUID.

_Copies automatically the output to the clipboard._

##### Usage:
```
~ zx uuid
d9ed16f4-f315-44e4-8e57-b00516d73420
```

##### Flags:
```
  -h, --help              help for uuid
```

## Contributing

- Fork it
- Download your fork to your PC (git clone https://github.com/your_username/zx && cd zx)
- Create your feature branch (git checkout -b my-new-feature)
- Make changes and add them (git add .)
- Commit your changes (git commit -m 'Add some feature')
- Push to the branch (git push origin my-new-feature)
- Create new pull request

## License

**zx** is released under the MIT license. See [LICENSE](https://github.com/martinusso/zx/blob/master/LICENSE)
