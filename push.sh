commit=$1

push() {
	git add .
	git commit -m "${commit}"
	git push origin master
}

if [ "${commit}" ]
then
	push
else
	echo "please add commit and excute again!"
fi
