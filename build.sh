project=$1

build() {
   	if [ ! -d "code/${project}" ]
	then
   		git clone git@github.com:pomo16/${project}.git code/${project}
	else
        	rm -rf code/${project}
       		git clone git@github.com:pomo16/${project}.git code/${project}
	fi
	echo "${project} code ready..."

	cd code/${project}/
	sh build.sh
	cd ../../

	if [ ! -d "service/${project}" ]
	then
		cp -R code/${project}/output service/${project}
		cp service/.gitkeep service/${project}/${project}_log/
	else
		rm -rf service/${project}
		cp -R code/${project}/output service/${project}
		cp service/.gitkeep service/${project}/${project}_log/
	fi
	echo "${project} build finish!"
}

if [[ ${project} = "go_spider" || ${project} = "go_monitor" ]]
then
	build
else
	echo "project ${project} not exist!"
fi
