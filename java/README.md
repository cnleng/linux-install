Install Java Development Kit and Runtime

Using package manager
---------------------
sudo apt-get install openjdk-8-jdk

Manual install
--------------
sudo wget https://download.java.net/java/ga/jdk11/openjdk-11_linux-x64_bin.tar.gz
sudo sha256sum openjdk-11_linux-x64_bin.tar.gz
sudo tar xzvf openjdk-11_linux-x64_bin.tar.gz
sudo mkdir /usr/lib/jvm
sudo mv jdk-11 /usr/lib/jvm/openjdk-11-manual-installation/
sudo update-alternatives --install /usr/bin/java java /usr/lib/jvm/openjdk-11-manual-installation/bin/java 1
sudo update-alternatives --install /usr/bin/javac javac /usr/lib/jvm/openjdk-11-manual-installation/bin/javac 1

update-alternatives --config java
