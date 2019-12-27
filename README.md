# zx

zx is a set of handy commands to make some daily tasks easier and more fun.

## Installing

### Linux

##### Prerequisites

- Check that Go is installed correctly
- curl or wget should be installed
- git should be installed

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
zx is a set of handy commands to make some daily tasks easier and more fun.
version: v0.1.0
```

### CPF

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


### CNPJ

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

### password

```
~ zx password
0n&r82i$
```

```
~ zx password 16
F@+pqBwkSH1907=2
```

### salt

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
