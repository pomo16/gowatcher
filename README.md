## gowatcher

This repo is gowatcher's workspace, which help to deploy the project to the server.



##### sh build.sh + ${project}

+ This script will execute `git clone` , build the project and copy the output to the service dir.



##### sh clean.sh + ${project}

+ This script will clean the code and the output dir. In the meantime, you have to confirm remove twice.



Due to the slow download speed of Alibaba cloud and the gojieba does not support cross compilation, git method will be abandoned.

##### ~~sh push.sh + ${commit}~~

+ ~~This script will push this repo to github.~~



##### Projects

- [go_spider](https://github.com/pomo16/go_spider)
- [go_monitor](https://github.com/pomo16/go_monitor)

