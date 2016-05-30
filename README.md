# git-orchestra
## Execute git commands for multiple local repositories

I developed this tool primarily to update all git repositories I have installed. Before I always forgot that I made some minor changes on other devices or even the server and so I often bothered with merge conflicts etc.
Nevertheless it is also possible to run $git commit$ and $git push$.

## How git-orchestra works

These are the commands:

    gito add respository_subfolder ... add a already initialized git repository in a subfolder
    gito rm respository_subfolder ... deleting it
    gito list ... lists all added git paths
    gito pull ... iterates through every added git path and pulls them separately
    gito commit ... the same but commits them – you will be asked for commit messages
    gito push ... the same but pushes them


Git orchestra creates a JSON config file in the current folder and expects the repositories as direct subfolders.


Hope this will help you!
