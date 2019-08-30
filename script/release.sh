set -e

if [[ $(git status --porcelain | wc -c) -ne 0 ]]; then
  echo "Cannot release - uncommitted changes found!"
  exit 1
fi

git push

VERSION=v`grep "VERSION =" secrets.go | awk '{print $4}' | tr -d '"'`
read -rp "Enter changelog to release version $VERSION: " CHANGELOG
read -sp "Enter GitHub password: " PASSWORD
RESPONSE=`http -ba jarmo:$PASSWORD POST "https://api.github.com/repos/jarmo/secrets/releases" tag_name="$VERSION" draft:=true name="secrets $VERSION" body="$CHANGELOG"`

rm -rf dist
mkdir -p dist

for file in `find bin -type file`; do
  DIST_FILE_BASE=`echo $file | awk -F "/" '{name=$3 "-" $2; print name}'`
  DIST_FILE_PATH=dist/$DIST_FILE_BASE-$VERSION.zip
  zip -j $DIST_FILE_PATH $file
  shasum -a 512 $DIST_FILE_PATH > $DIST_FILE_PATH.sha512
done

RELEASE_ID=`echo $RESPONSE | jq -r .id`
for file in `ls -d dist/*`; do
  http -ba jarmo:$PASSWORD POST "https://uploads.github.com/repos/jarmo/secrets/releases/$RELEASE_ID/assets?name=`basename $file`" @$file > /dev/null
done

RESPONSE=`http -ba jarmo:$PASSWORD PATCH "https://api.github.com/repos/jarmo/secrets/releases/$RELEASE_ID" draft:=false`

echo "Release done:"
echo $RESPONSE | jq -r .html_url
