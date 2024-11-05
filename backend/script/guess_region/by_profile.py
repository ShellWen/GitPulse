import re

import httpx
from bs4 import BeautifulSoup
from loguru import logger

import timezone_util
import flag
import pycountry

from by_email import guess_by_profile_email


def guess_by_profile_timezone(handle: str):
    url = f"https://github.com/{handle}"
    html_content = httpx.get(url).text
    soup = BeautifulSoup(html_content, 'html.parser')
    profile_timezone = soup.find('profile-timezone')
    if profile_timezone is None:
        return None
    timezone = profile_timezone.text
    m = re.findall(r'([+-]\d{2}:\d{2})', timezone)
    if m is None or len(m) != 1:
        return None
    timezone: str = m[0].replace(':', '')
    return timezone_util.get_countries_by_timezone(timezone)

def guess_by_profile_country_name(handle: str):
    url = f"https://github.com/{handle}"
    html_content = httpx.get(url).text
    soup = BeautifulSoup(html_content, 'html.parser')
    profile_name = soup.find('span', class_='p-name vcard-fullname d-block overflow-hidden')
    profile_bio = soup.find('div', class_='p-note user-profile-bio mb-3 js-user-profile-bio f4')
    profile_readme = soup.find('turbo-frame', class_='user-profile-frame')
    if profile_name is None or profile_bio is None or profile_readme is None:
        return None
    text = profile_name.text + profile_bio.text + profile_readme.text
    # use pycountry to extract country names in text
    countries = find_country(text)

    return countries

def guess_by_profile_country_flag(handle: str):
    url = f"https://github.com/{handle}"
    html_content = httpx.get(url).text
    soup = BeautifulSoup(html_content, 'html.parser')
    profile_name = soup.find('span', class_='p-name vcard-fullname d-block overflow-hidden')
    profile_bio = soup.find('div', class_='p-note user-profile-bio mb-3 js-user-profile-bio f4')
    profile_readme = soup.find('article', class_='markdown-body entry-content container-lg f5')
    text = ""
    if profile_name is not None:
        text += profile_name.text
    if profile_bio is not None:
        text += profile_bio.text
    if profile_readme is not None:
        text += profile_readme.text
    if text is None:
        return None
    flags = [c for c in text if is_flag_emoji(c)]
    if len(flags) == 0 or len(flags) % 2 != 0:
        return None
    flags = ["".join(flags[i:i + 2]) for i in range(0, len(flags), 2)]
    country_name = [flag.dflagize(the_flag) for the_flag in flags]
    country_name = [name[1:3].lower() for name in country_name]
    country_name_set = set(country_name)

    return country_name_set

def is_flag_emoji(c):
    return "\U0001F1E6\U0001F1E8" <= c <= "\U0001F1FF\U0001F1FC" or c in ["\U0001F3F4\U000e0067\U000e0062\U000e0065\U000e006e\U000e0067\U000e007f", "\U0001F3F4\U000e0067\U000e0062\U000e0073\U000e0063\U000e0074\U000e007f", "\U0001F3F4\U000e0067\U000e0062\U000e0077\U000e006c\U000e0073\U000e007f"]

def find_country(text):
    full_list =[]
    countries = sorted(pycountry.countries, key=lambda x: -len(x))
    for country in countries:
        if country.alpha_2.lower() in text.lower():
            full_list.append(country)
    return full_list

if __name__ == '__main__':
    logger.info(f"Guessed countries by profile: {guess_by_profile_timezone('ShellWen')}")
    logger.info(f"Guessed countries by profile: {guess_by_profile_country_name('ShellWen')}")
    logger.info(f"Guessed countries by profile: {guess_by_profile_country_flag('ShellWen')}")
