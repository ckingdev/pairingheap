# Thanks to http://stackoverflow.com/a/21142256

echo "mode: set" > acc.out
for Dir in $(find ./* -maxdepth 10 -type d );
do
        if ls $Dir/*.go &> /dev/null;
        then
            go test -coverprofile=profile.out $Dir
            if [ -f profile.out ]
            then
                cat profile.out | grep -v "mode: set" >> acc.out
            fi
fi
done
goveralls -coverprofile=acc.out -repotoken $REPO_TOKEN
rm -rf ./profile.out
rm -rf ./acc.out