import argparse

from collections import defaultdict


prefix = '''package collection

import (
	"github.com/ravil23/baristaschool/telegrambot/entity"
)

var Quiz = entity.NewQuiz(map[entity.Question]entity.Answer{'''

suffix = '''})'''

def convert_file(file: str):
    header = True
    results = defaultdict(list)
    with open(file) as f:
        for line in f.readlines():
            if header:
                header = False
                continue
            line = line.strip()
            if len(line) == 0:
                continue
            parts = line.split(';')
            if len(parts) <= 3:
                raise RuntimeError(f'invalid line: {line}')
            category, question, correct_option = parts[0], parts[1], parts[2]
            invalid_options = parts[3:]
            if len(invalid_options[0]) == 0:
                raise RuntimeError(f'invalid answers not found: {line}')
            invalid_options_string = '", "'.join([option for option in invalid_options if option != ""])
            converted_string = f'''    "{question}": {{
        CorrectOption: "{correct_option}",
        InvalidOptions: []string{{"{invalid_options_string}"}},
    }},'''
            results[category].append(converted_string)

    print(prefix)
    for category in sorted(results.keys()):
        print('    // ', category)
        print('\n'.join(sorted(results[category])))
        print()
    print(suffix)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Convert raw quiz to GoLang collection.')
    parser.add_argument('files', metavar='F', type=str, nargs='+', help='Input files')

    args = parser.parse_args()
    for file in args.files:
        convert_file(file)
