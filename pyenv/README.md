Install libraries (not necessary)
---------------------------------
sudo apt install curl git-core gcc make zlib1g-dev libbz2-dev libreadline-dev libsqlite3-dev libssl-dev (ubuntu)


install pyenv
-------------
sudo git clone https://github.com/pyenv/pyenv.git $HOME/.pyenv


Edit file
---------
sudo vim $HOME/.bashrc


Add following text
------------------
export PYENV_ROOT="$HOME/.pyenv"
export PATH="$PYENV_ROOT/bin:$PATH"

if command -v pyenv 1>/dev/null 2>&1; then
  eval "$(pyenv init -)"
fi


Activate source
---------------
source $HOME/.bashrc


Install virtual-env
-------------------
sudo git clone https://github.com/pyenv/pyenv-virtualenv.git $(pyenv root)/plugins/pyenv-virtualenv


Edit file
---------
sudo vim $HOME/.bashrc


Edit section
------------
if command -v pyenv 1>/dev/null 2>&1; then
  eval "$(pyenv init -)"
  eval "$(pyenv virtualenv-init -)"
fi


Activate source
---------------
source $HOME/.bashrc



Useful commands
---------------
List all versions available: 
pyenv install -l

Install a specific version: 
pyenv install <version>

List installed versions:
pyenv versions

Create a virtual environment: 
pyenv virtualenv <version> my-virtual-env


Activate/Deactivate an environment: 
pyenv activate my-virtual-env
pyenv deactivate my-virtual-env

Uninstall environment: 
pyenv uninstall my-virtual-env
