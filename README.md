<h1 align="center">
  Zathuraconf
</h1>

<h4 align="center">Configure Zathura theme!</h4>
<h3 align="center">
  <a href="#Installation">Installation</a> •
  <a href="#Usage">Usage</a> •
  <a href="#Examples">Examples</a>
</h3>

Create custom configurations for [zathura](https://pwmt.org/projects/zathura/) in one line: provide a colour configuration as `json` et voilà!

## Installation
Go get it!
```
go get -u -v github.com/gennaro-tedesco/zathuraconf
```

## Usage
Choose a colour configuration [from the examples](https://github.com/gennaro-tedesco/zathuraconf/tree/gomodule/examples) or make your own `config.json`:
```
{
  "page": "#3c3836", <-- change colours here
  "text": "#bdae93",
  "background": "#282828",
  "highlight": "#fabd2f",
  "highlight_active": "#98971a",
  "error": "#d65d0e"
}
```
then simply run
```
zathuraconf config.json
```

![demo](https://user-images.githubusercontent.com/15387611/114107808-f327f100-98d1-11eb-885f-27ff76b2504d.gif)

This creates your custom zathura configuration file in the default path `~/.config/zathura/zathurarc`; if your zathura settings are elsewhere, you can provide a different path via `zathuraconf config.json -p /path/to/config/`. 

See `zathuraconf -h` for details.


### Change defaults
The colour schemes settings are appended to a default configuration of zathura commands [specified here](https://github.com/gennaro-tedesco/zathuraconf/blob/6b64a7814737bdb930e7f44e53a0a407c6ab3a01/cmd/config.go#L16-L39): these defaults are standard and natural, however should you feel the need to change them, modify those lines at will.

## Examples
<p align="center">
  <img height="600" src="examples/solarized.png">
</p>

<p align="center">
  <img height="600" src="examples/onedark.png">
</p>


## Tests
Run unit tests with
```
go test -v ./cmd
```

## Feedback
If you find this application useful consider awarding it a ⭐, it is a great way to give feedback! Otherwise, any additional suggestions or merge request is warmly welcome!

