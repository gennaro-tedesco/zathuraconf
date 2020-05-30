from setuptools import setup

setup(
    name='zathura_conf',
    version='1.0',
	author='Gennaro Tedesco',
	author_email='gennarotedesco@gmail.com',
	description='change zathura configuration dynamically',
	url='https://github.com/gennaro-tedesco/zathuraconf',
	install_requires=['click'],
    entry_points={
        'console_scripts': [
            'zathura_conf=zathura_conf:parse_colour_schema'
        ]
    }
)
