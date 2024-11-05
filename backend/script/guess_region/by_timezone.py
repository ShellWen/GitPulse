import re

from github import Github
from github.Requester import Requester
from loguru import logger
import timezone_util


def get_latest_push_events(handle: str, max_number: int, gh: Github):
    events = gh.get_user(handle).get_events()
    retval = list()
    current_page = 0
    while len(retval) < max_number:
        page = events.get_page(current_page)
        for event in page:
            if event.type == 'PushEvent':
                retval.append(event)
            if len(retval) >= max_number:
                break
        current_page += 1
    return retval


def get_commits_from_push_event(event):
    return event.payload['commits']


def evaluate_commit_timezone(commit_url: str, gh: Github):
    # Since PyGithub doesn't support getting a commit patch directly, we have to use the REST API
    requester: Requester = vars(gh)['_Github__requester']
    assert isinstance(requester, Requester)
    headers, data = requester.requestBlobAndCheck('GET', commit_url,
                                                  headers={'Accept': 'application/vnd.github.v3.patch'})
    response = data['data']
    date_line = None
    for line in response.split('\n'):
        if line.startswith('Date: '):
            date_line = line[6:]
            break
    # Fri, 11 Mar 2022 14:00:00 +0800
    assert date_line is not None
    m = re.match(
        r'(((Mon|Tue|Wed|Thu|Fri|Sat|Sun))[,]?\s[0-9]{1,2})\s(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)\s([0-9]{4})\s([0-9]{2}):([0-9]{2})(:([0-9]{2}))?\s([\+|\-][0-9]{4})\s?',
        date_line)
    tz_group = m.group(10)
    assert tz_group is not None
   # logger.info(f'Found timezone {tz_group}...')
    return tz_group


def guess_by_timezone(handle: str, gh: Github):
    #logger.info(f'Guessing timezone by commits\' timezone for {handle}...')
    events = get_latest_push_events(handle, 3, gh)
    commits = [commit for event in events for commit in get_commits_from_push_event(event)]
    if len(commits) == 0:
        return None
    #logger.info(f'Found {len(commits)} commits...')
    commit_urls = [commit['url'] for commit in commits]
    timezones = [evaluate_commit_timezone(commit_url, gh) for commit_url in commit_urls]
    timezone_counts = dict()
    for tz in timezones:
        if tz not in timezone_counts:
            timezone_counts[tz] = 0
        timezone_counts[tz] += 1
    max_count = 0
    max_tz = None
    for tz, count in timezone_counts.items():
        if count > max_count:
            max_count = count
            max_tz = tz
    timezones = timezone_util.TIMEZONE_OFFSETS[max_tz]
    countries = set()
    for tz in timezones:
        if tz in timezone_util.TZ_COUNTRIES:
            countries.update(timezone_util.TZ_COUNTRIES[tz])
    return countries


if __name__ == '__main__':
    from github import Auth
    import os

    auth = Auth.Token(os.environ['GITHUB_TOKEN'])
    with Github(auth=auth) as g:
        logger.info(f"Guessed countries by commits' timezone: {guess_by_timezone("ShellWen", g)}")
