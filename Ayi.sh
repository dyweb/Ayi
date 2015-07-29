#!/usr/bin/env bash
current_version=0.0.1

case $1 in
	# FIXME: Improve cooking method for mie
	mie)
		echo "Roast mie!"
		;;
	# TODO: the logic for check here
	check)
		echo "Ayi is busy, call her later"
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