Enable SSH Key Logon and Disable Password


if ssh sever is not installed:
sudo apt-get install openssh-server
sudo systemctl enable ssh
sudo systemctl start ssh


Create Authentication SSH-Kegen Keys on client
----------------------------------------------

ssh-keygen -t rsa


Upload generated Public Keys to remote
--------------------------------------

ssh-copy-id user@remote


Allow Key log on
----------------

sudo vim /tec/ssh/sshd_config

Then uncomment and change the lines to match the ones below. Make sure these lines are un-commented, meaning they don't have the (#) before it.

PubkeyAuthentication yes
AuthorizedKeyFile  .ssh/authorized_keys
PasswordAuthentication no
ChallengeResponseAuthentication no


Reload Service
--------------
service sshd reload


