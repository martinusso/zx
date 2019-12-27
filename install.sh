#!/bin/sh
#
# This script should be run via curl:
#   sh -c "$(curl -fsSL https://raw.githubusercontent.com/martinusso/zx/master/install.sh)"
# or wget:
#   sh -c "$(wget -qO- https://raw.githubusercontent.com/martinusso/zx/master/install.sh)"

REPO=${REPO:-martinusso/zx}
REMOTE=${REMOTE:-https://github.com/${REPO}.git}
BRANCH=${BRANCH:-master}

command_exists() {
	command -v "$@" >/dev/null 2>&1
}

error() {
	echo ${RED}"Error: $@"${RESET} >&2
}

setup_color() {
	# Only use colors if connected to a terminal
	if [ -t 1 ]; then
		RED=$(printf '\033[31m')
		GREEN=$(printf '\033[32m')
		YELLOW=$(printf '\033[33m')
		BLUE=$(printf '\033[34m')
		BOLD=$(printf '\033[1m')
		RESET=$(printf '\033[m')
	else
		RED=""
		GREEN=""
		YELLOW=""
		BLUE=""
		BOLD=""
		RESET=""
	fi
}

setup_zx() {
  umask g-w,o-w

	echo "${BLUE}Cloning zx${RESET}"

	command_exists git || {
		error "git is not installed"
		exit 1
	}

	if [ "$OSTYPE" = cygwin ] && git --version | grep -q msysgit; then
		error "Windows/MSYS Git is not supported on Cygwin"
		error "Make sure the Cygwin git package is installed and is first on the \$PATH"
		exit 1
	fi

	git clone -c core.eol=lf -c core.autocrlf=false \
		--depth=1 --branch "$BRANCH" "$REMOTE" || {
		error "git clone of zx repo failed"
		exit 1
	}
  
  cd zx
  make install

	echo
}

main() {
	setup_color

	if command_exists zx; then
		cat <<-EOF
			${YELLOW}You already have zx installed.${RESET}
			You'll need to remove zx if you want to reinstall.
		EOF
		exit 1
	fi

	setup_zx

	printf "$GREEN"
	cat <<-'EOF'
     ___________  ___
     \___   /\  \/  /
      /    /  >    < 
     /_____ \/__/\_ \
           \/      \/   ... is now installed!

	EOF
	printf "$RESET"

	exec zx info
}

main "$@"
