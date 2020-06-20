import configparser
import click
import os
import pkg_resources
from typing import Dict

zathurarc_file = os.path.expanduser('~/.config/zathura/zathurarc')

@click.command()
@click.argument('ini_file', type=click.Path(exists=True))
@click.argument('zathurarc', type=click.Path(exists=True), default=zathurarc_file)
def parse_colour_schema(ini_file: str, zathurarc: str) -> Dict:
	if not ini_file.endswith('.ini'):
		print('Select .ini file')
		return None
	
	config = configparser.ConfigParser()
	config.read(ini_file)
	try:
		d = dict(config._sections)
		d = d['colours']
	except KeyError as ke:
		print('no colours mapping found')
		return None

	basic_conf_path = pkg_resources.resource_filename(__name__, 'basic.conf')
	with open(basic_conf_path, 'r') as basic_conf:
		with open(zathurarc, 'w') as f:
			[f.write(line) for line in basic_conf]
			f.write('\n# -----------------------')
			f.write('\n# ---- COLOUR SCHEME ----')
			f.write('\n# -----------------------\n')
			f.write('set notification-error-bg	  {}\n'.format(d['bg']))
			f.write('set notification-error-fg	  {}\n'.format(d['error']))
			f.write('set notification-warning-bg  {}\n'.format(d['bg']))
			f.write('set notification-warning-fg  {}\n'.format(d['error']))
			f.write('set notification-bg		  {}\n'.format(d['bg']))
			f.write('set notification-fg		  {}\n'.format(d['error']))
			f.write('set completion-group-bg	  {}\n'.format(d['bg']))
			f.write('set completion-group-fg	  {}\n'.format(d['highlight_active']))
			f.write('set completion-bg			  {}\n'.format(d['page']))
			f.write('set completion-fg			  {}\n'.format(d['text']))
			f.write('set completion-highlight-bg  {}\n'.format(d['highlight']))
			f.write('set completion-highlight-fg  {}\n'.format(d['bg']))
			f.write('set inputbar-bg			  {}\n'.format(d['bg']))
			f.write('set inputbar-fg			  {}\n'.format(d['text']))
			f.write('set statusbar-bg			  {}\n'.format(d['bg']))
			f.write('set statusbar-fg			  {}\n'.format(d['text']))
			f.write('set highlight-color		  {}\n'.format(d['highlight']))
			f.write('set highlight-active-color   {}\n'.format(d['highlight_active']))
			f.write('set default-bg				  {}\n'.format(d['bg']))
			f.write('set default-fg				  {}\n'.format(d['text']))
			f.write('set recolor				  {}\n'.format('true'))
			f.write('set recolor-lightcolor		  {}\n'.format(d['page']))
			f.write('set recolor-darkcolor		  {}\n'.format(d['text']))
			f.write('set index-bg				  {}\n'.format(d['page']))
			f.write('set index-fg                 {}\n'.format(d['text']))

if __name__ == "__main__":
	parse_colour_schema()
