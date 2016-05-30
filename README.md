# git-orchestra
## Execute git commands for multiple local repositories

I developed this tool primarily to update all git repositories I have installed. Before I always forgot that I made some minor changes on other devices or even the server and so I often bothered with merge conflicts etc.
Nevertheless it is also possible to run ```git commit``` and ```git push```.

##Installing
Fetching the repository:

    git clone github.com/matschiner/git-orchestra

Git Orchestra is written in Go lang. Therefore building the utility on a mac could look like this:

    go build -o /usr/local/bin/gito
    
Alternatively you can also copy the precompiled binary from the bin folder to your Path, which is in default /usr/local/bin on Mac   

    cp bin/gito_mac /usr/local/bin/gito # on a mac
    cp bin/gito_linux /bin/gito


## How git-orchestra works

These are the commands:

```gito <respository_subfolder>``` .... add a already initialized git repository in a subfolder

```gito rm <respository_subfolder>``` ..... deleting it

```gito list``` ......................... lists all added git paths

```gito pull``` ......................... iterates through every added git path and pulls them separately

```gito commit``` ....................... the same but commits them – you will be asked for commit messages

```gito push``` ......................... the same but pushes them


Git orchestra creates a JSON config file named ```.gitpullconfig``` in the current folder and expects the repositories as direct subfolders.


Hope this will help you!
