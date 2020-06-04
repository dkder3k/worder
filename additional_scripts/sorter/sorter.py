import click
import json
from string import ascii_lowercase


@click.command()
@click.argument('input_file', type=click.Path(exists=True))
@click.argument('output_file', type=click.Path(exists=False))
def sort(input_file, output_file):

    words_dict = dict()
    for c in ascii_lowercase:
        words_dict[c] = list()

    with open(input_file, 'r') as f:
        for line in f:
            words_dict[line[0]].append(line.rstrip())

    with open(output_file, 'w') as f:
        f.write(str(json.dumps(words_dict)))
