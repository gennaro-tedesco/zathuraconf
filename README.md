<h1 align="center">
  zathuraconf
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
Choose a colour configuration [from the examples](https://github.com/gennaro-tedesco/zathuraconf/tree/master/examples) or make your own `config.json`:
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
The colour schemes settings are appended to a default configuration of zathura commands [specified here](https://github.com/gennaro-tedesco/zathuraconf/blob/470c5d12378c8b29eff85b58818e0daa844edff7/cmd/config.go#L14-L41): these defaults are standard and natural, however should you feel the need to change them, modify those lines at will.

## Examples

<p align="center">
  <img alt="solarized" src="https://user-images.githubusercontent.com/15387611/114108427-5cf4ca80-98d3-11eb-8b39-99600dd42807.png" width="45%">
&nbsp; &nbsp; &nbsp; &nbsp;
  <img alt="onedark" src="https://user-images.githubusercontent.com/15387611/114108475-6ed66d80-98d3-11eb-9a45-4bd992d33b29.png" width="45%">
</p>


## Tests
Run unit tests with
```
go test -v ./cmd
```

## Feedback
If you find this application useful consider awarding it a ⭐, it is a great way to give feedback! Otherwise, any additional suggestions or merge request is warmly welcome!

