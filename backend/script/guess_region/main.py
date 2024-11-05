import os
import json
import sys
from traceback import print_list

from github import Github, Auth
from loguru import logger

from by_profile import guess_by_profile_email, guess_by_profile_timezone, guess_by_profile_country_flag, guess_by_profile_country_name
from by_timezone import guess_by_timezone

EMAIL_GUESS_WEIGHT = float(0.15)
PROFILE_TIMEZONE_GUESS_WEIGHT = float(0.20)
COUNTRY_NAME_GUESS_WEIGHT = float(0.25)
COUNTRY_FLAG_GUESS_WEIGHT =float(0.05)
TIMEZONE_GUESS_WEIGHT = float(0.15)

assert EMAIL_GUESS_WEIGHT + PROFILE_TIMEZONE_GUESS_WEIGHT + COUNTRY_NAME_GUESS_WEIGHT + COUNTRY_FLAG_GUESS_WEIGHT + TIMEZONE_GUESS_WEIGHT == 0.8

auth = Auth.Token(os.environ['GITHUB_API_TOKEN'])

def main():
    argv = sys.argv

    if len(argv) != 2:
        logger.error("Usage: python main.py USERNAME")
        return

    USERNAME = argv[1]

    confidence_dict = dict()
    with Github(auth=auth) as g:
        email_guess = guess_by_profile_email(USERNAME, g)
        if email_guess is not None:
            for country in email_guess:
                confidence_dict[country] = confidence_dict.get(country, 0) + EMAIL_GUESS_WEIGHT

        profile_timezone_guess = guess_by_profile_timezone(USERNAME)
        if profile_timezone_guess is not None:
            for country in profile_timezone_guess:
                confidence_dict[country] = confidence_dict.get(country, 0) + PROFILE_TIMEZONE_GUESS_WEIGHT

        country_name_guess = guess_by_profile_country_name(USERNAME)
        if country_name_guess is not None:
            for country in country_name_guess:
                confidence_dict[country] = confidence_dict.get(country, 0) + COUNTRY_NAME_GUESS_WEIGHT

        country_flag_guess = guess_by_profile_country_flag(USERNAME)
        if country_flag_guess is not None:
            for country in country_flag_guess:
                confidence_dict[country] = confidence_dict.get(country, 0) + COUNTRY_FLAG_GUESS_WEIGHT

        timezone_guess = guess_by_timezone(USERNAME, g)
        if timezone_guess is not None:
            for country in timezone_guess:
                confidence_dict[country] = confidence_dict.get(country, 0) + TIMEZONE_GUESS_WEIGHT

        print(json.dumps(
                    confidence_dict,
                    separators=(',', ': ')
                ))


if __name__ == '__main__':
    main()
