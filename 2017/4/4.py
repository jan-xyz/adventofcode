#!/usr/bin/env python


def is_valid_passphrase(passphrase):
    return len(passphrase) == len(set(passphrase)) and len(passphrase) > 0


def is_valid_no_palindrome(passphrase):
    passphrase = [''.join(sorted(word)) for word in passphrase]
    print(passphrase)
    return len(passphrase) == len(set(passphrase)) and len(passphrase) > 0


if __name__ == '__main__':

    with open('./puzzle_input') as file:
        lines = file.readlines()

    valid_passwords = 0
    for line in lines:
        words = line.strip().split(" ")
        if is_valid_no_palindrome(words):
            valid_passwords += 1
        else:
            print(words)

    print(valid_passwords)
