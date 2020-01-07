project=$1

clean() {
    if [ -d "code/${project}" ]
	then
        rm -r code/${project}
	fi
	echo "${project} code removed..."

	if [ -d "service/${project}" ]
	then
        rm -r service/${project}
	fi
	echo "${project} remove finish!"
}

if [[ ${project} = "go_spider" || ${project} = "go_monitor" ]]
then
	clean
else
	echo "project ${project} not exist!"
fi
