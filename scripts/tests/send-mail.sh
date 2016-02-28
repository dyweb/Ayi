#!/bin/sh

# The script will test sending mail to the specific email address
# Usage: ./send-mail.sh <email address>

set -o errexit
set -o nounset
set -o pipefail

if [[ $# -eq 0 ]] ; then
    echo 'Usage: ./send-mail.sh <email address>'
    exit 1
fi

ROOT=$(dirname "${BASH_SOURCE}")/../..

cd ${ROOT}
./Ayi m s $1 "Greet from the bash" ./fixture/test.html
cd - > /dev/null
