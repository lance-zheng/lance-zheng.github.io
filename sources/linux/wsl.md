<!-- customize-category: Linux -->

# WSL - Arch

- [WSL - Arch](#wsl---arch)
  - [Pacman mirrors](#pacman-mirrors)
  - [Install Zsh](#install-zsh)
  - [docker](#docker)
  - [Others](#others)

> <https://github.com/yuk7/ArchWSL#archwsl>

## Pacman mirrors

```sh
vim /etc/pacman.d/mirrorlist

#中科大源
Server = https://mirrors.ustc.edu.cn/archlinux/$repo/os/$arch

sudo pacman -Syyu
```

## Install Zsh

```bash
sudo pacman -S zsh
chsh -s /bin/zsh # 切換shell为zsh
```

```sh
# NodeJS
curl -L https://bit.ly/n-install | bash

# zplug
curl -sL --proto-redir -all,https https://raw.githubusercontent.com/zplug/installer/master/installer.zsh | zsh

# on my zsh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

## docker

```sh
sudo pacman -S docker
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -a -G docker $USER
sudo systemctl restart docker
```

## Others

Start Tmux in Windows Termianl WSL startup

```sh
tmux new-session -A -D -s main -n main
```

<img width=500 src='/assets/image/Snipaste_2024-02-13_14-03-40.png'/>
