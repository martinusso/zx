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

### zx info

`➜ zx info`

**output**
```
zx is a set of handy commands to make some daily tasks easier and more fun.
version: v0.1.0
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
