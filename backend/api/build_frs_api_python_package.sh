#!/bin/bash -e
# The scripts creates the Feast Registry Server (FRS) API python package.
# Requirements: jq and Java
# To install the prerequisites run the following:
#
# # Debian / Ubuntu:
# sudo apt-get install --no-install-recommends -y -q default-jdk jq
#
# # OS X
# brew tap caskroom/cask
# brew cask install caskroom/versions/java8
# brew install jq

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" > /dev/null && pwd)"
REPO_ROOT="$DIR/../"
VERSION="0.0.1"
if [ -z "$VERSION" ]; then
    echo "ERROR: $REPO_ROOT/VERSION is empty"
    exit 1
fi

codegen_file=/tmp/openapi-generator-cli.jar
# Browse all versions in: https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/
codegen_uri="https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.4.0/openapi-generator-cli-5.4.0.jar"
if ! [ -f "$codegen_file" ]; then
    curl -L "$codegen_uri" -o "$codegen_file"
fi

pushd "$(dirname "$0")"

CURRENT_DIR="$(pwd)"
DIR="$CURRENT_DIR/python_http_client"
swagger_file="$CURRENT_DIR/swagger/frs_api_single_file.swagger.json"

echo "Removing old content in DIR first."
rm -rf "$DIR"

echo "Generating python code from swagger json in $DIR."
java -jar "$codegen_file" generate -g python -i "$swagger_file" -o "$DIR" -c <(echo '{
    "packageName": "frs_api",
    "packageVersion": "'"$VERSION"'",
    "packageUrl": "https://github.com/feast-dev/feast"
}')

echo "Building the python package in $DIR."
pushd "$DIR"
python3 setup.py --quiet sdist
popd

popd
