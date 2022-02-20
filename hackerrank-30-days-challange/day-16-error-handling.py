# Day 16: Exceptions - String to Integer

# Task Read a string, S, and print its integer value; if cannot be converted to an integer, print Bad String.

#!/bin/python3

import math
import os
import random
import re
import sys

if __name__ == '__main__':
    S = input()
    try:
        S = int(S)
        print(S)
    except:
        print("Bad String")
