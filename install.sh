#!/bin/bash
Red='\033[1;31m'
Blue='\033[1;34m'
echo "Installing ck8s.... "

sudo make
sudo mkdir -p ${HOME}/.ck8s
sudo mv -v ./ck8s /usr/local/bin/ck8s

echo -e "${Blue} Installed Successfully!"
echo " "
echo  "
        __     ______         
  ____ |  | __/  __  \  ______
_/ ___\|  |/ />      < /  ___/
\  \___|    </   --   \\___ \ 
 \___  >__|_ \______  /____  >
     \/     \/      \/     \/ 
"
