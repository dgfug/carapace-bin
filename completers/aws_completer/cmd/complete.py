#!/usr/bin/env python
# Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.

# Licensed under the Apache License, Version 2.0 (the "License"). You
# may not use this file except in compliance with the License. A copy of
# the License is located at

#     http://aws.amazon.com/apache2.0/

# or in the "license" file accompanying this file. This file is
# distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF
# ANY KIND, either express or implied. See the License for the specific
# language governing permissions and limitations under the License.
import asyncio
import json
import os
import sys

from pprint import pprint

if os.environ.get('LC_CTYPE', '') == 'UTF-8':
    os.environ['LC_CTYPE'] = 'en_US.UTF-8'
from awscli.autocomplete.main import create_autocompleter
from awscli.clidriver import create_clidriver


def main():
    # bash exports COMP_LINE and COMP_POINT, tcsh COMMAND_LINE only
    command_line = (
        os.environ.get('COMP_LINE') or os.environ.get('COMMAND_LINE') or ''
    )
    command_index = int(os.environ.get('COMP_POINT') or len(command_line))

    try:
        args = str.split(command_line)
        completer = create_autocompleter(driver=create_clidriver(args))

        results = completer.autocomplete(command_line, command_index)
       
        loop = asyncio.get_event_loop()
        parsed = completer._parser.parse(command_line, None)
        groups = [None] * len(completer._completers)
        for num, completer in enumerate(completer._completers):
            groups[num] = asyncio.gather(complete(completer, parsed))
        all_groups = asyncio.gather(*groups)
        results = loop.run_until_complete(all_groups)
        loop.close()

        sys.stdout.write(json.dumps(
            [{'name': result.name, 'help_text': _get_display_meta(result)} for result in flatten(flatten(results))]))
    except KeyboardInterrupt:
        # If the user hits Ctrl+C, we don't want to print
        # a traceback to the user.
        pass

async def complete(completer, parsed):
    result = completer.complete(parsed)
    if result is not None:
        return result
    return []

def flatten(t):
    return [item for sublist in t for item in sublist]

def _get_display_meta(completion):
    display_meta = ''
    type_name = getattr(completion, 'cli_type_name', None)
    help_text = getattr(completion, 'help_text', None)
    if type_name:
        display_meta += f'[{type_name}] '
    if help_text:
        display_meta += f'{help_text}'
    return display_meta


if __name__ == '__main__':
    main()
