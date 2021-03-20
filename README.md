<h1 align="center">
  Zathuraconf
</h1>

<h2 align="center">
  <img alt="Python" src="https://img.shields.io/badge/python%20-%231572B6.svg?&style=flat&logo=python&logoColor=white"/>
</h2>

<h4 align="center">Configure Zathura theme!</h4>
<h3 align="center">
  <a href="#Installation">Installation</a> •
  <a href="#Usage">Usage</a> •
  <a href="#Examples">Examples</a>
</h3>

Creates colour schemes configuration files for [zathura](https://pwmt.org/projects/zathura/) starting from simple colour palettes provided in form of `.ini` files. To the basic configuration attributes (as indicated in `basic.conf`), we append the colour schemes by parsing the arguments given in the `.ini`.

## Installation
```
git clone https://github.com/gennaro-tedesco/zathuraconf.git 
cd zathuraconf
pip3 install -e .
```

## Usage
```
zathura_conf /path/to/colour/scheme 
```
Colour schemes must be provided following the examples in `config/*.ini`: make your own or use the default ones. Such `.ini` files must be passed as arguments to the script 
```
zathura_conf config/solarized.ini 
```
Basic configuration attributes are provided in `basic.conf`: feel free to change at will. We assume the zathura configuration file that we are creating is in its standard directory `~/.config/zathura/zathurarc` as per default installation. If not, pass its location on your local machine as second command line argument to the setup script
```
zathura_conf config/solarized.ini /path/to/different/zathurarc
```

## Examples
<p align="center">
  <img height="600" src="examples/solarized.png">
</p>

<p align="center">
  <img height="600" src="examples/onedark.png">
</p>

