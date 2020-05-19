# Zathuraconf
Create colour schemes configuration files for zathura starting from simple colour palettes provided in form of `.ini` files. The configuration consists of a basic frame configuration (as indicated in `basic.conf`, to which we append the colour schemes by parsing the arguments given in the `.ini` configuration.

See examples below for how to use.

## Installation and requirements
The script is written in Python 3.X; the only additional package needed is the command line interface `click`.  
```
git clone <git repo>
cd <git repo>

pip install click # or use your favourite package manager
```

## Usage
Basic configuration attributes are provided in `basic.conf`: feel free to change at will. 

Colour schemes must be provided following the examples that you can find in `colour_config/*.ini`: make your own or use the default ones. Such `.ini` files must be passed as arguments to the `setup.py` script 
```
python setup.py colour_config/solarized.ini 
```
