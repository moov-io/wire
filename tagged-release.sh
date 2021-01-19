#!/bin/bash

[ $# -eq 0 ] && { echo "Usage: $0 version_tag"; exit 1; }

newVersionTag=$1
if ! [[ "$newVersionTag" =~ v[0-9]+\.[0-9]+\.[0-9]+ ]]
then
  echo "invalid version tag $newVersionTag"
  exit 1
fi

# make sure this is a new tag
if git show-ref --tags | grep -q "$newVersionTag"; then
  echo "$newVersionTag already exists!"
  exit
fi

# aborts if the git command fails
if ! current=$(git describe --tags --abbrev=0); then
    exit 1
fi
parts=( $( echo "$current" | grep -o -E '[0-9]+') )
major=${parts[0]}
minor=${parts[1]}
bug=${parts[2]}

newMajorVersion=v$((major+1)).0.0
newMinorVersion=v$major.$((minor+1)).0
newBugVersion=v$major.$minor.$((bug+1))

if [[ "$newVersionTag" == "$newMajorVersion" ]]; then
  echo "releasing new major version $newVersionTag (currently on $current)"
elif [[ "$newVersionTag" == "$newMinorVersion" ]]; then
  echo "releasing new minor version $newVersionTag (currently on $current)"
elif [[ "$newVersionTag" == "$newBugVersion" ]]; then
  echo "releasing new bug fix $newVersionTag (currently on $current)"
else
  echo cannot increment version from "$current" to "$newVersionTag"
  exit 1
fi

read -p "Do you want to continue? (y/n) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]
then
    echo "aborting release"
    exit 0
fi

# releases should be created from master branch
if [ "$(git branch --show-current)" != "master" ]
then
    echo "switching to branch 'master'"
    git checkout master > /dev/null 2>&1
fi

# only 2 files should be changed
status=$(git status --porcelain)
if [ "$(echo "$status" | wc -l)" -ne 2 ]
then
  echo "Only version.go and CHANGELOG.md should be updated!"
  if [ -n "$status" ]
  then
    printf "Pending changes:\n%s\n" "$status"
  fi
  exit
elif ! echo "$status" | grep -q "CHANGELOG.md" && ! echo "$status" | grep -q "version.go"
then
    echo "version.go and changelog.md must be updated to proceed"
    exit
fi

firstLine=$(head -n 1 CHANGELOG.md)
if ! echo "$firstLine" | grep -q "$newVersionTag"; then
  echo "new tag ($newVersionTag) doesn't match CHANGELOG ($firstLine)"
  exit
fi

expectedHeader=$(printf "## $newVersionTag (Released %s)" "$(date +"%Y-%m-%d")")
if [ "$firstLine" != "$expectedHeader" ]
then
  echo "Did you update the CHANGELOG's header? Expected \"$expectedHeader\", found \"$firstLine\""
  exit
fi

git add CHANGELOG.md version.go
git commit -m "release $newVersionTag"
git tag "$newVersionTag"
git push origin master
git push origin "$newVersionTag"
