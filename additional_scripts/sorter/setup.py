from setuptools import setup

setup(
    name='sorter',
    version='0.1',
    py_modules=['sorter'],
    install_requires=[
        'Click',
    ],
    entry_points='''
        [console_scripts]
        sorter=sorter:sort
    ''',
)
