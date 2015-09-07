#!/usr/bin/env bash
current_version=0.0.1

if [ -n "${BASH_SOURCE[0]}" ]
then
	AYI_ROOT=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
else
	AYI_ROOT=$( cd "$( dirname "$0" )" && pwd )
fi

case $1 in
	# FIXME: Improve cooking method for mie
	mie)
		echo "Roast mie!"
		;;
	# TODO: the logic for check here
	check)
		. ${AYI_ROOT}/lib/check
		;;
	install)
		echo "You can you install, no can no bb"
		;;
	update)
		version=$(wget https://dyweb.github.io/Ayi/ver -q -O -)
		# FIXME: this is not the right logic....
		echo "Latest version is "$version
		# NOTE: the compare operator can be interperated as stream redirection
		if [[ "$current_version" < "$version" ]]; then
			echo "Let's update, but I don't know how..."
		else
			echo "You already have the latest version"
		fi
		;;
esac

